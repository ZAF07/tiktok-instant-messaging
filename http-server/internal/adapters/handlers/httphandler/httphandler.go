package httphandler

import (
	"log"
	"net/http"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
	httpdomain "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/domain/http_domain"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/* Driver Adapter (Driven by clients outside the application)
This package implements the interface (port) for clients (other services or end users) to interact with our core services.

This exposes a method for end users to interact with our services

This can have different flavours, for example, HTTP, CLI, or RPC to mention a few...

*/

type HTTPHandler struct {
	service ports.IHTTPService
}

func NewHTTPHandler(s ports.IHTTPService) *HTTPHandler {
	return &HTTPHandler{
		service: s,
	}
}

func (h *HTTPHandler) Push(c *gin.Context) {
	msg := httpdomain.Message{
		Text:   "Test message",
		Sender: "Test Sender",
	}

	res := h.service.Push(msg)
	c.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}
func (h *HTTPHandler) Pull(c *gin.Context) {
	res := h.service.Pull()
	c.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}

func (h *HTTPHandler) InitRoutes(r *gin.Engine) {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error loading application config file: %v", err)
	}
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     config.GetCorsAllowOrigins(),
		AllowHeaders:     config.GetHTTPAllowMethods(),
		// AllowAllOrigins:  true,
		// AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		// AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	messages := r.Group("messages")
	{
		messages.GET("/push", h.Push)
		messages.GET("/pull", h.Pull)
	}
}
