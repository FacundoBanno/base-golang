package application

import (
"github.com/gin-gonic/gin"
"net/http"
)

func Endpoints(router *gin.Engine) {
	testing(router)
}

func testing(router *gin.Engine){
	router.GET("/ping", func(g *gin.Context){
		g.String(http.StatusOK, "pong")
	})
}