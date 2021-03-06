package rtfdoc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"
)

// AddPicture adds picture
func (par *Paragraph) AddPicture(source []byte, format string) *Picture {
	var pic = Picture{
		paragraphWidth: par.maxWidth,
	}
	var err error

	formatFound := false
	for _, i := range []string{ImageFormatJpeg, ImageFormatPng} {
		if format == i {
			formatFound = true
			break
		}
	}
	if !formatFound {
		log.Println("Unknown format")
		return &pic
	}
	pic.updateMaxWidth()

	pic.src = source

	pic.format = format

	pic.scaleX = 100
	pic.scaleY = 100

	//Calculating dimensions
	pic.height, pic.width, err = getImageDimensions(pic.src)
	if err != nil {
		pic.height = 100
		pic.width = 100
	}
	if pic.width > getPixelsFromTwips(pic.maxWidth) {
		newWidth := getPixelsFromTwips(pic.maxWidth)
		pic.height = int(float64(pic.height) / (float64(pic.width) / float64(newWidth)))
		pic.width = newWidth
	}

	par.content = append(par.content, &pic)
	return &pic
}

func (pic *Picture) updateMaxWidth() *Picture {
	pic.maxWidth = pic.paragraphWidth
	return pic
}

func (pic *Picture) SetWidth(width int) *Picture {
	if getTwipsFromPixels(width) > pic.maxWidth {
		pic.width = getPixelsFromTwips(pic.maxWidth)
		return pic
	}
	pic.width = width
	return pic
}
func (pic *Picture) SetHeight(height int) *Picture {
	pic.height = height
	return pic
}

func (pic *Picture) SetScaleX(scaleX int) *Picture {
	pic.scaleX = scaleX
	return pic
}

func (pic *Picture) SetScaleY(scaleY int) *Picture {
	pic.scaleY = scaleY
	return pic
}

func (pic *Picture) SetCropLeft(cropL int) *Picture {
	pic.cropL = cropL
	return pic
}

func (pic *Picture) SetCropRight(cropR int) *Picture {
	pic.cropR = cropR
	return pic
}

func (pic *Picture) SetCropTop(cropT int) *Picture {
	pic.cropT = cropT
	return pic
}

func (pic *Picture) SetCropBottom(cropB int) *Picture {
	pic.cropB = cropB
	return pic
}

func (pic *Picture) compose() string {
	res := fmt.Sprintf("\n{\\*\\shppict{ \\pict\\picscalex%d\\picscaley%d\\piccropl%d\\piccropr%d\\piccropt%d\\piccropb%d\\picw%d\\pich%d\\picwgoal%d\\pichgoal%d\\%sblip",
		pic.scaleX, pic.scaleY,
		pic.cropL, pic.cropR, pic.cropT, pic.cropB,
		pic.width, pic.height,
		pic.width*15, pic.height*15,
		pic.format,
	)

	res += "\n" + hex.EncodeToString(pic.src)
	res += "\n}}"
	return res
}

func getImageDimensions(img []byte) (int, int, error) {
	i, _, err := image.DecodeConfig(bytes.NewReader(img))

	return i.Height, i.Width, err
}
