package v1

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/valyala/fastjson"
)

type PCBRender struct {
	Size     PCBSize         `json:"size"`
	Position PCBPositionAxis `json:"position"`
	Offset   PositionAxis    `json:"offset"`
	gg       *gg.Context
	pcb      *fastjson.Value
	opt      *RenderOptions
}
type RenderOptions struct {
	Zoom       float64     `json:"zoom"`       // the default is 100(1mm=100px)
	DarkColor  color.Color `json:"darkColor"`  // Object color, default is black
	ClearColor color.Color `json:"clearColor"` // Empty color, default is full transparent
}
type PCBPosition struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
type PCBPositionAxis struct {
	Start PositionAxis `json:"start"`
	End   PositionAxis `json:"end"`
}
type PCBSize struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Render if the zoom is 0, the default is 100(1mm=100px)
func (r *PCBRender) Render() *gg.Context {

	r.gg = gg.NewContext(int(r.Size.Width*r.opt.Zoom), int(r.Size.Height*r.opt.Zoom))

	// r.gg.InvertY()
	r.gg.SetColor(r.opt.ClearColor)
	r.gg.Clear()
	children := r.pcb.GetArray("children")

	for _, data := range children {
		polarity := Polarity(data.GetStringBytes("polarity"))
		if polarity == PolarityDark {
			r.gg.SetColor(r.opt.DarkColor)
		} else {
			r.gg.SetColor(r.opt.ClearColor)
		}
		childType := ChildModeType(data.GetStringBytes("type"))

		switch childType {
		case ChildModeTypeRegion:
			r.RenderRegion(data)
		case ChildModeTypeShape:
			r.RenderShape(data)
		case ChildModeTypePath:
			r.RenderPath(data)
		default:
			panic("unknown child type: " + childType)
		}

	}

	return r.gg
}

func (r *PCBRender) toImagePosX(pos *fastjson.Value) float64 {
	return pos.GetFloat64()*r.opt.Zoom + r.Offset.X*r.opt.Zoom
}
func (r *PCBRender) toImagePosY(pos *fastjson.Value) float64 {
	return pos.GetFloat64()*r.opt.Zoom + r.Offset.Y*r.opt.Zoom
}
func (r *PCBRender) toImagePosOtherValue(pos *fastjson.Value) float64 {
	return pos.GetFloat64() * r.opt.Zoom * r.opt.Zoom
}
func (r *PCBRender) toImagePosOther(pos float64) float64 {
	return pos * r.opt.Zoom
}
func (r *PCBRender) drawSegmentPath(data *fastjson.Value) {
	segments := data.GetArray("segments")
	r.gg.ClearPath()
	startFrom := segments[0].GetArray("start")
	r.gg.MoveTo(r.toImagePosX(startFrom[0]), r.toImagePosY(startFrom[1]))

	for _, segment := range segments {
		start := segment.GetArray("start")
		end := segment.GetArray("end")
		center := segment.GetArray("center")
		segmentType := PathSegmentType(segment.GetStringBytes("type"))
		if segmentType == PathSegmentTypeLine {
			r.gg.LineTo(r.toImagePosX(end[0]), r.toImagePosY(end[1]))
		} else if segmentType == PathSegmentTypeArc {
			r.gg.DrawArc(r.toImagePosX(center[0]), r.toImagePosY(center[1]), r.toImagePosOther(segment.GetFloat64("radius")), start[2].GetFloat64(), end[2].GetFloat64())
		} else {
			panic("unknown segments type: " + segmentType)
		}
	}
}
func (r *PCBRender) RenderPath(data *fastjson.Value) {
	r.drawSegmentPath(data)
	r.gg.Stroke()
}
func (r *PCBRender) RenderRegion(data *fastjson.Value) {
	r.drawSegmentPath(data)
	r.gg.Fill()
}

func (r *PCBRender) RenderShape(data *fastjson.Value) {
	shape := data.Get("shape")
	shapeType := ShapeType(shape.GetStringBytes("type"))

	if shapeType == ShapeTypeRect {
		r.RenderShapeRect(shape)
		// r.gg.DrawRectangle(r.toImagePosX(shape.Get("x")), r.toImagePosY(shape.Get("y")), r.toImagePosOther(shape.GetFloat64("xSize")), r.toImagePosOther(shape.GetFloat64("ySize")))
	} else if shapeType == ShapeTypeCircle {
		r.RenderShapeCircle(shape)
		// r.gg.DrawCircle(r.toImagePosX(shape.Get("cx")), r.toImagePosY(shape.Get("cy")), r.toImagePosOther(shape.GetFloat64("r")))
	} else if shapeType == ShapeTypeLayered {
		r.RenderShapeLayered(shape)
	} else {

		fmt.Println(string(data.MarshalTo(nil)))
		panic("unknown shape type: " + shapeType)
	}

}

func (r *PCBRender) RenderShapeRect(shape *fastjson.Value) {
	radius := shape.GetFloat64("r")
	if radius != 0 {
		r.gg.DrawRoundedRectangle(r.toImagePosX(shape.Get("x")), r.toImagePosY(shape.Get("y")), r.toImagePosOther(shape.GetFloat64("xSize")), r.toImagePosOther(shape.GetFloat64("ySize")), r.toImagePosOther(radius))

	} else {
		r.gg.DrawRectangle(r.toImagePosX(shape.Get("x")), r.toImagePosY(shape.Get("y")), r.toImagePosOther(shape.GetFloat64("xSize")), r.toImagePosOther(shape.GetFloat64("ySize")))

	}
	r.gg.Fill()
}

func (r *PCBRender) RenderShapeCircle(shape *fastjson.Value) {
	fmt.Println("shape circle", shape.Get("cx"), shape.Get("cy"), shape.GetFloat64("r"))
	r.gg.DrawCircle(r.toImagePosX(shape.Get("cx")), r.toImagePosY(shape.Get("cy")), r.toImagePosOther(shape.GetFloat64("r")))
	r.gg.Fill()

}

func (r *PCBRender) RenderShapeLayered(data *fastjson.Value) {
	shapes := data.GetArray("shapes")
	for _, shape := range shapes {
		shapeType := ShapeType(shape.GetStringBytes("type"))
		// r.gg.ClearPath()

		if shape.GetBool("erase") {
			r.gg.SetColor(r.opt.ClearColor)
		} else {
			r.gg.SetColor(r.opt.DarkColor)
		}

		if shapeType == ShapeTypePolygon {
			points := shape.GetArray("points")
			start := points[0].GetArray()
			r.gg.MoveTo(r.toImagePosX(start[0]), r.toImagePosY(start[1]))
			for _, point := range points {
				p := point.GetArray()
				r.gg.LineTo(r.toImagePosX(p[0]), r.toImagePosY(p[1]))
			}
			r.gg.Fill()
		} else if shapeType == ShapeTypeCircle {
			r.RenderShapeCircle(shape)
		} else if shapeType == ShapeTypeRect {
			r.RenderShapeRect(shape)
		} else {
			panic("unknown shape type: " + shapeType)
		}
	}

}
