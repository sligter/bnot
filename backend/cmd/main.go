package main

import (
	"log"

	"bnot/backend/internal/api"
	"bnot/backend/internal/auth"
	"bnot/backend/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	auth.SetupRoutes(r)
	api.SetupRoutes(r)

	r.Run(":8080")
}
