package main

import (
	"encoding/json"
	"fmt"
	admor "npia/pkg/adminorigin"
	kuberead "npia/pkg/kuberead"
	pkgutils "npia/pkg/utils"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/fatih/color"
)

func run() error {

	check_app_origin := admor.CheckAppOrigin()

	if check_app_origin == "WARNRC" {

		yn := "y"

		fmt.Println("No namespace and corresponding repositry, registry urls aren't set")
		fmt.Println("Setting them is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			fmt.Println("Abort.")

			return nil

		}

	} else if check_app_origin == "WARNRE" {

		yn := "y"

		fmt.Println("Either registry info or repository info is not set")
		fmt.Println("Setting them is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			fmt.Println("Abort.")
			return nil

		}

	} else if check_app_origin == "WARNNS" {

		yn := "y"

		fmt.Println("Target namespace is not set")
		fmt.Println("Setting it is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			fmt.Println("Abort.")

			return nil

		}

	} else if check_app_origin != "OKAY" {

		return fmt.Errorf("failed load app origin: %s", check_app_origin)

	}

	//cmd := exec.Command("docker-compose", "up", "-d", "--build", "-f", "./ADM/docker-compose.yaml")

	//cmd.Run()

	fmt.Println("Initiated")

	evelp := 0

	code := ""

	fmt.Println("For help, type [ list ]")
	fmt.Println("To terminate, type [ trm ]")

	for evelp == 0 {

		color.Green("TARGET : /*")
		fmt.Scanln(&code)

		switch code {
		case "read":

			fmt.Println("Reading cloud resoure...")

			if evelp_lower, err := read(); err != nil {

				return fmt.Errorf("read: %s", err.Error())

			} else {

				evelp = evelp_lower
			}

		case "write":

			fmt.Println("Writing cloud resource...")

			if evelp_lower, err := write(); err != nil {

				return fmt.Errorf("write: %s", err.Error())

			} else {
				evelp = evelp_lower
			}

		case "cicd":

			fmt.Println("Managing CICD process...")

			if evelp_lower, err := read(); err != nil {

				return fmt.Errorf("read: %s", err.Error())

			} else {
				evelp = evelp_lower
			}

		case "list":

			list_all()

		case "trm":

			evelp = terminate()

		default:

			fmt.Println("Invalid command")

		}
	}

	//cmd = exec.Command("docker-compose", "down", "-f", "./ADM/docker-compose.yaml")

	//cmd.Run()

	fmt.Println("npia session has been successfully terminated")

	fmt.Println("Bye")

	return nil

}

