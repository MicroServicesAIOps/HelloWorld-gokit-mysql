package main

import (
	"HelloWorld-gokit-mysql/api"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.Println("init")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

}

func connectMysql() *sql.DB {
	MysqlDb, err := sql.Open("mysql", "user:123456@tcp(127.0.0.1:3306)/user?charset=utf8")
	if err != nil {
		log.Printf("connect mysql database faied, error:[%v]", err.Error())
		return nil
	}
	MysqlDb.SetMaxOpenConns(10)
	MysqlDb.SetMaxIdleConns(5)

	if err != nil {
		log.Println("open mysql failed,", err)
	}
	return MysqlDb
}

func queryData(MysqlDb *sql.DB) {

	rows, err := MysqlDb.Query("select id,name ,age from student")
	if err != nil {
		log.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Println("get data failed:", err.Error())
		}
		log.Println(id, name, age)
	}
}

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	Db := connectMysql()
	defer Db.Close()
	queryData(Db)

	log.Println("This is a Demo")

	s := api.MyService{}

	e := api.MakeEndpoints(s)

	h := api.MakeHttpHandler(ctx, e)

	go func() {
		log.Println("Http Server start， Port:8080")
		errChan <- http.ListenAndServe(":8080", h)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(<-errChan)
}
