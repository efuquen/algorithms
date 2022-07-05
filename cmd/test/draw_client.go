package main

import (
	"github.com/efuquen/algorithms/pkg/algs4/draw"
	"github.com/efuquen/algorithms/pkg/algs4/random"
	"math"
)

func randomBarGraph() {
	d := draw.NewDraw(draw.DefaultWindowTitle)
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
		d.FilledRectangle(x, y, rw, rh)
	}
}

func pointGraphs() {
	d := draw.NewDraw()
	n := 100
	d.SetXScale(0.0, float32(n))
	d.SetYScale(0.0, float32(n*n))
	d.SetPenRadius(0.01)

	for i := float32(1); i <= float32(n); i++ {
		d.Point(i, i)
		d.Point(i, i*i)
		d.Point(i, i*float32(math.Log(float64(i))))
	}
}

func main() {
	//randomBarGraph()
	pointGraphs()
	draw.Run()
}
