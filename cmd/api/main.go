package main

import (
	"arturogood17/local-manga-tracking/internal/api/controllers"
	"net/http"
	"time"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := viper.Get("APP_PORT")

	s := &http.Server{
		Addr:         port.(string), //se necesita hacer type assertion porque viper.Get() devuelve Any
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.GET("/api", test)
	router.POST("/users", controllers.CreateUser)

	s.ListenAndServe()
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It is alive!",
	})
}
