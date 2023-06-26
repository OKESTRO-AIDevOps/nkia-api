package runtimefs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func InitUsrTarget(repoaddr string) error {

	var app_origin AppOrigin

	if _, err := os.Stat(".usr/target"); err == nil {

		cmd := exec.Command("rm", "-r", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		cmd = exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()
	} else {
		cmd := exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

	}

	cmd := exec.Command("git", "-C", ".usr/target", "init")

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	file_byte, err := LoadAdmOrigin()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &app_origin)

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	addr_found, rec_repoid, rec_repopw := GetRepoInfo(app_origin.REPOS, repoaddr)

	if !addr_found {

		return fmt.Errorf("failed to init target: %s", "repo info not found")

	}

	insert := "%s:%s@"

	prt_idx := strings.Index(repoaddr, "://")

	prt_idx += 3

	repo_url := repoaddr[:prt_idx] + insert + repoaddr[prt_idx:]

	repo_url = fmt.Sprintf(repo_url, rec_repoid, rec_repopw)

	cmd = exec.Command("git", "-C", ".usr/target", "pull", repo_url)

	_, err = cmd.Output()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	if _, err := os.Stat(".usr/target/docker-compose.yml"); err == nil {

		cmd = exec.Command("mv", ".usr/target/docker-compose.yml", ".usr/target/docker-compose.yaml")

		cmd.Run()

	}

	if _, err := os.Stat(".usr/target/docker-compose.yaml"); err != nil {

		cmd = exec.Command("rm", "-r", ".usr/target")

		cmd.Run()

		return fmt.Errorf("failed to init target: %s", err.Error())

	}

	return nil

}

func CreateUsrTargetOperationSource(LIBIF_BIN_KOMPOSE string) error {

	if _, err := os.Stat(".usr/target"); err != nil {

		return fmt.Errorf("failed to create ops src: %s", err.Error())

	}

	if _, err := os.Stat(".usr/target/docker-compose.yaml"); err != nil {

		return fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	cmd := exec.Command(LIBIF_BIN_KOMPOSE, "convert", "-f", ".usr/target/docker-compose.yaml", "--stdout")

	out, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	err = os.WriteFile(".usr/ops_src.yaml", out, 0644)

	if err != nil {
		return fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	return nil
}
