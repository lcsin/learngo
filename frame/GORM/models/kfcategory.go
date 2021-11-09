package models

import "time"

type KfCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:int;comment:主键id"`
	Name      string    `json:"name" gorm:"type:varchar(30);comment:类名"` //
	UpdateBy  uint      `json:"updateBy" gorm:"type:int;comment:修改人"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;comment:创建日期"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;comment:修改日期"`
	Image     string    `json:"image" gorm:"type:varchar(150);comment:Image"`         //
	ParentId  uint      `json:"parentId" gorm:"type:int;comment:父类id"`                //
	IsDelete  uint      `json:"isDelete" gorm:"type:tinyint;comment:IsDelete"`        //
	Sort      int       `json:"sort" gorm:"type:int;comment:Sort"`                    //
	ForceH5   int       `json:"forceH5" gorm:"type:tinyint;comment:0不设置 1强制为移动端浏览模式"` //
	Host      string    `json:"host" gorm:"type:varchar(64);comment:Host"`            //
	Hidden    int       `json:"hidden" gorm:"type:tinyint;comment:首页隐藏"`              //
}

func (KfCategory) TableName() string {
	return "category"
}
