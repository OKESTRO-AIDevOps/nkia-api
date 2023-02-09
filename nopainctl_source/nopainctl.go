package main

import (
	"fmt"
	"os"
)

var SERVER_URL string = "http://localhost:7331"

var SERVER_SUB_URL string = "http://localhost:7331/build"

func up() {

}

func down() {

}

func exec() {

}

func main() {

	action := os.Args[1]

	if action == "up" {

		up()

	} else if action == "down" {

		down()

	} else if action == "exec" {

		exec()

	} else {

		fmt.Println("Option Not Available")
	}

}
