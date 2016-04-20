package routes


import (
	"github.com/manjeshpv/bb/api/user/routes"
	"github.com/manjeshpv/bb/api/chain/routes"
	"github.com/manjeshpv/bb/api/hotel/routes"
	"github.com/manjeshpv/bb/api/department/routes"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	userroutes.Init(r)
	chainroutes.Init(r)
	hotelroutes.Init(r)
	departmentroutes.Init(r)
}
