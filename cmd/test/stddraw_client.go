package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/draw"
	"github.com/efuquen/algorithms/pkg/algs4/random"
	"math"
)

func randomBarGraph() {
	n := 50
	var a []float64
	for i := 0; i < n; i++ {
		a = append(a, random.Uniform())
	}
	for i := 0; i < n; i++ {
		x := 1.0 * float32(i) / float32(n)
		y := float32(a[i] / 2.0)
		rw := 0.5 / float32(n)
		rh := float32(a[i] / 2.0)
		fmt.Printf("x: %f y: %f rw: %f rh: %f\n", x, y, rw, rh)
		draw.FilledRectangle(x, y, rw, rh)
	}
}

func pointGraphs() {
	n := 100
	draw.SetXScale(0.0, float32(n))
	draw.SetYScale(0.0, float32(n*n))
	draw.SetPenRadius(0.01)

	for i := float32(1); i <= float32(n); i++ {
		draw.Point(i, i)
		draw.Point(i, i*i)
		draw.Point(i, i*float32(math.Log(float64(i))))
	}
}

func main() {
	//randomBarGraph()
	pointGraphs()
	draw.ShowAndRun()
}
