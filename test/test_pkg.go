package main

import (
	"fmt"
	"npia/pkg/builtinresource"
)

func Test_pkg() error {

	if !TEST_PKG {
		return nil
	}

	if err := test_kuberesource(); err != nil {

		return fmt.Errorf("kuberesource: %s", err.Error())

	}
	return nil

}

func test_kuberesource() error {

	if err := builtinresource.HorizontalPodAutoscaler_test(); err != nil {

		return fmt.Errorf("hpa: %s", err.Error())

	}

	if err := builtinresource.Ingress_test(); err != nil {

		return fmt.Errorf("ingress: %s", err.Error())

	}

	if err := builtinresource.Service_test(); err != nil {

		return fmt.Errorf("service: %s", err.Error())

	}

	if err := builtinresource.Deployment_test(); err != nil {

		return fmt.Errorf("deployment: %s", err.Error())

	}

	return nil

}
