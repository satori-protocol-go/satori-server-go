package event

type EventOpCallback func(eventOp EventOp) error

type EventInstance interface {
	ChannelEventInstance
}

type EventInstanceImpl struct {
	*channelEventInstacneImpl
}

func NewEventInstance() EventInstance {
	return &EventInstanceImpl{
		channelEventInstacneImpl: newChannelEventInstanceImpl(),
	}
}
