package gc_input

import "gcraft/gc_common/gc_enum"

type MouseEvent struct {
	HandlerEvent
	mouseButton gc_enum.MouseButton
}

func (e *MouseEvent) KeyMod() gc_enum.KeyMod {
	return e.HandlerEvent.keyMod
}

func (e *MouseEvent) ButtonMod() gc_enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

func (e *MouseEvent) X() int {
	return e.HandlerEvent.x
}

func (e *MouseEvent) Y() int {
	return e.HandlerEvent.y
}

func (e *MouseEvent) Button() gc_enum.MouseButton {
	return e.mouseButton
}
