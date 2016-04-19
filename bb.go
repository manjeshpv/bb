package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/manjeshpv/bb/routes"
	"github.com/manjeshpv/bb/config"
	"github.com/manjeshpv/bb/api/user/model"
	"github.com/manjeshpv/bb/api/chain/model"
	"github.com/manjeshpv/bb/api/hotel/model"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	createTables()

	r.Use(Cors())

	routes.Init(r)

	r.Run(":8001")
}

func createTables(){
	dbmap := dbconfig.Init()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(usermodel.User{}, "user").SetKeys(true, "Id")
	dbmap.AddTableWithName(chainmodel.Chain{}, "chain").SetKeys(true, "Id")
	dbmap.AddTableWithName(hotelmodel.Hotel{}, "hotel").SetKeys(true, "Id")
	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}