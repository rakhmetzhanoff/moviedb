package main

import (
	"github.com/gin-gonic/gin"
	"moviedb/internal/db"
	"moviedb/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
