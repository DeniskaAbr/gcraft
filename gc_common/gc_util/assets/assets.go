package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	"log"

	"golang.org/x/image/bmp"
)

const (
	CharWidth  = 8
	CharHeight = 16
)

func CreateTextImage() image.Image {
	s, err := gzip.NewReader(bytes.NewReader(CompressedDebugText))
	if err != nil {
		panic(fmt.Sprintf("assets: gzip.NewReader failed: %v", err))
	}

	debugBmp, err := bmp.Decode(s)
	if err != nil {
		panic(fmt.Sprintf("assets: bmp.Decode failed: %v", err))
	}

	err = s.Close()
	if err != nil {
		log.Print(err)
	}

	return debugBmp
}
