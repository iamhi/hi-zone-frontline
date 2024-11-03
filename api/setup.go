package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamhi/frontline/config"
	"github.com/iamhi/frontline/db/postgres"
)

func StartService() {
	config.Load()

	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())
	gin_engine.Use(gin.Logger())

	gin_engine.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Accept", "User-Agent"},
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8083", "https://api.ibeenhi.com", "https://hi-zone.ibeenhi.com", "*"},
		AllowWildcard:    true,
	}))

	initialize(gin_engine)

	postgres.Setup()

	gin_engine.Run("localhost:8083")

	fmt.Println("Server started")
}
