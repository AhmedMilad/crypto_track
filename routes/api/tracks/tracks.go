package tracks

import (
	"crypto_track/controllers"
	"github.com/gin-gonic/gin"
)

func Handle(r *gin.RouterGroup) {

	g := r.Group("/tracks")
	{
		g.POST("/", controllers.CreateTrack)
	}

}
