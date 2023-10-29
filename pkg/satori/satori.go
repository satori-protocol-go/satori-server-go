package satori

import (
	"github.com/dezhishen/satori-server-go/internal/event"
)

// API事件处理者
type APIHandler interface {
	*event.EventInstance
}

// API调用者
type APICaller interface {
}
