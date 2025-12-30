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

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建标签
func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagReq) (resp *types.CreateTagResp, err error) {
	// 1. 参数验证
	if err := l.validateReq(req); err != nil {
		return nil, err
	}

	// 2. 检查名称是否已存在
	existing, err := l.svcCtx.TagModel.FindByName(l.ctx, req.Name)
	if err != nil {
		l.Errorf("查询标签失败: %v", err)
		return nil, fmt.Errorf("查询标签失败: %w", err)
	}
	if existing != nil {
		return nil, errorx.New(errorx.ErrCodeTagAlreadyExists)
	}

	// 3. 设置默认颜色
	color := req.Color
	if color == "" {
		color = tag.DefaultColor
	}

	// 4. 构建数据
	tagData := &tag.Tag{
		Name:        req.Name,
		Description: req.Description,
		Color:       color,
		Status:      tag.StatusEnabled,
		CreatedBy:   1, // TODO: 从上下文获取用户ID
	}

	// 5. 插入数据库
	result, err := l.svcCtx.TagModel.Insert(l.ctx, tagData)
	if err != nil {
		l.Errorf("创建标签失败: %v", err)
		return nil, fmt.Errorf("创建标签失败: %w", err)
	}

	l.Infof("标签创建成功: id=%d, name=%s", result.Id, result.Name)

	return &types.CreateTagResp{
		Id: result.Id,
	}, nil
}

// validateReq 验证请求参数
func (l *CreateTagLogic) validateReq(req *types.CreateTagReq) error {
	if req.Name == "" {
		return errorx.NewWithMsg(errorx.ErrCodeParamInvalid, "标签名称不能为空")
	}
	if len(req.Name) < 2 {
		return errorx.NewWithMsg(errorx.ErrCodeParamInvalid, "标签名称至少2个字符")
	}
	if len(req.Name) > 50 {
		return errorx.NewWithMsg(errorx.ErrCodeParamInvalid, "标签名称最多50个字符")
	}
	if len(req.Description) > 200 {
		return errorx.NewWithMsg(errorx.ErrCodeParamInvalid, "标签描述最多200个字符")
	}
	return nil
}
