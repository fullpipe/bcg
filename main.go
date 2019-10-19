package main

import "fmt"
import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	fmt.Println("vim-go")
	imagick.Initialize()
	defer imagick.Terminate()
	//	var err error

	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()
	pw.SetColor("black")

	mw.NewImage(640, 480, pw)

	dw.SetFont("Source Code Pro Light")
	dw.SetFontSize(30)
	pw.SetColor("white")
	dw.SetFillColor(pw)
	//pw.SetColor("black")
	dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	dw.SetGravity(imagick.GRAVITY_NORTH)

	textHeight := 0
	for {
		mw.AnnotateImage(dw, 0, float64(textHeight), 0, "THE ART TEACHER IS FAT, \nNOT PREGNANT")
		textHeight += 80
		if textHeight > 480 {
			break
		}
	}
	mw.AnnotateImage(dw, 0, 0, 0, "THE ART TEACHER IS FAT, \nNOT PREGNANT")
	mw.AnnotateImage(dw, 0, 80, 0, "THE ART TEACHER IS FAT, \nNOT PREGNANT")
	//mw.(dw, 0, 0, 0, "THE ART TEACHER IS FAT, NOT PREGNANT")

	mw.WriteImage("out.png")
	return
	//	pw := imagick.NewPixelWand()
	//	pw.SetColor("black")
	//	pw.SetAlpha(0.5)
	//
	//	err = mw.ReadImage("logo:")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	width := mw.GetImageWidth()
	//	height := mw.GetImageHeight()
	//
	//	hWidth := float64(width / 2)
	//	hHeight := float64(height / 2)
	//
	//	fmt.Println(hWidth, hHeight)
	//
	//	err = mw.ResizeImage(uint(hWidth), uint(hHeight), imagick.FILTER_LANCZOS)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)
	//	mw.DistortImage(imagick.DISTORTION_PERSPECTIVE, []float64{
	//		0, 0, 0, 0,
	//		0, hHeight, 0, hHeight,
	//		hWidth, hHeight, hWidth, hHeight - 90,
	//		hWidth, 0, hWidth, 90,
	//	}, true)
	//
	//	tw := imagick.NewMagickWand()
	//	tw.NewImage(uint(hWidth), uint(hHeight), pw)
	//
	//	tw.ResetIterator()
	//	tw.SetFirstIterator()
	//	//tw.AddImage(mw)
	//	//tw = tw.AppendImages(true)
	//	tw.CompositeImage(mw, imagick.COMPOSITE_OP_ATOP, true, 0, 0)
	//	err = mw.SetImageCompressionQuality(95)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	out := "out.png"
	//
	//	if err = tw.WriteImage(out); err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println("done:", out)

}
