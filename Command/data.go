package Command

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
)

func Querycount(c *gin.Context) {
	sql := "select time ,count(*) from starLog " +
		"group by time order by time DESC limit 14;"
	Rows, err := Global.DB.Query(sql)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	var ret []querycountVo
	for Rows.Next() {
		var date, count string
		err := Rows.Scan(&date, &count)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(date) != 0 {
			ret = append(ret, querycountVo{
				Date:  date,
				Count: count,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": ret}) //输出
}

type querycountVo struct {
	Date  string `json:"date"`
	Count string `json:"count"`
}

func Starcount(c *gin.Context) {
	sql := "select consName, count(*) from " + Global.GetMysql().Database + ".starLog group by consName order by consName limit 10"
	Rows, err := Global.DB.Query(sql)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	var ret []starcountVo
	for Rows.Next() {
		var name, count string
		err := Rows.Scan(&name, &count)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(name) != 0 {
			ret = append(ret, starcountVo{
				Name:  name,
				Count: count,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": ret}) //输出
}

type starcountVo struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}
