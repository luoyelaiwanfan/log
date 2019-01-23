package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:emc123123:@tcp(47.104.66.33:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success")
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil{
		log.Fatalln(err)
	}

}
