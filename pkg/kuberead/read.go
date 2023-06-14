package kuberead

import (
	"fmt"

	admor "github.com/seantywork/x0f_npia/pkg/adminorigin"

	ioman "github.com/seantywork/x0f_npia/pkg/iomanager"

	"os/exec"
)

func ReadPod() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "pods")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadService() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	fmt.Println(strout)

	api_o.BODY = strout

	return api_o, nil

}

func ReadDeployment() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadNode() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "get", "nodes")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadEvent() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "events")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadResource() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "get", "all")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadNamespace() (ioman.API_OUTPUT, error) {

	var api_o ioman.API_OUTPUT

	kcfg_path, _ := admor.GetKubeConfigAndTargetNameSpace()

	cmd := exec.Command("kubectl", "--kubeconfig", kcfg_path, "get", "namespaces")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}
