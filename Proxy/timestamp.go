package Proxy

import (
	"crypto/md5"
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
	stamp, _ := strconv.ParseInt(c.Request.Header.Get("timestamp"), 10, 64)
	now := time.Now().UnixMilli()
	if now-stamp > 15000 || now-stamp < -10000 {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "签名已失效"})
		// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
		return
	}
	rand := c.Request.Header.Get("rand")
	sign := c.Request.Header.Get("sign")

	md5session := fmt.Sprintf("%x", md5.Sum([]byte(session)))
	sum := sha256.Sum256([]byte(md5session + strconv.FormatInt(stamp, 10) + Global.GetWeb().Key + rand))
	sums := fmt.Sprintf("%x", sum)
	if strings.Compare(sums, sign) != 0 {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "签名不匹配"})
		return
	}
	c.Next()

}
