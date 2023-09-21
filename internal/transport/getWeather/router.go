package handlerGetWeather

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/traningWithPgxPackage/balance"
)

func NewGetWeatherRouter(handler *gin.Engine, redis *redis.Client, balancer *balance.LoadBalancer) {
	// Options
	handler.Use(gin.Logger())   // logging data about the request
	handler.Use(gin.Recovery()) //send 500 if there is a panic
	g := &getWeatherRouted{redis, balancer}

	h := handler.Group("/v1") // add prefix to Group
	{
		h.GET("weather", g.doGetWeather)
	}
}
