package httpresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestErr(c *gin.Context, msg, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":   msg,
		"error": err,
	})
}

func InternalErr(c *gin.Context, msg, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		msg:     msg,
		"error": err,
	})
}
