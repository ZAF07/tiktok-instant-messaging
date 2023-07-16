package service

import (
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/ports"
)

type HTTPService struct {
	cache ports.ICacheStore
}

func NewHTTPService(c ports.ICacheStore) *HTTPService {
	return &HTTPService{
		cache: c,
	}
}

func (h *HTTPService) Pull() string {
	return "Pulling Service"
}

func (h *HTTPService) Push() string {
	h.cache.Save()
	return "Pushing Service"
}
