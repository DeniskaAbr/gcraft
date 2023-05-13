package gc_gui

import "gcraft/gc_common/gc_interface"

type MouseHandler func(gc_interface.MouseEvent)

type MouseMoveHandler func(gc_interface.MouseMoveEvent)

type widget interface {
	render(target gc_interface.Surface)
	advance(elapsed float64) error

	onMouseMove(event gc_interface.MouseMoveEvent) bool
	onMouseEnter(event gc_interface.MouseMoveEvent) bool
	onMouseLeave(event gc_interface.MouseMoveEvent) bool
	onMouseOver(event gc_interface.MouseMoveEvent) bool
	onMouseButtonDown(event gc_interface.MouseEvent) bool
	onMouseButtonUp(event gc_interface.MouseEvent) bool
	onMouseButtonClick(event gc_interface.MouseEvent) bool

	getPosition() (int, int)
	setOffset(x, y int)
	SetScreenPos(x, y int)
	ScreenPos() (x, y int)
	getSize() (int, int)
	getLayer() int
	isVisible() bool
	isExpanding() bool
}

type widgetBase struct {
	x         int
	y         int
	Sx        int
	Sy        int
	layer     int
	visible   bool
	expanding bool

	offsetX int
	offsetY int

	mouseEnterHandler MouseMoveHandler
	mouseLeaveHandler MouseMoveHandler
	mouseClickHandler MouseHandler
}

func (w *widgetBase) SetPosition(x, y int) {
	w.x = x
	w.y = y
}

func (w *widgetBase) GetPosition() (x, y int) {
	return w.x, w.y
}

func (w *widgetBase) GetOffset() (x, y int) {
	return w.offsetX, w.offsetY
}

func (w *widgetBase) SetScreenPos(x, y int) {
	w.Sx, w.Sy = x, y
}

func (w *widgetBase) ScreenPos() (x, y int) {
	return w.Sx, w.Sy
}

func (w *widgetBase) setOffset(x, y int) {
	w.offsetX = x
	w.offsetY = y
}

func (w *widgetBase) SetLayer(layer int) {
	w.layer = layer
}

func (w *widgetBase) SetVisible(visible bool) {
	w.visible = visible
}

func (w *widgetBase) SetExpanding(expanding bool) {
	w.expanding = expanding
}

func (w *widgetBase) SetMouseEnterHandleer(handler MouseMoveHandler) {
	w.mouseEnterHandler = handler
}

func (w *widgetBase) SetMouseLeaveHandler(handler MouseMoveHandler) {
	w.mouseLeaveHandler = handler
}

func (w *widgetBase) SetMouseClickHandler(handler MouseHandler) {
	w.mouseClickHandler = handler
}

func (w *widgetBase) getPosition() (x, y int) {
	return w.x, w.y
}

func (w *widgetBase) getSize() (width, height int) {
	return 0, 0
}

func (w *widgetBase) getLayer() int {
	return w.layer
}

func (w *widgetBase) isVisible() bool {
	return w.visible
}

func (w *widgetBase) isExpanding() bool {
	return w.expanding
}

func (w *widgetBase) render(_ gc_interface.Surface) { /* NOOP */ }

func (w *widgetBase) advance(_ float64) error {
	return nil
}

func (w *widgetBase) onMouseEnter(event gc_interface.MouseMoveEvent) bool {
	if w.mouseEnterHandler != nil {
		w.mouseEnterHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseLeave(event gc_interface.MouseMoveEvent) bool {
	if w.mouseLeaveHandler != nil {
		w.mouseLeaveHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseButtonClick(event gc_interface.MouseEvent) bool {
	if w.mouseClickHandler != nil {
		w.mouseClickHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseMove(_ gc_interface.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseOver(_ gc_interface.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonDown(_ gc_interface.MouseEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonUp(_ gc_interface.MouseEvent) bool {
	return false
}
