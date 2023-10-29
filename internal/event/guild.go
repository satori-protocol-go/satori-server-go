package event

import (
	"github.com/dezhishen/satori-model-go/pkg/guild"
)

type GuildEventInstance interface {
	EmitGuildAdded(g guild.Guild) error
	ListenGuildAdded(callback EventOpCallback) error
	EmitGuildUpdated(g guild.Guild) error
	ListenGuildUpdated(callback EventOpCallback) error
	EmitGuildRemoved(g guild.Guild) error
	ListGuildRemoved(callback EventOpCallback) error
	EmitGuildRequest(g guild.Guild) error
	ListGuildRequest(callback EventOpCallback) error
}

func newGuildEventInstanceImpl() GuildEventInstance {
	return &guildEventInstanceImpl{}
}

type guildEventInstanceImpl struct {
}

func (gI *guildEventInstanceImpl) EmitGuildAdded(g guild.Guild) error {
	return nil
}
func (gI *guildEventInstanceImpl) ListenGuildAdded(callback EventOpCallback) error {
	return nil
}
func (gI *guildEventInstanceImpl) EmitGuildUpdated(g guild.Guild) error {
	return nil
}
func (gI *guildEventInstanceImpl) ListenGuildUpdated(callback EventOpCallback) error {
	return nil
}
func (gI *guildEventInstanceImpl) EmitGuildRemoved(g guild.Guild) error {
	return nil
}
func (gI *guildEventInstanceImpl) ListGuildRemoved(callback EventOpCallback) error {
	return nil
}
func (gI *guildEventInstanceImpl) EmitGuildRequest(g guild.Guild) error {
	return nil
}
func (gI *guildEventInstanceImpl) ListGuildRequest(callback EventOpCallback) error {
	return nil
}
