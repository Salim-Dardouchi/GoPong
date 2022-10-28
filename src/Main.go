package main

import (
	"fmt"

	"GoPong/App"
)

func main() {
	fmt.Println("Entering GoPong")

	app, err := App.New(1280, 720, "GoPong")

	if err != nil {
		return
	}

	app.Run()

	fmt.Println("Exiting GoPong")
}
