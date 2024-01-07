package main

import (
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/RestApi/pkg/api"
)

func main() {
	router := api.NewRouter([]api.Route{}, []gin.HandlerFunc{})
	router.Run("8089")
}
