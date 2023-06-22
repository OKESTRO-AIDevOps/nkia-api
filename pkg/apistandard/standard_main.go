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

		b_out, cmd_err := kuberead.ReadNode(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-PDS":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPod(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-PLOG":

		ns := std_cmd["ns"]

		pod_name := std_cmd["podnm"]

		b_out, cmd_err := kuberead.ReadPodLog(ns, pod_name)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-SVC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadService(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-DPL":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadDeployment(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-IMGLI":
	case "RESOURCE-EVNT":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadEvent(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-RSRC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadResource(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NSPC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadNamespace(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-PRJPRB":
	case "RESOURCE-PSCH":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPodScheduled(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-PUNSCH":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPodUnscheduled(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CCPU":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerCPUUsage(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CMEM":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerMemUsage(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CFSR":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerFSRead(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CFSW":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerFSWrite(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CNETR":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerNetworkReceive(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-CNETT":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerNetworkTransmit(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-VOLAVAIL":

		b_out, cmd_err := kuberead.ReadKubeletVolumeAvailable()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-VOLCAP":

		b_out, cmd_err := kuberead.ReadKubeletVolumeCapacity()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-VOLUSD":

		b_out, cmd_err := kuberead.ReadKubeletVolumeUsed()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NTEMP":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsius()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NTEMPCH":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsiusChange()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NTEMPAV":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsiusAverage()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NPROCS":

		b_out, cmd_err := kuberead.ReadNodeProcessRunning()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NCORES":

		b_out, cmd_err := kuberead.ReadNodeCPUCores()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NMEM":

		b_out, cmd_err := kuberead.ReadNodeMemActive()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NMEMTOT":

		b_out, cmd_err := kuberead.ReadNodeMemTotal()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NDISKR":

		b_out, cmd_err := kuberead.ReadNodeDiskRead()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NDISKW":

		b_out, cmd_err := kuberead.ReadNodeDiskWrite()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NNETR":

		b_out, cmd_err := kuberead.ReadNodeNetworkReceive()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NNETT":
		b_out, cmd_err := kuberead.ReadNodeNetworkTransmit()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "RESOURCE-NDISKWT":

		b_out, cmd_err := kuberead.ReadNodeDiskWrittenTotal()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = b_out

	case "APPLY-SETREPO":
	case "APPLY-SETREG":
	case "APPLY-REGSEC":
		ns := std_cmd["ns"]

		str_out, cmd_err := kubewrite.WriteSecret(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = []byte(str_out)
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
