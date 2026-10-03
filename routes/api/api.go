package api

import (
	"crypto_track/middlewares"
	"crypto_track/routes/api/auth"
	"crypto_track/routes/api/tracks"

	"github.com/gin-gonic/gin"
)

func Handle(r *gin.Engine) {

	g := r.Group("/api")
	auth.Handle(g)

	g.Use(middlewares.JWTAuthMiddleware())
	tracks.Handle(g)

}
