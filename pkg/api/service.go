package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/turbovladimir/RestApi/pkg/metrics"
	"github.com/turbovladimir/RestApi/pkg/middleware"
	"net/http"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type Router struct {
	engine *gin.Engine
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type ResponseData struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

func NewRouter(routes []Route) *Router {
	engine := gin.Default()
	engine.ForwardedByClientIP = true
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		panic("Cannot set trusted proxies")
	}

	engine.Use(middleware.RequestMiddleware(metrics.New()))

	routes = append(routes,
		Route{
			Method: MethodGet,
			Path:   "/ping",
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, ResponseData{
					Data: "pong",
				})
			},
		},
		Route{
			Method:  MethodGet,
			Path:    "/metrics",
			Handler: gin.WrapH(promhttp.Handler()),
		},
	)

	r := &Router{engine: engine}
	r.addRoutes(routes)

	return r
}

func (r Router) addRoutes(routes []Route) {
	for _, route := range routes {
		if route.Method == MethodGet {
			r.engine.GET(route.Path, route.Handler)
		}

		if route.Method == MethodPost {
			r.engine.POST(route.Path, route.Handler)
		}
	}
}

func (r Router) Run(port string) {
	err := r.engine.Run(":" + port)

	if err != nil {
		panic(fmt.Sprintf("Cannot run on port %v. Message: %v", port, err.Error()))
	}
}
