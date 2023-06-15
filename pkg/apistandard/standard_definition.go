package apistandard

type EXCHANGE struct {
	ACD  string
	CMD  string
	DATA string
	CNT  string
	MSG  string
}

type FORM struct {
	FORM API_OUTPUT
}

type API_OUTPUT struct {
	HEADER string

	BODY string
}

type API_INPUT map[string][]string

var API_DEFINITION string = "SUBMIT:none" + "\n" +
	"CALLME:none" + "\n" +
	"SETTING-CRTNS:ns" + "\n" +
	"SETTING-CRTNSVOL:ns,vol_server" + "\n" +
	"GITLOG:" + "\n" +
	"PIPEHIST:" + "\n" +
	"PIPE:" + "\n" +
	"PIPELOG:" + "\n" +
	"BUILD:" + "\n" +
	"BUILDLOG:" + "\n" +
	"RESOURCE:" + "\n" +
	"APPLY:" + "\n" +
	"ADMIN:" + "\n" +
	"DELND:" + "\n" +
	"EXIT:"

func _CONSTRUCT_API_INPUT() {

}
