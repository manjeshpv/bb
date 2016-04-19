package chainroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjeshpv/bb/api/chain/controller"
)

func Init(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/chains", chaincontroller.GetChains)
		v1.GET("/chains/:id", chaincontroller.GetChain)
		v1.POST("/chains", chaincontroller.PostChain)
		v1.PUT("/chains/:id", chaincontroller.UpdateChain)
		v1.DELETE("/chains/:id", chaincontroller.DeleteChain)
		v1.OPTIONS("/chains", chaincontroller.OptionsChain)     // POST
		v1.OPTIONS("/chains/:id", chaincontroller.OptionsChain) // PUT, DELETE
	}
}
