package gc_interface

type InputEventHandler interface{}

type KeyDownHandler interface {
	OnKeyDown(event KeyEvent) (preventPropagation bool)
}

type KeyRepeatHandler interface {
	OnKeyRepeat(event KeyEvent) (preventPropagation bool)
}

type KeyUpHandler interface {
	OnKeyUp(event KeyEvent) (preventPropagation bool)
}
type KeyCharsHandler interface {
	OnKeyChars(event KeyCharsEvent) (preventPropagation bool)
}

type MouseButtonDownHandler interface {
	OnMouseButtonDown(event MouseEvent) (preventPropagation bool)
}

type MouseButtonRepeatHandler interface {
	OnMouseButtonRepeat(event MouseEvent) (preventPropagation bool)
}

type MouseButtonUpHandler interface {
	OnMouseButtonUp(event MouseEvent) (preventPropagation bool)
}

type MouseMoveHandler interface {
	OnMouseMove(event MouseMoveEvent) (preventPropagation bool)
}
