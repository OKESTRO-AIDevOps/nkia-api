package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func StringifyAllJSONFields() {

	test_json_file, _ := os.ReadFile("./data/cpu_usage.json")

	var f interface{}

	json.Unmarshal(test_json_file, &f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
