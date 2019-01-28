package main

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	//只能连接存在的数据库
	db, err := sql.Open("mysql", "root:123456@tcp(47.104.66.33:3306)/test?charset=utf8")
	//db, err := sql.Open("mysql", "root:@tcp(47.104.66.33:3306)/test?charset=utf8")
	if err != nil {
		beego.Info(err)
	}

	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()


	// 创建数据库
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil{
		beego.Info(err)
	}


	rs, err := db.Exec("INSERT INTO test.hello(world) VALUES ('hello world')")
	if err != nil{
		beego.Info(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil{
		beego.Info(err)
	}
	beego.Info("inserted rows", rowCount)


	rows, err := db.Query("SELECT world FROM test.hello")
	if err != nil{
		beego.Info(err)
	}

	for rows.Next(){
		var s string
		err = rows.Scan(&s)
		if err !=nil{
			beego.Info(err)
		}
		beego.Info("found row containing ", s)
	}
	rows.Close()


}
