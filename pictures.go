package rtfdoc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"io"
	"log"

	_ "image/jpeg"
	_ "image/png"
)

const (
	// JPGFORMAT - jpeg format
	JPEGFORMAT = "jpeg"
	// PNGFORMAT - png format
	PNGFORMAT = "png"
)

// Main Picture struct
type Picture struct {
	format string // EMF, PNG, JPEG
	src    []byte
	scaleX int
	scaleY int
	cropL  int
	cropR  int
	cropT  int
	cropB  int
	height int
	width  int
}

// AddPicture adds picture
func (par *Paragraph) AddPicture(source io.Reader, format string) *Picture {
	var pic Picture
	var err error

	formatFound := false
	for _, i := range []string{JPEGFORMAT, PNGFORMAT} {
		if format == i {
			formatFound = true
			break
		}
	}
	if !formatFound {
		log.Println("Unknown format")
		return &pic
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(source)
	pic.src = buf.Bytes()

	pic.format = format

	pic.scaleX = 100
	pic.scaleY = 100

	//Calculating dimensions
	pic.height, pic.width, err = getImageDimensions(pic.src)
	if err != nil {
		pic.height = 100
		pic.width = 100
	}

	par.content = append(par.content, &pic)
	return &pic
}

func (pic *Picture) SetWidth(width int) *Picture {
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
