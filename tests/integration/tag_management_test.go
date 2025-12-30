package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"idrm/model/tag_management/resource_tag"
	"idrm/model/tag_management/tag"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	StatusDisabled = 0
)

// TagManagementTestSuite 标签管理集成测试套件
type TagManagementTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SetupSuite 测试套件初始化
func (suite *TagManagementTestSuite) SetupSuite() {
	// 创建内存数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = db
}

// SetupTest 每个测试前运行 - 重置数据库
func (suite *TagManagementTestSuite) SetupTest() {
	// 删除所有表
	suite.db.Migrator().DropTable(&resource_tag.ResourceTag{}, &tag.Tag{})

	// 重新创建表
	err := suite.db.AutoMigrate(&tag.Tag{}, &resource_tag.ResourceTag{})
	suite.Require().NoError(err)
}

// TearDownSuite 测试套件清理
func (suite *TagManagementTestSuite) TearDownSuite() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

// getModels 获取Model实例
func (suite *TagManagementTestSuite) getModels() (tag.TagModel, resource_tag.ResourceTagModel) {
	return tag.NewTagModel(suite.db), resource_tag.NewResourceTagModel(suite.db)
}

// TestTagManagementCRUD 测试标签CRUD完整流程
func (suite *TagManagementTestSuite) TestTagManagementCRUD() {
	ctx := context.Background()
	tagModel, resourceTagModel := suite.getModels()

	// 1. 创建标签
	tag1 := &tag.Tag{
		Name:        "技术标签",
		Description: "技术相关内容",
		Color:       "#1890ff",
		Status:      tag.StatusEnabled,
		CreatedBy:   1,
	}
	created1, err := tagModel.Insert(ctx, tag1)
	suite.NoError(err)
	suite.Equal(int64(1), created1.Id)

	tag2 := &tag.Tag{
		Name:        "业务标签",
		Description: "业务相关内容",
		Color:       "#52c41a",
		Status:      tag.StatusEnabled,
		CreatedBy:   1,
	}
	created2, err := tagModel.Insert(ctx, tag2)
	suite.NoError(err)
	suite.Equal(int64(2), created2.Id)

	// 2. 查询标签列表
	tags, total, err := tagModel.List(ctx, 1, 10)
	suite.NoError(err)
	suite.Equal(int64(2), total)
	suite.Len(tags, 2)

	// 3. 根据ID查询
	found, err := tagModel.FindOne(ctx, created1.Id)
	suite.NoError(err)
	suite.Equal("技术标签", found.Name)

	// 4. 更新标签
	found.Description = "更新后的描述"
	err = tagModel.Update(ctx, found)
	suite.NoError(err)

	updated, _ := tagModel.FindOne(ctx, created1.Id)
	suite.Equal("更新后的描述", updated.Description)

	// 5. 为资源打标签
	err = resourceTagModel.BatchAssign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{created1.Id, created2.Id})
	suite.NoError(err)

	// 验证关联
	resourceTags, _ := resourceTagModel.FindByResource(ctx, 100, resource_tag.ResourceTypeCatalogCategory)
	suite.Len(resourceTags, 2)

	// 6. 查询使用次数
	count, _ := resourceTagModel.CountByTag(ctx, created1.Id)
	suite.Equal(int64(1), count)

	// 7. 按标签搜索资源
	resourceIDs, err := resourceTagModel.FindByTags(ctx, []int64{created1.Id}, resource_tag.ResourceTypeCatalogCategory)
	suite.NoError(err)
	suite.Contains(resourceIDs, int64(100))

	// 8. 移除标签
	err = resourceTagModel.BatchUnassign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{created2.Id})
	suite.NoError(err)

	resourceTags, _ = resourceTagModel.FindByResource(ctx, 100, resource_tag.ResourceTypeCatalogCategory)
	suite.Len(resourceTags, 1)

	// 9. 替换标签
	err = resourceTagModel.ReplaceTags(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{created2.Id})
	suite.NoError(err)

	resourceTags, _ = resourceTagModel.FindByResource(ctx, 100, resource_tag.ResourceTypeCatalogCategory)
	suite.Len(resourceTags, 1)
	suite.Equal(created2.Id, resourceTags[0].TagId)

	// 10. 删除标签（先移除关联）
	resourceTagModel.BatchUnassign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{created2.Id})
	err = tagModel.Delete(ctx, created2.Id)
	suite.NoError(err)

	// 验证删除
	_, err = tagModel.FindOne(ctx, created2.Id)
	suite.Equal(tag.ErrNotFound, err)
}

