package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/DataQuery"
	"starWebserver/Proxy"
)



func main() {
	api := gin.Default()
	api.GET("/constellation/getAll",Proxy.Log, DataQuery.Star)
	api.StaticFS("/web",http.Dir("./static"))
	api.Any("/",func(c *gin.Context) {
	       c.Request.URL.Path = "/web"
		   api.HandleContext(c)
	    })
	fmt.Println("http://localhost:8080/")
	api.Run(":8080")
}


