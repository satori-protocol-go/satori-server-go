package event

import "github.com/dezhishen/satori-model-go/pkg/channel"

type EventOpCallback func(eventOp EventOp) error

type EventInstance interface {
	EmitChannelCreated(c channel.Channel) error
	ListenChannelCreated(callback EventOpCallback) error
	EmitChannelUpdated(c channel.Channel) error
	ListenChannelUpdated(callback EventOpCallback) error
	EmitChannelDeleted(channel_id string) error
	ListChannelDelete(callback EventOpCallback) error
	EmitUserChannelCreated(user_id, guild_id string) error
	ListUserChannelCreated(callback EventOpCallback) error
}

type EventInstanceImpl struct {
}

func NewEventInstance() EventInstance {
	return &EventInstanceImpl{}
}

func (eI *EventInstanceImpl) EmitChannelCreated(c channel.Channel) error {
	return nil
}

func (eI *EventInstanceImpl) ListenChannelCreated(callback EventOpCallback) error {
	return nil
}

func (eI *EventInstanceImpl) EmitChannelUpdated(c channel.Channel) error {
	return nil
}

func (eI *EventInstanceImpl) ListenChannelUpdated(callback EventOpCallback) error {
	return nil
}

func (eI *EventInstanceImpl) EmitChannelDeleted(channel_id string) error {
	return nil
}
func (eI *EventInstanceImpl) ListChannelDelete(callback EventOpCallback) error {
	return nil
}
func (eI *EventInstanceImpl) EmitUserChannelCreated(user_id, guild_id string) error {
	return nil
}
func (eI *EventInstanceImpl) ListUserChannelCreated(callback EventOpCallback) error {
	return nil
}
