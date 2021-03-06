/*
https://go.dev/tour/methods/25
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
*/

package main

import (
"golang.org/x/tour/pic"
    "image"
    "image/color"
)


type Image struct{
    width int
    height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	v := x*x*y*y
    return color.RGBA{uint8(v), uint8(v), 255, 255}
}


func main() {
    m := Image{256, 64}
    pic.ShowImage(m)
}

