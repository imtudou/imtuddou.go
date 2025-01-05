package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (h *HTMLHandle) IndexHandle(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	r.ParseForm()
	pageIndex, _ := strconv.Atoi(r.Form.Get("page"))
	if pageIndex == 0 {
		pageIndex = 1
	}
	slug := strings.TrimPrefix(r.URL.Path, "/")
	data, err := service.GetAllIndexInfo(slug, pageIndex, 10)
	if err != nil {
		index.WriteError(w, errors.New("系统错误,请联系管理员："+err.Error()))
		return
	}
	index.WriteData(w, data)
}
