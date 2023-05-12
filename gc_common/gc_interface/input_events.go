package gc_interface

import "gcraft/gc_common/gc_enum"

type HandlerEvent interface {
	KeyMod() gc_enum.KeyMod
	ButtonMod() gc_enum.MouseButtonMod
	X() int
	Y() int
}

type KeyEvent interface {
	HandlerEvent
	Key() gc_enum.Key
	Duration() int
}

type KeyCharsEvent interface {
	HandlerEvent
	Chars() []rune
}

type MouseEvent interface {
	HandlerEvent
	Button() gc_enum.MouseButton
}

type MouseMoveEvent interface {
	HandlerEvent
}
