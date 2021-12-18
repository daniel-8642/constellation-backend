package Proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
	"strings"
	"time"
)

func Log(c *gin.Context) {
	//请求之前
	consName := c.Query("consName")
	if len(consName) != 9 || !strings.HasSuffix(consName, "座") {
		c.Abort()
		c.JSON(http.StatusAccepted, gin.H{"message": "这是不存在的星座名称"})
		return
	}
	ip := getRequestIP(c)
	times := time.Now().Format("2006-01-02")
	c.Next()
	value, exists := c.Get("log")
	var b = value.(bool)
	if exists && b {
		go func(conName string, ip string, times string) {
			fmt.Printf("log")
			sqlStr := "insert into starLog(consName, ip,time) values (?,?,?)"
			_, err := Global.DB.Exec(sqlStr, consName, ip, times)
			if err != nil {
				fmt.Printf("insert failed, err:%v\n", err)
				return
			}
		}(consName, ip, times)
	}
}

func getRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
