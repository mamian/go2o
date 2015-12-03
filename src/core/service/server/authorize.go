/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package server

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jsix/gof/net/jsv"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/partner"
	"go2o/src/core/service/dps"
	"strconv"
)

func Verify(m *jsv.Args) (memberId int, err error) {
	member_id, token := (*m)["member_id"].(string), (*m)["token"].(string)

	if memberId, err = strconv.Atoi(member_id); err != nil || token == "" {

		jsv.Println(err)
		return memberId, errors.New("invalid parameter")
	} else {
		//todo:  用户令牌更改，需修改。
		return memberId, nil
	}

	rds := Redis().Get()
	defer rds.Close()

	sessKey := fmt.Sprintf("dps:session:m%d", memberId)
	servToken, err := redis.String(rds.Do("GET", sessKey))
	if err != nil {
		return memberId, member.ErrInvalidSession
	}

	//	if jsv.Context.Debug() {
	//		jsv.Println("[Member][Verify]", memberId, token, servToken)
	//	}

	if servToken != token {
		return memberId, member.ErrSessionTimeout
	} else {
		rds.Do("SETEX", sessKey, 3600*3, token) //更新回话并延长时间
	}
	return memberId, nil
}

func VerifyPartner(m *jsv.Args) (partnerId int, err error, p *partner.ValuePartner) {
	partnerId, err = strconv.Atoi((*m)["partner_id"].(string))
	postSecret := (*m)["secret"].(string)
	if postSecret == "" {
		return partnerId, errors.New("missing token secret!"), nil
	}

	p, err = dps.PartnerService.GetPartner(partnerId)
	if p == nil {
		return partnerId, errors.New("no such partner"), nil
	}

	pa := dps.PartnerService.GetApiInfo(partnerId)
	if pa.ApiSecret != postSecret {
		return partnerId, errors.New("not authorized"), nil
	}

	return partnerId, err, p
}
