package main

import (
	"fmt"
)

var TEST_PKG = true

func main() {

	if err := Test_pkg(); err != nil {

		_ = fmt.Errorf("test failure: pkg: %s", err.Error())

		return

	} else {

		fmt.Println("test success: pkg")

	}

}
