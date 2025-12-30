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

type AssignTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 为数据打标签
func NewAssignTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignTagsLogic {
	return &AssignTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignTagsLogic) AssignTags(req *types.AssignTagsReq) (resp *types.AssignTagsResp, err error) {
	// 1. 验证所有标签ID存在
	for _, tagID := range req.TagIds {
		_, err := l.svcCtx.TagModel.FindOne(l.ctx, tagID)
		if err != nil {
			if err == tag.ErrNotFound {
				return nil, errorx.NewWithMsg(errorx.ErrCodeParamInvalid, fmt.Sprintf("标签ID %d 不存在", tagID))
			}
			return nil, fmt.Errorf("验证标签失败: %w", err)
		}
	}

	// 2. 批量关联标签
	if err := l.svcCtx.ResourceTagModel.BatchAssign(
		l.ctx,
		req.ResourceId,
		req.ResourceType,
		req.TagIds,
	); err != nil {
		l.Errorf("批量关联标签失败: %v", err)
		return nil, fmt.Errorf("批量关联标签失败: %w", err)
	}

	return &types.AssignTagsResp{
		Success:       true,
		AssignedCount: len(req.TagIds),
	}, nil
}
