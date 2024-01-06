package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"main/internal/app/metrics"
	"main/internal/app/middleware"
	"net/http"
)

func main() {
	r := SetupRouter()
	r.GET("/ping", func(c *gin.Context) {
		pong(c)
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	err := r.Run(":80")

	if err != nil {
		panic(err)
	}
}

type ResponseData struct {
	Message string `json:"message"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(middleware.RequestMiddleware(metrics.New()))

	if err != nil {
		return nil
	}

	return r
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ResponseData{
		Message: "pong",
	})
}
