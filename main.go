package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"pick_em_league/handlers"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	oddsApi := router.Group("/odds")
	{
		oddsApi.GET("/sports", handlers.GetSports)
		oddsApi.GET("/spreads/nfl", handlers.GetNFLSpreads)
	}

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
