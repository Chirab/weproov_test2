package main

import (
	"log"
	"weproov/cmd"
	"weproov/config"
)

// with go modules disabled

func main() {
	config, err := config.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }
	
	e := cmd.NewServer(config.PORT)
	e.Execute()
}
