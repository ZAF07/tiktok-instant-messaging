package service

import (
	"fmt"
	"log"

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

func (h *HTTPService) Pull() (string, error) {
	res, err := h.cache.Get("test-value")
	if err != nil {
		return "", err
	}

	return res, nil
}

func (h *HTTPService) Push(msg httpdomain.Message) (string, error) {
	if err := h.cache.Save(msg); err != nil {
		log.Printf("error saving to cache: %v", err)
		return "", err
	}
	res := fmt.Sprintf("Pushing service is done saving: %+v, from %s", msg.Text, msg.Sender)
	return res, nil
}
