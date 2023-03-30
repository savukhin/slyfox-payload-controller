package analyzers

import (
	"image"
	"math"

	"gocv.io/x/gocv"
)

type CameraPoint image.Point
type CameraResolution = CameraPoint
type Clusters []CameraPoint

type ImageAnalyzer struct {
	min_area int
	max_area int
}

func NewImageAnalyzer(min_area int, max_area int) *ImageAnalyzer {
	return &ImageAnalyzer{
		min_area: min_area,
		max_area: max_area,
	}
}

func (analyzer *ImageAnalyzer) Analyze(img gocv.Mat) Clusters {
	contours := gocv.FindContours(img, gocv.RetrievalTree, gocv.ChainApproxSimple)
	gocv.CvtColor(img, &img, gocv.ColorGrayToBGRA)

	clusters := make(Clusters, 0)

	for i := range contours.ToPoints() {
		rect := gocv.BoundingRect(contours.At(i))
		area := math.Pow(float64(rect.Dx()+rect.Dy()), 2)

		if area < float64(analyzer.min_area) || area > float64(analyzer.max_area) {
			continue
		}

		point := image.Point{
			X: rect.Max.X - rect.Dx()/2,
			Y: rect.Max.Y - rect.Dy()/2,
		}

		clusters = append(clusters, CameraPoint(point))
	}

	return clusters
}
