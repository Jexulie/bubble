package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

// TO8BIT x
const TO8BIT = 257

// ImgProps x
type ImgProps struct {
	width       int
	height      int
	colorMatrix []Point
}

// Point p
type Point struct {
	x     int
	y     int
	color RGBCOLOR
}

func (i ImgProps) colorAt(x, y int) RGBCOLOR {
	for _, v := range i.colorMatrix {
		if v.x == x && v.y == y {
			return v.color
		}
	}
	return RGBCOLOR{0, 0, 0}
}

func readImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	imgData, _, err1 := image.Decode(file)
	if err1 != nil {
		panic(err1)
	}
	return imgData
}

// SetProps p
func SetProps(img image.Image) ImgProps {
	var iProps ImgProps
	r := img.Bounds()
	iProps.width = r.Size().X
	iProps.height = r.Size().Y
	for x := 0; x <= iProps.width; x++ {
		for y := 0; y <= iProps.height; y++ {
			red, green, blue, _ := img.At(x, y).RGBA()
			color := RGBCOLOR{float64(red / TO8BIT), float64(green / TO8BIT), float64(blue / TO8BIT)}
			point := Point{x, y, color}
			iProps.colorMatrix = append(iProps.colorMatrix, point)
		}
	}
	return iProps
}

func main() {
	img := readImage("kawaii.jpg")
	p := SetProps(img)
	// for _, k := range p.colorMatrix {
	// 	fmt.Println(ConvertToHSL(k.color))
	// }

	l := ScaleImage(10, p)
	f, _ := os.Create("2xScaled.png")
	png.Encode(f, l)

	// im := image.NewRGBA(image.Rect(0, 0, p.width, p.height))
	// for x := 0; x < p.width; x++ {
	// 	for y := 0; y < p.height; y++ {
	// 		c := p.colorAt(x, y)
	// 		c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
	// 		im.Set(x, y, c1)
	// 	}
	// }

	// f, _ := os.Create("clone.png")
	// png.Encode(f, im)

}
