package libinterface

import (
	"io/fs"
	"os"
	"path/filepath"
)

type LibIface map[string]string

func _CONSTRUCTLIBIFACE(root string, files []fs.FileInfo, dir *os.File) LibIface {

	tmp_libif := make(LibIface)

	for _, f := range files {

		iface_name := f.Name()

		full_path := filepath.Join(root, iface_name)

		tmp_libif[iface_name] = full_path
	}

	dir.Close()

	return tmp_libif
}

var _ROOT, _ = filepath.Abs("../../lib")

var _DIR, _ = os.Open(_ROOT)

var _FILES, _ = _DIR.Readdir(-1)

var LIBIF LibIface = _CONSTRUCTLIBIFACE(_ROOT, _FILES, _DIR)
