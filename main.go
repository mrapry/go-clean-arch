package main

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"go-clean-arch/config/env"
	"os"
	"sync"
)

func main()  {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		err := errors.New(".env is not loaded properly")
		log.Error(err.Error())
		os.Exit(1)
	}

	env.InitEnvironment() // Load env config to in memory

	echoServer, err := MainHttp()

	if err != nil {
		log.Error("Echo Server initial, "+err.Error())
		os.Exit(2)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		echoServer.Run()
	}()

	wg.Wait()
}