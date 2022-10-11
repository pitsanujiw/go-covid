package covidserv

import (
	"github.com/gin-gonic/gin"

	"github.com/pitsanujiw/go-covid/internal/covid/handler"
	"github.com/pitsanujiw/go-covid/internal/covid/service"
	"github.com/pitsanujiw/go-covid/pkg/graceful"
	"github.com/pitsanujiw/go-covid/pkg/httpclient"
)

func setup() *gin.Engine {
	// new gin framework
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	// new http client
	http := httpclient.New()

	// new get covid summary service
	covidSer := service.NewCovidServ(http)

	// group api
	covidRoute := router.Group("/covid")

	// new handler
	handler.New(covidRoute, covidSer)

	return router
}

func Run() {
	runner := setup()
	// graceful shutdown
	graceful.StartServerWithGracefulShutdown(runner)
}
