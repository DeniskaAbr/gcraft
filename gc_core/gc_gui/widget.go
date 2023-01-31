package gc_gui

import "gcraft/gc_common/gc_interface"

type widget interface {
	render(target gc_interface.Surface)
	advance(elapsed float64) error
}

type widgetBase struct {
}
