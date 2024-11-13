package config

import (
	_ "embed"
	"github.com/kamioair/quick-utils/qconfig"
)

// Config 自定义配置
var Config = struct {
}{}

func Init(module string) {
	qconfig.Load(module, &Config)
}
