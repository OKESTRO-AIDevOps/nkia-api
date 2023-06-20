package kubewrite

import (
	"fmt"

	"github.com/seantywork/x0f_npia/pkg/dotfs"

	"os/exec"

	"encoding/json"
)

func WriteSecret(main_ns string) (string, error) {

	var app_origin dotfs.AppOrigin

	file_content, _ := dotfs.LoadAdmOrigin()

	err := json.Unmarshal(file_content, &app_origin)

	if err != nil {
		return "", fmt.Errorf(": %s", err.Error())
	}

	_, reg_url := dotfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if reg_url == "N" {

		return "", fmt.Errorf(": %s", "reg url not set")

	}

	reg_id, reg_pw := dotfs.GetRegInfo(app_origin.REGS, reg_url)

	if reg_id == "N" || reg_pw == "N" {

		return "", fmt.Errorf(": %s", "reg info not complete")

	}

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "secret", "docker-secret", "--no-headers", "-o", "custom-columns=:metadata.name")

	_, err = cmd.Output()

	docker_server := "--docker-server="

	docker_username := "--docker-username="

	docker_password := "--docker-password="

	docker_server += reg_url

	docker_username += reg_id

	docker_password += reg_pw

	if err != nil {

		cmd = exec.Command("kubectl", "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

		out, err := cmd.Output()

		if err != nil {
			return "", fmt.Errorf(": %s", err.Error())
		}

		strout := string(out)

		return strout, nil

	} else {

		cmd = exec.Command("kubectl", "-n", main_ns, "delete", "secret", "docker-secret")

		_, err = cmd.Output()

		if err != nil {
			return "", fmt.Errorf(": %s", err.Error())
		}

		cmd = exec.Command("kubectl", "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

		out, err := cmd.Output()

		if err != nil {
			return "", fmt.Errorf(": %s", err.Error())
		}

		strout := string(out)

		return strout, nil

	}

}

/*
   if code == "secret" {

   	var app_origin dotfs.AppOrigin

   	adm_origin_json := LIBIF.GetLibComponentPath(".etc", "ADM_origin.json")

   	file_content, _ := os.ReadFile(adm_origin_json)

   	_ = json.Unmarshal(file_content, &app_origin)

   	kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

   	_, reg_url := dotfs.GetRecordInfo(app_origin.RECORDS, main_ns)

   	if reg_url == "N" {

   		fmt.Println("Registry information has not been set as a record")
   		fmt.Println("Aborting")

   		return

   	}

   	reg_id, reg_pw := dotfs.GetRegInfo(app_origin.REGS, reg_url)

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

   	kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

   	kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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
*/

/*
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

		kcfg_path, main_ns, err := dotfs.GetKubeConfigAndTargetNameSpace()

		}

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		//list_all()

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		kcfg_path, main_ns := dotfs.GetKubeConfigAndTargetNameSpace()

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

		//list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}

}

*/
