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

//登录接口
func Login(c *gin.Context) {
	name := c.Param("name")
	upass := c.Param("upass")
	sqlStr := "select uid from user " +
		"where uname = ? and upass = ? limit 1"
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

//更改密码接口
func Setuserpass(c *gin.Context) {
	session := c.Param("session")
	oldupass := c.Param("oldupass")
	newupass := c.Param("newupass")
	//更新账户数据
	result, err := Global.DB.Exec("update user set upass = ? where uid= ("+
		" select uid from session where session = ? limit 1 "+
		") and upass= ? ;", newupass, session, oldupass)
	if err != nil {
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}

	if rowsaffected != 0 {
		c.String(http.StatusOK, fmt.Sprintf("{%s}", "ok")) //输出
	} else {
		c.String(http.StatusBadGateway, fmt.Sprintf("{%s}", "error")) //输出
	}
	//清空session

}

//func Deluser(c *gin.Context) {
//	name:=c.Param("name")
//	upass:=c.Param("upass")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//func Getuserauth(c *gin.Context) {
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
