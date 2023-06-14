package kuberead

import (
	"fmt"
	admor "npia/pkg/adminorigin"

	api "npia/pkg/api"

	"os/exec"
)

func ReadPod() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadService() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadDeployment() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadNode() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadEvent() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadResource() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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

func ReadNamespace() (api.API_OUTPUT, error) {

	var api_o api.API_OUTPUT

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
