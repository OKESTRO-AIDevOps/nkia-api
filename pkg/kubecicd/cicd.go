package kubecicd

/*

	if code == "build" {

		_, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url, reg_url := admor.GetRecordInfo(app_origin.RECORDS, main_ns)

		if repo_url == "N" || reg_url == "N" {

			fmt.Println("Associated repository or registry information doesn't exist as a record")

			return

		}

		repo_id, repo_pw := admor.GetRepoInfo(app_origin.REPOS, repo_url)

		reg_id, reg_pw := admor.GetRegInfo(app_origin.REGS, reg_url)

		if repo_id == "N" || repo_pw == "N" {

			fmt.Println("Associated repository credential is imperfect")

			return

		}

		if reg_id == "N" || reg_pw == "N" {

			fmt.Println("Associated registry credential is imperfect")

			return

		}

		if _, err := os.Stat("./ADM/target"); err == nil {

			cmd := exec.Command("rm", "-r", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			cmd = exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()
		} else {
			cmd := exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

		}

		os.Chdir("./ADM/target")

		cmd := exec.Command("git", "init")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		insert := "%s:%s@"

		prt_idx := strings.Index(repo_url, "://")

		prt_idx += 3

		repo_url = repo_url[:prt_idx] + insert + repo_url[prt_idx:]

		repo_url = fmt.Sprintf(repo_url, repo_id, repo_pw)

		cmd = exec.Command("git", "pull", repo_url)

		cmd.Run()

		dcfile_nm := ""

		if _, err := os.Stat("docker-compose.yaml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yaml"

		}

		if _, err := os.Stat("docker-compose.yml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yml"

		}

		if dcfile_nm == "" {

			os.Chdir("../../")

			str_stdout := "File [ docker-compose.yaml | docker-compose.yml ] Does Not Exist"

			fmt.Println(str_stdout)

			return

		}

		cmd = exec.Command("docker-compose", "up", "-d", "--build")

		bld_out, _ := cmd.StdoutPipe()

		bld_err, _ := cmd.StderrPipe()

		go func() {
			merged := io.MultiReader(bld_out, bld_err)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println(msg)
			}
		}()

		cmd.Run()

		os.Chdir("../../")

		cmd = exec.Command("/bin/cp", "-rf", dcfile_nm, "./ADM/docker-compose.yaml")

		cmd.Run()

		cmd = exec.Command("python3", "./ADM/yaml_handle.py", "push", reg_url, reg_id, reg_pw)

		bld_out, _ = cmd.StdoutPipe()

		bld_err, _ = cmd.StderrPipe()

		go func() {
			merged := io.MultiReader(bld_out, bld_err)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println(msg)
			}
		}()

		cmd.Run()

		cmd = exec.Command("rm", "-r", "./ADM/target")

		cmd.Run()

	} else if code == "deploy" {

		kcfg_path, main_ns := admor.GetKubeConfigAndTargetNameSpace()

		var app_origin admor.AppOrigin

		file_content, _ := os.ReadFile("./ADM/origin.json")

		_ = json.Unmarshal(file_content, &app_origin)

		repo_url, reg_url := admor.GetRecordInfo(app_origin.RECORDS, main_ns)

		if repo_url == "N" || reg_url == "N" {

			fmt.Println("Associated repository or registry information doesn't exist as a record")

			return

		}

		repo_id, repo_pw := admor.GetRepoInfo(app_origin.REPOS, repo_url)

		reg_id, reg_pw := admor.GetRegInfo(app_origin.REGS, reg_url)

		if repo_id == "N" || repo_pw == "N" {

			fmt.Println("Associated repository credential is imperfect")

			return

		}

		if reg_id == "N" || reg_pw == "N" {

			fmt.Println("Associated registry credential is imperfect")

			return

		}

		if _, err := os.Stat("./ADM/target"); err == nil {

			cmd := exec.Command("rm", "-r", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			cmd = exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()
		} else {
			cmd := exec.Command("mkdir", "./ADM/target")

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

		}

		os.Chdir("./ADM/target")

		cmd := exec.Command("git", "init")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		insert := "%s:%s@"

		prt_idx := strings.Index(repo_url, "://")

		prt_idx += 3

		repo_url = repo_url[:prt_idx] + insert + repo_url[prt_idx:]

		repo_url = fmt.Sprintf(repo_url, repo_id, repo_pw)

		cmd = exec.Command("git", "pull", repo_url)

		cmd.Run()

		dcfile_nm := ""

		if _, err := os.Stat("docker-compose.yaml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yaml"

		}

		if _, err := os.Stat("docker-compose.yml"); err == nil {

			dcfile_nm = "./ADM/target/docker-compose.yml"

		}

		if dcfile_nm == "" {

			os.Chdir("../../")

			str_stdout := "File [ docker-compose.yaml | docker-compose.yml ] Does Not Exist"

			fmt.Println(str_stdout)

			return

		}

		os.Chdir("../../")

		cmd = exec.Command("/bin/cp", "-rf", dcfile_nm, "./ADM/docker-compose.yaml")

		cmd.Run()

		cmd = exec.Command("python3", "./ADM/yaml_handle.py", "deploy", reg_url, reg_id, reg_pw)

		cmd.Run()

		cmd = exec.Command("kubectl", "--kubeconfig", kcfg_path, "-n", main_ns, "apply", "-f", "./ADM/ops_src.yaml")

		out, err := cmd.Output()

		if err != nil {

			strerr := err.Error()

			fmt.Println(strerr)

			return

		}

		strout := string(out)

		fmt.Println(strout)

	} else if code == "list" {

		list_all()

	} else if code == "back" {

		return

	} else {

		fmt.Println("Invalid command")
	}
*/
