package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//creating a gin router harimo default middle ware like: logging,recovery,etc
	rimwe := gin.Default()

	//defining a simple get request to /bite path
	rimwe.GET("/bite", func(c *gin.Context) {
		//responding with a json object
		c.JSON(200, gin.H{
			"Ubutumwa": "Hii  :) !!",
			"status":   "Byakunze, can't u tell pls",
		})
	})

	//starting the server on given port
	rimwe.Run(":8000")
}