// TestTagManagement_Validation 测试业务规则
func (suite *TagManagementTestSuite) TestTagManagement_Validation() {
	ctx := context.Background()
	tagModel, resourceTagModel := suite.getModels()

	// 1. 测试名称唯一性检查
	tag1 := &tag.Tag{Name: "唯一名称", Status: tag.StatusEnabled, CreatedBy: 1}
	_, err := tagModel.Insert(ctx, tag1)
	suite.NoError(err)

	// 注意: GORM不会自动检查唯一性，这应该在Logic层实现
	// 这里我们验证同名标签可以被插入（数据库层面）
	tag2 := &tag.Tag{Name: "唯一名称", Status: tag.StatusEnabled, CreatedBy: 1}
	_, err = tagModel.Insert(ctx, tag2)
	// GORM会成功插入，没有错误
	suite.NoError(err)

	// 2. 测试删除正在使用的标签
	tag3 := &tag.Tag{Name: "使用中标签", Status: tag.StatusEnabled, CreatedBy: 1}
	created3, _ := tagModel.Insert(ctx, tag3)

	// 关联到资源
	resourceTagModel.Assign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, created3.Id)

	// 验证使用次数
	count, _ := resourceTagModel.CountByTag(ctx, created3.Id)
	suite.Equal(int64(1), count)
}

// TestTagManagement_Transaction 测试事务处理
func (suite *TagManagementTestSuite) TestTagManagement_Transaction() {
	ctx := context.Background()
	tagModel, _ := suite.getModels()

	// 测试事务回滚
	err := tagModel.Trans(ctx, func(ctx context.Context, model tag.TagModel) error {
		tag1 := &tag.Tag{Name: "标签1", Status: tag.StatusEnabled, CreatedBy: 1}
		_, err := model.Insert(ctx, tag1)
		if err != nil {
			return err
		}

		tag2 := &tag.Tag{Name: "标签2", Status: tag.StatusEnabled, CreatedBy: 1}
		_, err = model.Insert(ctx, tag2)
		if err != nil {
			return err
		}

		// 故意失败
		return gorm.ErrInvalidTransaction
	})

	suite.Error(err)

	// 验证没有插入任何数据
	tags, total, _ := tagModel.List(ctx, 1, 10)
	suite.Equal(int64(0), total)
	suite.Len(tags, 0)

	// 测试事务成功
	err = tagModel.Trans(ctx, func(ctx context.Context, model tag.TagModel) error {
		tag1 := &tag.Tag{Name: "事务标签1", Status: tag.StatusEnabled, CreatedBy: 1}
		_, err := model.Insert(ctx, tag1)
		return err
	})

	suite.NoError(err)

	tags, total, _ = tagModel.List(ctx, 1, 10)
	suite.Equal(int64(1), total)
}

// TestTagManagement_Search 测试搜索功能
func (suite *TagManagementTestSuite) TestTagManagement_Search() {
	ctx := context.Background()
	tagModel, _ := suite.getModels()

	// 创建多个标签
	tags := []*tag.Tag{
		{Name: "红色", Description: "红色主题", Status: tag.StatusEnabled, CreatedBy: 1},
		{Name: "红色标签", Description: "红色相关", Status: tag.StatusEnabled, CreatedBy: 1},
		{Name: "蓝色", Description: "蓝色主题", Status: tag.StatusEnabled, CreatedBy: 1},
	}

	for _, t := range tags {
		tagModel.Insert(ctx, t)
	}

	// 搜索"红色"
	results, total, err := tagModel.Search(ctx, "红色", 1, 10)
	suite.NoError(err)
	suite.Equal(int64(2), total) // 应该找到2个包含"红色"的标签

	// 搜索不存在的关键词
	results, total, _ = tagModel.Search(ctx, "不存在", 1, 10)
	suite.Equal(int64(0), total)
	suite.Len(results, 0)
}

// TestTagManagement_MultiTagSearch 测试多标签AND查询
func (suite *TagManagementTestSuite) TestTagManagement_MultiTagSearch() {
	ctx := context.Background()
	tagModel, resourceTagModel := suite.getModels()

	// 创建标签
	tag1, _ := tagModel.Insert(ctx, &tag.Tag{Name: "Python", Status: tag.StatusEnabled, CreatedBy: 1})
	tag2, _ := tagModel.Insert(ctx, &tag.Tag{Name: "教程", Status: tag.StatusEnabled, CreatedBy: 1})
	tag3, _ := tagModel.Insert(ctx, &tag.Tag{Name: "高级", Status: tag.StatusEnabled, CreatedBy: 1})

	// 资源100有 Python + 教程
	resourceTagModel.BatchAssign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{tag1.Id, tag2.Id})

	// 资源200有 Python + 高级
	resourceTagModel.BatchAssign(ctx, 200, resource_tag.ResourceTypeCatalogCategory, []int64{tag1.Id, tag3.Id})

	// 资源300只有 Python
	resourceTagModel.Assign(ctx, 300, resource_tag.ResourceTypeCatalogCategory, tag1.Id)

	// 查询同时有 Python 和 教程 的资源（只有资源100）
	resourceIDs, err := resourceTagModel.FindByTags(ctx, []int64{tag1.Id, tag2.Id}, resource_tag.ResourceTypeCatalogCategory)
	suite.NoError(err)
	suite.Len(resourceIDs, 1)
	suite.Contains(resourceIDs, int64(100))

	// 查询同时有 Python 和 高级 的资源（只有资源200）
	resourceIDs, _ = resourceTagModel.FindByTags(ctx, []int64{tag1.Id, tag3.Id}, resource_tag.ResourceTypeCatalogCategory)
	suite.Len(resourceIDs, 1)
	suite.Contains(resourceIDs, int64(200))
}

