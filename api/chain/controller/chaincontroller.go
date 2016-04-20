package chaincontroller

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manjeshpv/bb/api/chain/model"
	"github.com/manjeshpv/bb/config"
	"os"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetChains(c *gin.Context) {
	var chains []chainmodel.Chain
	dbmap := dbconfig.Init()
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	_, err := dbmap.Select(&chains, "SELECT * FROM chain")

	if err == nil {
		c.JSON(200, chains)
	} else {
		//log.Fatal(err)
		c.JSON(404, gin.H{"error": "no chain(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/chains
}

func GetChain(c *gin.Context) {
	id := c.Params.ByName("id")
	var chain chainmodel.Chain
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&chain, "SELECT * FROM chain WHERE id=? LIMIT 1", id)

	if err == nil {
		chain_id, _ := strconv.ParseInt(id, 0, 64)

		content := &chainmodel.Chain{
			Id:        chain_id,
			Name: chain.Name,
			HotelId: chain.HotelId,

		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "chain not found"})
	}

	// curl -i http://localhost:8080/api/v1/chains/1
}

func PostChain(c *gin.Context) {
	var chain chainmodel.Chain
	c.Bind(&chain)

	log.Println(chain)
	dbmap := dbconfig.Init()
	if chain.Name != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO chain (name,hotel_id) VALUES (?,?)`, chain.Name, chain.HotelId); insert != nil {
			chain_id, err := insert.LastInsertId()
			if err == nil {
				content := &chainmodel.Chain{
					Id:        chain_id,
					Name: chain.Name,
					HotelId: chain.HotelId,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/chains
}

func UpdateChain(c *gin.Context) {
	id := c.Params.ByName("id")
	var chain chainmodel.Chain
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&chain, "SELECT * FROM chain WHERE id=?", id)

	if err == nil {
		var json chainmodel.Chain
		c.Bind(&json)

		chain_id, _ := strconv.ParseInt(id, 0, 64)

		chain := chainmodel.Chain{
			Id:        chain_id,
			Name: json.Name,
			HotelId: chain.HotelId,
		}

		if chain.Name != ""  {
			_, err = dbmap.Update(&chain)

			if err == nil {
				c.JSON(200, chain)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "chain not found"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/chains/1
}

func DeleteChain(c *gin.Context) {
	id := c.Params.ByName("id")

	var chain chainmodel.Chain
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&chain, "SELECT * FROM chain WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&chain)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "chain not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/chains/1
}

func OptionsChain(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
