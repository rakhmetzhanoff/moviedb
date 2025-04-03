package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"moviedb/internal/models"
	"moviedb/internal/routes"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://movie_user:movie_password@localhost:5444/movie_database?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
