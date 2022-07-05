package draw

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"math"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}
func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 0
	}
	return theme.DefaultTheme().Size(name)
}

const DefaultWindowTitle = "Standard Draw"

const DefaultSize float32 = 512
const DefaultPenRadius float32 = 0.002

var Black color.Color = color.Black
var Blue color.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var Red color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var Green color.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

var DefaultPenColor color.Color = Black

const Border = 0.00
const DefaultXMin = 0.0
const DefaultXMax = 1.0
const DefaultYMin = 0.0
const DefaultYMax = 1.0

var xmin, xmax, ymin, ymax float32

var penRadius = DefaultPenRadius
var penColor color.Color = DefaultPenColor
var width = DefaultSize
var height = DefaultSize

var a = app.New()
var w = a.NewWindow(DefaultWindowTitle)
var content = container.NewWithoutLayout()

func init() {
	w.Resize(fyne.NewSize(DefaultSize, DefaultSize))
	w.CenterOnScreen()
	w.SetContent(content)
	a.Settings().SetTheme(&myTheme{})

	SetXScale(DefaultXMin, DefaultXMax)
	SetYScale(DefaultYMin, DefaultYMax)
}

func SetPenRadius(r float32) {
	penRadius = r
}

func SetPenColor(color color.Color) {
	penColor = color
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

func factorX(w float32) float32 {
	return w * width / float32(math.Abs(float64(xmax-xmin)))
}

func factorY(h float32) float32 {
	return h * height / float32(math.Abs(float64(ymax-ymin)))
}

func Point(x, y float32) {
	xs := scaleX(x)
	ys := scaleY(y)
	scaledPenRadius := penRadius * DefaultSize

	circle := canvas.NewCircle(penColor)
	circle.Move(fyne.NewPos(xs-scaledPenRadius/2, ys-scaledPenRadius/2))
	circle.Resize(fyne.NewSize(scaledPenRadius, scaledPenRadius))
	content.Add(circle)
}

func Circle(x, y, radius float32) {
	xs := scaleX(x)
	ys := scaleY(y)
	ws := factorX(2 * radius)
	hs := factorY(2 * radius)

	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = penColor
	circle.StrokeWidth = penRadius * DefaultSize
	circle.Move(fyne.NewPos(xs-ws/2, ys-hs/2))
	circle.Resize(fyne.NewSize(ws, hs))
	content.Add(circle)
}

func FilledRectangle(x, y, halfWidth, halfHeight float32) {
	xs := scaleX(x)
	ys := scaleY(y)
	ws := factorX(2 * halfWidth)
	hs := factorY(2 * halfHeight)

	rect := canvas.NewRectangle(penColor)
	rect.Move(fyne.NewPos(xs-ws/2, ys-hs/2))
	rect.Resize(fyne.NewSize(ws, hs))
	content.Add(rect)
}

func ShowAndRun() {
	w.ShowAndRun()
}
