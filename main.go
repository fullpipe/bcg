package main

import "fmt"
import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	fmt.Println("vim-go")
	imagick.Initialize()
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage("logo:")
	if err != nil {
		panic(err)
	}

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		panic(err)
	}

	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	out := "out.png"

	if err = mw.WriteImage(out); err != nil {
		panic(err)
	}

	fmt.Println("done:", out)

}
