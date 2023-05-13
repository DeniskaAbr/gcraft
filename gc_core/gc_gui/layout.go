package gc_gui

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_math"
)

const layoutDebug = false

type PositionType int

const (
	PositionTypeAbsolute PositionType = iota
	PositionTypeVertical
	PositionTypeHorizontal
)

var _ widget = &Layout{}

type Layout struct {
	widgetBase

	renderer gc_interface.Renderer

	width  int
	height int

	positionType PositionType
	entries      []*layoutEntry
}

// render implements widget
func (l *Layout) render(target gc_interface.Surface) {
	l.AdjustEntryPlacement()

	for _, entry := range l.entries {
		if !entry.widget.isVisible() {
			continue
		}

		l.renderEntry(entry, target)

		if layoutDebug {
			l.renderEntryDebug(entry, target)
		}
	}
}

func (l *Layout) renderEntry(entry *layoutEntry, target gc_interface.Surface) {
	target.PushTranslation(entry.x, entry.y)
	defer target.Pop()

	entry.widget.render(target)
}

func (l *Layout) renderEntryDebug(entry *layoutEntry, target gc_interface.Surface) {
	target.PushTranslation(entry.x, entry.y)
	defer target.Pop()

}

func (l *Layout) advance(elapsed float64) error {
	for _, entry := range l.entries {
		if err := entry.widget.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

func (l *Layout) AdjustEntryPlacement() {

}

func (l *Layout) getContentSize() (width, height int) {
	for _, entry := range l.entries {
		x, y := entry.widget.getPosition()
		w, h := entry.widget.getSize()

		switch l.positionType {
		case PositionTypeVertical:
			width = gc_math.MaxInt(width, w)
			height += h
		case PositionTypeHorizontal:
			width += w
			height = gc_math.MaxInt(height, h)
		case PositionTypeAbsolute:
			width = gc_math.MaxInt(width, x+w)
			height = gc_math.MaxInt(height, y+h)
		}
	}

	return width, height
}

func (l *Layout) getSize() (width, height int) {
	width, height = l.getContentSize()
	return gc_math.MaxInt(width, l.width), gc_math.MaxInt(height, l.height)
}
