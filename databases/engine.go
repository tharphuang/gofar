package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TharpHuang/gofar/migration"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DBContainer struct {
	engine *xorm.Engine
	tables []interface{}
}

var dbContainer = DBContainer{}

type configuration struct {
	MysqlUser     string
	MysqlPassword string
	MysqlPath     string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%v",conf)
	//mysqlSource := fmt.Sprintf("%s:%s@%s", conf.MysqlUser, conf.MysqlPassword, conf.MysqlPath)

	engine, err := xorm.NewEngine("mysql","docker:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	dbContainer.engine = engine
	dbContainer.tables = append(dbContainer.tables, &migration.User{})

	engine.Sync2(dbContainer.tables[0])
	if err != nil {
		panic(err)
	}
	fmt.Print(engine.DBMetas())
}
