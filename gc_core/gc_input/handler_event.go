package gc_input

import (
	"gcraft/gc_common/gc_enum"
)

type HandlerEvent struct {
	keyMod    gc_enum.KeyMod
	buttonMod gc_enum.MouseButtonMod
	x         int
	y         int
}

func (e *HandlerEvent) KeyMod() gc_enum.KeyMod {
	return e.keyMod
}

func (e *HandlerEvent) ButtonMod() gc_enum.MouseButtonMod {
	return e.buttonMod
}

func (e *HandlerEvent) X() int {
	return e.x
}

func (e *HandlerEvent) Y() int {
	return e.y
}
