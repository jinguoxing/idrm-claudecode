package tag_management

import (
	"context"
	"errors"
	"testing"
	"time"

	"api/internal/logic/tag_management/mocks"
	"api/internal/svc"
	"api/internal/types"

	"idrm/model/tag_management/tag"
	"idrm/pkg/errorx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateTagLogic_CreateTag_Success 测试成功创建标签
func TestCreateTagLogic_CreateTag_Success(t *testing.T) {
	// 准备mock
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindByName 返回不存在（名称可用）
	mockTagModel.On("FindByName", ctx, "测试标签").Return((*tag.Tag)(nil), nil)

	// Mock Insert 返回成功
	insertedTag := &tag.Tag{Id: 1, Name: "测试标签"}
	mockTagModel.On("Insert", ctx, mock.AnythingOfType("*tag.Tag")).Return(insertedTag, nil)

	// 创建Logic
	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewCreateTagLogic(ctx, svcCtx)

	// 执行
	req := &types.CreateTagReq{
		Name:        "测试标签",
		Description: "测试描述",
		Color:       "#1890ff",
	}
	resp, err := logic.CreateTag(req)

	// 验证
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.Id)

	mockTagModel.AssertExpectations(t)
}

// TestCreateTagLogic_CreateTag_NameExists 测试标签名已存在
func TestCreateTagLogic_CreateTag_NameExists(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindByName 返回已存在
	existingTag := &tag.Tag{Id: 1, Name: "已存在标签"}
	mockTagModel.On("FindByName", ctx, "已存在标签").Return(existingTag, nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewCreateTagLogic(ctx, svcCtx)

	req := &types.CreateTagReq{
		Name: "已存在标签",
	}
	resp, err := logic.CreateTag(req)

	// 验证
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, errorx.ErrCodeTagAlreadyExists, err.(*errorx.CodeError).GetCode())

	mockTagModel.AssertExpectations(t)
}

// TestCreateTagLogic_CreateTag_ValidationError 测试参数验证失败
func TestCreateTagLogic_CreateTag_ValidationError(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()
	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}

	tests := []struct {
		name        string
		req         *types.CreateTagReq
		expectedErr string
	}{
		{
			name:        "名称为空",
			req:         &types.CreateTagReq{Name: ""},
			expectedErr: "标签名称不能为空",
		},
		{
			name:        "名称太短",
			req:         &types.CreateTagReq{Name: "A"},
			expectedErr: "标签名称至少2个字符",
		},
		{
			name:        "名称太长",
			req:         &types.CreateTagReq{Name: string(make([]byte, 51))},
			expectedErr: "标签名称最多50个字符",
		},
		{
			name:        "描述太长",
			req:         &types.CreateTagReq{Name: "正常", Description: string(make([]byte, 201))},
			expectedErr: "标签描述最多200个字符",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := NewCreateTagLogic(ctx, svcCtx)
			resp, err := logic.CreateTag(tt.req)

			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

// TestListTagsLogic_ListTags_Success 测试标签列表成功
func TestListTagsLogic_ListTags_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock List 返回
	tags := []*tag.Tag{
		{Id: 1, Name: "标签1", Description: "描述1", Color: "#1890ff", Status: 1, CreatedAt: testTime()},
		{Id: 2, Name: "标签2", Description: "描述2", Color: "#52c41a", Status: 1, CreatedAt: testTime()},
	}
	mockTagModel.On("List", ctx, 1, 20).Return(tags, int64(2), nil)

	// Mock CountByTag
	mockResourceTagModel.On("CountByTag", ctx, int64(1)).Return(int64(5), nil)
	mockResourceTagModel.On("CountByTag", ctx, int64(2)).Return(int64(3), nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewListTagsLogic(ctx, svcCtx)

	req := &types.ListTagsReq{Page: 1, PageSize: 20}
	resp, err := logic.ListTags(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(2), resp.Total)
	assert.Len(t, resp.List, 2)
	assert.Equal(t, int64(5), resp.List[0].UsageCount)
	assert.Equal(t, int64(3), resp.List[1].UsageCount)

	mockTagModel.AssertExpectations(t)
	mockResourceTagModel.AssertExpectations(t)
}

// TestGetTagLogic_GetTag_Success 测试获取标签详情成功
func TestGetTagLogic_GetTag_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindOne
	foundTag := &tag.Tag{Id: 1, Name: "测试标签", Description: "描述", Color: "#1890ff", Status: 1, CreatedAt: testTime()}
	mockTagModel.On("FindOne", ctx, int64(1)).Return(foundTag, nil)

	// Mock CountByTag
	mockResourceTagModel.On("CountByTag", ctx, int64(1)).Return(int64(10), nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewGetTagLogic(ctx, svcCtx)

	resp, err := logic.GetTag(1)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.Id)
	assert.Equal(t, "测试标签", resp.Name)
	assert.Equal(t, int64(10), resp.UsageCount)

	mockTagModel.AssertExpectations(t)
	mockResourceTagModel.AssertExpectations(t)
}

// TestGetTagLogic_GetTag_NotFound 测试标签不存在
func TestGetTagLogic_GetTag_NotFound(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindOne 返回错误
	mockTagModel.On("FindOne", ctx, int64(999)).Return((*tag.Tag)(nil), tag.ErrNotFound)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewGetTagLogic(ctx, svcCtx)

	resp, err := logic.GetTag(999)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, errorx.ErrCodeTagNotFound, err.(*errorx.CodeError).GetCode())

	mockTagModel.AssertExpectations(t)
}

