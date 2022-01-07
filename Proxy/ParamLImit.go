package Proxy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParamMinLenLImit(lenth int, param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param(param)
		if len(param) < lenth {
			c.Abort()
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "参数长度过短"})
			return
		}
		c.Next()
	}
}
