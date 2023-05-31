package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle { //实现Image包中生成图片边界的方法
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color { //实现Image包中生成图像某个点的方法
	return color.RGBA{R: uint8(x), G: uint8(y), B: uint8(255), A: uint8(255)}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
