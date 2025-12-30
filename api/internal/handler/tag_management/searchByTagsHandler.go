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

// 按标签搜索数据
func SearchByTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchByTagsResp
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag_management.NewSearchByTagsLogic(r.Context(), svcCtx)
		err := l.SearchByTags(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
