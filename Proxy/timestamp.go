package Proxy

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
	"strconv"
	"strings"
	"time"
)

func SessionTimestamp(c *gin.Context) {
	session := c.Param("session")
	Hsession := c.Request.Header.Get("session")
	if session != Hsession {
		// 验证不通过，不再调用后续的函数处理
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "参数不匹配"})
		// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
		return
	}
	stamp, _ := strconv.ParseInt(c.Request.Header.Get("timestamp"), 10, 64)
	now := time.Now().UnixMilli()
	if now-stamp > 3000 || now-stamp < -300 {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "签名已失效"})
		// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
		return
	}

	rand := c.Request.Header.Get("rand")
	sign := c.Request.Header.Get("sign")

	sum := sha256.Sum256([]byte(session + strconv.FormatInt(stamp, 10) + Global.GetWeb().Key + rand))
	sums := fmt.Sprintf("%x", sum)
	if strings.Compare(sums, sign) != 0 {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "签名不匹配"})
		return
	}
	c.Next()

}

func AccountTimestamp(c *gin.Context) {
	//session := c.Param("session")
	//Hsession := c.Request.Header.Get("session")
	//if session != Hsession {
	//	// 验证不通过，不再调用后续的函数处理
	//	c.Abort()
	//	c.JSON(http.StatusUnauthorized,gin.H{"message":"参数不匹配"})
	//	// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
	//	return
	//}
	//stamp, _ := strconv.ParseInt(c.Request.Header.Get("timestamp"),10,64)
	//now:= time.Now().UnixMilli()
	//if now-stamp>3000 || now-stamp< -300 {
	//	c.Abort()
	//	c.JSON(http.StatusUnauthorized,gin.H{"message":"签名已失效"})
	//	// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
	//	return
	//}
	//
	//rand := c.Request.Header.Get("rand")
	//sign := c.Request.Header.Get("sign")
	//
	//sum := sha256.Sum256([]byte(session + strconv.FormatInt(stamp,10) + Global.GetWeb().Key + rand))
	//sums:= fmt.Sprintf("%x", sum)
	//if strings.Compare(sums,sign)!=0 {
	//	c.Abort()
	//	c.JSON(http.StatusUnauthorized,gin.H{"message":"签名不匹配"})
	//	return
	//}
	//c.Next()

}
