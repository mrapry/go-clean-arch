package main

import (
	"fmt"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"go-mysql/config/env"
	"go-mysql/middleware"
	"net/http"
	"time"
)

// EchoServer structure
type EchoServer struct {
}

// time now struct
type timenow struct {
	Datetime string `json:"datetime"`
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger, mid.CORS())

	if env.GetEnv().Development == "1" {
		e.Debug = true
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service up and running !!")
	})

	e.GET("/api/timenow", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &timenow{
			Datetime: time.Now().Format(time.RFC3339),
		})
	})

	//dilGroupV1 := e.Group("/v1/dil")
	//dilGroupV1.Use()
	//s.dilV1.Mount(dilGroupV1)

	// set listener port
	listenerPort := fmt.Sprintf(":%d", env.GetEnv().HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

func MainHttp()(*EchoServer, error)  {
	return &EchoServer{}, nil
}