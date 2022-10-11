package covidserv

import (
	"github.com/gin-gonic/gin"

	"github.com/pitsanujiw/go-covid/internal/covid/handler"
	"github.com/pitsanujiw/go-covid/internal/covid/service"
	"github.com/pitsanujiw/go-covid/pkg/graceful"
	"github.com/pitsanujiw/go-covid/pkg/httpclient"
)

func setup() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	http := httpclient.New()

	covidSer := service.NewCovidServ(http)

	covidRoute := router.Group("/covid")

	handler.New(covidRoute, covidSer)

	return router
}

func Run() {
	runner := setup()

	graceful.StartServerWithGracefulShutdown(runner)
}
