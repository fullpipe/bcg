package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path[1:])

		text := string(r.URL.Path[1:])
		file, err := os.Open(generateChalkBoard(text))
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "image/png")
		http.ServeContent(w, r, "tmp/"+text+".png", time.Now().AddDate(0, 0, -1), file)
	})

	log.Fatalln(http.ListenAndServe(":8089", nil))
	//generateChalkBoard("asdasdasd")
}

func generateChalkBoard(text string) string {
	hash := md5.Sum([]byte(text))
	log.Println(hex.EncodeToString(hash[:]))

	file := "tmp/" + hex.EncodeToString(hash[:]) + ".png"
	if fileExists(file) {
		return file
	}

	text = strings.ToUpper(text)
	board := newBoard(text)
	cw := imagick.NewMagickWand()
	cw.ReadImage("Chalk_Gag_Season_1_Epicsode_7_.png")
	cw.SharpenImage(0, 2)
	cw.CompositeImage(board, imagick.COMPOSITE_OP_DST_OVER, true, -1, -116)

	pw := imagick.NewPixelWand()
	pw.SetColor("white")

	bg := imagick.NewMagickWand()
	bg.NewImage(cw.GetImageWidth(), cw.GetImageHeight(), pw)
	cw.CompositeImage(bg, imagick.COMPOSITE_OP_DST_OVER, true, 0, 0)

	cw.SetImageCompressionQuality(100)
	cw.WriteImage(file)

	return file
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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

	dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	dw.SetGravity(imagick.GRAVITY_NORTH)

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
