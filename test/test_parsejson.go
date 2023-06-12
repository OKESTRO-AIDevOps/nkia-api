package main

import (
	"fmt"
	"os"

	gabs "github.com/Jeffail/gabs/v2"
)

func ParseJSONGabs(file_name string) {

	file_byte, _ := os.ReadFile(file_name)

	json_parsed, err := gabs.ParseJSON(file_byte)

	if err != nil {

		fmt.Println("not okay!")

		return

	} else {

		fmt.Println("okay!")
	}

	ss := json_parsed.Path("series.0.refId")

	fmt.Println(ss)
}
