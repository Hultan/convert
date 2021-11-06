package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func main() {
	// Get files
	filePaths, err := filepath.Glob("/softtube/thumbnails/*.webp")
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		return
	}

	// Convert found files
	for _, filePath := range filePaths {
		if err = convertWebp2Jpg(filePath); err != nil {
			// We failed to convert using dwebp
		}
	}
}

func convertWebp2Jpg(webpFilePath string) error {
	jpgFilePath := getJpgFilename(webpFilePath)

	// Run dwebp to convert the image
	_, err := exec.Command("dwebp", webpFilePath, "-o", jpgFilePath).Output()
	if err != nil {
		return err
	}

	// Remove webp file (if the jpg exists)
	if _, err := os.Stat(jpgFilePath); err==nil {
		if err = os.Remove(webpFilePath); err != nil {
			return err
		}
	}

	return nil
}

func getJpgFilename(filePath string) string {
	var extension = filepath.Ext(filePath)
	return path.Join(filePath[0:len(filePath)-len(extension)]) + ".jpg"
}
