package tag

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
	err = db.AutoMigrate(&Tag{})
	if err != nil {
		t.Fatalf("数据库迁移失败: %v", err)
	}

	return db
}

// TestTagDao_Insert 测试插入标签
func TestTagDao_Insert(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	tag := &Tag{
		Name:        "测试标签",
		Description: "测试描述",
		Color:       "#1890ff",
		Status:      StatusEnabled,
		CreatedBy:   1,
	}

	result, err := dao.Insert(ctx, tag)
	if err != nil {
		t.Fatalf("插入失败: %v", err)
	}

	if result.Id == 0 {
		t.Error("ID不应为0")
	}
	if result.Name != "测试标签" {
		t.Errorf("期望名称=测试标签, 实际=%s", result.Name)
	}
}

// TestTagDao_FindOne 测试根据ID查询
func TestTagDao_FindOne(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	// 先插入
	tag := &Tag{Name: "测试标签", Status: StatusEnabled, CreatedBy: 1}
	dao.Insert(ctx, tag)

	// 查询
	result, err := dao.FindOne(ctx, tag.Id)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}

	if result.Name != "测试标签" {
		t.Errorf("期望名称=测试标签, 实际=%s", result.Name)
	}
}

// TestTagDao_FindByName 测试根据名称查询
func TestTagDao_FindByName(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	tag := &Tag{Name: "唯一标签", Status: StatusEnabled, CreatedBy: 1}
	dao.Insert(ctx, tag)

	// 查存在的
	result, err := dao.FindByName(ctx, "唯一标签")
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if result == nil {
		t.Error("应该找到记录")
	}

	// 查不存在的
	result, err = dao.FindByName(ctx, "不存在")
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if result != nil {
		t.Error("不应该找到记录")
	}
}

// TestTagDao_Update 测试更新
func TestTagDao_Update(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	tag := &Tag{Name: "原始名称", Status: StatusEnabled, CreatedBy: 1}
	dao.Insert(ctx, tag)

	// 更新
	tag.Name = "更新后名称"
	tag.Description = "更新后描述"
	err := dao.Update(ctx, tag)
	if err != nil {
		t.Fatalf("更新失败: %v", err)
	}

	// 验证
	result, _ := dao.FindOne(ctx, tag.Id)
	if result.Name != "更新后名称" {
		t.Errorf("期望名称=更新后名称, 实际=%s", result.Name)
	}
}

// TestTagDao_Delete 测试删除
func TestTagDao_Delete(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	tag := &Tag{Name: "待删除", Status: StatusEnabled, CreatedBy: 1}
	dao.Insert(ctx, tag)

	// 删除
	err := dao.Delete(ctx, tag.Id)
	if err != nil {
		t.Fatalf("删除失败: %v", err)
	}

	// 验证已删除
	_, err = dao.FindOne(ctx, tag.Id)
	if err != ErrNotFound {
		t.Error("删除后查询应该返回ErrNotFound")
	}
}

// TestTagDao_List 测试分页查询
func TestTagDao_List(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	// 插入测试数据
	for i := 1; i <= 25; i++ {
		tag := &Tag{
			Name:      fmt.Sprintf("标签%d", i),
			Status:    StatusEnabled,
			CreatedBy: 1,
		}
		dao.Insert(ctx, tag)
	}

	// 测试分页
	results, total, err := dao.List(ctx, 1, 10)
	if err != nil {
		t.Fatalf("分页查询失败: %v", err)
	}
	if total != 25 {
		t.Errorf("期望总数=25, 实际=%d", total)
	}
	if len(results) != 10 {
		t.Errorf("期望返回10条, 实际=%d", len(results))
	}
}

// TestTagDao_Search 测试搜索
func TestTagDao_Search(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()
	// 插入测试数据
	dao.Insert(ctx, &Tag{Name: "红色标签", Description: "红色", Status: StatusEnabled, CreatedBy: 1})
	dao.Insert(ctx, &Tag{Name: "蓝色标签", Description: "蓝色", Status: StatusEnabled, CreatedBy: 1})
	dao.Insert(ctx, &Tag{Name: "绿色标签", Description: "绿色", Status: StatusEnabled, CreatedBy: 1})

	// 搜索"红色"
	results, total, err := dao.Search(ctx, "红色", 1, 10)
	if err != nil {
		t.Fatalf("搜索失败: %v", err)
	}
	if total != 1 {
		t.Errorf("期望找到1条, 实际=%d", total)
	}
	if len(results) != 1 || results[0].Name != "红色标签" {
		t.Error("搜索结果不正确")
	}
}

// TestTagDao_Trans 测试事务
func TestTagDao_Trans(t *testing.T) {
	db := setupTestDB(t)
	dao := &tagDao{db: db}

	ctx := context.Background()

	// 测试事务成功
	err := dao.Trans(ctx, func(ctx context.Context, model TagModel) error {
		tag1 := &Tag{Name: "标签1", Status: StatusEnabled, CreatedBy: 1}
		if _, err := model.Insert(ctx, tag1); err != nil {
			return err
		}

		tag2 := &Tag{Name: "标签2", Status: StatusEnabled, CreatedBy: 1}
		if _, err := model.Insert(ctx, tag2); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("事务执行失败: %v", err)
	}

	// 验证两条记录都插入了
	results, _, _ := dao.List(ctx, 1, 10)
	if len(results) != 2 {
		t.Errorf("期望插入2条, 实际=%d", len(results))
	}
}
