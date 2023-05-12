package gc_input

import "gcraft/gc_common/gc_enum"

type MouseMoveEvent struct {
	HandlerEvent
}

// KeyMod represents the key mod
func (e *MouseMoveEvent) KeyMod() gc_enum.KeyMod {
	return e.HandlerEvent.keyMod
}

// ButtonMod represents the button mod
func (e *MouseMoveEvent) ButtonMod() gc_enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

// X represents the X position
func (e *MouseMoveEvent) X() int {
	return e.HandlerEvent.x
}

// Y represents the Y position
func (e *MouseMoveEvent) Y() int {
	return e.HandlerEvent.y
}
