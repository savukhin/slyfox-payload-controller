package utils

import (
	"math"
)

func Lerp(val, a_src, b_src, a_dst, b_dst float64) float64 {
	a1 := math.Min(a_src, b_src)
	b1 := math.Max(a_src, b_src)

	a2 := math.Min(a_dst, b_dst)
	b2 := math.Max(a_dst, b_dst)

	len_src := (b1 - a1)
	len_dst := (b2 - a2)

	return ((val-a1)/len_src)*len_dst + a_dst
}
