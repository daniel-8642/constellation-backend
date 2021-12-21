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
	api.StaticFS(Web.WebUrl, http.Dir(Web.StaticWeb))
	api.StaticFS(Web.BackendUrl, http.Dir(Web.StaticBackend))
	//转发根目录
	api.Any("/", func(c *gin.Context) {
		c.Request.URL.Path = Web.WebUrl
		api.HandleContext(c)
	})

	//打印结果开始服务
	fmt.Println("http://localhost:" + Web.Port + Web.WebUrl)
	fmt.Println("http://localhost:" + Web.Port + Web.BackendUrl)
	_ = api.Run(":" + Web.Port)
}
