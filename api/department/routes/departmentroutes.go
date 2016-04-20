package departmentroutes

import (
	"github.com/manjeshpv/bb/api/department/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/departments", departmentcontroller.GetDepartments)
		v1.GET("/departments/:id", departmentcontroller.GetDepartment)
		v1.POST("/departments", departmentcontroller.PostDepartment)
		v1.PUT("/departments/:id", departmentcontroller.UpdateDepartment)
		v1.DELETE("/departments/:id", departmentcontroller.DeleteDepartment)
		v1.OPTIONS("/departments", departmentcontroller.OptionsDepartment)     // POST
		v1.OPTIONS("/departments/:id", departmentcontroller.OptionsDepartment) // PUT, DELETE
	}
}
