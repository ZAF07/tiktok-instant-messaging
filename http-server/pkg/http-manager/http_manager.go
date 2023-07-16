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
		ReadTimeout:  time.Duration(c.GetReadTimeoutHTTP()),
		WriteTimeout: time.Duration(c.GetWriteTimeoutHTTP()),
		Handler:      gin.Default(),
		Addr:         c.GetPortHTTP(),
	}

	return &HTTPManager{
		server: s,
	}
}

func (h *HTTPManager) GetServer() *http.Server {
	h.server.ListenAndServe()
	return h.server
}
