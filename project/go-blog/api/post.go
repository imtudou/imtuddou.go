package api

import (
	"errors"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/untils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*APIHandle) Save(w http.ResponseWriter, r *http.Request) {
	var p = &models.Post{}
	// 获取参数
	token := r.Header.Get("Authorization")
	_, claims, err := untils.ParseToken(token)
	if err != nil {
		common.SetError(w, errors.New("登陆已过期"))
		return
	}

	params := common.GetRequestJsonParam(r)
	p.Title = params["title"].(string)
	p.Slug = params["slug"].(string)
	p.Content = params["content"].(string)
	p.Markdown = params["markdown"].(string)
	p.CategoryId = params["categoryId"].(int)
	p.UserId = claims.Uid
	p.Type = params["type"].(int)

	// 根据请求类型判断是新增还是修改
	method := r.Method
	switch method {
	case http.MethodPost:
		p.CreateAt = time.Now()
		if err := dao.SavePost(p); err != nil {
			common.SetError(w, err)
			return
		}
	case http.MethodPut:
		p.Pid = params["pid"].(int)
		p.UpdateAt = time.Now()
		if err := dao.UpdatePost(p); err != nil {
			common.SetError(w, err)
			return
		}
	}
	common.SetSuccess(w, p)
	return
}

// GetPost
func (*APIHandle) GetPost(w http.ResponseWriter, r *http.Request) {
	pidStr := strings.TrimPrefix(r.URL.Path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.SetError(w, errors.New("请求路径错误"))
		return
	}
	data, err := service.GetPostDetailByPid(pid)
	if err != nil {
		common.SetError(w, err)
	}
	common.SetSuccess(w, data)
	return
}

func (*APIHandle) SearchPost(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("val")
	data, err := service.SearchPost(val)
	if err != nil {
		common.SetError(w, err)
	}
	common.SetSuccess(w, data)
	return

}
