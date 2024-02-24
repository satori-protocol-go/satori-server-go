package satori

import "github.com/satori-protocol-go/satori-server-go/pkg/application"

// Event触发
type EventEmitter interface {
	StartWithCtx(instacne application.Instance) error
}
