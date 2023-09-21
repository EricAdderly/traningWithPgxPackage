package handlerGetWeather

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/traningWithPgxPackage/balance"
	getWeatherServices "github.com/traningWithPgxPackage/internal/services/getWeather"
	"github.com/traningWithPgxPackage/logger"
)

// Взаимодействие с внешним миром
type getWeatherRouted struct {
	client   *redis.Client
	balancer *balance.LoadBalancer
}

func (g *getWeatherRouted) doGetWeather(c *gin.Context) {
	g.balancer.ReverseProxyHandler(c.Writer, c.Request)

	town := c.Query("town")
	if len(town) == 0 {
		logger.ErrorLog("GetWeatherHandler", "invalid request body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
	}

	weather, err := getWeatherServices.GetMeWeather(town, g.client)
	if err != nil {
		logger.ErrorLogWithError("GetWeatherHandler", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "town not found")
	}

	c.JSON(http.StatusOK, weather)
}
