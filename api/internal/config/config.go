// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"idrm/pkg/config"
)

type Config struct {
	rest.RestConf

	// 数据库配置
	Database config.DatabaseConfig
}
