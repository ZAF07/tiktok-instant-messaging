package httpdomain

type Message struct {
	Id        string `json:"id,omitempty"`
	ChatId    string `json:"chatId,omitempty" validate:"required"`
	SenderId  string `json:"senderId,omitempty" validate:"required"`
	Content   string `json:"content,omitempty" validate:"required"`
	TimeStamp int32  `json:"timeStamp,omitempty"`
}
