package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Config"
	"starWebserver/DataQuery"
	"starWebserver/Proxy"
)

func main() {
	api := gin.Default()
	Web := Config.GetWeb()
	api.GET("/constellation/getAll", Proxy.Log, DataQuery.Star)
	api.StaticFS(Web.BaseUrl, http.Dir(Web.Static))
	api.Any("/", func(c *gin.Context) {
		c.Request.URL.Path = Web.BaseUrl
		api.HandleContext(c)
	})
	fmt.Println("http://localhost:" + Web.Port + "/")
	_ = api.Run(":" + Web.Port)
}
