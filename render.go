package webGerber

import (
	"image/color"

	v1 "github.com/Kirizu-Official/webgerber-render-go/v1"
)

func NewPCBRender(plotData []byte, opt *v1.RenderOptions) (*v1.PCBRender, error) {
	if opt == nil {
		opt = &v1.RenderOptions{
			Zoom:       100,
			DarkColor:  color.Black,
			ClearColor: &color.RGBA{R: 0, G: 0, B: 0, A: 0},
		}
	}
	if opt.Zoom == 0 {
		opt.Zoom = 100
	}
	if opt.ClearColor == nil {
		opt.ClearColor = &color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}
	if opt.DarkColor == nil {
		opt.DarkColor = color.Black
	}
	return v1.NewPCBRender(plotData, opt)
}
