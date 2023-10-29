package event

type EventOpCallback func(eventOp EventOp) error

type EventInstance interface {
	ChannelEventInstance
	GuildEventInstance
	GuildMemberEventInstance
}

type EventInstanceImpl struct {
	ChannelEventInstance
	GuildEventInstance
	GuildMemberEventInstance
}

func NewEventInstance() EventInstance {
	return &EventInstanceImpl{
		ChannelEventInstance:     newChannelEventInstanceImpl(),
		GuildEventInstance:       newGuildEventInstanceImpl(),
		GuildMemberEventInstance: newGuildMemberEventInstanceImpl(),
	}
}
