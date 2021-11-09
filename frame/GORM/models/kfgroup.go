package models

import "time"

type KfGroup struct {
	ID            uint      `json:"id" gorm:"primaryKey;type:int;comment:主键id"`
	Name          string    `json:"name" gorm:"type:varchar(20);comment:姓名,用于搜索"`          //
	IsAccept      int       `json:"isAccept" gorm:"type:tinyint;comment:是否开启自动接单 0关闭 1开启"` //
	MaxAccept     int       `json:"maxAccept" gorm:"type:int;comment:最大接单量"`               //
	AcceptStartAt int       `json:"acceptStartAt" gorm:"type:int;comment:自动接单开始时间"`        //
	AcceptEndAt   int       `json:"acceptEndAt" gorm:"type:int;comment:自动接单结束时间"`          //
	CreatedBy     uint      `json:"createdBy" gorm:"type:int;comment:创建人"`                 //
	CreatedAt     time.Time `json:"createdAt" gorm:"type:timestamp;comment:创建日期"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"type:timestamp;comment:修改日期"`
}

func (KfGroup) TableName() string {
	return "group"
}
