package departmentcontroller

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manjeshpv/bb/api/department/model"
	"github.com/manjeshpv/bb/config"

	//"fmt"
	"fmt"
)

//
//
//func testDb(){
//	db := dbconfig.InitGORM()
//	// Create
//	db.Create(&Product{Code: "L1212", Price: 1000})
//
//	var product Product
//	db.First(&product, 1) // find product with id 1
//	db.First(&product, "code = ?", "L1212") // find product with code l1212
//	fmt.Printf("%+v\n",product)
//
//	// Update - update product's price to 2000
//	db.Model(&product).Update("Price", 2000)
//
//	// Delete - delete product
//	db.Delete(&product)
//
//}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetDepartments(c *gin.Context) {
	db := dbconfig.InitGORM()
	var departments []departmentmodel.Department
	err := db.Find(&departments).GetErrors()
	fmt.Printf("departments '%v'", departments)
	if err == nil {

		//for i := range departments {
		//	fmt.Println(departments[i])
		//
		//}
		response := make([]departmentmodel.Department, 0, len(departments))
		for k := range departments {
			department := departmentmodel.Department{
				Id:departments[k].Id,
				Name:departments[k].Name,
			}
			fmt.Println(department)
			response = append(response, department)
		}
		c.JSON(200, response)
	} else {

		c.JSON(404, gin.H{"error": "no department(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/departments
}

func GetDepartment(c *gin.Context) {
	id := c.Params.ByName("id")
	var department departmentmodel.Department
	db := dbconfig.InitGORM()
	err := db.First(&department, id).GetErrors()

	if err == nil {
		//department_id, _ := strconv.ParseInt(id, 0, 64)
		//
		//content := &departmentmodel.Department{
		//	Id:        department_id,
		//	Name: department.Name,
		//
		//}
		//c.JSON(200, content)
		 c.JSON(200, department)
	} else {
		c.JSON(404, gin.H{"error": "department not found"})
	}

	// curl -i http://localhost:8080/api/v1/departments/1
}

func PostDepartment(c *gin.Context) {
	var department departmentmodel.Department
	c.Bind(&department)

	log.Println(department)
	db := dbconfig.InitGORM()
	err  := db.Create(&department).GetErrors()
	if err == nil {
		c.JSON(201, department)
	} else {
		fmt.Printf("Error when looking up Table, the error is '%v'", err)
		c.JSON(400, gin.H{"message": "error"})
	}

	//dbmap := dbconfig.Init()
	//if department.Name != "" {
	//
	//	if insert, _ := dbmap.Exec(`INSERT INTO department (name) VALUES (?)`, department.Name); insert != nil {
	//		department_id, err := insert.LastInsertId()
	//		if err == nil {
	//			content := &departmentmodel.Department{
	//				Id:        department_id,
	//				Name: department.Name,
	//
	//			}
	//			c.JSON(201, content)
	//		} else {
	//			checkErr(err, "Insert failed")
	//		}
	//	}
	//
	//} else {
	//	c.JSON(400, gin.H{"error": "Fields are empty"})
	//}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/departments
}

func UpdateDepartment(c *gin.Context) {
	id := c.Params.ByName("id")
	var department departmentmodel.Department
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&department, "SELECT * FROM department WHERE id=?", id)

	if err == nil {
		var json departmentmodel.Department
		c.Bind(&json)

		department_id, _ := strconv.ParseInt(id, 0, 64)

		department := departmentmodel.Department{
			Id:        department_id,
			Name: json.Name,
		}

		if department.Name != ""  {
			_, err = dbmap.Update(&department)

			if err == nil {
				c.JSON(200, department)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "department not found"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/departments/1
}

func DeleteDepartment(c *gin.Context) {
	id := c.Params.ByName("id")

	var department departmentmodel.Department
	dbmap := dbconfig.Init()
	err := dbmap.SelectOne(&department, "SELECT * FROM department WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&department)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "department not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/departments/1
}

func OptionsDepartment(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
