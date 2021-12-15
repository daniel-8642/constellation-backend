package Command

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"starWebserver/Global"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	name := c.Param("name")
	upass := c.Param("upass")
	sqlStr := "select uid from user where uname = ? and upass = ? limit 1"
	Row := Global.DB.QueryRow(sqlStr, name, upass)
	var uid string
	err := Row.Scan(&uid)
	if err != nil {
		//todo:编写错误信息
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	Now := time.Now().Format("2006-01-02 15:04:05")
	session := strconv.FormatUint(rand.Uint64(), 16)
	sqlStr = "insert into session (uid,session,lasttime) " +
		"values ( ? , ? , ? ) " +
		"on duplicate key update session= ? ,lasttime = ? ;"
	_, err = Global.DB.Exec(sqlStr, uid, session, Now, session, Now)
	if err != nil {
		//todo:编写错误信息
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s", session)) //输出
}

//func Adduser(c *gin.Context) {
//	name:=c.Param("name")
//	upass:=c.Param("upass")
//	uauth:=c.Param("uauth")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//
//func AdminAdduser(c *gin.Context) {
//	name:=c.Param("name")
//	upass:=c.Param("upass")
//	uauth:=c.Param("uauth")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//
//func Setuserpass(c *gin.Context) {
//	oldupass:=c.Param("oldupass")
//	newupass:=c.Param("newupass")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//func Deluser(c *gin.Context) {
//	name:=c.Param("name")
//	upass:=c.Param("upass")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//func Getuserauth(c *gin.Context) {
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
