package ports

import (
	message "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message_dto"

	messageProto "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message-service/proto"
)

// Interfaces which Adapters will imolement in order to interact with the core

type IHTTPService interface {
	Push(msg *message.Message) error
	Pull() (*messageProto.Messages, error)
}

type ICacheStore interface {
	Save(msg *messageProto.Message) error
	Get(string) ([]byte, error)
}
