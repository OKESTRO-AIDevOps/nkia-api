package apistandard

import (
	"fmt"

	"github.com/seantywork/014_npia/pkg/kuberead"
	"github.com/seantywork/014_npia/pkg/kubewrite"
)

func (asgi API_STD) Run(std_cmd API_INPUT) (API_OUTPUT, error) {

	var ret_api_out API_OUTPUT

	if verified := asgi.Verify(std_cmd); verified != nil {

		return ret_api_out, fmt.Errorf("run failed: %s", verified.Error())
	}

	cmd_id := std_cmd["id"]

	switch cmd_id {
	case "SUBMIT":
	case "CALLME":
	case "SETTING-CRTNS":
	case "SETTING-CRTNSVOL":
	case "SETTING-CRTVOL":
	case "SETTING-CRTMON":
	case "SETTING-DELNS":
	case "GITLOG":
	case "PIPEHIST":
	case "PIPE":
	case "PIPELOG":
	case "BUILD":
	case "BUILDLOG":
	case "RESOURCE-NDS":

		ns := std_cmd["ns"]

		str_out, cmd_err := kuberead.ReadNode(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = str_out

	case "RESOURCE-PDS":
	case "RESOURCE-PLOG":
	case "RESOURCE-SVC":
	case "RESOURCE-DPL":
	case "RESOURCE-IMGLI":
	case "RESOURCE-EVNT":
	case "RESOURCE-RSRC":
	case "RESOURCE-NSPC":
	case "RESOURCE-PRJPRB":
	case "RESOURCE-PSCH":
	case "RESOURCE-PUNSCH":
	case "RESOURCE-CCPU":
	case "RESOURCE-CMEM":
	case "RESOURCE-CFSR":
	case "RESOURCE-CFSW":
	case "RESOURCE-CNETR":
	case "RESOURCE-CNETT":
	case "RESOURCE-VOLAVAIL":
	case "RESOURCE-VOLCAP":
	case "RESOURCE-VOLUSD":
	case "RESOURCE-NTEMP":
	case "RESOURCE-NTEMPCH":
	case "RESOURCE-NTEMPAV":
	case "RESOURCE-NPROCS":
	case "RESOURCE-NCORES":
	case "RESOURCE-NMEM":
	case "RESOURCE-NMEMTOT":
	case "RESOURCE-NDISKR":
	case "RESOURCE-NDISKW":
	case "RESOURCE-NNETR":
	case "RESOURCE-NNETT":
	case "RESOURCE-NDISKWT":
	case "APPLY-SETREPO":
	case "APPLY-SETREG":
	case "APPLY-REGSEC":
		ns := std_cmd["ns"]

		str_out, cmd_err := kubewrite.WriteSecret(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = str_out
	case "APPLY-DIST":
	case "APPLY-CRTOPSSRC":
	case "APPLY-RESTART":
	case "APPLY-ROLLBACK":
	case "APPLY-KILL":
	case "APPLY-NETRESH":
	case "APPLY-HPA":
	case "APPLY-HPAUN":
	case "APPLY-QOS":
	case "APPLY-QOSUN":
	case "APPLY-GETINGR":
	case "APPLY-INGR":
	case "APPLY-INGRUN":
	case "APPLY-GETNDPORT":
	case "APPLY-NDPORT":
	case "APPLY-NDPORTUN":
	case "ADMIN-ADMRMTCHK":
	case "ADMIN-ADMRMTLDHA":
	case "ADMIN-ADMRMTLDMV":
	case "ADMIN-ADMRMTMSR":
	case "ADMIN-ADMRMTLDWRK":
	case "ADMIN-ADMRMTWRK":
	case "ADMIN-ADMRMTSTR":
	case "ADMIN-ADMRMTLOG":
	case "ADMIN-ADMRMTSTATUS":
	case "ADMIN-LEAD":
	case "ADMIN-MSR":
	case "ADMIN-LDVOL":
	case "ADMIN-WRK":
	case "ADMIN-STR":
	case "ADMIN-LOG":
	case "ADMIN-STATUS":
	case "ADMIN-UP":
	case "ADMIN-DOWN":
	case "DELND":
	case "EXIT":
	default:

		return ret_api_out, fmt.Errorf("failed to run api: %s", "invalid command id")

	}

	return ret_api_out, nil

}
