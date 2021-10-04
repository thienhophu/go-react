package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", GetAlbums)

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.Run(":3000")
	// router.Run()
}
