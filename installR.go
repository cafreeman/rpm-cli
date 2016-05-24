package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func installR(installerFile, installPath string) string {
	installPath = createInstallDir(installerFile, installPath)
	cmd := exec.Command(installerFile, "/SILENT", fmt.Sprintf(`/DIR=%v`, installPath))
	err := cmd.Run()
	if err == nil {
		fmt.Printf("Install completed. Cleaning up installer file.\n\n")
		err := os.Remove(installerFile)
		errCheck(err)
	} else {
		log.Fatal(err)
	}

	// Create a forward-slash version of the path to the installer so that we can actually use it inside R
	installDir := filepath.ToSlash(installPath)
	return installDir
}

func createInstallDir(installerFile, installPath string) string {
	fileName := filepath.Base(installerFile)
	rVersion := strings.TrimSuffix(fileName, "-win.exe")
	installDir := convertToWindowsPath(filepath.Join(installPath, rVersion))
	return installDir
}
