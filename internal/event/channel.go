package event

import "github.com/dezhishen/satori-model-go/pkg/channel"

type ChannelEventInstance interface {
	EmitChannelCreated(c channel.Channel) error
	ListenChannelCreated(callback EventOpCallback) error
	EmitChannelUpdated(c channel.Channel) error
	ListenChannelUpdated(callback EventOpCallback) error
	EmitChannelDeleted(channel_id string) error
	ListChannelDeleted(callback EventOpCallback) error
	EmitUserChannelCreated(user_id, guild_id string) error
	ListUserChannelCreated(callback EventOpCallback) error
}

func newChannelEventInstanceImpl() ChannelEventInstance {
	return &channelEventInstacneImpl{}
}

type channelEventInstacneImpl struct {
}

func (eI *channelEventInstacneImpl) EmitChannelCreated(c channel.Channel) error {
	return nil
}

func (eI *channelEventInstacneImpl) ListenChannelCreated(callback EventOpCallback) error {
	return nil
}

func (eI *channelEventInstacneImpl) EmitChannelUpdated(c channel.Channel) error {
	return nil
}

func (eI *channelEventInstacneImpl) ListenChannelUpdated(callback EventOpCallback) error {
	return nil
}

func (eI *channelEventInstacneImpl) EmitChannelDeleted(channel_id string) error {
	return nil
}
func (eI *channelEventInstacneImpl) ListChannelDeleted(callback EventOpCallback) error {
	return nil
}
func (eI *channelEventInstacneImpl) EmitUserChannelCreated(user_id, guild_id string) error {
	return nil
}
func (eI *channelEventInstacneImpl) ListUserChannelCreated(callback EventOpCallback) error {
	return nil
}
