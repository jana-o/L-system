package turtlego

import (
	"code/lsystem-v2/ls"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

type TurtleGo struct {
	Image    *image.RGBA
	Pos      Position
	Rotation float64
	Color    color.Color
	Draw     bool
}

type Position struct {
	X, Y float64
}

func NewTurtleGo(i *image.RGBA, start Position) (t *TurtleGo) {
	t = &TurtleGo{
		Image:    i,
		Pos:      start,
		Rotation: 0.0,
		Color:    color.Gray{0xA9},
		Draw:     true,
	}

	return
}

func (t *TurtleGo) Forward(dist float64) {
	for i := 0; i < int(dist); i++ {
		if t.Draw {
			t.Image.Set(int(t.Pos.X), int(t.Pos.Y), t.Color)
		}

		x := 1.0 * math.Sin(t.Rotation)
		y := 1.0 * -math.Cos(t.Rotation)

		t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
	}
}

func (t *TurtleGo) Backward(dist float64) {
	t.Forward(-dist)
}

func (t *TurtleGo) Right(radians float64) {
	t.Rotation += radians
}

func (t *TurtleGo) Left(radians float64) {
	t.Right(-radians)
}

func (t *TurtleGo) PenUp() {
	t.Draw = false
}

func (t *TurtleGo) PenDown() {
	t.Draw = true
}

func ToImage(l *ls.Lsystem) image.Image {

	image := image.NewRGBA(image.Rect(0, 0, 300, 300))
	pos := Position{150.0, 150.0}
	t := NewTurtleGo(image, pos)

	for i := 0; i < 2; i++ {
		fields := strings.Fields("F F - F")
		for j := 0; j < len(fields); j++ {
			switch fields[j] {
			case "F":
				t.Forward(40.0)
			case "+":
				t.Right(math.Pi / 6)
			case "-":
				t.Left(math.Pi / 6)
			default:
				fmt.Println("unkown:" + fields[j])
			}
		}
	}
	return image
}

func saveImage(image image.Image, path string) {

	myfile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer myfile.Close()

	png.Encode(myfile, image)
	fmt.Println("newfile created ", path)
}

//ToPNG saves image as png
func ToPNG(l *ls.Lsystem, path string) {
	saveImage(ToImage(l), path)
}
