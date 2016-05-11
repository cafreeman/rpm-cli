package main

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

// Manifest represents the R package manifest
type Manifest []*ManifestRecord

// ManifestRecord represents an individual row in the packages.csv manifest
type ManifestRecord struct {
	Package  string `csv:"Package"`
	Version  string `csv:"Version"`
	Status   string `csv:"Status"`
	Priority string `csv:"Priority"`
	Built    string `csv:"Built"`
}

// Create a Manifest struct from a packages.csv file
func readManifest(path string) Manifest {
	packageManifest, err := os.Open(path)
	errCheck(err)
	defer packageManifest.Close()

	packages := Manifest{}

	if err := gocsv.UnmarshalFile(packageManifest, &packages); err != nil {
		panic(err)
	}

	return packages
}

// Extract a list of non-default, non-Alteryx packages from a Manifest
func (m *Manifest) extractCRANPackages() []string {
	cranPackages := make([]string, len(*m))
	for i, pkg := range *m {
		if pkg.Priority == "NA" && !strings.HasPrefix(pkg.Package, "Alteryx") {
			cranPackages[i] = pkg.Package
		}
	}

	return cranPackages
}
