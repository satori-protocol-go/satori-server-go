package satori

import (
	"github.com/dezhishen/satori-application-go/pkg/application"
	"github.com/dezhishen/satori-model-go/pkg/channel"
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
