package plugin

type Instance interface {
	// 插件类型
	Type() string
	// 插件ID
	Id() string
	// 路径
	Path() string
}
