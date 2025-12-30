// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tag_management

import (
	"context"
	"fmt"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 按标签搜索数据
func NewSearchByTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByTagsLogic {
	return &SearchByTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchByTagsLogic) SearchByTags(req *types.SearchByTagsReq) (resp *types.SearchByTagsResp, err error) {
	// 1. 查询包含所有指定标签的资源ID列表
	resourceIDs, err := l.svcCtx.ResourceTagModel.FindByTags(
		l.ctx,
		req.TagIds,
		req.ResourceType,
	)
	if err != nil {
		l.Errorf("按标签搜索资源失败: %v", err)
		return nil, fmt.Errorf("按标签搜索资源失败: %w", err)
	}

	// 2. 根据资源ID查询资源详情
	// TODO: 这里需要调用具体的资源服务来获取资源详情
	// 目前返回占位数据
	resources := make([]types.ResourceInfo, 0, len(resourceIDs))
	for _, resourceID := range resourceIDs {
		resources = append(resources, types.ResourceInfo{
			Id:   resourceID,
			Name: fmt.Sprintf("Resource-%d", resourceID),
			Type: req.ResourceType,
		})
	}

	return &types.SearchByTagsResp{
		Total:     int64(len(resourceIDs)),
		Resources: resources,
	}, nil
}
