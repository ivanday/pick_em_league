package main

import (
	"github.com/gin-gonic/gin"
	"pick_em_league/handlers"
)

func main() {
	router := gin.Default()

	oddsApi := router.Group("/odds")
	{
		oddsApi.GET("/", handlers.GetOdds)
	}
}
