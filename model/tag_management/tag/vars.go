package tag

import "errors"

// 常量定义
const (
	// 默认颜色
	DefaultColor = "#1890ff"

	// 状态值
	StatusDisabled = 0 // 禁用
	StatusEnabled  = 1 // 启用
)

// 错误定义
var (
	ErrNotFound         = errors.New("标签不存在")
	ErrAlreadyExists    = errors.New("标签名称已存在")
	ErrInvalidName      = errors.New("标签名称格式错误")
	ErrInvalidStatus    = errors.New("标签状态无效")
	ErrNameRequired     = errors.New("标签名称不能为空")
	ErrNameTooLong      = errors.New("标签名称过长")
	ErrNameTooShort     = errors.New("标签名称过短")
	ErrDescriptionTooLong = errors.New("标签描述过长")
	ErrInvalidColor     = errors.New("标签颜色格式错误")
)
