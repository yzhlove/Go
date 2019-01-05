package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const imageSize = 300

func main() {

	pic := image.NewGray(image.Rect(0, 0, imageSize, imageSize))

	for i := 0; i < imageSize; i++ {
		for j := 0; j < imageSize; j++ {
			pic.SetGray(i, j, color.Gray{Y: 255})
		}
	}

	for x := 0; x < imageSize; x++ {
		s := float64(x) * 2 * math.Pi / imageSize
		y := imageSize/2 - math.Sin(s)*imageSize/2
		pic.SetGray(x, int(y), color.Gray{Y: 0})
	}

	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatalln(err)
	}
	png.Encode(file, pic)
	file.Close()
}
