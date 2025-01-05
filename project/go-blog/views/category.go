package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (h *HTMLHandle) CategoryHandle(w http.ResponseWriter, r *http.Request) {
	category := common.Template.Category
	cidStr := strings.TrimPrefix(r.URL.Path, "/c/")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		category.WriteError(w, errors.New("请求路径错误"))
		return
	}
	r.ParseForm()
	pageIndex, _ := strconv.Atoi(r.Form.Get("page"))
	if pageIndex == 0 {
		pageIndex = 1
	}
	data, err := service.GetPostPageByCategoryID(cid, pageIndex, 10)
	if err != nil {
		category.WriteError(w, errors.New("系统错误,请联系管理员："+err.Error()))
		return
	}
	category.WriteData(w, data)
}
