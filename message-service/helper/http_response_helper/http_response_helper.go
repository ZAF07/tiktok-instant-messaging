package httpresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, msg, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":   msg,
		"error": err,
	})

}