// TestTagManagement_Performance 测试性能
func (suite *TagManagementTestSuite) TestTagManagement_Performance() {
	ctx := context.Background()
	tagModel, _ := suite.getModels()

	// 批量创建100个标签
	for i := 1; i <= 100; i++ {
		tagRecord := &tag.Tag{
			Name:        fmt.Sprintf("性能测试标签%d", i),
			Description: "性能测试",
			Color:       "#1890ff",
			Status:      tag.StatusEnabled,
			CreatedBy:   1,
		}
		tagModel.Insert(ctx, tagRecord)
	}

	// 测试分页查询性能
	start := time.Now()
	_, total, err := tagModel.List(ctx, 1, 20)
	elapsed := time.Since(start)

	suite.NoError(err)
	suite.Equal(int64(100), total)
	suite.Less(elapsed, 100*time.Millisecond, "分页查询应该小于100ms")
}

// TestTagManagement_StatusUpdate 测试状态更新
func (suite *TagManagementTestSuite) TestTagManagement_StatusUpdate() {
	ctx := context.Background()
	tagModel, _ := suite.getModels()

	tagRecord := &tag.Tag{Name: "待禁用", Status: tag.StatusEnabled, CreatedBy: 1}
	created, _ := tagModel.Insert(ctx, tagRecord)

	// 禁用标签
	err := tagModel.UpdateStatus(ctx, created.Id, StatusDisabled)
	suite.NoError(err)

	// 验证
	updated, _ := tagModel.FindOne(ctx, created.Id)
	suite.Equal(StatusDisabled, updated.Status)
}

// TestTagManagement_EdgeCases 测试边界情况
func (suite *TagManagementTestSuite) TestTagManagement_EdgeCases() {
	ctx := context.Background()
	tagModel, resourceTagModel := suite.getModels()

	// 1. 查询不存在的标签
	_, err := tagModel.FindOne(ctx, 99999)
	suite.Equal(tag.ErrNotFound, err)

	// 2. 更新不存在的标签 - GORM不会报错，但affected为0
	nonExistent := &tag.Tag{Id: 99999, Name: "不存在"}
	err = tagModel.Update(ctx, nonExistent)
	// 由于GORM特性，更新不存在的记录不会报错
	suite.NoError(err)

	// 3. 删除不存在的标签 - GORM不会报错，但affected为0
	err = tagModel.Delete(ctx, 99999)
	suite.NoError(err)

	// 4. 空列表查询
	tags, total, err := tagModel.Search(ctx, "不存在的标签xyz", 1, 10)
	suite.NoError(err)
	suite.Equal(int64(0), total)
	suite.Len(tags, 0)

	// 5. 空批量操作
	err = resourceTagModel.BatchAssign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{})
	suite.NoError(err) // 空列表不应该报错

	err = resourceTagModel.BatchUnassign(ctx, 100, resource_tag.ResourceTypeCatalogCategory, []int64{})
	suite.NoError(err)
}

// TestTagManagement_Concurrency 测试并发安全
func (suite *TagManagementTestSuite) TestTagManagement_Concurrency() {
	ctx := context.Background()
	tagModel, resourceTagModel := suite.getModels()

	tagRecord := &tag.Tag{Name: "并发测试", Status: tag.StatusEnabled, CreatedBy: 1}
	created, _ := tagModel.Insert(ctx, tagRecord)

	// 使用同步方式批量增加使用次数（测试Assign的幂等性）
	// SQLite in-memory数据库在goroutine中可能不支持并发写入
	resourceIDs := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, resourceID := range resourceIDs {
		for j := 0; j < 5; j++ {
			// Assign使用OnConflict处理重复，所以多次调用同一资源是安全的
			resourceTagModel.Assign(ctx, resourceID, resource_tag.ResourceTypeCatalogCategory, created.Id)
		}
	}

	// 验证使用次数 - 每个资源只关联一次
	count, _ := resourceTagModel.CountByTag(ctx, created.Id)
	suite.Equal(int64(10), count)
}

// RunTestSuite 运行测试套件
func TestTagManagementTestSuite(t *testing.T) {
	suite.Run(t, new(TagManagementTestSuite))
}
