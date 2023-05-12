package gc_input

type KeyCharsEvent struct {
	HandlerEvent
	chars []rune
}

func (e *KeyCharsEvent) Chars() []rune {
	return e.chars
}
