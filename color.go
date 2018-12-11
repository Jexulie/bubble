package main

import "math"

func getMax(num ...float64) float64 {
	var max float64
	for _, n := range num {
		if n > max {
			max = n
		}
	}
	return max
}

func getMin(num ...float64) float64 {
	min := num[0]
	for _, n := range num {
		if n < min {
			min = n
		}
	}
	return min
}

func calibrate(num float64) float64 {
	if num < 0 {
		return num + 1
	} else if num > 1 {
		return num - 1
	} else {
		return num
	}
}

func test(num, temp, temp2 float64) float64 {
	if (6 * num) < 1 {
		num = temp2 + ((temp - temp2) * (6 * num))
	} else if (2 * num) < 1 {
		num = temp
	} else if (3 * num) < 2 {
		num = temp2 + ((temp - temp2) * ((0.666 - num) * 6))
	} else {
		num = temp2
	}
	return num
}

// RGBCOLOR c
type RGBCOLOR struct {
	R float64
	G float64
	B float64
}

// HSLCOLOR x
type HSLCOLOR struct {
	H float64
	S float64
	L float64
}

// Blend x
func Blend(first RGBCOLOR, second RGBCOLOR) RGBCOLOR {
	return RGBCOLOR{
		first.R - (first.R-second.R)/2,
		first.G - (first.G-second.G)/2,
		first.B - (first.B-second.B)/2,
	}
}

// Darken x
func Darken(first RGBCOLOR) RGBCOLOR {
	return RGBCOLOR{
		first.R - (first.R-0)/2,
		first.G - (first.G-0)/2,
		first.B - (first.B-0)/2,
	}
}

// Brigthen x
func Brigthen(first RGBCOLOR) RGBCOLOR {
	return RGBCOLOR{
		first.R - (first.R-255)/2,
		first.G - (first.G-255)/2,
		first.B - (first.B-255)/2,
	}
}

// ConvertToHSL x
func ConvertToHSL(rgb RGBCOLOR) HSLCOLOR {

	var newColor HSLCOLOR

	r := rgb.R / 255
	g := rgb.G / 255
	b := rgb.B / 255

	max := getMax(r, g, b)
	min := getMin(r, g, b)

	c := max - min

	if c == 0 {
		newColor.H = 0
	} else if max == r {
		newColor.H = float64(int((g-b)/c) % 6)
	} else if max == g {
		newColor.H = ((b - r) / c) + 2
	} else if max == b {
		newColor.H = ((r - g) / c) + 4
	}

	newColor.H *= 60

	newColor.L = (max + min) * 0.5

	if c == 0 {
		newColor.S = 0
	} else if newColor.L > 0.5 {
		newColor.S = c / (2 - max - min)
	} else {
		newColor.S = c / (min + max)
	}

	newColor.H = math.Round(newColor.H)
	newColor.S = math.Round(newColor.S * 100)
	newColor.L = math.Round(newColor.L * 100)

	if newColor.H < 0 {
		newColor.H *= -1
	}

	if newColor.S < 0 {
		newColor.S *= -1
	}

	if newColor.L < 0 {
		newColor.L *= -1
	}

	return newColor
}

// ConvertToRGB may not need ?
func ConvertToRGB(hsl HSLCOLOR) RGBCOLOR {

	h := hsl.H / 360
	s := hsl.S / 100
	l := hsl.L / 100

	if s == 0 {
		return RGBCOLOR{
			R: l * 255,
			G: l * 255,
			B: l * 255,
		}
	}

	var temp float64
	var temp2 float64

	if l < 0.5 {
		temp = l * (1 + s)
	} else {
		temp = (l + s) * (l * s)
	}

	temp2 = (2 * l) - temp

	tempR := calibrate(h + 0.333)
	tempG := calibrate(h)
	tempB := calibrate(h - 0.333)

	return RGBCOLOR{
		R: math.Round(test(tempR, temp, temp2) * 255),
		G: math.Round(test(tempG, temp, temp2) * 255),
		B: math.Round(test(tempB, temp, temp2) * 255),
	}
}
