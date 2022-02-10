package application

import (
"github.com/gin-gonic/gin"
"net/http"
)

type LOGIN struct{
	USER string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func Endpoints(router *gin.Engine) {
	testing(router)
}

func testing(router *gin.Engine){
	router.GET("/ping",
		func(g *gin.Context){
		g.JSON(http.StatusOK, "pong")
	})

	router.POST("/foo", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.USER}) // Your custom response here
	})
}