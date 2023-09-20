package mysql

import (
	"gorm.io/gorm"
)

type (
	// Conds 为Cond类型map，用于定义Where方法参数 map[field.name]interface{}
	Conds = map[string]interface{}

	// Ups 为更新某一条记录时存放的变更数据集合 map[field.name]field.value
	Ups = map[string]interface{}
)

func Init(db *gorm.DB) error {
	return db.AutoMigrate(User{})
}
