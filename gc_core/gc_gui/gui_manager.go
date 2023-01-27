package gc_gui

import "gcraft/gc_common/gc_util"

const logPrefix = "GUI Manager"

type GuiManager struct {
	layout  *Layout
	loading bool

	*gc_util.Logger
}

func CreateGuiManager() (*GuiManager, error) {
	manager := &GuiManager{}

	return manager, nil
}

func (m *GuiManager) Advance(elapsed float64) error {
	if !m.loading && m.layout != nil {
		if err := m.layout.advance(elapsed); err != nil {
			return err
		}
	}
	return nil
}
