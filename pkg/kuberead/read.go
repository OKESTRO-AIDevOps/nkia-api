package kuberead

import (
	"fmt"

	apist "github.com/seantywork/x0f_npia/pkg/apistandard"

	"os/exec"
)

func ReadPod(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "pods")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadService(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	fmt.Println(strout)

	api_o.BODY = strout

	return api_o, nil

}

func ReadDeployment(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadNode(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "get", "nodes")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadEvent(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "events")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadResource(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "all")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}

func ReadNamespace(main_ns string) (apist.API_OUTPUT, error) {

	var api_o apist.API_OUTPUT

	cmd := exec.Command("kubectl", "get", "namespaces")

	out, err := cmd.Output()

	if err != nil {

		return api_o, fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	api_o.BODY = strout

	return api_o, nil

}
