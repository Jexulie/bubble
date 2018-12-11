package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

const TO8BIT = 257

type ImgProps struct {
	width       int
	height      int
	colorMatrix []Point
}

type Point struct {
	x     int
	y     int
	color RGBCOLOR
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

func setProps(img image.Image) ImgProps {
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
	p := setProps(img)
	for _, k := range p.colorMatrix {
		fmt.Println(ConvertToHSL(k.color))
	}
}
