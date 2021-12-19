package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
	"starWebserver/Routers"
)

func main() {
	//开启api服务
	api := gin.Default()
	Routers.SetUpRouter(api)
	//服务静态网页
	Web := Global.GetWeb()
	api.StaticFS(Web.BaseUrl, http.Dir("./static/web"))
	api.StaticFS("/backend", http.Dir("./static/backend"))
	//转发根目录
	api.Any("/", func(c *gin.Context) {
		c.Request.URL.Path = Web.BaseUrl
		api.HandleContext(c)
	})

	//打印结果开始服务
	fmt.Println("http://localhost:" + Web.Port + "/")
	_ = api.Run(":" + Web.Port)
}
