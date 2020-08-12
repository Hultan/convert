package main

import (
	"fmt"
	"golang.org/x/image/webp"
	"image/jpeg"
	"os"
	"path"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("/softtube/thumbnails/*.webp")
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		return
	}

	for _,file := range files {
		convertWebp2Jpg(file)
	}
}

func convertWebp2Jpg(filePath string) error {
	f, err := os.Open(filePath)
	if err!=nil {
		fmt.Printf("%s", err.Error())
	}
	img, err:=webp.Decode(f)
	if err!=nil {
		fmt.Printf("%s", err.Error())
	}

	out,err := os.Create(getJpgFilename(filePath))
	jpeg.Encode(out, img, nil)

	err=os.Remove(filePath)
	if err!=nil {
		fmt.Printf("%s", err.Error())

	}

	return nil
}

func getJpgFilename(filePath string) string {
	var extension = filepath.Ext(filePath)
	return path.Join(filePath[0:len(filePath)-len(extension)]) + ".jpg"
}