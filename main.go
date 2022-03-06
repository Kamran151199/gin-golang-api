package main

import (
	"github.com/Kamran151199/gin-golang-api/core/endpoints"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/albums", endpoints.GetAlbums)
	router.POST("/albums", endpoints.PostAlbums)
	router.GET("/albums/:id", endpoints.RetrieveAlbum)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
