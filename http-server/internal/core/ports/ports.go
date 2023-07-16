package ports

// Interfaces which Adapters will imolement in order to interact with the core

type IHTTPService interface {
	Push() string
	Pull() string
}

type ICacheStore interface {
	Save()
}
