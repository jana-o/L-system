package turtlego

import (
	"code/lsystem-v6/ls"
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
	Stack    []float64
}

type Position struct {
	X, Y float64
}

//Push saves x,y coordinates and angle on stack
func (t *TurtleGo) Push() {
	t.Stack = append(t.Stack, t.Pos.X, t.Pos.Y, t.Rotation)
}

//Pop loads x,y coordinates and angle from stack
func (t *TurtleGo) Pop() {
	if len(t.Stack) <= 3 {
		t.Pos = Position{t.Stack[0], t.Stack[1]}
		t.Rotation = t.Stack[2]
	} else {
		t.Pos = Position{t.Stack[len(t.Stack)-3], t.Stack[len(t.Stack)-2]}
		t.Rotation = t.Stack[len(t.Stack)-1]
	}
}

//Forward moves forward
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

//Backward moves backwards
func (t *TurtleGo) Backward(dist float64) {
	t.Forward(-dist)
}

//Rotate changes direction of next move by radians degrees
func (t *TurtleGo) Rotate(radians float64) {
	t.Rotation += radians
}

func (t *TurtleGo) PenUp() {
	t.Draw = false
}

func (t *TurtleGo) PenDown() {
	t.Draw = true
}

//NewTurtleGo creates instance of TurtleGo
func NewTurtleGo(i *image.RGBA, start Position) (t *TurtleGo) {
	t = &TurtleGo{
		Image:    i,
		Pos:      start,
		Rotation: 0.0,
		Color:    color.Gray{0xA9},
		Draw:     true,
		Stack:    []float64{},
	}
	return
}

//ToImage translates generated string into geometric structure and creates image of lsystem
func ToImage(l *ls.Lsystem) image.Image {

	image := image.NewRGBA(image.Rect(0, 0, 300, 300))
	pos := Position{150.0, 300.0}
	t := NewTurtleGo(image, pos)
	r := l.Result[len(l.Result)-1]

	fmt.Println("enter toimage", len(r))

	for i := 0; i < 1; i++ {
		fields := strings.Split(r, "")
		if len(r) == 0 {
			t.Forward(40.0)
		} else {
			for j := 0; j < len(fields); j++ {
				switch fields[j] {
				case "F":
					t.PenDown()
					t.Forward(40.0)
				case "+":
					t.PenDown()
					t.Rotate(math.Pi / 6)
				case "-":
					t.PenDown()
					t.Rotate(-math.Pi / 6)
				case "[":
					t.Push()
				case "]":
					t.PenUp()
					t.Pop()
				default:
					fmt.Println("unkown:" + fields[j])
				}
			}
		}
	}
	return image
}

//saveImage creates image file
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
