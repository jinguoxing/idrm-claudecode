package tag

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type tagDao struct {
	db *gorm.DB
}

func init() {
	RegisterGormFactory(newTagDao)
}

// newTagDao 创建tagDao实例
func newTagDao(db *gorm.DB) TagModel {
	return &tagDao{db: db}
}

// Insert 插入新记录
func (d *tagDao) Insert(ctx context.Context, data *Tag) (*Tag, error) {
	if err := d.db.WithContext(ctx).Create(data).Error; err != nil {
		return nil, fmt.Errorf("插入标签失败: %w", err)
	}
	return data, nil
}

// FindOne 根据ID查询
func (d *tagDao) FindOne(ctx context.Context, id int64) (*Tag, error) {
	var result Tag
	err := d.db.WithContext(ctx).
		Where("id = ?", id).
		First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("查询标签失败: %w", err)
	}
	return &result, nil
}

// FindByName 根据名称查询
func (d *tagDao) FindByName(ctx context.Context, name string) (*Tag, error) {
	var result Tag
	err := d.db.WithContext(ctx).
		Where("name = ?", name).
		First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("根据名称查询标签失败: %w", err)
	}
	return &result, nil
}

// Update 更新记录
func (d *tagDao) Update(ctx context.Context, data *Tag) error {
	err := d.db.WithContext(ctx).
		Model(&Tag{}).
		Where("id = ?", data.Id).
		Updates(data).Error
	if err != nil {
		return fmt.Errorf("更新标签失败: %w", err)
	}
	return nil
}

// Delete 删除记录
func (d *tagDao) Delete(ctx context.Context, id int64) error {
	err := d.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&Tag{}).Error
	if err != nil {
		return fmt.Errorf("删除标签失败: %w", err)
	}
	return nil
}

// FindAll 查询所有记录
func (d *tagDao) FindAll(ctx context.Context) ([]*Tag, error) {
	var results []*Tag
	err := d.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("查询所有标签失败: %w", err)
	}
	return results, nil
}

// List 分页查询
func (d *tagDao) List(ctx context.Context, page, pageSize int) ([]*Tag, int64, error) {
	var results []*Tag
	var total int64

	offset := (page - 1) * pageSize

	// 查询总数
	if err := d.db.WithContext(ctx).Model(&Tag{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询标签总数失败: %w", err)
	}

	// 分页查询
	err := d.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&results).Error
	if err != nil {
		return nil, 0, fmt.Errorf("分页查询标签失败: %w", err)
	}

	return results, total, nil
}

// Search 关键词搜索
func (d *tagDao) Search(ctx context.Context, keyword string, page, pageSize int) ([]*Tag, int64, error) {
	var results []*Tag
	var total int64

	offset := (page - 1) * pageSize
	query := d.db.WithContext(ctx).Model(&Tag{})

	if keyword != "" {
		pattern := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", pattern, pattern)
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("搜索标签总数失败: %w", err)
	}

	// 分页查询
	err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&results).Error
	if err != nil {
		return nil, 0, fmt.Errorf("搜索标签失败: %w", err)
	}

	return results, total, nil
}

// UpdateStatus 更新状态
func (d *tagDao) UpdateStatus(ctx context.Context, id int64, status int) error {
	err := d.db.WithContext(ctx).
		Model(&Tag{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("更新标签状态失败: %w", err)
	}
	return nil
}

// WithTx 设置事务
func (d *tagDao) WithTx(tx interface{}) TagModel {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return d
	}
	return &tagDao{db: db}
}

// Trans 事务处理
func (d *tagDao) Trans(ctx context.Context, fn func(ctx context.Context, model TagModel) error) error {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txModel := &tagDao{db: tx}
		return fn(ctx, txModel)
	})
	if err != nil {
		return fmt.Errorf("事务执行失败: %w", err)
	}
	return nil
}
