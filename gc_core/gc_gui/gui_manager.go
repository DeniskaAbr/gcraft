package gc_gui

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_util"
)

const logPrefix = "GUI Manager"

type GuiManager struct {
	layout *Layout

	cursorX       int
	cursorY       int
	cursorVisible bool

	loading bool

	*gc_util.Logger
}

func CreateGuiManager(l gc_util.LogLevel, inputManager gc_interface.InputManager) (*GuiManager, error) {
	manager := &GuiManager{
		cursorVisible: true,
	}

	manager.Logger = gc_util.NewLogger()
	manager.Logger.SetPrefix(logPrefix)
	manager.Logger.SetLevel(l)

	manager.clear()

	if err := inputManager.BindHandler(manager); err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *GuiManager) SetLayout(layout *Layout) {
	m.layout = layout
	if m.layout != nil {
		m.layout.AdjustEntryPlacement()
	}
}

func (m *GuiManager) Advance(elapsed float64) error {
	if !m.loading && m.layout != nil {
		if err := m.layout.advance(elapsed); err != nil {
			return err
		}
	}
	return nil
}

func (m *GuiManager) HideLoadScreen() {
	m.loading = false
}

func (m *GuiManager) clear() {
	m.SetLayout(nil)
	m.HideLoadScreen()
}
