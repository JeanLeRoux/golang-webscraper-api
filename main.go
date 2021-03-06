package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/Tech", getLatestTech)
	router.GET("/CryptoMeta", getCryptoMetadata)
	router.GET("/CryptoNews", getCryptoNews)
	router.GET("/CryptoChart", getCryptoChartData)
	router.GET("/CryptoDetails", getCryptoDetails)
	router.Run()
}
