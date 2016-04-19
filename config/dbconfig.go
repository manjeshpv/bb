package dbconfig

import (

)
import (
	"os"
	"log"
	"gopkg.in/gorp.v1"
	"database/sql"
	"github.com/manjeshpv/bb/api/user/model"
	"fmt"
	"github.com/manjeshpv/bb/api/chain/model"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func DBUrl() string {
	dburl := os.Getenv("BB_SQL")
	fmt.Println(dburl)
	if dburl == "" {
		dburl = "ayyayo_beatle:beatle1234@tcp(188.166.210.246:3306)/ayyayo_beatle"
	}

	return dburl
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", DBUrl())
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(usermodel.User{}, "user").SetKeys(true, "Id")
	dbmap.AddTableWithName(chainmodel.Chain{}, "chain").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func Init() *gorp.DbMap {
	var dbmap = initDb()
	return dbmap
}