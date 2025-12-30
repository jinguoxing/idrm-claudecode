package tag

import "context"

// TagModel 标签数据访问接口
type TagModel interface {
	// Insert 插入新记录
	Insert(ctx context.Context, data *Tag) (*Tag, error)

	// FindOne 根据ID查询
	FindOne(ctx context.Context, id int64) (*Tag, error)

	// FindByName 根据名称查询
	FindByName(ctx context.Context, name string) (*Tag, error)

	// Update 更新记录
	Update(ctx context.Context, data *Tag) error

	// Delete 删除记录
	Delete(ctx context.Context, id int64) error

	// FindAll 查询所有记录
	FindAll(ctx context.Context) ([]*Tag, error)

	// List 分页查询
	List(ctx context.Context, page, pageSize int) ([]*Tag, int64, error)

	// Search 关键词搜索
	Search(ctx context.Context, keyword string, page, pageSize int) ([]*Tag, int64, error)

	// UpdateStatus 更新状态
	UpdateStatus(ctx context.Context, id int64, status int) error

	// WithTx 设置事务
	WithTx(tx interface{}) TagModel

	// Trans 事务处理
	Trans(ctx context.Context, fn func(ctx context.Context, model TagModel) error) error
}
