package gc_ui

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_util"
)

type UIManager struct {
	renderer gc_interface.Renderer
	widgets  []Widget

	*gc_util.Logger
}

func (ui *UIManager) Initialize() {
}

func (ui *UIManager) Reset() {
}

func (ui *UIManager) addWidget(widget Widget) {
	ui.widgets = append(ui.widgets, widget)

	widget.bindManager(ui)
}

func (ui *UIManager) Renderer() gc_interface.Renderer {
	return ui.renderer
}

func (ui *UIManager) Advance(elapsed float64) {
}

func (ui *UIManager) Render(target gc_interface.Surface) {
}
