package service

import (
	"errors"
	"log"

	messageProto "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message-service/proto"
	message "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message_dto"
	"github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/ports"
	"github.com/golang/protobuf/proto"
)

type HTTPService struct {
	cache ports.ICacheStore
}

func NewHTTPService(c ports.ICacheStore) *HTTPService {
	return &HTTPService{
		cache: c,
	}
}

func (h *HTTPService) Pull() (*message.Message, error) {
	// Get key:value from cache
	res, err := h.cache.Get("test-key")
	if err != nil {
		return nil, err
	}

	// Unmarshal into Proto
	result := &messageProto.Message{}
	uErr := proto.Unmarshal(res, result)
	if uErr != nil {
		log.Fatalf("error unmarshalling into struct: %v", uErr)
	}

	// Convert into DTO
	foundMsg := &message.Message{
		Id:        result.Id,
		ChatId:    result.ChatId,
		SenderId:  result.SenderId,
		Content:   result.Content,
		TimeStamp: result.TimeStamp,
	}
	return foundMsg, nil
}

func (h *HTTPService) Push(msg *message.Message) error {
	mockTextID := "mock_text_id"
	// Validate the new message
	if msg.Content == "" {
		return errors.New("text content cannot be empty")
	}

	msg.Id = mockTextID

	nm := &messageProto.Message{
		Id:        msg.Id,
		ChatId:    msg.ChatId,
		SenderId:  msg.SenderId,
		Content:   msg.Content,
		TimeStamp: msg.TimeStamp,
	}

	if err := h.cache.Save(nm); err != nil {
		log.Printf("error saving to cache: %v", err)
	}

	return nil
}
