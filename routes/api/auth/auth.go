package auth

import (
	"crypto_track/controllers"
	"github.com/gin-gonic/gin"
)

func Handle(r *gin.RouterGroup) {

	g := r.Group("/auth")
	{
		g.POST("/sign-up", controllers.SignUp)
		g.POST("/login", controllers.Login)
	}

}
