package gc_enum

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonMin = MouseButtonLeft
	MouseButtonMax = MouseButtonRight
)

type MouseButtonMod int

const (
	MouseButtonModLeft MouseButtonMod = 1 << iota
	MouseButtonModMiddle
	MouseButtonModRight
)
