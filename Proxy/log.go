package Proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"starWebserver/Global"
	"time"
)

func Log(c *gin.Context) {
	//请求之前
	consName := c.Query("consName")
	ip := getRequestIP(c)
	times := time.Now().Format("2006-01-02")
	go func(conName string, ip string, times string) {
		sqlStr := "insert into starLog(consName, ip,time) values (?,?,?)"
		_, err := Global.DB.Exec(sqlStr, consName, ip, times)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}(consName, ip, times)
	c.Next()

}

func getRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