// TestUpdateTagLogic_UpdateTag_Success 测试更新标签成功
func TestUpdateTagLogic_UpdateTag_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindOne 返回现有标签
	existingTag := &tag.Tag{Id: 1, Name: "旧名称", Description: "旧描述", Color: "#1890ff", Status: 1}
	mockTagModel.On("FindOne", ctx, int64(1)).Return(existingTag, nil)

	// Mock FindByName 新名称不存在
	mockTagModel.On("FindByName", ctx, "新名称").Return((*tag.Tag)(nil), nil)

	// Mock Update 成功
	mockTagModel.On("Update", ctx, mock.AnythingOfType("*tag.Tag")).Return(nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewUpdateTagLogic(ctx, svcCtx)

	req := &types.UpdateTagReq{
		Id:          1,
		Name:        "新名称",
		Description: "新描述",
		Color:       "#52c41a",
		Status:      1,
	}
	resp, err := logic.UpdateTag(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)

	mockTagModel.AssertExpectations(t)
}

// TestDeleteTagLogic_DeleteTag_Success 测试删除标签成功
func TestDeleteTagLogic_DeleteTag_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindOne 返回存在
	existingTag := &tag.Tag{Id: 1, Name: "待删除"}
	mockTagModel.On("FindOne", ctx, int64(1)).Return(existingTag, nil)

	// Mock CountByTag 返回未使用
	mockResourceTagModel.On("CountByTag", ctx, int64(1)).Return(int64(0), nil)

	// Mock Delete 成功
	mockTagModel.On("Delete", ctx, int64(1)).Return(nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewDeleteTagLogic(ctx, svcCtx)

	resp, err := logic.DeleteTag(1)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)

	mockTagModel.AssertExpectations(t)
	mockResourceTagModel.AssertExpectations(t)
}

// TestDeleteTagLogic_DeleteTag_InUse 测试删除正在使用的标签
func TestDeleteTagLogic_DeleteTag_InUse(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	existingTag := &tag.Tag{Id: 1, Name: "使用中"}
	mockTagModel.On("FindOne", ctx, int64(1)).Return(existingTag, nil)

	// Mock CountByTag 返回正在使用
	mockResourceTagModel.On("CountByTag", ctx, int64(1)).Return(int64(5), nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewDeleteTagLogic(ctx, svcCtx)

	resp, err := logic.DeleteTag(1)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, errorx.ErrCodeTagInUse, err.(*errorx.CodeError).GetCode())

	mockTagModel.AssertExpectations(t)
	mockResourceTagModel.AssertExpectations(t)
}

