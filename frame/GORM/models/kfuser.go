package models

import "time"

type KfUsers struct {
	ID        uint      `json:"id" gorm:"primaryKey;type:int;comment:主键id"`
	RoleID    uint      `json:"RoleID" gorm:"type:int;comment:角色ID"`
	Username  string    `json:"username" gorm:"type:varchar(150);comment:Username"` //
	Password  string    `json:"password" gorm:"type:varchar(150);comment:Password"` //
	ApiToken  string    `json:"apiToken" gorm:"type:varchar(100);comment:ApiToken"` //
	Status    uint      `json:"status" gorm:"type:tinyint unsigned;comment:Status"` //
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;comment:创建日期"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;comment:修改日期"`
	DeletedAt time.Time `json:"deletedAt" gorm:"type:timestamp;comment:删除日期;default:NULL"`
}

func (KfUsers) TableName() string {
	return "users"
}
