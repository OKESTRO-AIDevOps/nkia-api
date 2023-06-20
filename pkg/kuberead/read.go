package kuberead

import (
	"fmt"

	"os/exec"
)

func ReadPod(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "pods")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadService(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadDeployment(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadNode(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "get", "nodes")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadEvent(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "events")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadResource(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "all")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}

func ReadNamespace(main_ns string) (string, error) {

	cmd := exec.Command("kubectl", "get", "namespaces")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf(": %s", err.Error())
	}

	strout := string(out)

	return strout, nil

}
