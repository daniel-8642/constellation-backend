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
	api.GET("/constellation/getAll", Proxy.Log, Command.Star)
}
func setUpUser(api *gin.Engine) {
	api.GET("/user/login/:name/:upass", Command.Login)
	//api.POST("/user/:name/upass/:uauth",Command.Adduser)
	//api.POST("/user/:session/:name/:upass/:uauth",Command.AdminAdduser)
	api.PUT("/user/:session/:oldupass/:newupass", Command.Setuserpass)
	//api.DELETE("/user/:session/:name/:upass",Command.Deluser)
	//api.GET("/user/auth/:session",Command.Getuserauth)
}
func setUpData(api *gin.Engine) {
	//api.GET("/data/querycount/:session",Command.Querycount)
	api.GET("/data/starcount/:session", Proxy.SessionAuth, Proxy.SessionTimestamp, Command.Starcount)
}
