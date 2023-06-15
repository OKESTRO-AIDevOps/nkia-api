package main

import (
	"fmt"
	"os/user"

	"github.com/seantywork/x0f_npia/pkg/dotfs"
	"github.com/seantywork/x0f_npia/pkg/kuberead"
	"github.com/seantywork/x0f_npia/pkg/kubewrite"

	"github.com/fatih/color"
)

func run() error {

	check_app_origin, err := dotfs.CheckAppOrigin()

	if err != nil {

		return fmt.Errorf("run failed: %s", err.Error())

	}

	if check_app_origin == "WARNRC" {

		yn := "y"

		fmt.Println("No namespace and corresponding repository, registry urls aren't set")
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

	} else if check_app_origin != "OKAY" {

		return fmt.Errorf("failed load app origin: %s", check_app_origin)

	}

	err = checkRuntimeParamNS()

	if err != nil {

		yn := "y"

		fmt.Println("Runtime target namespace is not set")
		fmt.Println("Setting them is possible in later stages")
		fmt.Println("Are you sure you want to proceed? [ y | n ]")

		fmt.Scanln(&yn)

		if yn == "n" {

			fmt.Println("Abort.")
			return nil

		}

	}

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

		case "origin":

			evelp, err = origin_set()

			if err != nil {

				return fmt.Errorf("origin: %s", err.Error())

			}

		case "runtime":

			evelp, err = runtime_set()

			if err != nil {

				return fmt.Errorf("runtime: %s", err.Error())

			}

		case "list":

			list_all()

		case "trm":

			evelp = terminate()

		default:

			fmt.Println("Invalid command")

		}
	}

	fmt.Println("npia session has been successfully terminated")

	fmt.Println("Bye")

	return nil

}

func read() (int, error) {

	code := ""

	color.Green("TARGET : /read/*")
	fmt.Scanln(&code)

	evelp := 0

	var err error

	for evelp == 0 {

		switch code {

		case "pod":

			color.Blue("RUN: /read/pod")

			if api_o, err := kuberead.ReadPod(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("pod: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "service":

			color.Blue("RUN: /read/service")

			if api_o, err := kuberead.ReadService(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("service: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "deployment":

			color.Blue("RUN: /read/deployment")

			if api_o, err := kuberead.ReadDeployment(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("deployment: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "node":

			color.Blue("RUN: /read/node")

			if api_o, err := kuberead.ReadNode(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("node: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "event":

			color.Blue("RUN: /read/event")

			if api_o, err := kuberead.ReadEvent(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "resource":

			color.Blue("RUN: /read/resource")

			if api_o, err := kuberead.ReadResource(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "namespace":

			color.Blue("RUN: /read/namespace")

			if api_o, err := kuberead.ReadNamespace(RPARAM["NS"]); err != nil {

				return 1, fmt.Errorf("event: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

		case "origin":

			evelp, err = origin_set()

			if err != nil {

				return 1, fmt.Errorf("origin: %s", err.Error())

			}

		case "runtime":

			evelp, err = runtime_set()

			if err != nil {

				return 1, fmt.Errorf("runtime: %s", err.Error())

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

func write() (int, error) {

	code := ""

	color.Green("TARGET : /write/*")
	fmt.Scanln(&code)

	evelp := 0

	var err error

	for evelp == 0 {

		switch code {

		case "secret":

			color.Blue("RUN: /write/secret")

			permit_overwrite := 1

			check := "y"

			fmt.Println("Permit overwrite if secret exists? [ y | n ] :")
			fmt.Scanln(&check)

			if check == "n" {
				permit_overwrite = 0
			}

			if api_o, err := kubewrite.WriteSecret(RPARAM["NS"], permit_overwrite); err != nil {

				return 1, fmt.Errorf("secret: %s", err.Error())

			} else {

				fmt.Println(api_o.BODY)

			}

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

			evelp, err = origin_set()

			if err != nil {

				return 1, fmt.Errorf("origin: %s", err.Error())

			}

		case "runtime":

			evelp, err = runtime_set()

			if err != nil {

				return 1, fmt.Errorf("runtime: %s", err.Error())

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

	return 0, nil

}

func cicd() (int, error) {

	code := ""

	color.Green("TARGET : /cicd/*")
	fmt.Scanln(&code)

	evelp := 0

	var err error

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

			evelp, err = origin_set()

			if err != nil {

				return 1, fmt.Errorf("origin: %s", err.Error())

			}

		case "runtime":

			evelp, err = runtime_set()

			if err != nil {

				return 1, fmt.Errorf("runtime: %s", err.Error())

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
	fmt.Println("[ origin ] : sets up origin file ")
	fmt.Println("[ runtime ] : sets up runtime parameters")
	fmt.Println("[ back ] : steps back to the previous stage")
	fmt.Println("[ list ] : lists all available commands")
	fmt.Println("[ trm ] : ends nopainctl session")
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

		err_final := fmt.Errorf("Error: %s", err.Error())

		fmt.Println(err_final.Error())

	}

}
