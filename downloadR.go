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

func downloadInstaller(url *string, installerPath *string) {
	response, err := http.Get(*url)
	errCheck(err)
	if response.StatusCode != http.StatusOK {
		msg := fmt.Sprintf(`There was an issue downloading the installer from %s
      The URL returned the following response:
      %v`,
			*url,
			response.Status,
		)
		log.Fatal(msg)
	}

	defer response.Body.Close()
	bar := createProgressBar(response)

	saveInstaller(response, installerPath, bar)
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
	bar.SetWidth(120)

	return bar
}
