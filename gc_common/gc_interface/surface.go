package gc_interface

import "image/color"

type Surface interface {
	Clear(color color.Color)
	PushTranslation(x, y int)
	DrawTextf(format string, params ...interface{})
}
