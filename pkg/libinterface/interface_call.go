package libinterface

import (
	"path/filepath"
	//. "github.com/goccy/go-yaml"
)

func (libif LibIface) GetLibComponentPath(iface_name string, component string) string {

	iface_path := libif[iface_name]

	iface_component_path := filepath.Join(iface_path, component)

	return iface_component_path

}
