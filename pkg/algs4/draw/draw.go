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
	if name == theme.ColorNameBackground {
		return White
	}
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
var White color.Color = color.White
var Blue color.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var Red color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var Green color.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

var DefaultPenColor color.Color = Black

const DefaultXMin = 0.0
const DefaultXMax = 1.0
const DefaultYMin = 0.0
const DefaultYMax = 1.0

var a = app.New()

func init() {
	a.Settings().SetTheme(&myTheme{})
}

type Draw struct {
	window                 fyne.Window
	content                *fyne.Container
	penRadius              float32
	penColor               color.Color
	width                  float32
	height                 float32
	xmin, xmax, ymin, ymax float32
}

func Run() {
	a.Run()
}

func NewDraw(title string) Draw {
	d := Draw{
		a.NewWindow(title),
		container.NewWithoutLayout(),
		DefaultPenRadius,
		DefaultPenColor,
		DefaultSize,
		DefaultSize,
		DefaultXMin,
		DefaultXMax,
		DefaultYMin,
		DefaultYMax,
	}
	d.window.Resize(fyne.NewSize(DefaultSize, DefaultSize))
	d.window.CenterOnScreen()
	d.window.SetContent(d.content)

	return d
}

func (d *Draw) SetPenRadius(r float32) {
	d.penRadius = r
}

func (d *Draw) SetPenColor(color color.Color) {
	d.penColor = color
}

func (d *Draw) SetXScale(min float32, max float32) {
	d.xmin = min
	d.xmax = max
}

func (d *Draw) SetYScale(min float32, max float32) {
	d.ymin = min
	d.ymax = max
}

func (d *Draw) scaleX(x float32) float32 {
	return d.width * (x - d.xmin) / (d.xmax - d.xmin)
}

func (d *Draw) scaleY(y float32) float32 {
	return d.height * (d.ymax - y) / (d.ymax - d.ymin)
}

func (d *Draw) factorX(w float32) float32 {
	return w * d.width / float32(math.Abs(float64(d.xmax-d.xmin)))
}

func (d *Draw) factorY(h float32) float32 {
	return h * d.height / float32(math.Abs(float64(d.ymax-d.ymin)))
}

func (d *Draw) Point(x, y float32) {
	xs := d.scaleX(x)
	ys := d.scaleY(y)
	scaledPenRadius := d.penRadius * DefaultSize

	circle := canvas.NewCircle(d.penColor)
	circle.Move(fyne.NewPos(xs-scaledPenRadius/2, ys-scaledPenRadius/2))
	circle.Resize(fyne.NewSize(scaledPenRadius, scaledPenRadius))
	d.content.Add(circle)
	d.window.Show()
}

func (d *Draw) Circle(x, y, radius float32) {
	xs := d.scaleX(x)
	ys := d.scaleY(y)
	ws := d.factorX(2 * radius)
	hs := d.factorY(2 * radius)

	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = d.penColor
	circle.StrokeWidth = d.penRadius * DefaultSize
	circle.Move(fyne.NewPos(xs-ws/2, ys-hs/2))
	circle.Resize(fyne.NewSize(ws, hs))
	d.content.Add(circle)
	d.window.Show()
}

func (d *Draw) FilledRectangle(x, y, halfWidth, halfHeight float32) {
	xs := d.scaleX(x)
	ys := d.scaleY(y)
	ws := d.factorX(2 * halfWidth)
	hs := d.factorY(2 * halfHeight)

	rect := canvas.NewRectangle(d.penColor)
	rect.Move(fyne.NewPos(xs-ws/2, ys-hs/2))
	rect.Resize(fyne.NewSize(ws, hs))
	d.content.Add(rect)
	d.window.Show()
}
