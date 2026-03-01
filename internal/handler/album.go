package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Album represents data about a record album.
type Album struct {
    ID     string  `json:"id" example:"1"`
    Title  string  `json:"title" example:"Blue Train"`
    Artist string  `json:"artist" example:"John Coltrane"`
    Price  float64 `json:"price" example:"56.99"`
}

// albums slice to seed record album data.
var albums = []Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// RegisterRoutes registers the album-related routes on the gin engine.
func RegisterRoutes(router *gin.Engine) {
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
}

// getAlbums responds with the list of all albums as JSON.
// @Summary      Get all albums
// @Description  Responds with the list of all albums as JSON.
// @Tags         albums
// @Produce      json
// @Success      200  {array}   Album
// @Router       /albums [get]
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
// @Summary      Post a new album
// @Description  Adds an album from JSON received in the request body.
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        album  body      Album  true  "Album JSON"
// @Success      201    {object}  Album
// @Router       /albums [post]
func postAlbums(c *gin.Context) {
    var newAlbum Album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter.
// @Summary      Get an album by ID
// @Description  Locates the album whose ID value matches the id parameter.
// @Tags         albums
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      200  {object}  Album
// @Failure      404  {object}  map[string]string
// @Router       /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
