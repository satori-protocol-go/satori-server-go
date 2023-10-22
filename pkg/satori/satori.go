package satori

import (
	"github.com/dezhishen/satori-sdk-go/pkg/api"
)

// API事件处理者
type APIHandler interface {
	api.SatoriApi
}

// API调用者
type APICaller interface {
	api.SatoriApi
}
