package gc_ui

import "gcraft/gc_common/gc_interface"

const (
	logPrefix = "UI Manager"
)

func NewUIManager(
	renderer gc_interface.Renderer,
) *UIManager {
	ui := &UIManager{
		renderer: renderer,
	}

	return ui
}
