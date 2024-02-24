package event

import "github.com/satori-protocol-go/satori-model-go/pkg/message"

type MessageEventInstance interface {
	EmitMessageCreated(msg message.Message) error
	ListenMessageCreated(callback EventOpCallback) error
	EmitMessageUpdated(msg message.Message) error
	ListenMessageUpdated(callback EventOpCallback) error
	EmitMessageDeleted(msg message.Message) error
	ListenMessageDeleted(callback EventOpCallback) error
}

type messageEventInstanceImpl struct {
}

func newMessageEventInstanceImpl() MessageEventInstance {
	return &messageEventInstanceImpl{}
}

// EmitMessageCreated implements MessageEventInstance.
func (*messageEventInstanceImpl) EmitMessageCreated(msg message.Message) error {
	panic("unimplemented")
}

// EmitMessageDeleted implements MessageEventInstance.
func (*messageEventInstanceImpl) EmitMessageDeleted(msg message.Message) error {
	panic("unimplemented")
}

// EmitMessageUpdated implements MessageEventInstance.
func (*messageEventInstanceImpl) EmitMessageUpdated(msg message.Message) error {
	panic("unimplemented")
}

// ListenMessageCreated implements MessageEventInstance.
func (*messageEventInstanceImpl) ListenMessageCreated(callback EventOpCallback) error {
	panic("unimplemented")
}

// ListenMessageDeleted implements MessageEventInstance.
func (*messageEventInstanceImpl) ListenMessageDeleted(callback EventOpCallback) error {
	panic("unimplemented")
}

// ListenMessageUpdated implements MessageEventInstance.
func (*messageEventInstanceImpl) ListenMessageUpdated(callback EventOpCallback) error {
	panic("unimplemented")
}
