package service

type HTTPService struct {
}

func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

func (h *HTTPService) Pull() string {
	return "Pulling Service"
}

func (h *HTTPService) Push() string {
	return "Pushing Service"
}
