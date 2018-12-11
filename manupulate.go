package main

import (
	"image"
	"image/color"
	"log"
	"os"
)

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// ConvertToGrayscale a
// func ConvertToGrayscale(img image.Image) image.Image {

// }

// ScaleImage s
func ScaleImage(times int, img ImgProps) image.Image {
	// if its square
	// what if not ? :(
	isEven := isEven(times)
	pixelsEach := (times - 1)
	if times < 0 {
		log.Fatalln("Times must be positive")
		os.Exit(1)
	}
	im := image.NewRGBA(image.Rect(0, 0, img.width*times, img.height*times))
	for x := 0; x < img.width*times; x++ {
		for y := 0; y < img.height*times; y++ {
			if isEven {
				if x%2 == 0 && y%2 == 0 {
					if x == 0 || y == 0 {
						c := img.colorAt(0, 0)
						c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
						im.Set(x, y, c1)
					} else {
						c := img.colorAt(x/times, y/times)
						c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
						im.Set(x, y, c1)
					}
				} else {
					c := img.colorAt(x/times, y/times)
					c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
					im.Set(x, y, c1)

					for h := 1; h < pixelsEach; h++ {
						im.Set(x+h, y, c1)
						im.Set(x, y+h, c1)
						im.Set(x+h, y+h, c1)
					}
				}
			} else {
				if x%2 != 0 && y%2 != 0 {
					if x == 0 || y == 0 {
						c := img.colorAt(0, 0)
						c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
						im.Set(x, y, c1)
					} else {
						c := img.colorAt(x/times, y/times)
						c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
						im.Set(x, y, c1)
					}
				} else {
					c := img.colorAt(x/times, y/times)
					c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
					im.Set(x, y, c1)

					for h := 1; h < pixelsEach; h++ {
						im.Set(x+h, y, c1)
						im.Set(x, y+h, c1)
						im.Set(x+h, y+h, c1)
					}
				}
			}
		}
	}
	return im
}

// CropImage x
func CropImage(Ox, Oy, Px, Py int, img ImgProps) *image.RGBA {
	if Ox+Px > img.width {
		log.Fatalln("Wrong Parameters")
		os.Exit(1)
	}

	if Oy+Py > img.height {
		log.Fatalln("Wrong Parameters")
		os.Exit(1)
	}
	im := image.NewRGBA(image.Rect(0, 0, Px, Py))
	for x := 0; x < Px; x++ {
		for y := 0; y < Py; y++ {
			c := img.colorAt(x+Ox, y+Oy)
			c1 := color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 0xff}
			im.Set(x, y, c1)
		}
	}
	return im
}
