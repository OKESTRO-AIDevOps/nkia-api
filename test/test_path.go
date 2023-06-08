package main

import (
	"fmt"
	"path/filepath"
)

func PathTest() {

	var LIB_ROOT, _ = filepath.Abs("../../lib")

	fmt.Println(LIB_ROOT)

	var LIB_BIN = filepath.Join(LIB_ROOT, "bin")

	fmt.Println(LIB_BIN)

	var LIB_SCRIPTS = filepath.Join(LIB_ROOT, "scripts")

	fmt.Println(LIB_SCRIPTS)

	var LIB_TEMPLATE = filepath.Join(LIB_ROOT, "template")

	fmt.Println(LIB_TEMPLATE)

}
