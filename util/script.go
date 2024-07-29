package util

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/dop251/goja"
)

func LoadScripts(dir string) (map[string]*goja.Program, error) {
	scripts := make(map[string]*goja.Program)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		scriptName, program, errComp := ComplieProgramFromFile(path)
		if errComp != nil {
			slog.Error("Complied script failed", "path", path)
		} else {
			scripts[scriptName] = program
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

func ComplieProgramFromFile(path string) (string, *goja.Program, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", nil, err
	}
	program, err := goja.Compile(path, string(bytes), false)
	if err != nil {
		return "", nil, err
	}
	filename := filepath.Base(path)
	scriptName := strings.TrimSuffix(filename, filepath.Ext(filename))
	return scriptName, program, nil
}
