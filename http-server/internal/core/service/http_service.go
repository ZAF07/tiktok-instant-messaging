package service

import (
	"fmt"

	httpdomain "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/domain/http_domain"
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

func (h *HTTPService) Push(msg httpdomain.Message) string {
	h.cache.Save(msg)
	res := fmt.Sprintf("Pushing service is saving: %+v, from %s", msg.Text, msg.Sender)
	return res
}
