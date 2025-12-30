// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tag_management

import (
	"net/http"

	"api/internal/logic/tag_management"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 为数据打标签
func AssignTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssignTagsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag_management.NewAssignTagsLogic(r.Context(), svcCtx)
		resp, err := l.AssignTags(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
