package v1

import (
	"fmt"
	"math"

	"github.com/valyala/fastjson"
)

func NewPCBRender(data []byte, opt *RenderOptions) (*PCBRender, error) {
	plotData, err := fastjson.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	if string(plotData.GetStringBytes("type")) != "image" {
		return nil, ErrInvalidPlotData
	}
	size := plotData.GetArray("size")
	if len(size) != 4 {
		return nil, ErrInvalidPCBSize
	}
	fmt.Println("offset", -size[0].GetFloat64(), -size[1].GetFloat64())
	render := &PCBRender{
		Size:     PCBSize{Width: math.Abs(size[2].GetFloat64() - size[0].GetFloat64()), Height: math.Abs(size[3].GetFloat64() - size[1].GetFloat64())},
		Position: PCBPositionAxis{Start: PositionAxis{X: size[0].GetFloat64(), Y: size[1].GetFloat64()}, End: PositionAxis{X: size[2].GetFloat64(), Y: size[3].GetFloat64()}},
		Offset:   PositionAxis{X: -size[0].GetFloat64(), Y: -size[1].GetFloat64()},
		pcb:      plotData,
		opt:      opt,
	}

	return render, err
}
