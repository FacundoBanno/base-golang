package pkg

import (
	"github.com/gin-gonic/gin"
	"my-driveapp/internal/application"
	"net/http"
)

type api struct {
	router *gin.Engine
}

type Server interface {
	Run() http.Handler
}

func New() Server {
	var (
		a = &api{}
		r = gin.Default()
	)

	a.router = r

	return a
}

func (a *api) Run() http.Handler {
	application.Endpoints(a.router)
	return a.router
}