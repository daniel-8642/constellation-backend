package Global

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB = initDataBase()

func initDataBase() sql.DB {
	Mysql := GetMysql()
	dataSourceName := Mysql.User + ":" + Mysql.Password + "@(" + Mysql.Ip + ":" + Mysql.Port + ")/" + Mysql.Database
	//开启时检查数据库连接
	fmt.Println(dataSourceName)
	db, _ := sql.Open(Mysql.DriverName, dataSourceName)
	////数据库连接
	err := db.Ping()
	if err != nil {
		panic("数据库链接失败")
	}
	return *db
}
