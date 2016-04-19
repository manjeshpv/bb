package routes


import (
	"github.com/manjeshpv/bb/api/user/routes"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	userroutes.Init(r)
}
