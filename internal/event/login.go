package event

import "github.com/dezhishen/satori-model-go/pkg/login"

type LoginEventInstance interface {
	// login-added
	EmitLoginAdded(l login.Login) error
	ListenLoginAdded(callback EventOpCallback) error
	// login-removed
	EmitLoginRemoved(l login.Login) error
	ListenLoginRemoved(callback EventOpCallback) error
	// login-updated
	EmitLoginUpdated(l login.Login) error
	ListenLoginUpdated(callback EventOpCallback) error
}

func newLoginEventInstanceImpl() LoginEventInstance {
	return &loginEventInstanceImpl{}
}

type loginEventInstanceImpl struct {
}

// EmitLoginAdded implements LoginEventInstance.
func (*loginEventInstanceImpl) EmitLoginAdded(l login.Login) error {
	panic("unimplemented")
}

// EmitLoginRemoved implements LoginEventInstance.
func (*loginEventInstanceImpl) EmitLoginRemoved(l login.Login) error {
	panic("unimplemented")
}

// EmitLoginUpdated implements LoginEventInstance.
func (*loginEventInstanceImpl) EmitLoginUpdated(l login.Login) error {
	panic("unimplemented")
}

// ListenLoginAdded implements LoginEventInstance.
func (*loginEventInstanceImpl) ListenLoginAdded(callback EventOpCallback) error {
	panic("unimplemented")
}

// ListenLoginRemoved implements LoginEventInstance.
func (*loginEventInstanceImpl) ListenLoginRemoved(callback EventOpCallback) error {
	panic("unimplemented")
}

// ListenLoginUpdated implements LoginEventInstance.
func (*loginEventInstanceImpl) ListenLoginUpdated(callback EventOpCallback) error {
	panic("unimplemented")
}
