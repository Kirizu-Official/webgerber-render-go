package v1

type PlotImageTree struct {
	Type     string         `json:"type"` // fixed: image
	Units    UnitType       `json:"units"`
	Size     []float64      `json:"size"` // [x1,y1,x2,y2]
	Children []ImageGraphic `json:"children"`
	// Tools    tools          `json:"tools"`
}

type ImageGraphic struct {
	Type     ChildModeType `json:"type"`
	Polarity Polarity      `json:"polarity"`
}

// =======================Shape======================

type ImageGraphicShape struct {
	Type  ShapeType   `json:"type"`
	Shape interface{} `json:"shape"`
	Erase bool        `json:"erase"` // for layered shape
}

type ImageGraphicShapeCircle struct {
	CX float64 `json:"cx"`
	CY float64 `json:"cy"`
	R  float64 `json:"r"`
}

type ImageGraphicShapeRect struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	XSize float64 `json:"xSize"`
	YSize float64 `json:"ySize"`
	R     float64 `json:"r"`
}
type ImageGraphicShapePolygon struct {
	Points []float64 `json:"points"` // [x,y]
}
type ImageGraphicShapeOutline struct {
	Segments []PathSegment `json:"segments"`
}

type ImageGraphicShapeLayered struct {
	Shapes []ImageGraphicShape `json:"shapes"`
}

type PathSegment struct {
	Type   PathSegmentType `json:"type"`
	Start  Position        `json:"start"`
	End    Position        `json:"end"`
	Center Position        `json:"center"` // for arc
	Radius float64         `json:"radius"` // for arc
}
type PathSegmentType string

const PathSegmentTypeLine PathSegmentType = "line"
const PathSegmentTypeArc PathSegmentType = "arc"

type Position []float64
type PositionAxis struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// =======================Path======================

type ImageGraphicPath struct {
	Type     ChildModeType `json:"type"` // fixed: imagePath
	Width    float64       `json:"width"`
	Segments []PathSegment `json:"segments"`
}

// =======================Region======================

type ImageGraphicRegion struct {
	Type     ChildModeType `json:"type"` // fixed: imageRegion
	Segments []PathSegment `json:"segments"`
}
