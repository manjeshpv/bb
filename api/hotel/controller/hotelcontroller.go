package hotelcontroller

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manjeshpv/bb/api/hotel/model"
	"github.com/manjeshpv/bb/config"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetHotels(c *gin.Context) {
	var hotels []hotelmodel.Hotel
	dbmap := dbconfig.Init()
	_, err := dbmap.Select(&hotels, "SELECT * FROM hotel")

	if err == nil {
		c.JSON(200, hotels)
	} else {
		c.JSON(404, gin.H{"error": "no hotel(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/hotels
}

func GetHotel(c *gin.Context) {
	id := c.Params.ByName("id")
	var hotel hotelmodel.Hotel
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&hotel, "SELECT * FROM hotel WHERE id=? LIMIT 1", id)

	if err == nil {
		hotel_id, _ := strconv.ParseInt(id, 0, 64)

		content := &hotelmodel.Hotel{
			Id:        hotel_id,
			Firstname: hotel.Firstname,
			Lastname:  hotel.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "hotel not found"})
	}

	// curl -i http://localhost:8080/api/v1/hotels/1
}

func PostHotel(c *gin.Context) {
	var hotel hotelmodel.Hotel
	c.Bind(&hotel)

	log.Println(hotel)
	dbmap := dbconfig.Init()
	if hotel.Firstname != "" && hotel.Lastname != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO hotel (firstname, lastname) VALUES (?, ?)`, hotel.Firstname, hotel.Lastname); insert != nil {
			hotel_id, err := insert.LastInsertId()
			if err == nil {
				content := &hotelmodel.Hotel{
					Id:        hotel_id,
					Firstname: hotel.Firstname,
					Lastname:  hotel.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/hotels
}

func UpdateHotel(c *gin.Context) {
	id := c.Params.ByName("id")
	var hotel hotelmodel.Hotel
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&hotel, "SELECT * FROM hotel WHERE id=?", id)

	if err == nil {
		var json hotelmodel.Hotel
		c.Bind(&json)

		hotel_id, _ := strconv.ParseInt(id, 0, 64)

		hotel := hotelmodel.Hotel{
			Id:        hotel_id,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}

		if hotel.Firstname != "" && hotel.Lastname != "" {
			_, err = dbmap.Update(&hotel)

			if err == nil {
				c.JSON(200, hotel)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "hotel not found"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/hotels/1
}

func DeleteHotel(c *gin.Context) {
	id := c.Params.ByName("id")

	var hotel hotelmodel.Hotel
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&hotel, "SELECT * FROM hotel WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&hotel)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "hotel not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/hotels/1
}

func OptionsHotel(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
