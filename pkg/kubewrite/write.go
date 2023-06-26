package kubewrite

import (
	"fmt"

	runfs "github.com/seantywork/014_npia/pkg/runtimefs"

	"github.com/seantywork/014_npia/pkg/libinterface"

	"os/exec"

	"encoding/json"
)

func WriteRepoInfo(main_ns string, repoaddr string, repoid string, repopw string) ([]byte, error) {

	var ret_byte []byte

	var app_origin runfs.AppOrigin

	adm_origin_byte, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	err = json.Unmarshal(adm_origin_byte, &app_origin)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, _, rec_regaddr := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	app_origin.RECORDS = runfs.SetRecordInfo(app_origin.RECORDS, main_ns, repoaddr, rec_regaddr)

	app_origin.REPOS = runfs.SetRepoInfo(app_origin.REPOS, repoaddr, repoid, repopw)

	err = runfs.UnloadAdmOrigin(app_origin)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte = []byte("repo info registered\n")

	return ret_byte, nil
}

func WriteRegInfo(main_ns string, regaddr string, regid string, regpw string) ([]byte, error) {

	var ret_byte []byte

	var app_origin runfs.AppOrigin

	adm_origin_byte, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	err = json.Unmarshal(adm_origin_byte, &app_origin)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, rec_repoaddr, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	app_origin.RECORDS = runfs.SetRecordInfo(app_origin.RECORDS, main_ns, rec_repoaddr, regaddr)

	app_origin.REGS = runfs.SetRegInfo(app_origin.REGS, regaddr, regid, regpw)

	err = runfs.UnloadAdmOrigin(app_origin)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte = []byte("reg infor registered\n")

	return ret_byte, nil
}

func WriteSecret(main_ns string) ([]byte, error) {

	var ret_byte []byte

	var app_origin runfs.AppOrigin

	adm_origin_byte, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	err = json.Unmarshal(adm_origin_byte, &app_origin)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, _, reg_url := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	if reg_url == "N" {

		return ret_byte, fmt.Errorf(": %s", "reg url not set")

	}

	addr_found, reg_id, reg_pw := runfs.GetRegInfo(app_origin.REGS, reg_url)

	if !addr_found {

		return ret_byte, fmt.Errorf(": %s", "reg info not complete")

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
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		ret_byte = out

		return ret_byte, nil

	} else {

		cmd = exec.Command("kubectl", "-n", main_ns, "delete", "secret", "docker-secret")

		_, err = cmd.Output()

		if err != nil {
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		cmd = exec.Command("kubectl", "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

		out, err := cmd.Output()

		if err != nil {
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		ret_byte = out

		return ret_byte, nil

	}

}

func WriteDeployment(main_ns string, repoaddr string, regaddr string) ([]byte, error) {

	var ret_byte []byte

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	_, err = libif.GetLibComponentAddress("bin", "kompose")

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	var app_origin runfs.AppOrigin

	adm_origin_byte, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	err = json.Unmarshal(adm_origin_byte, &app_origin)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, repoaddr, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	err = runfs.InitUsrTarget(repoaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	return ret_byte, nil

}

func WriteOperationSource() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteUpdateOrRestart() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteDeletion() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteNetworkRefresh() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteHPA() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteHPAUndo() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteQOS() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteIngress() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteIngressUndo() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteNodePort() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}

func WriteNodePortUndo() ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil

}
