package event

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/guild"
	"github.com/satori-protocol-go/satori-model-go/pkg/guildmember"
	"github.com/satori-protocol-go/satori-model-go/pkg/user"
)

type GuildMemberEventInstance interface {
	// GuildMemberAdded
	EmitGuildMemberAdded(g guild.Guild, gm guildmember.GuildMember, u user.User) error
	ListenGuildMemberAdded(callback EventOpCallback) error
	// GuildMemberUpdated
	EmitGuildMemberUpdated(g guild.Guild, gm guildmember.GuildMember, u user.User) error
	ListenGuildMemberUpdated(callback EventOpCallback) error
	// GuildMemberRemoved
	EmitGuildMemberRemoved(g guild.Guild, gm guildmember.GuildMember, u user.User) error
	ListenGuildMemberRemoved(callback EventOpCallback) error
	// GuildMemberRequest
	EmitGuildMemberRequest(g guild.Guild, gm guildmember.GuildMember, u user.User) error
	ListenGuildMemberRequest(callback EventOpCallback) error
}

func newGuildMemberEventInstanceImpl() GuildMemberEventInstance {
	return &guildMemberEventInstanceImpl{}
}

type guildMemberEventInstanceImpl struct {
}

// GuildMemberAdded
func (gI *guildMemberEventInstanceImpl) EmitGuildMemberAdded(g guild.Guild, gm guildmember.GuildMember, u user.User) error {
	return nil
}
func (gI *guildMemberEventInstanceImpl) ListenGuildMemberAdded(callback EventOpCallback) error {
	return nil
}

// GuildMemberUpdated
func (gI *guildMemberEventInstanceImpl) EmitGuildMemberUpdated(g guild.Guild, gm guildmember.GuildMember, u user.User) error {
	return nil
}
func (gI *guildMemberEventInstanceImpl) ListenGuildMemberUpdated(callback EventOpCallback) error {
	return nil
}

// GuildMemberRemoved
func (gI *guildMemberEventInstanceImpl) EmitGuildMemberRemoved(g guild.Guild, gm guildmember.GuildMember, u user.User) error {
	return nil
}
func (gI *guildMemberEventInstanceImpl) ListenGuildMemberRemoved(callback EventOpCallback) error {
	return nil
}

// GuildMemberRequest
func (gI *guildMemberEventInstanceImpl) EmitGuildMemberRequest(g guild.Guild, gm guildmember.GuildMember, u user.User) error {
	return nil
}
func (gI *guildMemberEventInstanceImpl) ListenGuildMemberRequest(callback EventOpCallback) error {
	return nil
}
