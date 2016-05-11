package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func installR(installerFile, installPath string) string {
	cmd := exec.Command(installerFile, "/SILENT", fmt.Sprintf(`/DIR=%v`, installPath), "/y")
	err := cmd.Run()
	errCheck(err)

	installDir := filepath.ToSlash(installPath)
	return installDir
}
