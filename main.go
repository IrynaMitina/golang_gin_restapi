package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	_"github.com/IrynaMitina/golang_gin_restapi/docs"  // load local docs/docs.go
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// album model
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title     Albums API
// @version         1.0
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.Run("localhost:8080")
}

// getAlbums     godoc
// @Summary      Get albums
// @Description  Responds with the list of all albums as JSON.
// @Tags         albums
// @Produce      json
// @Success      200  {array}  album
// @Router       /albums [get]
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums    godoc
// @Summary      Add a new album
// @Description  Takes an album JSON and store in DB. Return saved JSON.
// @Tags         albums
// @Produce      json
// @Param        album  body  album  true  "Album JSON"
// @Success      200   {object}  album
// @Router       /albums [post]
func postAlbums(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID  godoc
// @Summary      Get single album by id
// @Description  Returns the album whose ID value matches the id sent by the client.
// @Tags         albums
// @Produce      json
// @Param        id  path      int  true  "search album by id"
// @Success      200  {object}  album
// @Router       /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}