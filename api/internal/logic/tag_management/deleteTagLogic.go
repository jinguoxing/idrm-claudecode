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

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除标签
func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(id int64) (resp *types.DeleteTagResp, err error) {
	// 1. 验证标签存在
	_, err = l.svcCtx.TagModel.FindOne(l.ctx, id)
	if err != nil {
		if err == tag.ErrNotFound {
			return nil, errorx.New(errorx.ErrCodeTagNotFound)
		}
		return nil, fmt.Errorf("查询标签失败: %w", err)
	}

	// 2. 检查是否被使用
	count, err := l.svcCtx.ResourceTagModel.CountByTag(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("检查标签使用情况失败: %w", err)
	}
	if count > 0 {
		return nil, errorx.New(errorx.ErrCodeTagInUse)
	}

	// 3. 删除标签
	if err := l.svcCtx.TagModel.Delete(l.ctx, id); err != nil {
		l.Errorf("删除标签失败: %v", err)
		return nil, fmt.Errorf("删除标签失败: %w", err)
	}

	return &types.DeleteTagResp{
		Success: true,
	}, nil
}
