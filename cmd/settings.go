package main

import (
	"encoding/json"
	"fmt"

	runfs "github.com/seantywork/014_npia/pkg/runtimefs"

	"github.com/fatih/color"
)

func origin_set() (int, error) {

	code := ""

	fmt.Println("TARGET : origin")
	fmt.Scanln(&code)

	evelp := 0

	for evelp == 0 {

		switch code {

		case "namespace-new":
			ns := ""
			repo := ""
			reg := ""

			color.Blue("RUN: origin namespace-new")

			fmt.Println("New namespace:")
			fmt.Scanln(&ns)
			fmt.Println("New repo URL:")
			fmt.Scan(&repo)
			fmt.Println("New reg URL:")
			fmt.Scanln(&reg)

			if err := runfs.SetAdminOriginNewNS(ns, repo, reg); err != nil {

				return 1, fmt.Errorf("namespace-new: %s", err.Error())

			} else {

				fmt.Println("namespace-new: success")

			}

		case "namespace-main":

			ns := ""
			color.Blue("RUN: origin namespace-main")

			fmt.Println("Target namespace:")
			fmt.Scanln(&ns)

			if err := setRuntimeParamNS(ns); err != nil {

				err = fmt.Errorf("namespace-main: %s", err.Error())

				fmt.Println(err.Error())

			} else {

				fmt.Println("namespace-main: success")

			}

		case "origin-repo":

			ns := ""
			repo := ""
			repo_id := ""
			repo_pw := ""

			var app_origin runfs.AppOrigin

			color.Blue("RUN: origin origin-repo")

			fmt.Println("Target namespace:")
			fmt.Scanln(&ns)
			fmt.Println("Target repo URL:")
			fmt.Scanln(&repo)
			fmt.Println("Target repo ID:")
			fmt.Scanln(&repo_id)
			fmt.Println("Target repo PW:")
			fmt.Scanln(&repo_pw)

			file_byte, err := runfs.LoadAdmOrigin()

			if err != nil {

				return 1, fmt.Errorf("origin-repo: %s", err.Error())

			}

			err = json.Unmarshal(file_byte, &app_origin)

			if err != nil {

				return 1, fmt.Errorf("origin-repo: %s", err.Error())

			}

			app_origin.REPOS = runfs.SetRepoInfo(app_origin.REPOS, repo, repo_id, repo_pw)

			err = runfs.UnloadAdmOrigin(app_origin)

			if err != nil {
				return 1, fmt.Errorf("origin-repo: %s", err.Error())
			}

		case "origin-reg":

			ns := ""
			reg := ""
			reg_id := ""
			reg_pw := ""

			var app_origin runfs.AppOrigin

			color.Blue("RUN: origin origin-reg")

			fmt.Println("Target namespace:")
			fmt.Scanln(&ns)
			fmt.Println("Target reg URL:")
			fmt.Scanln(&reg)
			fmt.Println("Target reg ID:")
			fmt.Scanln(&reg_id)
			fmt.Println("Target reg PW:")
			fmt.Scanln(&reg_pw)

			file_byte, err := runfs.LoadAdmOrigin()

			if err != nil {

				return 1, fmt.Errorf("origin-reg: %s", err.Error())

			}

			err = json.Unmarshal(file_byte, &app_origin)

			if err != nil {

				return 1, fmt.Errorf("origin-reg: %s", err.Error())

			}

			app_origin.REGS = runfs.SetRegInfo(app_origin.REGS, reg, reg_id, reg_pw)

			err = runfs.UnloadAdmOrigin(app_origin)

			if err != nil {
				return 1, fmt.Errorf("origin-reg: %s", err.Error())
			}

		case "list":

			list_all()

		case "back":

			return 0, nil

		case "trm":

			evelp = terminate()

		default:

			fmt.Println("Invalid option")
			list_all()

		}
	}

	return evelp, nil

}

func runtime_set() (int, error) {

	code := ""

	fmt.Println("TARGET : runtime")
	fmt.Scanln(&code)

	evelp := 0

	for evelp == 0 {

		switch code {

		case "ns":

			ns := ""
			color.Blue("RUN: runtime ns")

			fmt.Println("Target namespace:")
			fmt.Scanln(&ns)

			if err := setRuntimeParamNS(ns); err != nil {

				err = fmt.Errorf("ns: %s", err.Error())

				fmt.Println(err.Error())

			} else {

				fmt.Println("ns: success")

			}

		case "list":

			list_all()

		case "back":

			return 0, nil

		case "trm":

			evelp = terminate()

		default:

			fmt.Println("Invalid option")
			list_all()

		}
	}

	return evelp, nil

}
