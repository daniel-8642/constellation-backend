package Proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
	"time"
)

func SessionAuth(c *gin.Context) {
	session := c.Param("session")
	sqlStr := "select lasttime from session " +
		"where session= ? limit 1"
	Row := Global.DB.QueryRow(sqlStr, session)
	var lasttime string
	err := Row.Scan(&lasttime)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", lasttime, time.Local)
	day := (time.Now().Unix() - t.Unix()) / (3600 * 24) //天
	if day > 7 {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录已过期"})
		return
	}
	c.Next()
}

func ConfigAuth(authNum int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.Param("session")
		sqlStr := "select uauth from user where uid =(" +
			"select uid from session where session= ? limit 1) limit 1;"
		Row := Global.DB.QueryRow(sqlStr, session)
		var auth int
		err := Row.Scan(&auth)
		if err != nil {
			fmt.Println(err)
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			return
		}
		if auth > authNum {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "权限不足"})
			return
		}
		c.Next()
	}
}
