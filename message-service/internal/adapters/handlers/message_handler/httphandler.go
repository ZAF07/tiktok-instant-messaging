package messagehandler

import (
	"log"
	"net/http"

	message "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message_dto"
	"github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/ports"

	// responseHelper "github.com/ZAF07/tiktok-instant-messaging/messaeg-service/helper/http_response_helper"
	responseHelper "github.com/ZAF07/tiktok-instant-messaging/message-service/helper/http_response_helper"
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

// When a user pushes a new message, i want to store it in the cache first (Write-through strategy)
func (h *HTTPHandler) Push(c *gin.Context) {
	newMsg := &message.Message{}
	if bindErr := c.ShouldBindJSON(newMsg); bindErr != nil {
		log.Fatalf("error binding request body to struct: %v", bindErr)
		responseHelper.InternalErr(c, "Internal error. Try again", bindErr.Error())
	}

	if err := newMsg.Validate(); err != nil {
		responseHelper.BadRequestErr(c, "Missing required fields for the new message", err.Error())
	}

	if err := h.service.Push(newMsg); err != nil {
		responseHelper.InternalErr(c, "failed to save message", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Sent message ✅",
	})
}
func (h *HTTPHandler) Pull(c *gin.Context) {
	// PROFILING ------------------

	// Start profiling
	// f, err := os.Create("performance.prof")
	// if err != nil {

	// 	fmt.Println(err)
	// 	return

	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// profiling ------------------

	res, err := h.service.Pull()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "failed to get key",
			"error:": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}
