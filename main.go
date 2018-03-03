package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()

	r.GET("/dispensers/:slot", func(c *gin.Context) {
		slot := c.Param("slot")
		log.Printf("Dispensing slot %s\n", slot)

		c.JSON(http.StatusOK, gin.H{
			"rc":  "ok",
			"msg": "",
		})
	})

	log.Println("**** Running http://localhost:5000")
	r.Run(":5000")
}
