package main

import (
	"fmt"
	"iot-go-api/controllers"
	"iot-go-api/core/authentication"
	"iot-go-api/settings"
	"iot-go-api/user_controllers"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {
	settings.Init()
	StartGin()
}

func StartGin() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	gin.Logger()
	//	router.Use(rateLimit, gin.Recovery())
	router.Use(gin.Logger())
	router.GET("/", MyBenchLogger(), index)
	router.GET("/auth", authentication.RequireTokenAuthentication(), index)
	router.POST("/test", controllers.Login)

	//mongodb user create
	uc := user_controllers.NewUserController(getSession())
	router.GET("/user", uc.GetUser)
	router.POST("/jwtowner", uc.JwtCreateOwner)
	router.GET("/jwtclient", uc.JwtCreateclilent)
	router.POST("/message", uc.GetMessage)

	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.RemoveUser)

	router.Run(":5001")

}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		fmt.Println(start, path)
		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		fmt.Println(latency)

	}
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
