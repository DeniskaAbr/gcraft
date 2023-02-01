package gc_ui

import (
	"gcraft/gc_common/gc_interface"
	"gcraft/gc_common/gc_util"
)

var _ Widget = &Label{}

type Label struct {
	*BaseWidget

	*gc_util.Logger
}

func (ui *UIManager) NewLabel() *Label {

	base := NewBaseWidget(ui)

	result := &Label{
		BaseWidget: base,
		Logger:     ui.Logger,
	}

	result.bindManager(ui)

	result.SetVisible(false)

	ui.addWidget(result)

	return result
}

func (l *Label) Render(target gc_interface.Surface) {

}

func (l *Label) Advance(elapsed float64) error {
	return nil
}
