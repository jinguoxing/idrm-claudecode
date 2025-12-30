package resource_tag

import "errors"

// 常量定义
const (
	// 资源类型
	ResourceTypeCatalogCategory = "catalog_category"
	ResourceTypeCatalogDataset  = "catalog_dataset"
	ResourceTypeDataView        = "data_view"
	ResourceTypeDataUnderstanding = "data_understanding"
)

// 错误定义
var (
	ErrNotFound      = errors.New("关联不存在")
	ErrAlreadyExists = errors.New("关联已存在")
	ErrInvalidParams = errors.New("参数无效")
)
