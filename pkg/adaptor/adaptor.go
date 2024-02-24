package adaptor

import (
	"context"
	"sync"

	"github.com/satori-protocol-go/satori-server-go/pkg/config"
	"github.com/satori-protocol-go/satori-server-go/pkg/satori"
)

// Instance 适配器实例
type Instance interface {
	satori.APIHandler
	satori.EventEmitter
	// 声明自身平台
	Platform() string
	// 生命自身ID
	SelfId() string
	// 初始化
	Init() error
	// 开始
	Start(ctx context.Context) error
}

type InstanceRegistry interface {
	Register(adaptor Instance) error
	GetInstance(platform, selfId string) (Instance, bool)
	GetAllPlatform() []string
	GetAllSelfIdByPlatform(platform string) []string
}

type simpleRWRegistry struct {
	*sync.RWMutex
	cfg      config.Cfg
	adaptors map[string]map[string]Instance
}

func NewSimpleRegistry(cfg config.Cfg) InstanceRegistry {
	return &simpleRWRegistry{
		cfg:      cfg,
		adaptors: make(map[string]map[string]Instance),
	}
}

func (h *simpleRWRegistry) Register(instance Instance) error {
	h.Lock()
	defer h.Unlock()
	instanceUnderSelfId, ok := h.adaptors[instance.Platform()]
	if !ok {
		h.adaptors[instance.Platform()] = make(map[string]Instance)
	}
	instanceUnderSelfId[instance.SelfId()] = instance
	return nil
}

func (h *simpleRWRegistry) GetInstance(platform, selfId string) (Instance, bool) {
	h.RLock()
	defer h.Unlock()
	instanceUnderSelfId, ok := h.adaptors[platform]
	if !ok {
		return nil, false
	}
	result, ok := instanceUnderSelfId[selfId]
	return result, ok
}
func (h *simpleRWRegistry) GetAllPlatform() []string {
	h.RLock()
	defer h.Unlock()
	if len(h.adaptors) == 0 {
		return nil
	}
	result := make([]string, len(h.adaptors))
	i := 0
	for k := range h.adaptors {
		result[i] = k
		i++
	}
	return result
}

func (h *simpleRWRegistry) GetAllSelfIdByPlatform(platform string) []string {
	h.RLock()
	defer h.Unlock()
	instanceUnderSelfId, ok := h.adaptors[platform]
	if !ok {
		return nil
	}
	if len(instanceUnderSelfId) == 0 {
		return nil
	}
	result := make([]string, len(instanceUnderSelfId))
	i := 0
	for k := range instanceUnderSelfId {
		result[i] = k
		i++
	}
	return result
}
