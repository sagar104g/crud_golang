package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sagar104g/crud_golang/handlers"
	_ "github.com/sagar104g/crud_golang/init"
)

func main() {
	g := gin.New()

	g.Use(gin.LoggerWithWriter(os.Stdout))

	PORT := ":" + os.Getenv("PORT")
	if PORT == ":" {
		PORT = ":3000"
	}
	handlers.Register(g)
	_ = g.Run(PORT)
	log.Println("app is live on port " + PORT)
}
