package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func installRPackages(rPath string, packageList *[]string, repoURL string) {
	// Construct path to R executable
	rScriptPath := filepath.Join(convertToWindowsPath(rPath), "bin", "Rscript.exe")

	fmt.Println(rScriptPath)

	for _, pkg := range *packageList {
		installCmdString := fmt.Sprintf("install.packages('%s', repos='%s')", pkg, repoURL)
		fmt.Println(installCmdString)
		installCmd := exec.Command(rScriptPath, "-e", installCmdString)
		out, err := installCmd.Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	}
}
