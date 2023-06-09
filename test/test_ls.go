package main

import (
	"fmt"
	"os"
)

func ReadDirectory() {

	dir_name := "../lib"

	files, _ := os.Open(dir_name)

	defer files.Close()

	fileInfos, _ := files.Readdir(-1)

	for _, f_info := range fileInfos {
		fmt.Println(f_info.Name()) //print the files from directory
	}

}
