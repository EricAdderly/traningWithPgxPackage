package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/traningWithPgxPackage/balance"
	"github.com/traningWithPgxPackage/config"
	"github.com/traningWithPgxPackage/db"
	handlerGetWeather "github.com/traningWithPgxPackage/internal/transport/getWeather"
	"github.com/traningWithPgxPackage/redis"
)

// это вход во внутреннюю логику сервиса

func Run(cfg *config.Config) {
	pg, err := db.CreateDbPool(cfg.DB.Url)
	if err != nil {
		log.Fatalf("Connection to db error: %s", err)
	}
	defer pg.Close()

	redis := redis.RedisNew(cfg.Redis.Address, cfg.Redis.Password, cfg.Redis.DBName)
	balancer := balance.NewLoadBalancer()

	handler := gin.New()
	handlerGetWeather.NewGetWeatherRouter(handler, redis, balancer)
	handler.Run(cfg.HTTP.Port)
}
