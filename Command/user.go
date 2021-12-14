package Command

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func Login(c *gin.Context) {
//	name:=c.Param("name")
//	upass:=c.Param("upass")
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}
//
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
