package event

import (
	"github.com/dezhishen/satori-model-go/pkg/guild"
	"github.com/dezhishen/satori-model-go/pkg/guildrole"
)

type GuildRoleEventInstance interface {
	EmitGuildRoleCreated(g guild.Guild, gr guildrole.GuildRole) error
	ListenGuildRoleCreated(callback EventOpCallback) error
	EmitGuildRoleUpdated(g guild.Guild, gr guildrole.GuildRole) error
	ListenGuildRoleUpdated(callback EventOpCallback) error
	EmitGuildRoleDelete(g guild.Guild, gr guildrole.GuildRole) error
	ListGuildRoleDelete(callback EventOpCallback) error
}

func newGuildRoleEventInstanceImpl() GuildRoleEventInstance {
	return &GuildRoleEventInstanceImpl{}
}

type GuildRoleEventInstanceImpl struct {
}

func (gI *GuildRoleEventInstanceImpl) EmitGuildRoleCreated(g guild.Guild, gr guildrole.GuildRole) error {
	return nil
}
func (gI *GuildRoleEventInstanceImpl) ListenGuildRoleCreated(callback EventOpCallback) error {
	return nil
}
func (gI *GuildRoleEventInstanceImpl) EmitGuildRoleUpdated(g guild.Guild, gr guildrole.GuildRole) error {
	return nil
}
func (gI *GuildRoleEventInstanceImpl) ListenGuildRoleUpdated(callback EventOpCallback) error {
	return nil
}
func (gI *GuildRoleEventInstanceImpl) EmitGuildRoleDelete(g guild.Guild, gr guildrole.GuildRole) error {
	return nil
}
func (gI *GuildRoleEventInstanceImpl) ListGuildRoleDelete(callback EventOpCallback) error {
	return nil
}
