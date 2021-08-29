package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {

	api := g.Group("/api")
	{
		api.POST("/addBook", addBook)
		api.GET("/getBook", getBook)
		api.PUT("/updateBook", updateBook)
		api.DELETE("/deleteBook", deleteBook)
	}

	g.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus((http.StatusOK))
	})

	g.NoRoute(lost)
}

func respond(c *gin.Context, status int, payload interface{}, err error) {
	if err != nil {
		log.Println("[ERROR]: ", err)
		c.JSON(status, map[string]interface{}{"error": err.Error()})
	} else {
		if payload != nil {
			resp, _ := json.Marshal(payload)
			log.Println("[INFO]: ", string(resp))
			c.Data(status, "application/json", resp)
		} else {
			log.Println("[INFO]: Status OK")
			c.String(status, "")
		}
	}
}
