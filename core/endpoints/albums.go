package endpoints

import (
	"github.com/Kamran151199/gin-golang-api/core/database"
	"github.com/Kamran151199/gin-golang-api/core/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum structs.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	database.Albums = append(database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func RetrieveAlbum(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range database.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
