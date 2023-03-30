package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

// type Center

func main() {
	img := gocv.IMRead("./suitable_img.png", gocv.IMReadGrayScale)

	fmt.Println(img)

	contours := gocv.FindContours(img, gocv.RetrievalTree, gocv.ChainApproxSimple)
	gocv.CvtColor(img, &img, gocv.ColorGrayToBGRA)

	clusters := make([]image.Point, contours.Size())

	for i := range contours.ToPoints() {
		rect := gocv.BoundingRect(contours.At(i))
		area := math.Pow(float64(rect.Dx()+rect.Dy()), 2)

		if area < 50 || area > 2000 {
			continue
		}

		clusters[i] = image.Point{
			X: (rect.Max.X + rect.Dx()) / 2,
			Y: (rect.Max.Y + rect.Dy()) / 2,
		}

		// for _, point := range contours.At(i).ToPoints() {
		// gocv.Circle(&img, point, 2, color.RGBA{R: 255, G: 0, B: 255, A: 255}, 1)
		// }

		// gocv.Rectangle(&img, rect, color., 2)
		gocv.Rectangle(&img, rect, color.RGBA{255, 0, 0, 255}, 1)
	}

	// fmt.Println(img.Cols(), img.Rows())
	// gocv.Rectangle(&img, image.Rectangle{
	// 	Min: image.Point{
	// 		100, 100,
	// 	},
	// 	Max: image.Point{500, 900},
	// }, color.RGBA{255, 0, 0, 255}, 3)
	// gocv.Rectangle(&img, image.Rect(100, 100, 500, 900), color.RGBA{255, 0, 0, 255}, 3)

	window := gocv.NewWindow("img")
	window.IMShow(img)
	window.WaitKey(0)

}
