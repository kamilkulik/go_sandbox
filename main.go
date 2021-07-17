package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Fear of the Dark", Artist: "Iron Maiden", Price: 49.99},
	{ID: "2", Title: "Back to Black", Artist: "ACDC", Price: 39.99},
	{ID: "3", Title: "Stone Sour", Artist: "Stone Sour", Price: 29.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}