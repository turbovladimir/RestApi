package main

import (
	"github.com/turbovladimir/RestApi/pkg/api"
)

func main() {
	router := api.NewRouter([]api.Route{})
	router.Run("8086")
}
