// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"api/internal/config"
	"api/internal/middleware"
	"fmt"
	"gorm.io/gorm"
	"idrm/model/tag_management/resource_tag"
	"idrm/model/tag_management/tag"
	"idrm/pkg/db"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	Auth             rest.Middleware
	DB               *gorm.DB
	TagModel         tag.TagModel
	ResourceTagModel resource_tag.ResourceTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	gormDB, err := initDB(c.Database)
	if err != nil {
		panic(fmt.Sprintf("初始化数据库失败: %v", err))
	}

	return &ServiceContext{
		Config:           c,
		Auth:             middleware.NewAuthMiddleware().Handle,
		DB:               gormDB,
		TagModel:         tag.NewTagModel(gormDB),
		ResourceTagModel: resource_tag.NewResourceTagModel(gormDB),
	}
}

// initDB 初始化数据库连接
func initDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// 将 DatabaseConfig 转换为 db.Config
	dbConfig := db.Config{
		// 默认配置
		Database:        "idrm",
		Username:        "root",
		Password:        "123456",
		Host:           "127.0.0.1",
		Port:           3306,
		Charset:        "utf8mb4",
		MaxOpenConns:   cfg.MaxOpenConns,
		MaxIdleConns:   cfg.MaxIdleConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdleTime: 600,
		LogLevel:       "warn",
	}

	// 如果 Source 不为空，解析连接字符串
	if cfg.Source != "" {
		dbConfig.Source = cfg.Source
	}

	return db.InitGorm(dbConfig)
}
