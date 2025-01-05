package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		// 耗时操作 使用协程处理
		templateDir := config.Cfg.SystemConfig.CurrentDir + "/template/"
		Template, err = models.InitTemplate(templateDir)
		if err != nil {
			panic(err)
		}
		w.Done() // 标识已完成
	}()
	w.Wait() // 等待

}

// GetRequestJsonParam 获取json参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var m map[string]interface{}
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, &m)
	return m
}

func SetSuccess(w http.ResponseWriter, data interface{}) {
	var ret models.Result
	ret.Code = http.StatusOK
	ret.Data = data
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	retStr, _ := json.Marshal(ret)
	w.Write(retStr)
}

func SetError(w http.ResponseWriter, err error) {
	var ret models.Result
	ret.Code = http.StatusInternalServerError
	ret.Error = err.Error()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	retStr, _ := json.Marshal(ret)
	w.Write(retStr)
}
