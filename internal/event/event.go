package event

type EventOpCallback func(eventOp EventOp) error

type EventInstance interface {
	ChannelEventInstance
	GuildEventInstance
	GuildMemberEventInstance
	GuildRoleEventInstance
	LoginEventInstance
	MessageEventInstance
}

type EventInstanceImpl struct {
	ChannelEventInstance
	GuildEventInstance
	GuildMemberEventInstance
	GuildRoleEventInstance
	LoginEventInstance
	MessageEventInstance
}

func NewEventInstance() EventInstance {
	return &EventInstanceImpl{
		ChannelEventInstance:     newChannelEventInstanceImpl(),
		GuildEventInstance:       newGuildEventInstanceImpl(),
		GuildMemberEventInstance: newGuildMemberEventInstanceImpl(),
		GuildRoleEventInstance:   newGuildRoleEventInstanceImpl(),
		LoginEventInstance:       newLoginEventInstanceImpl(),
		MessageEventInstance:     newMessageEventInstanceImpl(),
	}
}
