package ports

import httpdomain "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/domain/http_domain"

// Interfaces which Adapters will imolement in order to interact with the core

type IHTTPService interface {
	Push(msg httpdomain.Message) string
	Pull() string
}

type ICacheStore interface {
	Save(msg httpdomain.Message)
}
