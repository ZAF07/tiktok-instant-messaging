package router

import (
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/handlers/httphandler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h httphandler.HTTPHandler) *gin.Engine {
	// c, err := config.GetConfig()
	// if err != nil {
	// 	log.Fatalf("error getting config. error msg: %v", err)
	// }

	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins:  true,
		// ❌ Might want to only allow specific host for security
		// AllowOrigins: c.GetCorsAllowOrigins(),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	push := r.Group("push")
	{
		push.GET("", h.Push)
	}

	pull := r.Group("pull")
	{
		pull.GET("", h.Pull)
	}

	return r
}
