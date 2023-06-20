package main

import (
	"encoding/json"
	"fmt"

	"github.com/seantywork/014_npia/pkg/dotfs"
)

var RPARAM = map[string]string{
	"NS": "",
}

func setRuntimeParamNS(ns string) error {

	var app_origin dotfs.AppOrigin

	file_byte, err := dotfs.LoadAdmOrigin()

	if err != nil {

		return fmt.Errorf("runtime params: %s", err.Error())

	}

	err = json.Unmarshal(file_byte, &app_origin)

	if err != nil {

		return fmt.Errorf("runtime params: %s", err.Error())

	}

	repo, reg := dotfs.GetRecordInfo(app_origin.RECORDS, ns)

	if repo == "N" || reg == "N" {

		return fmt.Errorf("runtime params: %s", "incomplete ns record")

	}

	err = dotfs.CheckKubeNS(ns)

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
