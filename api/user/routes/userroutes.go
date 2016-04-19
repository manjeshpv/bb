package userroutes

import (
	"github.com/manjeshpv/bb/api/user/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", usercontroller.GetUsers)
		v1.GET("/users/:id", usercontroller.GetUser)
		v1.POST("/users", usercontroller.PostUser)
		v1.PUT("/users/:id", usercontroller.UpdateUser)
		v1.DELETE("/users/:id", usercontroller.DeleteUser)
		v1.OPTIONS("/users", usercontroller.OptionsUser)     // POST
		v1.OPTIONS("/users/:id", usercontroller.OptionsUser) // PUT, DELETE
	}
}
