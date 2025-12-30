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

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新标签
func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.UpdateTagReq) (resp *types.UpdateTagResp, err error) {
	// 1. 验证标签存在
	existing, err := l.svcCtx.TagModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if err == tag.ErrNotFound {
			return nil, errorx.New(errorx.ErrCodeTagNotFound)
		}
		return nil, fmt.Errorf("查询标签失败: %w", err)
	}

	// 2. 检查名称唯一性（排除自己）
	if req.Name != existing.Name {
		duplicate, err := l.svcCtx.TagModel.FindByName(l.ctx, req.Name)
		if err != nil {
			return nil, fmt.Errorf("检查名称唯一性失败: %w", err)
		}
		if duplicate != nil {
			return nil, errorx.New(errorx.ErrCodeTagAlreadyExists)
		}
	}

	// 3. 更新数据
	existing.Name = req.Name
	existing.Description = req.Description
	existing.Color = req.Color
	existing.Status = req.Status
	updatedBy := int64(1) // TODO: 从上下文获取用户ID
	existing.UpdatedBy = &updatedBy

	if err := l.svcCtx.TagModel.Update(l.ctx, existing); err != nil {
		l.Errorf("更新标签失败: %v", err)
		return nil, fmt.Errorf("更新标签失败: %w", err)
	}

	return &types.UpdateTagResp{
		Success: true,
	}, nil
}
