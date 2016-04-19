package hotelroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjeshpv/bb/api/hotel/controller"
)

func Init(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/hotels", hotelcontroller.GetHotels)
		v1.GET("/hotels/:id", hotelcontroller.GetHotel)
		v1.POST("/hotels", hotelcontroller.PostHotel)
		v1.PUT("/hotels/:id", hotelcontroller.UpdateHotel)
		v1.DELETE("/hotels/:id", hotelcontroller.DeleteHotel)
		v1.OPTIONS("/hotels", hotelcontroller.OptionsHotel)     // POST
		v1.OPTIONS("/hotels/:id", hotelcontroller.OptionsHotel) // PUT, DELETE
	}
}
