package service

import (
	"log"
	"time"

	messageProto "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message-service/proto"
	message "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message_dto"
	"github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/ports"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type HTTPService struct {
	cache ports.ICacheStore
}

func NewHTTPService(c ports.ICacheStore) *HTTPService {
	return &HTTPService{
		cache: c,
	}
}

func (h *HTTPService) Pull() (*messageProto.Messages, error) {
	// Get key:value from cache
	res, err := h.cache.Get("test-key")
	if err != nil {
		return nil, err
	}

	// Unmarshal into Proto
	result := &messageProto.Messages{}
	uErr := proto.Unmarshal(res, result)
	if uErr != nil {
		log.Fatalf("error unmarshalling into struct @ service.Pull: %v", uErr)
	}

	// Convert into DTO
	// foundMsg := &message.Message{
	// 	Id:        result.Id,
	// 	ChatId:    result.ChatId,
	// 	SenderId:  result.SenderId,
	// 	Content:   result.Content,
	// 	TimeStamp: result.TimeStamp,
	// }
	return result, nil
}

func (h *HTTPService) Push(msg *message.Message) error {
	textId := uuid.New().String()

	// Validate the new message
	if validateErr := msg.Validate(); validateErr != nil {
		return validateErr
	}

	// Set the fields to the new text that has to be generated on the server
	msg.TimeStamp = int32(time.Now().Unix()) // timestamp should be created on the client side actually
	msg.Id = textId

	nm := &messageProto.Message{
		Id:        msg.Id,
		ChatId:    msg.ChatId,
		SenderId:  msg.SenderId,
		Content:   msg.Content,
		TimeStamp: msg.TimeStamp,
	}

	if err := h.cache.Save(nm); err != nil {
		log.Printf("error saving to cache: %v", err)
		return err
	}

	return nil
}
