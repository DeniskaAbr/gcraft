package gc_input

import (
	"gcraft/gc_common/gc_enum"
)

type KeyEvent struct {
	HandlerEvent
	key gc_enum.Key

	duration int
}

func (e *KeyEvent) Key() gc_enum.Key {
	return e.key
}

func (e *KeyEvent) Duration() int {
	return e.duration
}
