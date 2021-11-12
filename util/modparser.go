// Package for parsing project's modfile
package util

import (
	"os"
	"sync"

	"golang.org/x/mod/modfile"
)

const (
	modFileName = "go.mod"
)

var (
	once                sync.Once
	parsedModFile       *modfile.File
	parseModFileErr     error
	modFileVersionFixer modfile.VersionFixer = func(path, version string) (string, error) {
		return version, nil
	}
)

func ParseModfile() (*modfile.File, error) {
	once.Do(func() {
		modFileBytes, parseModFileErr := os.ReadFile(modFileName)

		if parseModFileErr == nil {
			parsedModFile, parseModFileErr = modfile.Parse(modFileName, modFileBytes, modFileVersionFixer)
		}
	})

	return parsedModFile, parseModFileErr
}
