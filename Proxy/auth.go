package Proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
)

func SessionAuth(c *gin.Context) {
	session := c.Param("session")
	sqlStr := "select lasttime from session where session= '" + session + "' limit 1"
	Row := Global.DB.QueryRow(sqlStr)
	var lasttime string
	err := Row.Scan(&lasttime)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	fmt.Println(lasttime)
	c.Next()

}
func AccountAuth(c *gin.Context) {
	c.Param("name")
	c.Param("upass")
	//// 验证不通过，不再调用后续的函数处理
	//c.Abort()
	//c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
	//// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
	//return
	c.Next()

}
