package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thienhophu/go-react/greeting"
)

func main() {
	log.SetPrefix(("greeting: "))
	log.SetFlags(0)

	message, err := greeting.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	// r.Run()
	r.Run(":3000")
}
