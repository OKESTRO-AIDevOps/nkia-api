package dotfs

type AppOrigin struct {
	RECORDS []RecordInfo
	REPOS   []RepoInfo
	REGS    []RegInfo
}

type RecordInfo struct {
	NS        string
	REPO_ADDR string
	REG_ADDR  string
}

type RepoInfo struct {
	REPO_ADDR string
	REPO_ID   string
	REPO_PW   string
}

type RegInfo struct {
	REG_ADDR string
	REG_ID   string
	REG_PW   string
}
