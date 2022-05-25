package service

// Services is the interface for enlosing all the services
type Services interface {
}

type services struct {
}

// Init intializes the services
func Init() Services {
	return &services{}
}
