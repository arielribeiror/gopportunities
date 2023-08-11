package main

import (
	"fmt"

	"github.com/arielribeiror/gopportunities/config"
	"github.com/arielribeiror/gopportunities/router"
)

func main() {

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
