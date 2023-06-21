package apistandard

import "strings"

type API_OUTPUT struct {
	HEAD map[string]map[string]string

	BODY []byte
}

type API_INPUT map[string]string

type API_STD map[string][]string

var API_DEFINITION string = "" +
	//            id          :       keys
	"SUBMIT                   :id                                                           " + "\n" +
	"CALLME                   :id                                                           " + "\n" +
	"SETTING-CRTNS            :id, ns                                                       " + "\n" +
	"SETTING-CRTNSVOL         :id, ns, volserver                                            " + "\n" +
	"SETTING-CRTVOL           :id, ns, volserver                                            " + "\n" +
	"SETTING-CRTMON           :id, ns                                                       " + "\n" +
	"SETTING-DELNS            :id, ns                                                       " + "\n" +
	"GITLOG                   :id, ns, repoaddr                                             " + "\n" +
	"PIPEHIST                 :id, ns                                                       " + "\n" +
	"PIPE                     :id, ns, repoaddr, regaddr                                    " + "\n" +
	"PIPELOG                  :id                                                           " + "\n" +
	"BUILD                    :id, ns, repoaddr, regaddr                                    " + "\n" +
	"BUILDLOG                 :id, ns                                                       " + "\n" +
	"RESOURCE-NDS             :id, ns                                                       " + "\n" +
	"RESOURCE-PDS             :id, ns                                                       " + "\n" +
	"RESOURCE-PLOG            :id, ns, podnm                                                " + "\n" +
	"RESOURCE-SVC             :id, ns                                                       " + "\n" +
	"RESOURCE-DPL             :id, ns                                                       " + "\n" +
	"RESOURCE-IMGLI           :id, ns                                                       " + "\n" +
	"RESOURCE-EVNT            :id, ns                                                       " + "\n" +
	"RESOURCE-RSRC            :id, ns                                                       " + "\n" +
	"RESOURCE-NSPC            :id, ns                                                       " + "\n" +
	"RESOURCE-PRJPRB          :id, ns                                                       " + "\n" +
	"RESOURCE-PSCH            :id, ns                                                       " + "\n" +
	"RESOURCE-PUNSCH          :id, ns                                                       " + "\n" +
	"RESOURCE-CCPU            :id, ns                                                       " + "\n" +
	"RESOURCE-CMEM            :id, ns                                                       " + "\n" +
	"RESOURCE-CFSR            :id, ns                                                       " + "\n" +
	"RESOURCE-CFSW            :id, ns                                                       " + "\n" +
	"RESOURCE-CNETR           :id, ns                                                       " + "\n" +
	"RESOURCE-CNETT           :id, ns                                                       " + "\n" +
	"RESOURCE-VOLAVAIL        :id, ns                                                       " + "\n" +
	"RESOURCE-VOLCAP          :id, ns                                                       " + "\n" +
	"RESOURCE-VOLUSD          :id, ns                                                       " + "\n" +
	"RESOURCE-NTEMP           :id, ns                                                       " + "\n" +
	"RESOURCE-NTEMPCH         :id, ns                                                       " + "\n" +
	"RESOURCE-NTEMPAV         :id, ns                                                       " + "\n" +
	"RESOURCE-NPROCS          :id, ns                                                       " + "\n" +
	"RESOURCE-NCORES          :id, ns                                                       " + "\n" +
	"RESOURCE-NMEM            :id, ns                                                       " + "\n" +
	"RESOURCE-NMEMTOT         :id, ns                                                       " + "\n" +
	"RESOURCE-NDISKR          :id, ns                                                       " + "\n" +
	"RESOURCE-NDISKW          :id, ns                                                       " + "\n" +
	"RESOURCE-NNETR           :id, ns                                                       " + "\n" +
	"RESOURCE-NNETT           :id, ns                                                       " + "\n" +
	"RESOURCE-NDISKWT         :id, ns                                                       " + "\n" +
	"APPLY-SETREPO            :id, ns, repoaddr, repoid, repopw                             " + "\n" +
	"APPLY-SETREG             :id, ns, regaddr, regid, regpw                                " + "\n" +
	"APPLY-REGSEC             :id, ns                                                       " + "\n" +
	"APPLY-DIST               :id, ns, repoaddr, regaddr                                    " + "\n" +
	"APPLY-CRTOPSSRC          :id, ns, repoaddr, regaddr                                    " + "\n" +
	"APPLY-RESTART            :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-ROLLBACK           :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-KILL               :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-NETRESH            :id, ns                                                       " + "\n" +
	"APPLY-HPA                :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-HPAUN              :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-QOS                :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-QOSUN              :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-GETINGR            :id, ns                                                       " + "\n" +
	"APPLY-INGR               :id, ns, hostnm, svcnm                                        " + "\n" +
	"APPLY-INGRUN             :id, ns, hostnm, svcnm                                        " + "\n" +
	"APPLY-GETNDPORT          :id, ns                                                       " + "\n" +
	"APPLY-NDPORT             :id, ns, svcnm                                                " + "\n" +
	"APPLY-NDPORTUN           :id, ns, svcnm                                                " + "\n" +
	"ADMIN-ADMRMTCHK          :id, hostaddr, hostport, usernm, userpw, machinerole, acd     " + "\n" +
	"ADMIN-ADMRMTLDHA         :id, hostaddr, hostport, usernm, userpw, remoteip             " + "\n" +
	"ADMIN-ADMRMTLDMV         :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	"ADMIN-ADMRMTMSR          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	"ADMIN-ADMRMTLDWRK        :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	"ADMIN-ADMRMTWRK          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	"ADMIN-ADMRMTSTR          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	"ADMIN-ADMRMTLOG          :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	"ADMIN-ADMRMTSTATUS       :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	"ADMIN-LEAD               :id, token, targetip                                          " + "\n" +
	"ADMIN-MSR                :id, token, targetip                                          " + "\n" +
	"ADMIN-LDVOL              :id, token, targetip                                          " + "\n" +
	"ADMIN-WRK                :id, token, targetip                                          " + "\n" +
	"ADMIN-STR                :id, token, targetip                                          " + "\n" +
	"ADMIN-LOG                :id                                                           " + "\n" +
	"ADMIN-STATUS             :id, token, targetip                                          " + "\n" +
	"ADMIN-UP                 :id, token                                                    " + "\n" +
	"ADMIN-DOWN               :id, token                                                    " + "\n" +
	"DELND                    :id                                                           " + "\n" +
	"EXIT                     :id                                                           "

func _CONSTRUCT_API_INPUT() API_STD {

	apistd := make(API_STD)

	sanitized_def := strings.ReplaceAll(API_DEFINITION, " ", "")

	def_list := strings.Split(sanitized_def, "\n")

	for i := 0; i < len(def_list); i++ {

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		value_list := strings.Split(record_list[1], ",")

		key := record_list[0]

		apistd[key] = value_list

	}

	return apistd

}

var ASgi = _CONSTRUCT_API_INPUT()
