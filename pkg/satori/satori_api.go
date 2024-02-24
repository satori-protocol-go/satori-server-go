package satori

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/channel"
	"github.com/satori-protocol-go/satori-server-go/pkg/application"
)

// API处理
type APIHandler interface {
	channelAPIHandler
}

type channelAPIHandler interface {
	HandleChannelGet(instance application.Instance, channel_id string) (*channel.Channel, error)
	HandleChannelList(instance application.Instance, guild_id string, next string) (*channel.ChannelList, error)
	HandleChannelCreate(instance application.Instance, guild_id string, data channel.Channel) (*channel.Channel, error)
	HandleChannelUpdate(instance application.Instance, channel_id string, data channel.Channel) error
	HandleChannelDelete(instance application.Instance, channel_id string) error
	HandleUserChannelCreate(instance application.Instance, user_id string) (*channel.Channel, error)
}
