package adminorigin

import (
	"encoding/json"
	"os"
	"os/exec"

	. "github.io/seantywork/npia/pkg/libinterface"
)

var adm_origin_json = LIBIF.GetLibComponentPath(".etc", "ADM_origin.json")

func GetRecordInfo(records []RecordInfo, ns string) (string, string) {

	arr_leng := len(records)

	var repo_addr string = "N"

	var reg_addr string = "N"

	for i := 0; i < arr_leng; i++ {

		if records[i].NS == ns {

			repo_addr = records[i].REPO_ADDR

			reg_addr = records[i].REG_ADDR

			break

		}

	}

	return repo_addr, reg_addr

}

func SetRecordInfo(records []RecordInfo, ns string, repo_addr string, reg_addr string) []RecordInfo {

	exists := 0

	arr_leng := len(records)

	var new_record_info RecordInfo

	for i := 0; i < arr_leng; i++ {

		if records[i].NS == ns {

			exists = 1

			records[i].REPO_ADDR = repo_addr

			records[i].REG_ADDR = reg_addr

			break

		}
	}

	if exists != 1 {

		new_record_info.NS = ns

		new_record_info.REPO_ADDR = repo_addr

		new_record_info.REG_ADDR = reg_addr

		records = append(records, new_record_info)

	}

	return records

}

func GetRepoInfo(repos []RepoInfo, addr string) (string, string) {

	arr_leng := len(repos)

	var repo_id string = "N"

	var repo_pw string = "N"

	for i := 0; i < arr_leng; i++ {

		if repos[i].REPO_ADDR == addr {

			repo_id = repos[i].REPO_ID

			repo_pw = repos[i].REPO_PW

			break

		}
	}

	return repo_id, repo_pw

}

func SetRepoInfo(repos []RepoInfo, addr string, id string, pw string) []RepoInfo {

	exists := 0

	arr_leng := len(repos)

	repo_id := id

	repo_pw := pw

	var new_repo_info RepoInfo

	for i := 0; i < arr_leng; i++ {

		if repos[i].REPO_ADDR == addr {

			exists = 1

			repos[i].REPO_ID = repo_id

			repos[i].REPO_PW = repo_pw

			break

		}
	}

	if exists != 1 {

		new_repo_info.REPO_ADDR = addr

		new_repo_info.REPO_ID = repo_id

		new_repo_info.REPO_PW = repo_pw

		repos = append(repos, new_repo_info)

	}

	return repos

}

func GetRegInfo(regs []RegInfo, addr string) (string, string) {

	arr_leng := len(regs)

	var reg_id string = "N"

	var reg_pw string = "N"

	for i := 0; i < arr_leng; i++ {

		if regs[i].REG_ADDR == addr {

			reg_id = regs[i].REG_ID

			reg_pw = regs[i].REG_PW

			break

		}
	}

	return reg_id, reg_pw

}

func SetRegInfo(regs []RegInfo, addr string, id string, pw string) []RegInfo {

	exists := 0

	arr_leng := len(regs)

	reg_id := id

	reg_pw := pw

	var new_reg_info RegInfo

	for i := 0; i < arr_leng; i++ {

		if regs[i].REG_ADDR == addr {

			exists = 1

			regs[i].REG_ID = reg_id

			regs[i].REG_PW = reg_pw

			break

		}
	}

	if exists != 1 {

		new_reg_info.REG_ADDR = addr

		new_reg_info.REG_ID = reg_id

		new_reg_info.REG_PW = reg_pw

		regs = append(regs, new_reg_info)

	}

	return regs
}

func CheckAppOrigin() string {

	var app_origin AppOrigin

	file_content, err := os.ReadFile(adm_origin_json)

	if err != nil {

		return "Origin file is corrupted"

	}

	err = json.Unmarshal(file_content, &app_origin)

	if err != nil {

		return "Origin file is corrputed"
	}

	if app_origin.KCFG_PATH == "" {

		return "Origin path is not set"

	}

	kcfg := app_origin.KCFG_PATH

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg, "get", "nodes")

	_, err = cmd.Output()

	if err != nil {

		strerr := err.Error()

		return strerr

	}

	if len(app_origin.RECORDS) == 0 {

		return "WARNRC"

	}

	if len(app_origin.REGS) == 0 || len(app_origin.REPOS) == 0 {

		return "WARNRE"

	}

	if app_origin.MAIN_NS == "" {

		return "WARNNS"
	}

	return "OKAY"
}

func GetKubeConfigAndTargetNameSpace() (string, string) {

	var app_origin AppOrigin

	file_content, _ := os.ReadFile(adm_origin_json)

	_ = json.Unmarshal(file_content, &app_origin)

	kcfg_path := app_origin.KCFG_PATH

	main_ns := app_origin.MAIN_NS

	return kcfg_path, main_ns

}
