package httpmanager

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ZAF07/tiktok-instant-messaging/message-service/config"
	"github.com/gin-gonic/gin"
)

/*
	This package initialises a HTTP server
	Can be used by any package which needs it
k
*/

type HTTPManager struct {
	server *http.Server
}

// cmux -> httpserver(includes the gin engine) -> routers -> handlers
func NewHTTPServer() *HTTPManager {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error getting config. error msg:  %v", err)
	}

	s := &http.Server{
		ReadTimeout:  time.Duration(c.GetReadTimeoutHTTP()) * time.Second,
		WriteTimeout: time.Duration(c.GetWriteTimeoutHTTP()) * time.Second,
		Handler:      gin.Default(),
		Addr:         c.GetPortHTTP(),
	}

	return &HTTPManager{
		server: s,
	}
}

func (h *HTTPManager) GetServer() *http.Server {
	return h.server
}

func (h *HTTPManager) GetHandler() *gin.Engine {
	return h.server.Handler.(*gin.Engine)
}

func (h *HTTPManager) StartHTTPServer(l net.Listener) {

	if err := h.server.Serve(l); err != nil {
		log.Fatalf("error establishing http server")
	}

}
