/**
 * Copyright 2015 @ z3q.net.
 * name : shop_manager.go
 * author : jarryliu
 * date : 2016-05-28 12:13
 * description :
 * history :
 */
package shop

import (
	"go2o/core/domain/interface/enum"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/merchant/shop"
	"strings"
	"time"
)

var _ shop.IShopManager = new(shopManagerImpl)

type shopManagerImpl struct {
	_merchant merchant.IMerchant
	_rep      shop.IShopRep
	_shops    []shop.IShop
}

func NewShopManagerImpl(m merchant.IMerchant, rep shop.IShopRep) shop.IShopManager {
	return &shopManagerImpl{
		_merchant: m,
		_rep:      rep,
	}
}

// 新建商店
func (this *shopManagerImpl) CreateShop(v *shop.Shop) shop.IShop {
	v.CreateTime = time.Now().Unix()
	v.MerchantId = this._merchant.GetAggregateRootId()
	return newShop(this, v, this._rep)
}

// 获取所有商店
func (this *shopManagerImpl) GetShops() []shop.IShop {
	if this._shops == nil {
		shops := this._rep.GetShopsOfMerchant(this._merchant.GetAggregateRootId())
		this._shops = make([]shop.IShop, len(shops))
		for i, v := range shops {
			this._shops[i] = this.CreateShop(v)
		}
	}
	return this._shops
}

// 根据名称获取商店
func (this *shopManagerImpl) GetShopByName(name string) shop.IShop {
	name = strings.TrimSpace(name)
	for _, v := range this.GetShops() {
		if strings.TrimSpace(v.GetValue().Name) == name {
			return v
		}
	}
	return nil
}

// 获取营业中的商店
func (this *shopManagerImpl) GetBusinessInShops() []shop.IShop {
	var list []shop.IShop = make([]shop.IShop, 0)
	for _, v := range this.GetShops() {
		if v.GetValue().State == enum.ShopBusinessIn {
			list = append(list, v)
		}
	}
	return list
}

// 获取商店
func (this *shopManagerImpl) GetShop(shopId int) shop.IShop {
	//	v := this.rep.GetValueShop(this.GetAggregateRootId(), shopId)
	//	if v == nil {
	//		return nil
	//	}
	//	return this.CreateShop(v)
	shops := this.GetShops()

	for _, v := range shops {
		if v.GetValue().Id == shopId {
			return v
		}
	}
	return nil
}

// 删除门店
func (this *shopManagerImpl) DeleteShop(shopId int) error {
	//todo : 检测订单数量
	mchId := this._merchant.GetAggregateRootId()
	s := this.GetShop(shopId)
	if s != nil {
		this.Reload() //重新加载数据
		switch s.Type() {
		case shop.TypeOfflineShop:
			return this.deleteOfflineShop(mchId, s)

		case shop.TypeOnlineShop:
			return this.deleteOnlineShop(mchId, s)
		}
	}
	return nil
}

func (this *shopManagerImpl) deleteOnlineShop(mchId int, s shop.IShop) error {
	shopId := s.GetDomainId()
	err := this._rep.DeleteOnlineShop(mchId, shopId)
	return err
}

func (this *shopManagerImpl) deleteOfflineShop(mchId int, s shop.IShop) error {
	shopId := s.GetDomainId()
	err := this._rep.DeleteOfflineShop(mchId, shopId)
	return err
}

// 重新加载数据
func (this *shopManagerImpl) Reload() {
	this._shops = nil
	//todo:  如果系统后台和前台,无法同时清理缓存
}
