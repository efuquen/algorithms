package stddraw

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const DefaultWindowTitle = "Standard Draw"

const DefaultSize float32 = 512
const DefaultPenRadius float32 = 0.002

const Border = 0.00
const DefaultXMin = 0.0
const DefaultXMax = 1.0
const DefaultYMin = 0.0
const DefaultYMax = 1.0

var xmin, xmax, ymin, ymax float32

var penRadius = DefaultPenRadius
var width = DefaultSize
var height = DefaultSize

var drawApp = app.New()
var w = drawApp.NewWindow(DefaultWindowTitle)
var content = container.NewWithoutLayout()

func init() {
	w.Resize(fyne.NewSize(DefaultSize, DefaultSize))
	w.CenterOnScreen()
	w.SetContent(content)

	SetXScale(DefaultXMin, DefaultXMax)
	SetYScale(DefaultYMin, DefaultYMax)
}

func SetPenRadius(r float32) {
	penRadius = r
}

func SetXScale(min float32, max float32) {
	size := max - min
	xmin = min - Border*size
	xmax = max + Border*size
}

func SetYScale(min float32, max float32) {
	size := max - min
	ymin = min - Border*size
	ymax = max + Border*size
}

func scaleX(x float32) float32 {
	return width * (x - xmin) / (xmax - xmin)
}

func scaleY(y float32) float32 {
	return height * (ymax - y) / (ymax - ymin)
}

func Point(x, y float32) {
	xs := scaleX(x)
	ys := scaleY(y)
	scaledPenRadius := penRadius * DefaultSize

	circle := canvas.NewCircle(color.White)
	circle.Move(fyne.NewPos(xs-scaledPenRadius/2, ys-scaledPenRadius/2))
	circle.Resize(fyne.NewSize(scaledPenRadius, scaledPenRadius))
	content.Add(circle)
	w.ShowAndRun()
}
