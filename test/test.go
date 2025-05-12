package main

import (
	"image/color"
	"os"

	webGerber "github.com/Kirizu-Official/webgerber-render-go"
	v1 "github.com/Kirizu-Official/webgerber-render-go/v1"
)

func main() {
	file, err := os.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}

	pcb, err := webGerber.NewPCBRender(file, &v1.RenderOptions{
		Zoom:       0,
		DarkColor:  color.Black,
		ClearColor: color.White,
	})
	if err != nil {
		panic(err)
	}
	img := pcb.Render()
	err = img.SavePNG("test.png")
	if err != nil {
		panic(err)
	}
}
