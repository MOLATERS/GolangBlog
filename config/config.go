package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
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
type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

var Cfg *tomlConfig

func init() {
	//程序启动的时候会启动方法
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	path, _ := os.Getwd()
	Cfg.System.CurrentDir = path

	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
