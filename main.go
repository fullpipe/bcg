package main

import (
	"fmt"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	board := newBoard("I WILL NOT DRAW NAKED LADIES IN CLASS")
	cw := imagick.NewMagickWand()
	cw.ReadImage("Chalk_Gag_Season_1_Epicsode_7_.png")
	cw.SharpenImage(0, 2)
	cw.CompositeImage(board, imagick.COMPOSITE_OP_DST_OVER, true, 0, -116)

	cw.SetImageCompressionQuality(100)
	cw.WriteImage("out.png")
}

func newBoard(text string) *imagick.MagickWand {
	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()
	pw.SetColor("black")

	w := 860.0
	h := 690.0
	mw.NewImage(uint(w), uint(h), pw)

	dw.SetFont("./Pangolin-Regular.ttf")
	dw.SetFontSize(30)
	pw.SetColor("white")
	dw.SetFillColor(pw)
	//pw.SetColor("black")
	dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	dw.SetGravity(imagick.GRAVITY_NORTH)
	fonts := mw.QueryFonts("*")
	fmt.Println(fonts)
	textHeight := 100
	for {
		mw.AnnotateImage(dw, 0, float64(textHeight), 0, text)
		textHeight += 50
		if textHeight > int(h-100) {
			break
		}
	}

	ki, err := imagick.NewKernelInfo("Octagon:1")
	if err != nil {
		panic(err)
	}
	mw.MorphologyImage(imagick.MORPHOLOGY_ERODE_INTENSITY, 1, ki)

	mw.BlurImage(0, 2)
	mw.SharpenImage(0, 2)
	mw.SharpenImage(0, 2)
	mw.SharpenImage(0, 2)

	mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)
	mw.DistortImage(imagick.DISTORTION_PERSPECTIVE, []float64{
		0, 0, 0, 0,
		0, h, 0, h,
		w, h, w, h - 78,
		w, 0, w, 116,
	}, true)

	return mw
}
