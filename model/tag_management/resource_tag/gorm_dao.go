package resource_tag

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type resourceTagDao struct {
	db *gorm.DB
}

func init() {
	RegisterGormFactory(newResourceTagDao)
}

// newResourceTagDao 创建resourceTagDao实例
func newResourceTagDao(db *gorm.DB) ResourceTagModel {
	return &resourceTagDao{db: db}
}

// Assign 为资源关联单个标签
func (d *resourceTagDao) Assign(ctx context.Context, resourceID int64, resourceType string, tagID int64) error {
	// 使用FirstOrCreate避免重复插入
	assignment := &ResourceTag{
		ResourceId:   resourceID,
		ResourceType: resourceType,
		TagId:        tagID,
	}
	err := d.db.WithContext(ctx).
		Where("resource_id = ? AND resource_type = ? AND tag_id = ?", resourceID, resourceType, tagID).
		FirstOrCreate(assignment).Error
	if err != nil {
		return fmt.Errorf("关联标签失败: %w", err)
	}
	return nil
}

// Unassign 移除资源的单个标签关联
func (d *resourceTagDao) Unassign(ctx context.Context, resourceID int64, resourceType string, tagID int64) error {
	err := d.db.WithContext(ctx).
		Where("resource_id = ? AND resource_type = ? AND tag_id = ?", resourceID, resourceType, tagID).
		Delete(&ResourceTag{}).Error
	if err != nil {
		return fmt.Errorf("移除标签关联失败: %w", err)
	}
	return nil
}

// GetResourceTags 获取资源的所有标签ID
func (d *resourceTagDao) GetResourceTags(ctx context.Context, resourceID int64, resourceType string) ([]int64, error) {
	var tagIDs []int64
	err := d.db.WithContext(ctx).
		Model(&ResourceTag{}).
		Where("resource_id = ? AND resource_type = ?", resourceID, resourceType).
		Pluck("tag_id", &tagIDs).Error
	if err != nil {
		return nil, fmt.Errorf("获取资源标签失败: %w", err)
	}
	return tagIDs, nil
}

// BatchAssign 批量为资源关联标签
func (d *resourceTagDao) BatchAssign(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error {
	if len(tagIDs) == 0 {
		return nil
	}

	// 构建批量插入数据
	var assignments []*ResourceTag
	for _, tagID := range tagIDs {
		assignments = append(assignments, &ResourceTag{
			ResourceId:   resourceID,
			ResourceType: resourceType,
			TagId:        tagID,
		})
	}

	// 批量插入（忽略重复）
	err := d.db.WithContext(ctx).
		Create(&assignments).Error
	if err != nil {
		return fmt.Errorf("批量关联标签失败: %w", err)
	}
	return nil
}

// BatchUnassign 批量移除资源的标签关联
func (d *resourceTagDao) BatchUnassign(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error {
	if len(tagIDs) == 0 {
		return nil
	}

	err := d.db.WithContext(ctx).
		Where("resource_id = ? AND resource_type = ? AND tag_id IN ?", resourceID, resourceType, tagIDs).
		Delete(&ResourceTag{}).Error
	if err != nil {
		return fmt.Errorf("批量移除标签关联失败: %w", err)
	}
	return nil
}

// ReplaceTags 替换资源的所有标签
func (d *resourceTagDao) ReplaceTags(ctx context.Context, resourceID int64, resourceType string, tagIDs []int64) error {
	return d.Trans(ctx, func(ctx context.Context, model ResourceTagModel) error {
		txModel := model.(*resourceTagDao)

		// 先删除所有现有关联
		err := txModel.db.WithContext(ctx).
			Where("resource_id = ? AND resource_type = ?", resourceID, resourceType).
			Delete(&ResourceTag{}).Error
		if err != nil {
			return fmt.Errorf("清除现有标签失败: %w", err)
		}

		// 添加新关联
		if len(tagIDs) > 0 {
			var assignments []*ResourceTag
			for _, tagID := range tagIDs {
				assignments = append(assignments, &ResourceTag{
					ResourceId:   resourceID,
					ResourceType: resourceType,
					TagId:        tagID,
				})
			}
			err = txModel.db.WithContext(ctx).Create(&assignments).Error
			if err != nil {
				return fmt.Errorf("添加新标签失败: %w", err)
			}
		}

		return nil
	})
}

// FindByResource 查询资源的所有标签关联
func (d *resourceTagDao) FindByResource(ctx context.Context, resourceID int64, resourceType string) ([]*ResourceTag, error) {
	var results []*ResourceTag
	err := d.db.WithContext(ctx).
		Where("resource_id = ? AND resource_type = ?", resourceID, resourceType).
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("查询资源标签关联失败: %w", err)
	}
	return results, nil
}

// FindByTag 查询标签关联的所有资源
func (d *resourceTagDao) FindByTag(ctx context.Context, tagID int64) ([]*ResourceTag, error) {
	var results []*ResourceTag
	err := d.db.WithContext(ctx).
		Where("tag_id = ?", tagID).
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("查询标签关联资源失败: %w", err)
	}
	return results, nil
}

// FindByTags 查询包含所有指定标签的资源ID列表（AND关系）
func (d *resourceTagDao) FindByTags(ctx context.Context, tagIDs []int64, resourceType string) ([]int64, error) {
	if len(tagIDs) == 0 {
		return []int64{}, nil
	}

	// 统计每个资源出现的次数
	type ResourceCount struct {
		ResourceId int64
		Count      int64
	}

	var results []ResourceCount
	err := d.db.WithContext(ctx).
		Model(&ResourceTag{}).
		Select("resource_id, COUNT(*) as count").
		Where("tag_id IN ? AND resource_type = ?", tagIDs, resourceType).
		Group("resource_id").
		Having("COUNT(*) = ?", len(tagIDs)).
		Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("按标签查询资源失败: %w", err)
	}

	// 提取资源ID
	var resourceIDs []int64
	for _, r := range results {
		resourceIDs = append(resourceIDs, r.ResourceId)
	}

	return resourceIDs, nil
}

// CountByTag 统计标签被使用的次数
func (d *resourceTagDao) CountByTag(ctx context.Context, tagID int64) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).
		Model(&ResourceTag{}).
		Where("tag_id = ?", tagID).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("统计标签使用次数失败: %w", err)
	}
	return count, nil
}

// WithTx 设置事务
func (d *resourceTagDao) WithTx(tx interface{}) ResourceTagModel {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return d
	}
	return &resourceTagDao{db: db}
}

// Trans 事务处理
func (d *resourceTagDao) Trans(ctx context.Context, fn func(ctx context.Context, model ResourceTagModel) error) error {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txModel := &resourceTagDao{db: tx}
		return fn(ctx, txModel)
	})
	if err != nil {
		return fmt.Errorf("事务执行失败: %w", err)
	}
	return nil
}
