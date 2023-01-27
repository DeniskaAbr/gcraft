package gc_ui

import "gcraft/gc_common/gc_interface"

type UIManager struct {
	renderer gc_interface.Renderer
}

func (ui *UIManager) Initialize() {
}

func (ui *UIManager) Reset() {
}

func (ui *UIManager) Renderer() gc_interface.Renderer {
	return ui.renderer
}
