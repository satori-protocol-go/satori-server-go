package plugin

import "context"

type Config interface {
}

// 加载插件，且返回Instance列表
func Load(ctx context.Context, conf Config) ([]Instance, error) {
	return nil, nil
}
