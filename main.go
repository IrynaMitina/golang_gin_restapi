package main

import (
    "net/http"
    "strings"
    "strconv"
    "github.com/gin-gonic/gin"
	_"github.com/IrynaMitina/golang_gin_restapi/docs"  // load local docs/docs.go
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Albums API
// @version         1.0
func main() {
    ConnectDB()

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
    allAlbums, err := allAlbums()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
        return
    }
    c.IndentedJSON(http.StatusOK, allAlbums)
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
    id, err := addAlbum(newAlbum)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
        return
    }
    newAlbum.ID = id
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
    id, _ := strconv.Atoi(c.Param("id"))
    alb, err := albumById(id)
    if err != nil {
        if strings.Contains(err.Error(), "no such album") {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
        } else {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
        }
        return
    }
    c.IndentedJSON(http.StatusOK, alb)    
}