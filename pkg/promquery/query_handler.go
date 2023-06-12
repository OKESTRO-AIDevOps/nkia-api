package promquery

import (
	"fmt"

	gabs "github.com/Jeffail/gabs/v2"

	agph "github.com/guptarohit/asciigraph"
)

type PQOutputFormat struct {
}

func PromQueryStandardizer(raw_query_bytes []byte, to_terminal bool) (string, error) {

	ret_json_string := ""

	var value_float_arr []float64

	json_parsed, err := gabs.ParseJSON(raw_query_bytes)

	if err != nil {

		return "[]", fmt.Errorf("received query result format corrupted, abort: %s", err.Error())

	}

	_ = json_parsed.Path("data.result")

	if to_terminal {

		PromQueryResultASCIIGraph(value_float_arr)

	}

	return ret_json_string, nil

}

func PromQueryResultASCIIGraph(values []float64) {

	agph.Plot(values)

}
