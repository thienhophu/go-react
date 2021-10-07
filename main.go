package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/thienhophu/go-react/api"
)

func main() {
	router := gin.Default()

	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumById)
	router.POST("/albums", api.PostAlbums)

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	// router.Run(":3000")
	router.Run()
}
