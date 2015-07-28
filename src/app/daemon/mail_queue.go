/**
 * Copyright 2015 @ S1N1 Team.
 * name : mail_queue
 * author : jarryliu
 * date : 2015-07-27 17:06
 * description :
 * history :
 */
package daemon

import (
	"go2o/src/core/domain/interface/partner/mss"
	mssIns "go2o/src/core/infrastructure/mss"
	"go2o/src/core/variable"
	"time"
)

func startMailQueue() {
	if i, _ := gCTX.Storage().GetInt(variable.KvNewMailTask); i > 0 {
		var list = []*mss.MailTask{}
		gCTX.Db().GetOrm().Select(&list, "is_send = 0 OR is_failed = 1")
		mailChan = make(chan int, len(list))
		for _, v := range list {
			go func(ch chan int, t *mss.MailTask) {
				err := mssIns.SendMailWithDefaultConfig(t.Subject, []string{t.SendTo}, []byte(t.Body))
				if err != nil {
					gCTX.Log().PrintErr(err)
					t.IsFailed = 1
					t.IsSend = 1
				} else {
					t.IsSend = 1
					t.IsFailed = 0
				}
				t.SendTime = time.Now().Unix()
				gCTX.Db().GetOrm().Save(t.Id, t)
				mailChan <- 0
			}(mailChan, v)
			<-mailChan
		}
		gCTX.Storage().Set(variable.KvNewMailTask, 0)
	}

	time.Sleep(time.Second * 15)

	startMailQueue()
}