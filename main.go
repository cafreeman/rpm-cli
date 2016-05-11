package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

var (
	version            string
	installDestination string
)

func main() {
	app := cli.NewApp()
	app.Name = "rpm"
	app.Version = "0.1.0"
	app.Usage = "Build new R installers!"

	app.Commands = []cli.Command{
		{
			Name:    "build-installer",
			Aliases: []string{"build"},
			Usage:   "Build a new R installer from scratch",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "release, r",
					Value:       "",
					Usage:       "The `VERSION` of R you'd like to install. e.g. \"3.3.0\"",
					Destination: &version,
				},
				cli.StringFlag{
					Name:        "destination, d",
					Value:       "",
					Usage:       "The destination directory for your R install",
					Destination: &installDestination,
				},
			},
			Action: func(c *cli.Context) error {
				if version == "" {
					log.Fatal("You must provide a version of R to install. Use the --release flag. Type `rpm-cli help` for more information.")
				}
				if installDestination == "" {
					log.Fatal("You must provide a directory path for the R install. Use the --destination flag. Type `rpm-cli help` for more information.")
				}
				rootPath, _ := os.Getwd()

				// Create the download URL for the specific version of R
				url := fmt.Sprintf("https://cran.rstudio.com/bin/windows/base/R-%v-win.exe", version)

				// Pull the file name out of the URL to use when creating the file
				fileName := getFileName(url)

				// Check if the installer has already been downloaded
				isDownloaded, installerPath := checkInstaller(rootPath, fileName)

				// If the installer has not been downloaded, then download and save it to `fileName`
				if isDownloaded != true {
					createInstallDirectory(&installerPath)

					downloadInstaller(&url, &installerPath)

					fmt.Printf("Successfully downloaded %s. The installer is located at %s.\n", fileName, installerPath)
				} else {
					fmt.Println(fileName, "has already been downloaded. Skipping to install step\n")
				}

				rInstall := installR(installerPath, filepath.Join(rootPath, installDestination))

				manifest := readManifest(filepath.Join(rootPath, "packages.csv"))
				cranPackages := manifest.extractCRANPackages()

				repoURL := "https://cran.rstudio.com"

				// sampleList := cranPackages[:4]

				// installRPackages(rInstall, &sampleList, repoURL)
				installRPackages(rInstall, &cranPackages, repoURL)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
