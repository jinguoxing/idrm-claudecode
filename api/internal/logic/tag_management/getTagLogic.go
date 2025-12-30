// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tag_management

import (
	"context"
	"fmt"

	"api/internal/svc"
	"api/internal/types"

	"idrm/model/tag_management/tag"
	"idrm/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取标签详情
func NewGetTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagLogic {
	return &GetTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagLogic) GetTag(id int64) (resp *types.GetTagResp, err error) {
	// 查询标签
	result, err := l.svcCtx.TagModel.FindOne(l.ctx, id)
	if err != nil {
		if err == tag.ErrNotFound {
			return nil, errorx.New(errorx.ErrCodeTagNotFound)
		}
		l.Errorf("查询标签失败: %v", err)
		return nil, fmt.Errorf("查询标签失败: %w", err)
	}

	// 获取使用次数
	usageCount, _ := l.svcCtx.ResourceTagModel.CountByTag(l.ctx, result.Id)

	return &types.GetTagResp{
		TagInfo: types.TagInfo{
			Id:          result.Id,
			Name:        result.Name,
			Description: result.Description,
			Color:       result.Color,
			Status:      result.Status,
			UsageCount:  usageCount,
			CreatedAt:   result.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
