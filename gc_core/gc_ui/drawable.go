package gc_ui

import "gcraft/gc_common/gc_interface"

type Drawable interface {
	Render(target gc_interface.Surface)
	Advance(elapsed float64) error
	GetVisible() bool
	SetVisible(visible bool)
}
