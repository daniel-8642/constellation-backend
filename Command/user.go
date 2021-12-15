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
	c.JSON(http.StatusOK, gin.H{"session": session})
}

func Adduser(c *gin.Context) {
	name := c.Param("name")
	upass := c.Param("upass")
	sql := "insert into user (uname, upass, uauth) " +
		"values ( ? , ? ,50);"
	result, err := Global.DB.Exec(sql, name, upass)
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
}

func AdminAdduser(c *gin.Context) {
	name := c.Param("name")
	upass := c.Param("upass")
	uauth := c.Param("uauth")
	sql := "insert into user (uname, upass, uauth) " +
		"values ( ? , ? ,?);"
	result, err := Global.DB.Exec(sql, name, upass, uauth)
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
}

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
	delSessionForSession(session)
}

func Deluser(c *gin.Context) {
	session := c.Param("session")
	name := c.Param("name")
	upass := c.Param("upass")
	//更新账户数据
	sql := "delete from user where uid = (" +
		"select uid from session where session = ? limit 1" +
		") and uname= ? and upass = ? limit 1;"
	result, err := Global.DB.Exec(sql, session, name, upass)
	if err != nil {
		fmt.Println("DataBase err ,err:", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v \n", err)
		return
	}

	if rowsaffected != 0 {
		c.String(http.StatusOK, fmt.Sprintf("{%s}", "ok")) //输出
	} else {
		c.String(http.StatusBadGateway, fmt.Sprintf("{%s}", "error")) //输出
	}
	//清空session
	delSessionForSession(session)
}

func Getuserauth(c *gin.Context) {
	session := c.Param("session")
	sql := "select uauth from user where uid = (" +
		"select uid from session where session = ? limit 1" +
		") limit 1;"
	Row := Global.DB.QueryRow(sql, session)
	var auth int
	err := Row.Scan(&auth)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "内部错误"})
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%d", auth)) //输出
}

func delSessionForSession(session string) bool {
	//更新账户数据
	sql := "delete from session " +
		"where session = ? limit 1;"
	result, err := Global.DB.Exec(sql, session)
	if err != nil {
		fmt.Println("DataBase err ,err:", err)
		return false
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v \n", err)
		return false
	}
	return rowsaffected != 0
}
