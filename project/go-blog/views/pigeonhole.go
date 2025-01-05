package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HTMLHandle) PigeonholeHandle(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole
	data, err := service.GetPigeonhole()
	if err != nil {
		pigeonhole.WriteError(w, err)
		return
	}
	pigeonhole.WriteData(w, data)
}
