package main

import (
	"fmt"
	"os"
	"strings"

	goya "github.com/goccy/go-yaml"
	//"github.com/fatih/color"
)

func Test_local() error {

	if !TEST_LOCAL {

		return nil

	}

	return nil
}

func test_YAMLLoad(file_path string) {

	file_byte, _ := os.ReadFile(file_path)

	file_list := strings.Split(string(file_byte), "---")

	for _, yaml_file := range file_list {

		test_ReadFromYAML(yaml_file, "$.spec.ports[0].port")

	}

}

func test_ReadFromYAML(yaml_file string, yaml_path string) {

	ypath, _ := goya.PathString(yaml_path)

	var value int

	_ = ypath.Read(strings.NewReader(yaml_file), &value)

	fmt.Println(value)

}
