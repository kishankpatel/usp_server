package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kishankpatel/usp_server/routes"
)

func main() {
	g := gin.Default()

	routes.Handler(g)
	g.Run(":4040")
}
