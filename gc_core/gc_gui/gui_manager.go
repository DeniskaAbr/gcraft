package gc_gui

const logPrefix = "GUI Manager"

type GuiManager struct{}

func CreateGuiManager() (*GuiManager, error) {
	manager := &GuiManager{}

	return manager, nil
}

func (m *GuiManager) Advance(elapsed float64) error {
	return nil
}
