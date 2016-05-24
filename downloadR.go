package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
)

func getFileName(url string) string {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	return fileName
}

// Create installer path and check to see if the installer has already been downloaded
func checkInstaller(rootPath, fileName string) (bool, string) {
	installerPath := createDownloadPath(rootPath, fileName)

	if _, err := os.Stat(installerPath); err == nil {
		return true, installerPath
	}
	return false, installerPath
}

func createDownloadPath(rootPath, fileName string) string {
	absPath, err := filepath.Abs(rootPath)
	errCheck(err)
	return convertToWindowsPath(filepath.Join(absPath, fileName))
}

// Create install directory recursively, then create the installer file itself
func createInstallDirectory(path *string) {
	err := os.MkdirAll(filepath.Dir(*path), 0755)
	errCheck(err)
}

func createURLs(version *string) []string {
	currentVersionURL := fmt.Sprintf("https://cran.rstudio.com/bin/windows/base/R-%v-win.exe", *version)
	oldVersionURL := fmt.Sprintf("https://cran.rstudio.com/bin/windows/base/old/%[1]v/R-%[1]v-win.exe", *version)
	return []string{currentVersionURL, oldVersionURL}
}

// Returns the URL from a successful download
func downloadInstaller(version *string, installerPath *string) string {
	urls := createURLs(version)

	url := urls[0]

	// Attempt to download using currentVersionURL first
	response, err := http.Get(url)
	errCheck(err)
	// If the server returns an error status, optimistically assume we're just looking for an old
	// version and attempt to download using the oldVersionURL
	if response.StatusCode != http.StatusOK {
		fmt.Printf("No installer found at %s.\n"+
			"You might be attempting to download an older version of R.\n"+
			"Attempting to download from past releases.\n\n", url)
		// Download from oldVersionURL
		url = urls[1]
		response, err = http.Get(url)
		errCheck(err)
		// If the second download attempt fails, assume something else is wrong and exit
		if response.StatusCode != http.StatusOK {
			msg := fmt.Sprintf("There was an issue downloading the installer from %s\n"+
				"The URL returned the following response:\n"+
				"%v",
				url,
				response.Status,
			)
			log.Fatal(msg)
		}
	}

	defer response.Body.Close()
	bar := createProgressBar(response)

	saveInstaller(response, installerPath, bar)

	return url
}

func saveInstaller(response *http.Response, installerPath *string, bar *pb.ProgressBar) string {
	bar.Start()

	installerFile, err := os.Create(*installerPath)
	defer installerFile.Close()
	errCheck(err)

	// Create multi-writer for output destination and progress bar
	writer := io.MultiWriter(installerFile, bar)

	// Copy to output
	_, err = io.Copy(writer, response.Body)
	errCheck(err)
	bar.Finish()

	return installerFile.Name()
}

func createProgressBar(response *http.Response) *pb.ProgressBar {
	// Get the response size from the HTTP header for calculating download progress
	responseSize, _ := strconv.Atoi(response.Header.Get("Content-Length"))

	// Create progress bar
	bar := pb.New(int(responseSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	// bar.SetWidth(120)

	return bar
}
