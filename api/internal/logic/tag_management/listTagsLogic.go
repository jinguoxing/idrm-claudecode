// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tag_management

import (
	"context"
	"fmt"

	"api/internal/svc"
	"api/internal/types"

	"idrm/model/tag_management/resource_tag"
	"idrm/model/tag_management/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标签列表
func NewListTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagsLogic {
	return &ListTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagsLogic) ListTags(req *types.ListTagsReq) (resp *types.ListTagsResp, err error) {
	// 默认分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	var results []*tag.Tag
	var total int64

	// 根据是否有关键词选择查询方式
	if req.Keyword != "" {
		results, total, err = l.svcCtx.TagModel.Search(l.ctx, req.Keyword, page, pageSize)
	} else {
		results, total, err = l.svcCtx.TagModel.List(l.ctx, page, pageSize)
	}

	if err != nil {
		l.Errorf("查询标签列表失败: %v", err)
		return nil, fmt.Errorf("查询标签列表失败: %w", err)
	}

	// 转换为响应格式
	list := make([]types.TagInfo, 0, len(results))
	for _, t := range results {
		// 获取使用次数
		usageCount, _ := l.svcCtx.ResourceTagModel.CountByTag(l.ctx, t.Id)

		list = append(list, types.TagInfo{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Color:       t.Color,
			Status:      t.Status,
			UsageCount:  usageCount,
			CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.ListTagsResp{
		Total: total,
		List:  list,
	}, nil
}
