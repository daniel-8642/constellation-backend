package Routers

import (
	"github.com/gin-gonic/gin"
	"starWebserver/Command"
	"starWebserver/Proxy"
)

func SetUpRouter(api *gin.Engine) {
	setUpStar(api)
	setUpUser(api)
	setUpData(api)
}

func setUpStar(api *gin.Engine) {
	//公共查询星座运势接口
	api.GET("/constellation/getAll", Proxy.Log, Command.Star)
}
func setUpUser(api *gin.Engine) {
	//登录换取session接口
	api.GET("/user/login/:name/:upass", Command.Login)
	//注册用户接口
	api.POST("/user/:name/:upass", Command.Adduser)
	//管理员注册用户接口
	//api.POST("/user/:session/:name/:upass/:uauth",Command.AdminAdduser)
	//用户更改密码接口
	api.PUT("/user/:session/:oldupass/:newupass", Command.Setuserpass)
	//删除用户接口
	api.DELETE("/user/:session/:name/:upass", Command.Deluser)
	//获取用户权限接口
	api.GET("/user/auth/:session", Command.Getuserauth)
}
func setUpData(api *gin.Engine) {
	//查询近期访问历史接口
	api.GET("/data/querycount/:session", Command.Querycount)
	//查询各个星座人数接口
	api.GET("/data/starcount/:session", Proxy.SessionAuth, Proxy.SessionTimestamp, Command.Starcount)
}
