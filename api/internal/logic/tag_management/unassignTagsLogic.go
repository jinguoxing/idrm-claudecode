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

type UnassignTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 移除数据标签
func NewUnassignTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnassignTagsLogic {
	return &UnassignTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnassignTagsLogic) UnassignTags(req *types.UnassignTagsReq) (resp *types.UnassignTagsResp, err error) {
	// 批量移除标签关联
	if err := l.svcCtx.ResourceTagModel.BatchUnassign(
		l.ctx,
		req.ResourceId,
		req.ResourceType,
		req.TagIds,
	); err != nil {
		l.Errorf("批量移除标签失败: %v", err)
		return nil, fmt.Errorf("批量移除标签失败: %w", err)
	}

	return &types.UnassignTagsResp{
		Success: true,
	}, nil
}
