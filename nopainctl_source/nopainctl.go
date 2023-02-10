package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type appOrigin struct {
	KCFG_PATH string
	MAIN_NS   string
	REPOS     []repoInfo
	REGS      []regInfo
}

type repoInfo struct {
	REPO_ADDR string
	REPO_ID   string
	REPO_PW   string
}

type regInfo struct {
	REG_ADDR string
	REG_ID   string
	REG_PW   string
}

func getRepoInfo(repos []repoInfo, addr string) (string, string) {

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

func setRepoInfo(repos []repoInfo, addr string, id string, pw string) []repoInfo {

	exists := 0

	arr_leng := len(repos)

	repo_id := id

	repo_pw := pw

	var new_repo_info repoInfo

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

func getRegInfo(regs []regInfo, addr string) (string, string) {

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

func setRegInfo(regs []regInfo, addr string, id string, pw string) []regInfo {

	exists := 0

	arr_leng := len(regs)

	reg_id := id

	reg_pw := pw

	var new_reg_info regInfo

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

func checkAppOrigin() string {

	var app_origin appOrigin

	file_content, err := os.ReadFile("./ADM/origin.json")

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

	if len(app_origin.REGS) == 0 || len(app_origin.REPOS) == 0 {

		return "WARNRE"

	}

	if app_origin.MAIN_NS == "" {

		return "WARNNS"
	}

	return "OKAY"
}

func getBoth() (string, string) {

	var app_origin appOrigin

	file_content, _ := os.ReadFile("./ADM/origin.json")

	_ = json.Unmarshal(file_content, &app_origin)

	kcfg_path := app_origin.KCFG_PATH

	main_ns := app_origin.MAIN_NS

	return kcfg_path, main_ns

}

func checkIfAInStrList(a string, str_list []string) bool {

	hit := false

	for i := 0; i < len(str_list); i++ {

		if str_list[i] == a {

			hit = true

			return hit
		}

	}

	return hit

}

func run() {

	check_app_origin := checkAppOrigin()

	if check_app_origin == "WARNRE" {

		yn := "y"

		fmt.Println("Either registry info or repository info is not set")
		fmt.Println("Setting them is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			return

		}

	} else if check_app_origin == "WARNNS" {

		yn := "y"

		fmt.Println("Target namespace is not set")
		fmt.Println("Setting it is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			return

		}

	} else if check_app_origin != "OKAY" {

		fmt.Println(check_app_origin)

		return

	}

	fmt.Println("Initiaing.....")

	cmd := exec.Command("docker-compose", "up", "-d", "--build", "-f", "./ADM/docker-compose.yaml")

	cmd.Run()

	fmt.Println("Initiation Completed")

	evelp := 0

	code := ""

	fmt.Println("For help, type [ list ]")
	fmt.Println("To terminate, type [ trm ]")

	for evelp == 0 {

		fmt.Println("COMMAND : /<>")
		fmt.Scanln(&code)

		if code == "check" {

			fmt.Println("Checking cloud resoure...")

			check()

		} else if code == "set" {

			fmt.Println("Setting cloud resource...")

			set()

		} else if code == "cicd" {

			fmt.Println("Managing CICD process...")

			cicd()

		} else if code == "qos" {

			fmt.Println("Managing QOS setting...")

			qos()

		} else if code == "lifecycle" {

			fmt.Println("Managing application lifecycle...")

			lifecycle()

		} else if code == "list" {

			list_all()

		} else if code == "trm" {

			evelp = terminate()

		} else {

			fmt.Println("Invalid command")

		}

	}

	cmd = exec.Command("docker-compose", "down", "-f", "./ADM/docker-compose.yaml")

	cmd.Run()

	fmt.Println("nopainctl session has been successfully terminated")

	fmt.Println("Bye")

}

func check() {

	code := ""

	fmt.Println("COMMAND : /check/<>")
	fmt.Scanln(&code)

	if code == "pod" {

		kcfg_path, main_ns := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "pods")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "net" {

		kcfg_path, main_ns := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "services")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "dpl" {

		kcfg_path, main_ns := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "deployments")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "node" {

		kcfg_path, _ := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "get", "nodes")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "event" {

		kcfg_path, main_ns := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "events")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "resource" {

		kcfg_path, main_ns := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "all")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "namespace" {

		kcfg_path, _ := getBoth()

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "get", "namespaces")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func set() {

	code := ""

	fmt.Println("COMMAND : /set/<>")
	fmt.Scanln(&code)

	if code == "namespace-new" {

		ns := ""

		fmt.Println("New namespace to create : ")
		fmt.Scanln(&ns)

		cmd := exec.Command("kubectl", "create", "namespace", ns)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "namesapce-main" {

		new_main_ns := ""

		var app_origin appOrigin

		black_list := []string{"kube-node-lease", "kube-public", "kube-system", "local-path-storage", "", " "}

		cmd := exec.Command("kubectl", "get", "namespace", "--no-headers", "-o", "custom-columns=:metadata.name")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		fmt.Println("Available namespaces are ----- ")

		strout := string(out)

		strout_list := strings.Split(strout, "\n")

		str_ready_list := []string{}

		for i := 0; i < len(strout_list); i++ {

			a := strout_list[i]

			hit := checkIfAInStrList(a, black_list)

			if hit {

				continue

			}

			str_ready_list = append(str_ready_list, a)

			fmt.Println(a)

		}

		fmt.Println("Choose from the above : ")

		fmt.Scanln(&new_main_ns)

		hit := checkIfAInStrList(new_main_ns, str_ready_list)

		if !hit {

			fmt.Println("Not an available namespace")

			return

		}

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		app_origin.MAIN_NS = new_main_ns

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Target namespace has been set")

	} else if code == "origin-repo" {

		var app_origin appOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url := ""

		repo_id := ""

		repo_pw := ""

		fmt.Println("Type repository url : ")

		fmt.Scanln(&repo_url)

		fmt.Println("Type repository id : ")

		fmt.Scanln(&repo_id)

		fmt.Println("Type repository password: ")

		fmt.Scanln(&repo_pw)

		app_origin.REPOS = setRepoInfo(app_origin.REPOS, repo_url, repo_id, repo_pw)

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Repository info has been set")

	} else if code == "origin-reg" {

		var app_origin appOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		reg_url := ""

		reg_id := ""

		reg_pw := ""

		fmt.Println("Type registry url : ")

		fmt.Scanln(&reg_url)

		fmt.Println("Type registry id : ")

		fmt.Scanln(&reg_id)

		fmt.Println("Type registry password: ")

		fmt.Scanln(&reg_pw)

		app_origin.REGS = setRegInfo(app_origin.REGS, reg_url, reg_id, reg_pw)

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Registry info has been set")

	} else if code == "secret" {

		_, main_ns := getBoth()

		cmd := exec.Command("kubectl", "get", "secret", "-n", main_ns, "--no-headers", "-o", "custom-columns=:metadata.name")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		strout := string(out)

		if strout == "\n" || strout == "" {

			fmt.Println("No Pre-existing secret")

		} else {

			fmt.Println("Secret already exists")
			fmt.Println("Further action will overwrite the existing secret")

		}

		fmt.Println("Type the url of the registry of which you want ")

	} else if code == "hpa" {

	} else if code == "external-access" {

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func cicd() {

	code := ""

	fmt.Println("COMMAND : /cicd/<>")
	fmt.Scanln(&code)

	if code == "build" {

	} else if code == "deploy" {

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func qos() {

	code := ""

	fmt.Println("COMMAND : /qos/<>")
	fmt.Scanln(&code)

	if code == "highest" {

	} else if code == "higher" {

	} else if code == "default" {

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func lifecycle() {

	code := ""

	fmt.Println("COMMAND : /lifecycle/<>")
	fmt.Scanln(&code)

	if code == "update" {

	} else if code == "revert" {

	} else if code == "history" {

	} else if code == "kill" {

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func list_all() {
	fmt.Println("*COMMAND LIST*")
	fmt.Println("[ /check/pod ] : gets pods in a namespace")
	fmt.Println("[ /check/net ] : gets services in a namespace")
	fmt.Println("[ /check/dpl ] : gets deployments in a namespace")
	fmt.Println("[ /check/node ] : gets all nodes of the target cluster")
	fmt.Println("[ /check/event ] : gets all events in a namespace")
	fmt.Println("[ /check/resource ] : gets all resources in a namespace")
	fmt.Println("[ /check/namespace ] : gets all namespaces available of the target cluster")
	fmt.Println("[ /set/namespace-new ] : creates a namespace")
	fmt.Println("[ /set/namespace-main ] : uses a namespace")
	fmt.Println("[ /set/origin-repo ] : sets repository info")
	fmt.Println("[ /set/origin-reg ] : sets registry info")
	fmt.Println("[ /set/secret ] : sets cluster secret based on origin info")
	fmt.Println("[ /set/hpa ] : deploys HorizontalPodAutoscaler of a deployment in a namespace")
	fmt.Println("[ /set/external-access ] : deploys ingress of a service in a namespace")
	fmt.Println("[ /cicd/build ] : deploys the apps predefined in a user's source code repository")
	fmt.Println("[ /cicd/deploy ] : deploys the apps predefined in a user's source code repository")
	fmt.Println("[ /qos/highest ] : modifies a deployment's QoS policy in a namespace to Guaranteed")
	fmt.Println("[ /qos/higher ] : modifies a deployment's QoS policy in a namespace to Burstable")
	fmt.Println("[ /qos/default ] : modifies a deployment's QoS policy in a namespace to Best-Effort")
	fmt.Println("[ /lifecycle/update ] : updates (or restart) a deployment in a namespace")
	fmt.Println("[ /lifecycle/revert] : reverts a deployment in a namespace to a previous status")
	fmt.Println("[ /lifecycle/history] : gets revision history of a deployment in a namespace")
	fmt.Println("[ /lifecycle/kill ] : deletes a deployment in a namespace and a corresponding service")
	fmt.Println("[ /*/back ] : steps back to the previous stage")
	fmt.Println("[ /list, /*/list ] : lists all available commands")
	fmt.Println("[ /trm ] : ends nopainctl session")
}

func terminate() int {

	yn := "n"

	fmt.Println("Are you sure you want to quit? [ y | n ]")

	fmt.Scanln(&yn)

	if yn == "y" {

		return 1

	}

	return 0

}

func main() {

	action := os.Args[1]

	if action == "run" {

		run()

	} else if action == "set" {

		set()

	} else {

		fmt.Println("Option Not Available")
	}

}
