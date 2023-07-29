package httpdomain

import "github.com/go-playground/validator/v10"

type Message struct {
	Id        string `json:"id,omitempty"`
	ChatId    string `json:"chatId,omitempty" validate:"required"`
	SenderId  string `json:"senderId,omitempty" validate:"required"`
	Content   string `json:"content,omitempty" validate:"required"`
	TimeStamp int32  `json:"timeStamp,omitempty"`
}

func (m *Message) Validate() (err error) {
	validate := validator.New()
	err = validate.Struct(m)
	return
}

type MessagesDTO struct {
	Messeges []Message
}
