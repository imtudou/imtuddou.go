package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*APIHandle) LogIn(w http.ResponseWriter, r *http.Request) {

	// 接受用户名和密码 返回json数据
	params := common.GetRequestJsonParam(r)
	name := params["username"].(string)
	pwd := params["passwd"].(string)
	data, err := service.LogIn(name, pwd)
	if err != nil {
		common.SetError(w, err)
		return
	}
	common.SetSuccess(w, data)
}
