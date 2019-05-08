package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"encoding/json"
)
func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../client/public", true)))
	ping := router.Group("/path") 
	ping.GET("/", post)
	router.Run(":5000")
	
}

func post(c *gin.Context) {
	err := c.Request.Body.ParseForm()
    if err != nil {
        // Handle error here via logging and then return            
    }

    query := c.Request.FormValue("query")
    fmt.Println(query)
}