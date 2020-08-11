package main

import (
	"fmt"
	"image/jpeg"
	"golang.org/x/image/webp"
	"os"
	"path"
)

func main() {

}

func convertWebp2Jpg(directory, fileName string) error {
	f, err := os.Open(path.Join(directory, fileName))
	if err!=nil {
		fmt.Printf("%s", err.Error())
	}
	img, err:=webp.Decode(f)
	if err!=nil {
		fmt.Printf("%s", err.Error())
	}
	out,err := os.Create(path.Join(directory, fileName, ".jpg"))
	jpeg.Encode(out, img, nil)

	return nil
}