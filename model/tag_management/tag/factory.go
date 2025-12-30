package tag

import (
	"gorm.io/gorm"
)

var (
	gormFactory func(db *gorm.DB) TagModel
)

// RegisterGormFactory 注册GORM工厂函数
func RegisterGormFactory(fn func(db *gorm.DB) TagModel) {
	gormFactory = fn
}

// NewTagModel 创建TagModel实例
func NewTagModel(db *gorm.DB) TagModel {
	if gormFactory != nil {
		return gormFactory(db)
	}
	return nil
}
