package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	rootPath, _ := os.Getwd()
	version := "3.2.5"
	targetPath := filepath.Join(rootPath, "rpm-test")

	// Create the download URL for the specific version of R
	url := fmt.Sprintf("https://cran.rstudio.com/bin/windows/base/R-%v-win.exe", version)
	// Pull the file name out of the URL to use when creating the file
	fileName := getFileName(url)

	// Check if the installer has already been downloaded
	isDownloaded, installerPath := checkInstaller(targetPath, fileName)

	// If the installer has not been downloaded, then download and save it to `fileName`
	if isDownloaded != true {
		createInstallDirectory(&installerPath)

		downloadInstaller(&url, &installerPath)

		fmt.Printf("Successfully downloaded %s. The installer is located at %s.", fileName, installerPath)
	} else {
		fmt.Println(fileName, "has already been downloaded. Skipping to install step")
	}

	rInstall := installR(installerPath, filepath.Join(rootPath, "R-install"))

	manifest := readManifest(filepath.Join(rootPath, "packages.csv"))
	cranPackages := manifest.extractCRANPackages()

	sampleList := cranPackages[:4]

	fmt.Printf("%v\n", sampleList)

	repoURL := "https://cran.rstudio.com"

	installRPackages(rInstall, &sampleList, repoURL)
}
