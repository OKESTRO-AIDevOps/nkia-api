package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.io/seantywork/npia/pkg/promquery"

	agph "github.com/guptarohit/asciigraph"
)

func CheckIfEleInStrList(ele string, str_list []string) bool {

	hit := false

	for i := 0; i < len(str_list); i++ {

		if str_list[i] == ele {

			hit = true

			return hit
		}

	}

	return hit

}

func RenderASCIIGraph(render_target string) {

	var render_data promquery.PQOutputFormat

	file_byte, _ := os.ReadFile(render_target)

	json.Unmarshal(file_byte, &render_data)

	out := agph.Plot(render_data.Values)

	fmt.Println(out)

}
