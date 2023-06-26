package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"

	pkgutils "github.com/OKESTRO-AIDevOps/npia-api/pkg/utils"

	goya "github.com/goccy/go-yaml"
	//"github.com/fatih/color"
)

type AppOrigin struct {
	KCFG_PATH string
	MAIN_NS   string
	RECORDS   []RecordInfo
	REPOS     []RepoInfo
	REGS      []RegInfo
}

type RecordInfo struct {
	NS        string
	REPO_ADDR string
	REG_ADDR  string
}

type RepoInfo struct {
	REPO_ADDR string
	REPO_ID   string
	REPO_PW   string
}

type RegInfo struct {
	REG_ADDR string
	REG_ID   string
	REG_PW   string
}

func yamlLoad(file_path string) {

	file_byte, _ := os.ReadFile(file_path)

	file_list := strings.Split(string(file_byte), "---")

	for _, yaml_file := range file_list {

		readFromYAML(yaml_file, "$.spec.ports[0].port")

	}

}

func readFromYAML(yaml_file string, yaml_path string) {

	ypath, _ := goya.PathString(yaml_path)

	goya.JSON()

	var value int

	_ = ypath.Read(strings.NewReader(yaml_file), &value)

	fmt.Println(value)

}

func komposeTest() {

	cmd := exec.Command("../lib/bin/kompose", "convert", "-f", "../lib/bin/docker-compose.yaml", "--stdout")

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func writeToAdmOrigin() {

	var test_ao AppOrigin

	var test_ri RecordInfo

	var test_rep RepoInfo

	var test_reg RegInfo

	test_ao.RECORDS = append(test_ao.RECORDS, test_ri)

	test_ao.REPOS = append(test_ao.REPOS, test_rep)

	test_ao.REGS = append(test_ao.REGS, test_reg)

	file_byte, _ := json.Marshal(test_ao)

	_ = os.WriteFile("testadmorigin.json", file_byte, 0644)

}

func callApiDef() {

	ASgi.PrintPrettyDefinition()

}

func callApiDefStructure() {

	ASgi.PrintRawDefinition()

}

func sliceTest() {

	ret := pkgutils.InsertToSliceByIndex[string]([]string{"b", "c", "d"}, 0, "a")

	fmt.Println(ret)
}

func main() {

	//	callApiDefStructure()

	// sliceTest()

	komposeTest()
}
