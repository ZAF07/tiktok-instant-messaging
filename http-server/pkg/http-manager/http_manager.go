package httpmanager

import (
	"log"
	"net/http"
	"time"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
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

func (h *HTTPManager) StartServer() {
	if err := h.server.ListenAndServe(); err != nil {
		log.Fatalf("error starting HTTP server. error msg: %v", err)
	}
}
