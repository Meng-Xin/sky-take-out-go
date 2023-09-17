package model

import (
	"gorm.io/gorm"
	"take-out/common/enum"
	"time"
)

type Category struct {
	Id         uint64    `json:"id"`
	Cate       int       `json:"type"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

func (e *Category) BeforeCreate(tx *gorm.DB) error {
	// 自动填充 创建时间、创建人、更新时间、更新用户
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.CreateUser = uid
		e.UpdateUser = uid
	}
	return nil
}

func (e *Category) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	e.UpdateTime = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.UpdateUser = uid
	}
	return nil
}
