package Command

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starWebserver/Global"
)

//func Querycount(c *gin.Context) {
//	c.String(http.StatusOK, fmt.Sprintf("%s", body)) //输出
//}

func Starcount(c *gin.Context) {
	sql := "select consName, count(*) from " + Global.GetMysql().Database + ".starLog group by consName order by consName"
	Rows, err := Global.DB.Query(sql)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		return
	}
	ret := map[string]string{}
	for Rows.Next() {
		var name, count string
		err := Rows.Scan(&name, &count)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(name) != 0 {
			ret[name] = count
		}
	}
	marshal, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("%s", marshal)) //输出
}
