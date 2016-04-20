package main

import (
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manjeshpv/bb/config"
	"github.com/manjeshpv/bb/api/user/model"
	"github.com/manjeshpv/bb/api/chain/model"
	"github.com/manjeshpv/bb/api/hotel/model"
	//"github.com/manjeshpv/bb/api/department/model"
	"github.com/manjeshpv/bb/routes"
)

func main() {

	r := gin.Default()

	createTables()

	//createTablesGORM()

	r.Use(Cors())

	routes.Init(r)

	r.Run(":8001")
}


func createTablesGORM(){
	//db := dbconfig.InitGORM()
	//db.CreateTable(&departmentmodel.Department{})
	//.Table("Product").
	//db.CreateTable(&Product{})
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}