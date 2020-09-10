package main

import (
	"fmt"
	log "github.com/hultan/softteam/log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	logger *log.Logger
)

func main() {
	// Setup logging
	logger = log.NewLog(path.Join("/softtube/log/", "softtube.convert.log"))
	defer logger.Close()

	// Get files
	filePaths, err := filepath.Glob("/softtube/thumbnails/*.webp")
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		return
	}

	// Log number of found files
	msg := fmt.Sprintf("Found %d files that need converting.", len(filePaths))
	logger.Log(msg)

	// Convert found files
	for _, filePath := range filePaths {
		logger.Log(fmt.Sprintf("Converting file %s.", filePath))
		err = convertWebp2Jpg(filePath)
		if err != nil {
			// We failed to convert using dwebp
			logger.Log(fmt.Sprintf("Convert failed : %s",err.Error()))
		}
	}

	// Log that we are finished
	logger.Log("Finished converting.\n")
}

func convertWebp2Jpg(webpFilePath string) error {
	jpgFilePath := getJpgFilename(webpFilePath)

	// Run dwebp to convert the image
	_, err := exec.Command("dwebp", webpFilePath,"-o", jpgFilePath).Output()
	if err != nil {
		return err
	}

	// Remove webp file, if jpg exists
	if _, err := os.Stat(jpgFilePath); os.IsNotExist(err) {
		err = os.Remove(webpFilePath)
		if err != nil {
			return err
		}
		logger.Log(fmt.Sprintf("Removed file %s.", webpFilePath))
	}

	logger.Log(fmt.Sprintf("Successfully converted %s to %s.", webpFilePath, jpgFilePath))

	return nil
}

func getJpgFilename(filePath string) string {
	var extension = filepath.Ext(filePath)
	return path.Join(filePath[0:len(filePath)-len(extension)]) + ".jpg"
}
