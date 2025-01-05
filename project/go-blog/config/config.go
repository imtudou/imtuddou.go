package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer       Viewer
	SystemConfig SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动的时候会读取init方法
	Cfg = new(tomlConfig)
	Cfg.SystemConfig.AppName = "go-blog"
	Cfg.SystemConfig.Version = 1.0
	path, _ := os.Getwd()
	Cfg.SystemConfig.CurrentDir = path
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}

}
