package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:         ":4043",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.GET("/api", test)

	s.ListenAndServe()
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It is alive!",
	})
}
