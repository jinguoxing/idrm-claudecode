package resource_tag

import (
	"gorm.io/gorm"
)

var (
	gormFactory func(db *gorm.DB) ResourceTagModel
)

// RegisterGormFactory 注册GORM工厂函数
func RegisterGormFactory(fn func(db *gorm.DB) ResourceTagModel) {
	gormFactory = fn
}

// NewResourceTagModel 创建ResourceTagModel实例
func NewResourceTagModel(db *gorm.DB) ResourceTagModel {
	if gormFactory != nil {
		return gormFactory(db)
	}
	return nil
}
