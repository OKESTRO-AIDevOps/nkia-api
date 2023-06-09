package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

type LibIface map[string]string

func _ConstructLibIface(root string, files []fs.FileInfo, dir *os.File) LibIface {

	tmp_libif := make(LibIface)

	for _, f := range files {

		iface_name := f.Name()

		full_path := filepath.Join(root, iface_name)

		tmp_libif[iface_name] = full_path
	}

	dir.Close()

	return tmp_libif
}

var _ROOT, _ = filepath.Abs("../lib")

var _DIR, _ = os.Open(_ROOT)

var _FILES, _ = _DIR.Readdir(-1)

var LIBIF LibIface = _ConstructLibIface(_ROOT, _FILES, _DIR)

func (libif LibIface) GetIfaceComponentPath(iface_name string, component string) string {

	iface_path := libif[iface_name]

	iface_component_path := filepath.Join(iface_path, component)

	return iface_component_path

}
