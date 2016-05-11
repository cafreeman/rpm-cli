// package main
//
// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// 	"time"
//
// 	"github.com/cheggaaa/pb"
// )
//
// func downloadR(version, rootPath string) string {
// 	url := fmt.Sprintf("https://cran.rstudio.com/bin/windows/base/R-%v-win.exe", version)
//
// 	// Parse URL and create filename from last element
// 	tokens := strings.Split(url, "/")
// 	fileName := tokens[len(tokens)-1]
//
// 	// Check if file has already been downloaded. If not, create the file at the specified path
// 	installerPath := createDownloadPath(rootPath, fileName)
// 	println(installerPath)
//
// 	// Check to see if the installer has already been downloaded
// 	if _, err := os.Stat(installerPath); err == nil {
// 		fmt.Println(fileName, "already exists!")
// 		// return forward-slash installer path
// 		installerFile, err := os.Open(installerPath)
// 		errCheck(err)
// 		return installerFile.Name()
// 	}
//
// 	// Create install directory recursively, then create the installer file itself
// 	os.MkdirAll(filepath.Dir(installerPath), 0666)
// 	installerFile, err := os.Create(installerPath)
// 	errCheck(err)
// 	defer installerFile.Close()
//
// 	// Start the installer process
// 	fmt.Println("Downloading", url, "to", fileName)
//
// 	// Download the installer
// 	response, err := http.Get(url)
// 	errCheck(err)
// 	defer response.Body.Close()
//
// 	// Print the http response to stdout
// 	fmt.Println(response.Status)
//
// 	// Get the response size from the HTTP header for calculating download progress
// 	responseSize, _ := strconv.Atoi(response.Header.Get("Content-Length"))
//
// 	// Create progress bar
// 	bar := pb.New(int(responseSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
// 	bar.ShowSpeed = true
// 	bar.SetWidth(120)
// 	bar.Start()
//
// 	// Create multi-writer for output destination and progress bar
// 	writer := io.MultiWriter(installerFile, bar)
//
// 	// Copy to output
// 	_, err = io.Copy(writer, response.Body)
// 	errCheck(err)
// 	bar.Finish()
//
// 	fmt.Printf("Successfully downloaded %s. The installer is located at %s.", fileName, installerPath)
//
// 	return installerFile.Name()
// }
//
// func createDownloadPath(rootPath, fileName string) string {
// 	absPath, err := filepath.Abs(rootPath)
// 	errCheck(err)
// 	return convertToWindowsPath(filepath.Join(absPath, fileName))
// }
