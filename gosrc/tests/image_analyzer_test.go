package analyzers

import (
	"fmt"
	"slyfox-payload-controller/analyzers"
	"testing"

	"github.com/stretchr/testify/require"
	"gocv.io/x/gocv"
)

func TestImageAnalyzer(t *testing.T) {
	analyzer := analyzers.NewImageAnalyzer(1000, 400000)

	t.Log("Testing grayscale img with several clusters")
	{
		img := gocv.IMRead("./suitable_img.png", gocv.IMReadGrayScale)
		clusters := analyzer.Analyze(img)
		require.Equal(t, len(clusters), 11)
	}

	t.Log("Testing random img which is grayscaled")
	{
		img := gocv.IMRead("./img.png", gocv.IMReadGrayScale)
		clusters := analyzer.Analyze(img)
		require.Equal(t, len(clusters), 11)
	}

	t.Log("Testing random img with one cluster")
	{
		img := gocv.IMRead("./one_cluster_img.png", gocv.IMReadGrayScale)
		clusters := analyzer.Analyze(img)
		require.Equal(t, len(clusters), 1)
		fmt.Println(clusters)
	}
}
