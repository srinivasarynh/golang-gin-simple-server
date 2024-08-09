package main

import (
	"fmt"
	"gogintuts/middleware"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("this is main file")
	router := gin.Default()

	// gins basic authentication and pass it into authgroups
	auth := gin.BasicAuth(gin.Accounts{
		"gin": "gintuki",
	})

	// middleware
	router.Use(middleware.Authentication)

	// router.GET("/data", getData)
	// router.POST("/data", postData)

	// applying middleware to specific route
	router.GET("/queryData", middleware.Authentication, getquerydata)
	router.GET("/getdata/:name/:age", getparamsdata)

	// router groups
	admin := router.Group("/admin", middleware.Authentication, auth)
	{
		admin.GET("/data", getData)
	}
	client := router.Group("/client")
	{
		client.POST("/data", postData)
	}

	// built in gin
	// router.Run(":8000")

	// built in golang server configuring
	server := http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// running server
	server.ListenAndServe()
}

func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is get request",
	})
}

func postData(ctx *gin.Context) {
	body := ctx.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		log.Fatal("failed to read request boyd")
	}

	ctx.JSON(200, gin.H{
		"message": "this is post request",
		"value":   string(value),
	})
}

func getquerydata(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")

	ctx.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func getparamsdata(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")

	ctx.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
