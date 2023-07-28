package messagehandler

import (
	"fmt"
	"log"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) InitRoutes(r *gin.Engine) {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error loading application config file: %v", err)
	}
	fmt.Println("CORS:: ", config.GetCorsAllowOrigins())
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
