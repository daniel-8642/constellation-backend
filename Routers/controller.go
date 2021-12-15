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
	api.POST("/user/:name/:upass/:session/:uauth",
		Proxy.SessionTimestamp, Proxy.ConfigAuth(1), Command.AdminAdduser)
	//用户更改密码接口
	api.PUT("/user/:session/:oldupass/:newupass",
		Proxy.SessionAuth, Proxy.SessionTimestamp, Command.Setuserpass)
	//删除用户接口
	api.DELETE("/user/:session/:name/:upass",
		Proxy.SessionAuth, Proxy.SessionTimestamp, Command.Deluser)
	//获取用户权限接口
	api.GET("/user/auth/:session", Proxy.
		SessionAuth, Proxy.SessionTimestamp, Command.Getuserauth)
}
func setUpData(api *gin.Engine) {
	//查询近期访问历史接口
	api.GET("/data/querycount/:session",
		Proxy.SessionAuth, Proxy.ConfigAuth(50), Proxy.SessionTimestamp, Command.Querycount)
	//查询各个星座人数接口
	api.GET("/data/starcount/:session",
		Proxy.SessionAuth, Proxy.ConfigAuth(50), Proxy.SessionTimestamp, Command.Starcount)
}
