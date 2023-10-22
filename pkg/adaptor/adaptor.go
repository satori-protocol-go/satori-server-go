package adaptor

type Adaptor interface {
	Plaform() string
	SelfId() string
}

type adaptorHandler struct {
	adaptors []Adaptor
}

func NewHandler() *adaptorHandler {
	return &adaptorHandler{}
}

func (h *adaptorHandler) Register(adaptor Adaptor) error {
	h.adaptors = append(h.adaptors, adaptor)
	return nil
}
