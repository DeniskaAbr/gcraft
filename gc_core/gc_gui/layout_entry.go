package gc_gui

import (
	"gcraft/gc_common/gc_geom"
	"gcraft/gc_common/gc_interface"
)

type layoutEntry struct {
	widget widget

	x int
	y int

	width  int
	height int

	mouseOver bool
	mouseDown [3]bool
}

func (l *layoutEntry) IsIN(event gc_interface.HandlerEvent) bool {
	sx, sy := l.widget.ScreenPos()
	rect := gc_geom.Rectangle{Left: sx, Top: sy, Width: l.width, Height: l.height}

	return rect.IsInRect(event.X(), event.Y())
}
