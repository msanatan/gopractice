package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msanatan/gopractice/ginexample/models"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbums)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func createAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		panic(err)
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func removeAlbum(slice []models.Album, idx int) []models.Album {
	return append(slice[:idx], slice[idx+1:]...)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for idx, a := range albums {
		if a.ID == id {
			albums = removeAlbum(albums, idx)
			c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