// TestAssignTagsLogic_AssignTags_Success 测试关联标签成功
func TestAssignTagsLogic_AssignTags_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindAll 验证标签存在
	mockTagModel.On("FindOne", ctx, int64(1)).Return(&tag.Tag{Id: 1, Name: "标签1"}, nil)
	mockTagModel.On("FindOne", ctx, int64(2)).Return(&tag.Tag{Id: 2, Name: "标签2"}, nil)

	// Mock BatchAssign 成功
	mockResourceTagModel.On("BatchAssign", ctx, int64(100), "catalog_category", []int64{1, 2}).Return(nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewAssignTagsLogic(ctx, svcCtx)

	req := &types.AssignTagsReq{
		ResourceId:   100,
		ResourceType: "catalog_category",
		TagIds:       []int64{1, 2},
	}
	resp, err := logic.AssignTags(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Equal(t, 2, resp.AssignedCount)

	mockTagModel.AssertExpectations(t)
	mockResourceTagModel.AssertExpectations(t)
}

// TestSearchByTagsLogic_SearchByTags_Success 测试按标签搜索成功
func TestSearchByTagsLogic_SearchByTags_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindByTags 返回资源ID列表
	resourceIDs := []int64{100, 200, 300}
	mockResourceTagModel.On("FindByTags", ctx, []int64{1, 2}, "catalog_category").Return(resourceIDs, nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewSearchByTagsLogic(ctx, svcCtx)

	req := &types.SearchByTagsReq{
		TagIds:       []int64{1, 2},
		ResourceType: "catalog_category",
		Page:         1,
		PageSize:     20,
	}
	resp, err := logic.SearchByTags(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(3), resp.Total)
	assert.Len(t, resp.Resources, 3)

	mockResourceTagModel.AssertExpectations(t)
}

// testTime 辅助函数
func testTime() time.Time {
	return time.Now()
}

// TestUpdateTagLogic_UpdateTag_NotFound 测试更新不存在的标签
func TestUpdateTagLogic_UpdateTag_NotFound(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	mockTagModel.On("FindOne", ctx, int64(999)).Return((*tag.Tag)(nil), tag.ErrNotFound)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewUpdateTagLogic(ctx, svcCtx)

	req := &types.UpdateTagReq{
		Id:   999,
		Name: "新名称",
	}
	resp, err := logic.UpdateTag(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, errorx.ErrCodeTagNotFound, err.(*errorx.CodeError).GetCode())

	mockTagModel.AssertExpectations(t)
}

// TestUnassignTagsLogic_UnassignTags_Success 测试移除标签成功
func TestUnassignTagsLogic_UnassignTags_Success(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock BatchUnassign 成功
	mockResourceTagModel.On("BatchUnassign", ctx, int64(100), "catalog_category", []int64{1, 2}).Return(nil)

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewUnassignTagsLogic(ctx, svcCtx)

	req := &types.UnassignTagsReq{
		ResourceId:   100,
		ResourceType: "catalog_category",
		TagIds:       []int64{1, 2},
	}
	resp, err := logic.UnassignTags(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)

	mockResourceTagModel.AssertExpectations(t)
}

// TestCreateTagLogic_CreateTag_DBError 测试数据库错误
func TestCreateTagLogic_CreateTag_DBError(t *testing.T) {
	mockTagModel := new(mocks.MockTagModel)
	mockResourceTagModel := new(mocks.MockResourceTagModel)

	ctx := context.Background()

	// Mock FindByName 返回错误
	mockTagModel.On("FindByName", ctx, "测试标签").Return((*tag.Tag)(nil), errors.New("db error"))

	svcCtx := &svc.ServiceContext{
		TagModel:         mockTagModel,
		ResourceTagModel: mockResourceTagModel,
	}
	logic := NewCreateTagLogic(ctx, svcCtx)

	req := &types.CreateTagReq{
		Name: "测试标签",
	}
	resp, err := logic.CreateTag(req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	mockTagModel.AssertExpectations(t)
}
