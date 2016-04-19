package dbconfig

import (

)
import (
	"os"
	"log"
	"gopkg.in/gorp.v1"
	"database/sql"
	"fmt"
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

	// Will log all SQL statements + args as they are run
	// The first arg is a string prefix to prepend to all log messages
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	// Turn off tracing
	// dbmap.TraceOff()

	return dbmap
}

func Init() *gorp.DbMap {
	var dbmap = initDb()
	return dbmap
}