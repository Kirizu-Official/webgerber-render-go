package v1

type UnitType string

const UnitTypeMM UnitType = "mm"
const UnitTypeInch UnitType = "inch"

type ChildModeType string

const ChildModeTypeShape ChildModeType = "imageShape"
const ChildModeTypePath ChildModeType = "imagePath"
const ChildModeTypeRegion ChildModeType = "imageRegion"

type Polarity string

const PolarityDark Polarity = "dark"
const PolarityClear Polarity = "clear"

type ShapeType string

const ShapeTypeCircle ShapeType = "circle"
const ShapeTypeRect ShapeType = "rectangle"
const ShapeTypePolygon = "polygon"
const ShapeTypeOutline = "outline"
const ShapeTypeLayered ShapeType = "layeredShape"
