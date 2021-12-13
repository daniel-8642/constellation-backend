package Proxy

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"starWebserver/Config"
	"time"
)

var dataSourceName = "root:taotao@(127.0.0.1:3306)/starWeb"

func init() {

	//开启时检查数据库连接
	db, _ := sql.Open("mysql", dataSourceName)
	//数据库连接
	err := db.Ping()
	if err != nil {
		panic("数据库链接失败")
	}
	defer db.Close()

	////多行查询
	//rows, _ := db.Query("select * from starLog")
	//var id string
	//for rows.Next() {
	//	rows.Scan(&id)
	//	fmt.Println(id)
	//}
}

func Log(c *gin.Context) {
	//请求之前
	consName := c.Query("consName")
	ip := getRequestIP(c)
	times := time.Now().Format("2006-01-02 15:04:05")
	go func(conName string, ip string, times string) {
		fmt.Print(" consName: " + conName)
		fmt.Print(" ip: " + ip)
		fmt.Println(" time: " + times)

		db, _ := sql.Open("mysql", "root:taotao@(127.0.0.1:3306)/starWeb")
		defer db.Close()
		sqlStr := "insert into starLog(consName, ip,time) values (?,?,?)"
		/*ret*/
		_, err := db.Exec(sqlStr, consName, ip, times)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
		//theID, err := ret.LastInsertId() // 新插入数据的id
		//if err != nil {
		//	fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		//	return
		//}
		//fmt.Printf("insert success, the id is %d.\n", theID)
	}(consName, ip, times)
	c.Next()

}

func getRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
