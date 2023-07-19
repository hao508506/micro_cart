package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id        int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:int(11);comment:id"`
	ProductId int64      `json:"product_id" form:"name" binding:"required" gorm:"type:int(11);not null; comment:名称"`
	Num       int64      `json:"num" gorm:"type:int(11);unique;not null; comment:sku"`
	SizeId    int64      `json:"size_id" gorm:"type:int(11);comment:价格"`
	UserId    int64      `json:"uid" gorm:"type:int(11);comment:描述"`
	CreateAt  time.Time  `json:"createAt" gorm:"type:datetime;autoCreateTime"`
	UpdateAt  *time.Time `json:"updateAt" gorm:"type:datetime;"`
	DeleteAt  int64      `json:"deleteAt" gorm:"type:int(11);"`
}

func (u *Cart) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdateAt", time.Now())
	return nil
}
