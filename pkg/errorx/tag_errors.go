package errorx

// 标签管理错误码 (31000-31999)
const (
	// 标签相关错误 (31000-31999)
	ErrCodeTagNotFound      = 31001 // 标签不存在
	ErrCodeTagAlreadyExists = 31002 // 标签名称已存在
	ErrCodeTagNameInvalid   = 31003 // 标签名称格式错误
	ErrCodeTagInUse         = 31004 // 标签正在使用中
	ErrCodeTagStatusInvalid = 31005 // 标签状态无效

	// 标签关联错误 (32000-32999)
	ErrCodeResourceTagExists  = 32001 // 关联已存在
	ErrCodeResourceTagNotFound = 32002 // 关联不存在
	ErrCodeResourceTagInvalid = 32003 // 无效的关联
)

// 初始化时添加标签错误消息
func init() {
	errMsgMap[ErrCodeTagNotFound] = "标签不存在"
	errMsgMap[ErrCodeTagAlreadyExists] = "标签名称已存在"
	errMsgMap[ErrCodeTagNameInvalid] = "标签名称格式错误"
	errMsgMap[ErrCodeTagInUse] = "标签正在使用中，无法删除"
	errMsgMap[ErrCodeTagStatusInvalid] = "标签状态无效"

	errMsgMap[ErrCodeResourceTagExists] = "标签关联已存在"
	errMsgMap[ErrCodeResourceTagNotFound] = "标签关联不存在"
	errMsgMap[ErrCodeResourceTagInvalid] = "无效的标签关联"
}
