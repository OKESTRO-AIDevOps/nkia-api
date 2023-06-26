package main

import (
	"encoding/json"
	"fmt"

	runfs "github.com/seantywork/014_npia/pkg/runtimefs"
)

var RPARAM = map[string]string{
	"NS": "",
}

func setRuntimeParamNS(ns string) error {

	var app_origin runfs.AppOrigin

	file_byte, err := runfs.LoadAdmOrigin()

	if err != nil {

		return fmt.Errorf("runtime params: %s", err.Error())

	}

	err = json.Unmarshal(file_byte, &app_origin)

	if err != nil {

		return fmt.Errorf("runtime params: %s", err.Error())

	}

	ns_found, _, _ := runfs.GetRecordInfo(app_origin.RECORDS, ns)

	if !ns_found {

		return fmt.Errorf("runtime params: %s", "incomplete ns record")

	}

	err = runfs.CheckKubeNS(ns)

	if err != nil {

		return fmt.Errorf("set rparams: %s", err.Error())

	}

	return nil

}

func checkRuntimeParamNS() error {

	if RPARAM["NS"] == "" {

		return fmt.Errorf("check failed: %s", "ns not set")

	}

	return nil
}
