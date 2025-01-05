package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

func (*HTMLHandle) AboutHandle(w http.ResponseWriter, r *http.Request) {
	about := common.Template.About
	about.WriteData(w, config.Cfg.Viewer)
}
