package satori

import "github.com/dezhishen/satori-server-go/pkg/application"

// Event触发
type EventEmitter interface {
	StartWithCtx(instacne application.Instance) error
}
