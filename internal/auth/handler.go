package auth

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}