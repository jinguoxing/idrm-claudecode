package resource_tag

import "context"

// ResourceTagModel 资源标签关联数据访问接口
type ResourceTagModel interface {
	// Assign 为资源关联单个标签
	Assign(ctx context.Context, resourceID int64, resourceType string, tagID int64) error

	// Unassign 移除资源的单个标签关联
	Unassign(ctx context.Context, resourceID int64, resourceType string, tagID int64) error

	// GetResourceTags 获取资源的所有标签ID
	GetResourceTags(ctx context.Context, resourceID int64, resourceType string) ([]int64, error)

	// BatchAssign 批量为资源关联标签
	BatchAssign(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error

	// BatchUnassign 批量移除资源的标签关联
	BatchUnassign(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error

	// ReplaceTags 替换资源的所有标签
	ReplaceTags(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error

	// FindByResource 查询资源的所有标签关联
	FindByResource(ctx context.Context, resourceID int64, resourceType string) ([]*ResourceTag, error)

	// FindByTag 查询标签关联的所有资源
	FindByTag(ctx context.Context, tagID int64) ([]*ResourceTag, error)

	// FindByTags 查询包含所有指定标签的资源ID列表（AND关系）
	FindByTags(ctx context.Context, tagIDs []int64, resourceType string) ([]int64, error)

	// CountByTag 统计标签被使用的次数
	CountByTag(ctx context.Context, tagID int64) (int64, error)

	// WithTx 设置事务
	WithTx(tx interface{}) ResourceTagModel

	// Trans 事务处理
	Trans(ctx context.Context, fn func(ctx context.Context, model ResourceTagModel) error) error
}
