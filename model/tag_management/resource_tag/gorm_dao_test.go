package resource_tag

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB 创建测试数据库
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("无法创建测试数据库: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&ResourceTag{})
	if err != nil {
		t.Fatalf("数据库迁移失败: %v", err)
	}

	return db
}

// TestResourceTagDao_Assign 测试关联标签
func TestResourceTagDao_Assign(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	err := dao.Assign(ctx, 100, ResourceTypeCatalogCategory, 1)
	if err != nil {
		t.Fatalf("关联失败: %v", err)
	}

	// 验证
	results, err := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("期望1条关联, 实际=%d", len(results))
	}

	// 重复关联应该不报错
	err = dao.Assign(ctx, 100, ResourceTypeCatalogCategory, 1)
	if err != nil {
		t.Errorf("重复关联不应该报错: %v", err)
	}
}

// TestResourceTagDao_Unassign 测试移除关联
func TestResourceTagDao_Unassign(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	dao.Assign(ctx, 100, ResourceTypeCatalogCategory, 1)

	// 移除
	err := dao.Unassign(ctx, 100, ResourceTypeCatalogCategory, 1)
	if err != nil {
		t.Fatalf("移除关联失败: %v", err)
	}

	// 验证
	results, _ := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if len(results) != 0 {
		t.Error("移除后应该没有关联记录")
	}
}

// TestResourceTagDao_BatchAssign 测试批量关联
func TestResourceTagDao_BatchAssign(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	tagIDs := []int64{1, 2, 3}

	err := dao.BatchAssign(ctx, 100, ResourceTypeCatalogCategory, tagIDs)
	if err != nil {
		t.Fatalf("批量关联失败: %v", err)
	}

	// 验证
	results, _ := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if len(results) != 3 {
		t.Errorf("期望3条关联, 实际=%d", len(results))
	}
}

// TestResourceTagDao_BatchUnassign 测试批量移除
func TestResourceTagDao_BatchUnassign(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	tagIDs := []int64{1, 2, 3}
	dao.BatchAssign(ctx, 100, ResourceTypeCatalogCategory, tagIDs)

	// 批量移除
	err := dao.BatchUnassign(ctx, 100, ResourceTypeCatalogCategory, []int64{1, 2})
	if err != nil {
		t.Fatalf("批量移除失败: %v", err)
	}

	// 验证只剩1条
	results, _ := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if len(results) != 1 {
		t.Errorf("期望剩余1条, 实际=%d", len(results))
	}
}

// TestResourceTagDao_ReplaceTags 测试替换标签
func TestResourceTagDao_ReplaceTags(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	// 先关联3个
	dao.BatchAssign(ctx, 100, ResourceTypeCatalogCategory, []int64{1, 2, 3})

	// 替换成2个
	newTagIDs := []int64{4, 5}
	err := dao.ReplaceTags(ctx, 100, ResourceTypeCatalogCategory, newTagIDs)
	if err != nil {
		t.Fatalf("替换失败: %v", err)
	}

	// 验证
	results, _ := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if len(results) != 2 {
		t.Errorf("期望2条, 实际=%d", len(results))
	}
}

// TestResourceTagDao_FindByResource 测试查询资源的标签
func TestResourceTagDao_FindByResource(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	dao.BatchAssign(ctx, 100, ResourceTypeCatalogCategory, []int64{1, 2, 3})

	results, err := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("期望3条, 实际=%d", len(results))
	}
}

// TestResourceTagDao_FindByTag 测试查询标签的资源
func TestResourceTagDao_FindByTag(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	// 资源100和200都关联标签1
	dao.Assign(ctx, 100, ResourceTypeCatalogCategory, 1)
	dao.Assign(ctx, 200, ResourceTypeCatalogCategory, 1)

	results, err := dao.FindByTag(ctx, 1)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("期望2条, 实际=%d", len(results))
	}
}

// TestResourceTagDao_FindByTags 测试多标签AND查询
func TestResourceTagDao_FindByTags(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	// 资源100关联标签1,2,3
	dao.BatchAssign(ctx, 100, ResourceTypeCatalogCategory, []int64{1, 2, 3})
	// 资源200关联标签1,2
	dao.BatchAssign(ctx, 200, ResourceTypeCatalogCategory, []int64{1, 2})
	// 资源300只关联标签1
	dao.Assign(ctx, 300, ResourceTypeCatalogCategory, 1)

	// 查询同时包含标签1和2的资源
	resourceIDs, err := dao.FindByTags(ctx, []int64{1, 2}, ResourceTypeCatalogCategory)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	// 只有资源100和200同时包含标签1和2
	if len(resourceIDs) != 2 {
		t.Errorf("期望2个资源, 实际=%d", len(resourceIDs))
	}
}

// TestResourceTagDao_CountByTag 测试统计标签使用次数
func TestResourceTagDao_CountByTag(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()
	// 3个资源都关联标签1
	dao.Assign(ctx, 100, ResourceTypeCatalogCategory, 1)
	dao.Assign(ctx, 200, ResourceTypeCatalogCategory, 1)
	dao.Assign(ctx, 300, ResourceTypeCatalogCategory, 1)

	count, err := dao.CountByTag(ctx, 1)
	if err != nil {
		t.Fatalf("统计失败: %v", err)
	}
	if count != 3 {
		t.Errorf("期望使用次数=3, 实际=%d", count)
	}
}

// TestResourceTagDao_Trans 测试事务
func TestResourceTagDao_Trans(t *testing.T) {
	db := setupTestDB(t)
	dao := &resourceTagDao{db: db}

	ctx := context.Background()

	// 测试事务成功
	err := dao.Trans(ctx, func(ctx context.Context, model ResourceTagModel) error {
		if err := model.Assign(ctx, 100, ResourceTypeCatalogCategory, 1); err != nil {
			return err
		}
		if err := model.Assign(ctx, 100, ResourceTypeCatalogCategory, 2); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("事务执行失败: %v", err)
	}

	// 验证两条记录都插入了
	results, _ := dao.FindByResource(ctx, 100, ResourceTypeCatalogCategory)
	if len(results) != 2 {
		t.Errorf("期望插入2条, 实际=%d", len(results))
	}

	// 测试事务回滚
	err = dao.Trans(ctx, func(ctx context.Context, model ResourceTagModel) error {
		model.Assign(ctx, 200, ResourceTypeCatalogCategory, 1)
		model.Assign(ctx, 200, ResourceTypeCatalogCategory, 2)
		// 返回错误触发回滚
		return fmt.Errorf("故意失败")
	})
	if err == nil {
		t.Error("事务应该失败")
	}

	// 验证资源200没有插入
	results, _ = dao.FindByResource(ctx, 200, ResourceTypeCatalogCategory)
	if len(results) != 0 {
		t.Error("事务回滚后不应该有记录")
	}
}
