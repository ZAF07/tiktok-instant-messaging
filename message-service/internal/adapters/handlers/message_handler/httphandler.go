package messagehandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	message "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message_dto"
	"github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/ports"
	"github.com/go-playground/validator/v10"

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
	validate := validator.New()
	// Receive msg, do validation
	newMsg := &message.Message{}
	if bErr := c.ShouldBindJSON(newMsg); bErr != nil {
		log.Fatalf("error binding request body to struct: %v", bErr)
	}

	if err := validate.Struct(newMsg); err != nil {
		responseHelper.ReturnError(c, "Missing required fields for the new message", err.Error())
		return
	}

	newMsg.TimeStamp = int32(time.Now().Unix())

	err := h.service.Push(newMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "failed to save message",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Sent message ✅",
	})
}
func (h *HTTPHandler) Pull(c *gin.Context) {
	// PROFILING ------------------

	// Start profiling
	f, err := os.Create("performance.prof")
	if err != nil {

		fmt.Println(err)
		return

	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
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
