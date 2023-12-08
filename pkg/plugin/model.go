package plugin

import "github.com/dezhishen/satori-server-go/pkg/adaptor"

type PluginType string

const (
	PLUGIN_ADAPTOR PluginType = "adaptor"
)

type Instance interface {
	// 插件类型
	Type() PluginType
	// 插件ID
	Id() string
	// 路径
	Path() string
	// 适配器类型的插件会提供适配器
	Adatpor() adaptor.Instance
}
