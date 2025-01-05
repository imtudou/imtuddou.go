package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLHandle) PostDetailHandle(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	// 获取地址栏参数 pid
	pidStr := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/p/"), ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteError(w, errors.New("请求路径错误"))
		return
	}
	data, err := service.GetPostDetailByPid(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询数据异常"))
		return
	}
	detail.WriteData(w, data)
}
