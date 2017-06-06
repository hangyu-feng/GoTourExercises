/*自定义的 Image 类型，要实现必要的方法，并且调用 pic.ShowImage

Bounds 应当返回一个 image.Rectangle，例如 `image.Rect(0, 0, w, h)`

ColorModel 应当返回 color.RGBAModel

At 应当返回一个颜色；在这个例子里，在最后一个图片生成器的值 v 匹配 `color.RGBA{v, v, 255, 255}`*/
package main

import (
	"image"
	"image/color"
)

type Image struct {
	w, h    int // width and height
	picture map[image.Point]color.Color
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// TODO: ColorModel() color.Model  // should return color.RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// TODO: At(x int, y int) color.RGBA
func (i Image) At(x, y int) color.Color {
	point := image.Point{x, y}
	return i.picture[point]
}
