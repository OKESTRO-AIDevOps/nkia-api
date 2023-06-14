package main

import (
	"fmt"
)

var TEST_LOCAL = true

var TEST_PKG = false

func main() {

	if err := Test_local(); err != nil {

		_ = fmt.Errorf("test failure: local: %s", err.Error())

		return

	} else {

		fmt.Println("test success: local")

	}

	if err := Test_pkg(); err != nil {

		_ = fmt.Errorf("test failure: pkg: %s", err.Error())

		return

	} else {

		fmt.Println("test success: pkg")

	}

}