func read() (int, error) {

	code := ""

	color.Green("TARGET : /read/*")
	fmt.Scanln(&code)

	evelp := 0

	for evelp == 0 {

		switch code {

		case "pod":

			color.Blue("RUN: /read/pod")

			if api_o, err := kuberead.ReadPod(); err != nil {

				return 1, fmt.Errorf("pod: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "service":

			color.Blue("RUN: /read/service")

			if api_o, err := kuberead.ReadService(); err != nil {

				return 1, fmt.Errorf("service: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "deployment":

			color.Blue("RUN: /read/deployment")

			if api_o, err := kuberead.ReadDeployment(); err != nil {

				return 1, fmt.Errorf("deployment: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "node":

			color.Blue("RUN: /read/node")

			if api_o, err := kuberead.ReadNode(); err != nil {

				return 1, fmt.Errorf("node: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "event":

			color.Blue("RUN: /read/event")

			if api_o, err := kuberead.ReadEvent(); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "resource":

			color.Blue("RUN: /read/resource")

			if api_o, err := kuberead.ReadResource(); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "namespace":

			color.Blue("RUN: /read/namespace")

			if api_o, err := kuberead.ReadNamespace(); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "origin":

			origin_set()

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

func write() (int, error) {

	code := ""

	color.Green("TARGET : /write/*")
	fmt.Scanln(&code)

	evelp := 0

	for evelp == 0 {

		switch code {

		case "secret":

			color.Blue("RUN: /write/secret")

		case "hpa":

			color.Blue("RUN: /write/hpa")

		case "external-access":

			color.Blue("RUN: /write/external-access")

		case "internal-access":

			color.Blue("RUN: /write/internal-access")

		case "qos":

			color.Blue("RUN: /write/qos")

		case "update":

			color.Blue("RUN: /write/update")

		case "revert":

			color.Blue("RUN: /write/revert")

		case "history":

			color.Blue("RUN: /write/history")

		case "kill":

			color.Blue("RUN: /write/kill")

		case "origin":

			origin_set()

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

	return 0, nil

}

func cicd() (int, error) {

	code := ""

	color.Green("TARGET : /cicd/*")
	fmt.Scanln(&code)

	evelp := 0

	for evelp == 0 {

		switch code {

		case "build":
			color.Blue("TARGET : /cicd/build")

		case "deploy":
			color.Blue("TARGET : /cicd/deploy")

		case "pipe-start":
			color.Blue("TARGET : /cicd/pipe-start")

		case "pipe-history":
			color.Blue("TARGET : /cicd/pipe-history")

		case "git-log":
			color.Blue("TARGET : /cicd/git-log")

		case "origin":

			origin_set()

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

func qos() {

	code := ""

	fmt.Println("COMMAND : /qos/<>")
	fmt.Scanln(&code)

	if code == "highest" {

		if _, err := os.Stat("./ADM/ops_src.yaml"); err != nil {

			str_stdout := "Failed Because Ops Resource Doesn't Exist"

			fmt.Println(str_stdout)

			return
		}

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc := "deployment"

		rsc_nm := ""

		code := "highest"

		fmt.Println("Deployment to set highest priority : ")
		fmt.Scanln(&rsc_nm)

		cmd := exec.Command("python3", "./ADM/yaml_handle.py", "qos", rsc, rsc_nm, code, kcfg_path)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", "./ADM/qos.yaml")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "higher" {

		if _, err := os.Stat("./ADM/ops_src.yaml"); err != nil {

			str_stdout := "Failed Because Ops Resource Doesn't Exist"

			fmt.Println(str_stdout)

			return
		}

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc := "deployment"

		rsc_nm := ""

		code := "higher"

		fmt.Println("Deployment to set higher priority : ")
		fmt.Scanln(&rsc_nm)

		cmd := exec.Command("python3", "./ADM/yaml_handle.py", "qos", rsc, rsc_nm, code, kcfg_path)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", "./ADM/qos.yaml")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "default" {

		if _, err := os.Stat("./ADM/ops_src.yaml"); err != nil {

			str_stdout := "Failed Because Ops Resource Doesn't Exist"

			fmt.Println(str_stdout)

			return
		}

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc := "deployment"

		rsc_nm := ""

		fmt.Println("Deployment to set default priority : ")
		fmt.Scanln(&rsc_nm)

		cmd := exec.Command("python3", "./ADM/yaml_handle.py", "qosundo", rsc, rsc_nm)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", "./ADM/qos.yaml")

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

func lifecycle() {

	code := ""

	fmt.Println("COMMAND : /lifecycle/<>")
	fmt.Scanln(&code)

	if code == "update" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc_rscnm := ""

		fmt.Println("Deployment to update (or restart) : ")
		fmt.Scanln(&rsc_rscnm)

		rsc_rscnm = "deployment/" + rsc_rscnm

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "rollout", "restart", rsc_rscnm)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "revert" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc_rscnm := ""

		fmt.Println("Deployment to revert to previous version : ")
		fmt.Scanln(&rsc_rscnm)

		rsc_rscnm = "deployment/" + rsc_rscnm

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "rollout", "undo", rsc_rscnm)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "history" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc_rscnm := ""

		fmt.Println("Deployment to get history of : ")
		fmt.Scanln(&rsc_rscnm)

		rsc_rscnm = "deployment/" + rsc_rscnm

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "rollout", "history", rsc_rscnm)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "kill" {

		if _, err := os.Stat("./ADM/ops_src.yaml"); err != nil {

			str_stdout := "Failed Because Ops Resource Doesn't Exist"

			fmt.Println(str_stdout)

			return
		}

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc := "deployment"

		rsc_nm := ""

		fmt.Println("Deployment to delete: ")
		fmt.Scanln(&rsc_nm)

		cmd := exec.Command("python3", "./ADM/yaml_handle.py", "kill", rsc, rsc_nm)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "delete", "-f", "./ADM/delete_ops_src.yaml")

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

func list_all() {
	fmt.Println("*COMMAND LIST*")
	fmt.Println("[ /read/pod ] : gets pods in a namespace")
	fmt.Println("[ /read/service ] : gets services in a namespace")
	fmt.Println("[ /read/deployment ] : gets deployments in a namespace")
	fmt.Println("[ /read/node ] : gets all nodes of the target cluster")
	fmt.Println("[ /read/event ] : gets all events in a namespace")
	fmt.Println("[ /read/resource ] : gets all resources in a namespace")
	fmt.Println("[ /read/namespace ] : gets all namespaces available of the target cluster")
	fmt.Println("[ /write/secret ] : sets cluster secret based on origin info")
	fmt.Println("[ /write/hpa ] : deploys HorizontalPodAutoscaler of a deployment in a namespace")
	fmt.Println("[ /write/external-access ] : deploys ingress of a service in a namespace")
	fmt.Println("[ /write/internal-access ] : deploys nodeport of a service in a namespace")
	fmt.Println("[ /write/qos ]: modifies a deployment's QoS policy in a namespace to Burstable")
	fmt.Println("[ /write/update ]: updates (or restart) a deployment in a namespace")
	fmt.Println("[ /write/revert ]: reverts a deployment in a namespace to a previous status")
	fmt.Println("[ /write/history ]: gets revision history of a deployment in a namespace")
	fmt.Println("[ /write/kill ]: deletes a deployment in a namespace and a corresponding service")
	fmt.Println("[ /origin ] : sets up origin file ")
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

func origin() {

	fmt.Println("Initiated")

	evelp := 0

	code := ""

	fmt.Println("For help, type [ list ]")
	fmt.Println("To terminate, type [ trm ]")

	for evelp == 0 {

		fmt.Println("COMMAND : /<>")
		fmt.Scanln(&code)

		if code == "set" {

			fmt.Println("Setting cluster origin...")

			origin_set()

		} else if code == "run" {

			fmt.Println("Starting nopainctl ...")

			run()

		} else if code == "list" {

			origin_list_all()

		} else if code == "trm" {

			evelp = terminate()

		} else {

			fmt.Println("Invalid command")

		}

	}

	//cmd = exec.Command("docker-compose", "down", "-f", "./ADM/docker-compose.yaml")

	//cmd.Run()

	fmt.Println("nopainctl session has been successfully terminated")

	fmt.Println("Bye")

}

func origin_set() {

	code := ""

	fmt.Println("COMMAND : /set/<>")
	fmt.Scanln(&code)

	if code == "kubeconfig-path" {

		kcfg_path_in := ""

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		fmt.Println("Kube config file path (must be absolute path including the file name) : ")

		fmt.Scanln(&kcfg_path_in)

		root_idx := strings.Index(kcfg_path_in, "/")

		if root_idx != 0 {

			fmt.Println("Must be absolute path")

			return

		}

		if _, err := os.Stat(kcfg_path_in); err != nil {

			fmt.Println("Unable to find kube config file at the specified location")

			return

		}

		app_origin.KCFG_PATH = kcfg_path_in

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Kube config path has been set")

	} else if code == "namespace-new" {

		kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

		ns := ""
		repo_url_in := ""
		reg_url_in := ""

		var app_origin admor.AppOrigin

		fmt.Println("New namespace to create : ")
		fmt.Scanln(&ns)

		fmt.Println("Repository url to be used in this namespace : ")
		fmt.Scanln(&repo_url_in)

		fmt.Println("Registry url to be used in this namespace : ")
		fmt.Scanln(&reg_url_in)

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url, reg_url := admor.GetRecordInfo(app_origin.RECORDS, ns)

		if repo_url != "N" || reg_url != "N" {

			yorn := "y"

			fmt.Println("Associated repository or registry information already exists")

			fmt.Println("Further action will overwrite the previous information")

			fmt.Println("Are you sure want to proceed? [ y | n ]")

			fmt.Scanln(&yorn)

			if yorn == "n" {

				return

			}

		}

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "create", "namespace", ns)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return
		}

		strout := string(out)

		fmt.Println(strout)

		app_origin.RECORDS = admor.SetRecordInfo(app_origin.RECORDS, ns, repo_url_in, reg_url_in)

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Namespace record has been successfully set")

	} else if code == "namespace-main" {

		kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

		new_main_ns := ""

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		available_list := []string{}

		for i := 0; i < len(app_origin.RECORDS); i++ {

			available_list = append(available_list, app_origin.RECORDS[i].NS)

		}

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "get", "namespace", "--no-headers", "-o", "custom-columns=:metadata.name")

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

			hit := pkgutils.CheckIfEleInStrList(a, available_list)

			if !hit {

				continue

			}

			str_ready_list = append(str_ready_list, a)

			fmt.Println(a)

		}

		fmt.Println("Choose from the above : ")

		fmt.Scanln(&new_main_ns)

		hit := pkgutils.CheckIfEleInStrList(new_main_ns, str_ready_list)

		if !hit {

			fmt.Println("Not an available namespace")

			return

		}

		app_origin.MAIN_NS = new_main_ns

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Target namespace has been set")

	} else if code == "origin-repo" {

		var app_origin admor.AppOrigin

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

		app_origin.REPOS = admor.SetRepoInfo(app_origin.REPOS, repo_url, repo_id, repo_pw)

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Repository info has been set")

	} else if code == "origin-reg" {

		var app_origin admor.AppOrigin

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

		app_origin.REGS = admor.SetRegInfo(app_origin.REGS, reg_url, reg_id, reg_pw)

		app_origin_bytes, _ := json.Marshal(app_origin)

		_ = os.WriteFile("./ADM/origin.json", app_origin_bytes, 0644)

		fmt.Println("Registry info has been set")

	} else if code == "list" {

		origin_list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

func origin_list_all() {

	fmt.Println("*COMMAND LIST*")
	fmt.Println("[ /set/kubeconfig-path ] : sets the kubeconfig file path, must be absolute path")
	fmt.Println("[ /set/namespace-main ] : uses a namespace")
	fmt.Println("[ /set/namespace-new ] : creates a namespace")
	fmt.Println("[ /set/origin-repo ] : sets repository info")
	fmt.Println("[ /set/origin-reg ] : sets registry info")
	fmt.Println("[ /run ] : starts nopainctl ")
	fmt.Println("[ /*/back ] : steps back to the previous stage")
	fmt.Println("[ /list, /*/list ] : lists all available commands")
	fmt.Println("[ /trm ] : ends nopainctl session")

}

func main() {

	currentUser, err := user.Current()

	if err != nil {

		strerr := err.Error()

		fmt.Println(strerr)

		return

	}

	if currentUser.Username != "root" {

		fmt.Println("You're not running this process as root")
		fmt.Println("Use [ sudo ./npia ] instead")

		return

	}

	if err := run(); err != nil {

		_ = fmt.Errorf("Error: %s", err.Error())

	}

}
