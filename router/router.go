package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorwvm/rest_api/controller"
)

func StartRouter() {

	r := gin.Default()

	r.GET("/films", controller.GetFilms)

	r.Run("localhost:8080")

}
