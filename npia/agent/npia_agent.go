package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	admor "npia/pkg/adminorigin"
	. "npia/pkg/libinterface"
	pkgutils "npia/pkg/utils"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func run() {

	check_app_origin := admor.CheckAppOrigin()

	if check_app_origin == "WARNRC" {

		yn := "y"

		fmt.Println("No namespace and corresponding repositry, registry urls aren't set")
		fmt.Println("Setting them is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			return

		}

	} else if check_app_origin == "WARNRE" {

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

	//cmd := exec.Command("docker-compose", "up", "-d", "--build", "-f", "./ADM/docker-compose.yaml")

	//cmd.Run()

	fmt.Println("Initiated")

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

		} else if code == "origin" {

			fmt.Println("Setting cluster origin...")

			origin()

		} else if code == "list" {

			list_all()

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

func check() {

	code := ""

	fmt.Println("COMMAND : /check/<>")
	fmt.Scanln(&code)

	if code == "pod" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

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

	if code == "secret" {

		var app_origin admor.AppOrigin

		adm_origin_json := LIBIF.GetLibComponentPath(".etc", "ADM_origin.json")

		file_content, _ := os.ReadFile(adm_origin_json)

		_ = json.Unmarshal(file_content, &app_origin)

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		_, reg_url := admor.GetRecordInfo(app_origin.RECORDS, main_ns)

		if reg_url == "N" {

			fmt.Println("Registry information has not been set as a record")
			fmt.Println("Aborting")

			return

		}

		reg_id, reg_pw := admor.GetRegInfo(app_origin.REGS, reg_url)

		if reg_id == "N" || reg_pw == "N" {

			fmt.Println("Registry information is not complete")
			fmt.Println("Aborting")

			return

		}

		cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "secret", "docker-secret", "--no-headers", "-o", "custom-columns=:metadata.name")

		_, err := cmd.Output()

		docker_server := "--docker-server="

		docker_username := "--docker-username="

		docker_password := "--docker-password="

		docker_server += reg_url

		docker_username += reg_id

		docker_password += reg_pw

		if err != nil {

			fmt.Println("No Pre-existing secret")

			cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

			cmd.Run()

			fmt.Println("Secret has been successfully created")

		} else {

			yorn := "y"

			fmt.Println("Secret already exists")
			fmt.Println("Further action will overwrite the existing secret in the namespace")
			fmt.Println("-----")
			fmt.Println(main_ns)
			fmt.Println("-----")
			fmt.Println("Are you sure you want to proceed? [ y | n ] ")

			fmt.Scanln(&yorn)

			if yorn == "n" {

				return

			}

			cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "delete", "secret", "docker-secret")

			cmd.Run()

			cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

			cmd.Run()

			fmt.Println("Secret has been successfully modified")

		}

	} else if code == "hpa" {

		ops_src_yaml := LIBIF.GetLibComponentPath(".usr", "ops_src.yaml")

		if _, err := os.Stat(ops_src_yaml); err != nil {

			str_stdout := "Failed Because Ops Resource Doesn't Exist"

			fmt.Println(str_stdout)

			return
		}

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		rsc := "deployment"

		rsc_nm := ""

		fmt.Println("Deployment to autoscale : ")
		fmt.Scanln(&rsc_nm)

		yaml_handle_py := LIBIF.GetLibComponentPath("scripts", "yaml_handle.py")

		hpa_yaml := LIBIF.GetLibComponentPath(".usr", "hpa.yaml")

		cmd := exec.Command("python3", yaml_handle_py, "hpa", rsc, rsc_nm, kcfg_path)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", hpa_yaml)

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "external-access" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		host_nm := ""

		svc_nm := ""

		fmt.Println("Domain name you want to use : ")
		fmt.Scanln(host_nm)
		fmt.Println("Upstream network name you want to connect to : ")
		fmt.Scanln(svc_nm)

		yaml_handle_py := LIBIF.GetLibComponentPath("scripts", "yaml_handle.py")

		ingress_yaml := LIBIF.GetLibComponentPath(".usr", "ingress.yaml")

		cmd := exec.Command("python3", yaml_handle_py, "ingr", main_ns, host_nm, svc_nm, kcfg_path)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", ingress_yaml)

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

func cicd() {

	code := ""

	fmt.Println("COMMAND : /cicd/<>")
	fmt.Scanln(&code)

	if code == "build" {

		_, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url, reg_url := admor.GetRecordInfo(app_origin.RECORDS, main_ns)

		if repo_url == "N" || reg_url == "N" {

			fmt.Println("Associated repository or registry information doesn't exist as a record")

			return

		}

		repo_id, repo_pw := admor.GetRepoInfo(app_origin.REPOS, repo_url)

		reg_id, reg_pw := admor.GetRegInfo(app_origin.REGS, reg_url)

		if repo_id == "N" || repo_pw == "N" {

			fmt.Println("Associated repository credential is imperfect")

			return

		}

		if reg_id == "N" || reg_pw == "N" {

			fmt.Println("Associated registry credential is imperfect")

			return

		}

		if _, err := os.Stat("./ADM/target"); err == nil {

			cmd := exec.Command("rm", "-r", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			cmd = exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()
		} else {
			cmd := exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

		}

		os.Chdir("./ADM/target")

		cmd := exec.Command("git", "init")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		insert := "%s:%s@"

		prt_idx := strings.Index(repo_url, "://")

		prt_idx += 3

		repo_url = repo_url[:prt_idx] + insert + repo_url[prt_idx:]

		repo_url = fmt.Sprintf(repo_url, repo_id, repo_pw)

		cmd = exec.Command("git", "pull", repo_url)

		cmd.Run()

		dcfile_nm := ""

		if _, err := os.Stat("docker-compose.yaml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yaml"

		}

		if _, err := os.Stat("docker-compose.yml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yml"

		}

		if dcfile_nm == "" {

			os.Chdir("../../")

			str_stdout := "File [ docker-compose.yaml | docker-compose.yml ] Does Not Exist"

			fmt.Println(str_stdout)

			return

		}

		cmd = exec.Command("docker-compose", "up", "-d", "--build")

		bld_out, _ := cmd.StdoutPipe()

		bld_err, _ := cmd.StderrPipe()

		go func() {
			merged := io.MultiReader(bld_out, bld_err)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println(msg)
			}
		}()

		cmd.Run()

		os.Chdir("../../")

		cmd = exec.Command("/bin/cp", "-rf", dcfile_nm, "./ADM/docker-compose.yaml")

		cmd.Run()

		cmd = exec.Command("python3", "./ADM/yaml_handle.py", "push", reg_url, reg_id, reg_pw)

		bld_out, _ = cmd.StdoutPipe()

		bld_err, _ = cmd.StderrPipe()

		go func() {
			merged := io.MultiReader(bld_out, bld_err)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println(msg)
			}
		}()

		cmd.Run()

		cmd = exec.Command("rm", "-r", "./ADM/target")

		cmd.Run()

	} else if code == "deploy" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url, reg_url := admor.GetRecordInfo(app_origin.RECORDS, main_ns)

		if repo_url == "N" || reg_url == "N" {

			fmt.Println("Associated repository or registry information doesn't exist as a record")

			return

		}

		repo_id, repo_pw := admor.GetRepoInfo(app_origin.REPOS, repo_url)

		reg_id, reg_pw := admor.GetRegInfo(app_origin.REGS, reg_url)

		if repo_id == "N" || repo_pw == "N" {

			fmt.Println("Associated repository credential is imperfect")

			return

		}

		if reg_id == "N" || reg_pw == "N" {

			fmt.Println("Associated registry credential is imperfect")

			return

		}

		if _, err := os.Stat("./ADM/target"); err == nil {

			cmd := exec.Command("rm", "-r", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			cmd = exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()
		} else {
			cmd := exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

		}

		os.Chdir("./ADM/target")

		cmd := exec.Command("git", "init")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		insert := "%s:%s@"

		prt_idx := strings.Index(repo_url, "://")

		prt_idx += 3

		repo_url = repo_url[:prt_idx] + insert + repo_url[prt_idx:]

		repo_url = fmt.Sprintf(repo_url, repo_id, repo_pw)

		cmd = exec.Command("git", "pull", repo_url)

		cmd.Run()

		dcfile_nm := ""

		if _, err := os.Stat("docker-compose.yaml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yaml"

		}

		if _, err := os.Stat("docker-compose.yml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yml"

		}

		if dcfile_nm == "" {

			os.Chdir("../../")

			str_stdout := "File [ docker-compose.yaml | docker-compose.yml ] Does Not Exist"

			fmt.Println(str_stdout)

			return

		}

		os.Chdir("../../")

		cmd = exec.Command("/bin/cp", "-rf", dcfile_nm, "./ADM/docker-compose.yaml")

		cmd.Run()

		cmd = exec.Command("python3", "./ADM/yaml_handle.py", "deploy", reg_url, reg_id, reg_pw)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", "./ADM/ops_src.yaml")

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
	fmt.Println("[ /check/pod ] : gets pods in a namespace")
	fmt.Println("[ /check/net ] : gets services in a namespace")
	fmt.Println("[ /check/dpl ] : gets deployments in a namespace")
	fmt.Println("[ /check/node ] : gets all nodes of the target cluster")
	fmt.Println("[ /check/event ] : gets all events in a namespace")
	fmt.Println("[ /check/resource ] : gets all resources in a namespace")
	fmt.Println("[ /check/namespace ] : gets all namespaces available of the target cluster")
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
		fmt.Println("Use [ sudo ./nopainctl ] instead")

		return

	}

	arg_len := len(os.Args)

	if arg_len < 2 {

		fmt.Println("No argument")
		fmt.Println("There are only two options you can choose from")

		fmt.Println("1. [ sudo ./nopainctl run ] : starts nopainctl program")
		fmt.Println("2. [ sudo ./nopainctl origin] : sets up origin file")

		return

	} else if arg_len > 2 {

		fmt.Println("Too many arguments")
		fmt.Println("There are only two options you can choose from")

		fmt.Println("1. [ sudo ./nopainctl run ] : starts nopainctl program")
		fmt.Println("2. [ sudo ./nopainctl origin] : sets up origin file")

		return

	}

	action := os.Args[1]

	if action == "run" {

		run()

	} else if action == "origin" {

		origin()

	} else {

		fmt.Println("Wrong argument")
		fmt.Println("There are only two options you can choose from")

		fmt.Println("1. [ sudo ./nopainctl run ] : starts nopainctl program")
		fmt.Println("2. [ sudo ./nopainctl origin] : sets up origin file")

		return

	}

}
