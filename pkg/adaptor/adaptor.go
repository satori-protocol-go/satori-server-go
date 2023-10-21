package adaptor

type Adaptor interface {
}

func Register(adaptor Adaptor) error {
	return nil
}
