package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var (
	version            string
	installDestination string
	manifestPath       string
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
				cli.StringFlag{
					Name:        "manifest, m",
					Value:       "",
					Usage:       "The path to your package manifest (packages.csv)",
					Destination: &manifestPath,
				},
			},
			Action: func(c *cli.Context) error {
				if version == "" {
					log.Fatal("You must provide a version of R to install. Use the --release flag. Type `rpm-cli help` for more information.")
				}
				if installDestination == "" {
					log.Fatal("You must provide a directory path for the R install. Use the --destination flag. Type `rpm-cli help` for more information.")
				}
				if manifestPath == "" {
					log.Fatal("You must provide a file path to your package manifest. Use --manifest flag. Type `rpm-cli help` for more informaton.")
				}
				// Read package manifest first, so we can error out if the file path is invalid
				manifest := readManifest(filepath.Join(manifestPath))
				cranPackages := manifest.extractCRANPackages()

				repoURL := "https://cran.rstudio.com"

				// Get current working directory
				rootPath, _ := os.Getwd()

				// Pull the file name out of the URL to use when creating the file
				fileName := fmt.Sprintf("R-%v-win.exe", version)

				// Check if the installer has already been downloaded
				isDownloaded, installerPath := checkInstaller(rootPath, fileName)

				// If the installer has not been downloaded, then download and save it to `fileName`
				if isDownloaded != true {
					createInstallDirectory(&installerPath)

					downloadInstaller(&version, &installerPath)

					fmt.Printf("Successfully downloaded %s. The installer is located at %s.\n", fileName, installerPath)
				} else {
					fmt.Printf("%s has already been downloaded. Skipping to install step\n\n", fileName)
				}

				// Run the R installer and return the install path
				rInstall := installR(installerPath, filepath.Join(rootPath, installDestination))

				installRPackages(rInstall, &cranPackages, repoURL)

				return nil
			},
		},
	}

	app.Run(os.Args)
}
