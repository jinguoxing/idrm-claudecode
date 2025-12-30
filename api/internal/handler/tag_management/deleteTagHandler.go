// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tag_management

import (
	"net/http"

	"api/internal/logic/tag_management"
	"api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除标签
func DeleteTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag_management.NewDeleteTagLogic(r.Context(), svcCtx)
		// 从URL路径参数获取ID
		id, _ := httpx.Atoi(r.URL.Query().Get("id"))
		resp, err := l.DeleteTag(int64(id))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
