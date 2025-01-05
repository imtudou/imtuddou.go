package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HTMLHandle) WritingHadle(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	data, err := service.GetWriting()
	if err != nil {
		writing.WriteError(w, err)
		return
	}
	writing.WriteData(w, data)
}
