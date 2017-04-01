package imageproxy

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strconv"
)

func generateSolidColorRec(colorHex string, w int, h int) ([]byte, error) {
	r, _ := strconv.ParseInt(colorHex[0:2], 16, 0)
	g, _ := strconv.ParseInt(colorHex[2:4], 16, 0)
	b, _ := strconv.ParseInt(colorHex[4:6], 16, 0)
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, m)
	return buf.Bytes(), err
}
