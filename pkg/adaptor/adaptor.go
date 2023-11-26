package adaptor

import (
	"context"
	"fmt"
	"sync"
)

type AdaptorInstance interface {
	// 声明自身平台
	Platform() string
	// 生命自身ID
	SelfId() string
	// 初始化
	Init() error
	// 开始
	Start(ctx context.Context) error
}

type adaptorFactory struct {
	*sync.RWMutex
	adaptors map[string]map[string]AdaptorInstance
}

func NewGRPCAdaptorInstance() *adaptorFactory {
	return &adaptorFactory{
		adaptors: make(map[string]map[string]AdaptorInstance),
	}
}

func (h *adaptorFactory) Register(adaptor AdaptorInstance) error {
	h.Lock()
	defer h.Unlock()
	adaptorUnderSelfId, ok := h.adaptors[adaptor.Platform()]
	if !ok {
		adaptorUnderSelfId = make(map[string]AdaptorInstance)
		h.adaptors[adaptor.Platform()] = adaptorUnderSelfId
	}
	adaptorUnderSelfId[adaptor.SelfId()] = adaptor
	return nil
}

func (h *adaptorFactory) EmitEvent(platform, selfId string, eventType, eventName string, request interface{}, resp interface{}) error {
	instance, ok := h.getInstance(platform, selfId)
	if !ok {
		return fmt.Errorf("can't find adaptor instance by %s/%s", platform, selfId)
	}
	fmt.Print(instance.Platform(), instance.SelfId())
	return nil
}

func (h *adaptorFactory) getInstance(platform, selfId string) (AdaptorInstance, bool) {
	h.RLock()
	defer h.Unlock()
	adaptorUnderSelfId, ok := h.adaptors[platform]
	if !ok {
		return nil, false
	}
	result, ok := adaptorUnderSelfId[selfId]
	return result, ok
}
