package main

import (
	"log"
	"strconv"

	"github.com/fogleman/gg"
)

const (
	file      = "./TX Tileset Wall.png"
	tile_size = 32.0
)

func main() {
	src, err := gg.LoadImage(file)
	if err != nil {
		log.Fatalln(err.Error() + file)
	}

	dst := gg.NewContextForImage(src)
	if err := dst.LoadFontFace("04B_03__.TTF", 11); err != nil {
		log.Fatalln(err)
	}

	count := 1
	// get width and height from src img then divided by tile_size to get cols and rows
	col := src.Bounds().Dx() / tile_size
	height := src.Bounds().Dy() / tile_size
	dst.SetRGB255(255, 255, 255)
	for y := 0; y < height; y++ {
		for x := 0; x < col; x++ {
			// dst.Push()
			// dst.SetRGB255(255, 0, 255)
			// dst.DrawString(strconv.Itoa(count), float64(x)*tile_size+1, float64(y+1)*tile_size-2)
			// dst.Pop()
			dst.DrawString(strconv.Itoa(count), float64(x)*tile_size, float64(y+1)*tile_size-2)
			dst.DrawString(strconv.Itoa(count), float64(x)*tile_size, float64(y+1)*tile_size-1)
			count++
		}
	}

	// vẽ lưới
	for y := 1; y < height; y++ {
		var row float64 = float64(y) * tile_size
		dst.MoveTo(0, row)
		dst.LineTo(float64(dst.Width()), row)
		dst.Stroke()
	}
	for x := 1; x < height; x++ {
		var col float64 = float64(x) * tile_size
		dst.MoveTo(col, 0)
		dst.LineTo(col, float64(dst.Height()))
		dst.Stroke()
	}

	dst.SavePNG(file + "_numbered.png")
}
