package gc_ui

type Widget interface {
	Drawable
	bindManager(ui *UIManager)
}

type BaseWidget struct {
	manager *UIManager

	visible bool
}

func NewBaseWidget(manager *UIManager) *BaseWidget {
	return &BaseWidget{
		manager: manager,
		visible: true,
	}
}

func (b *BaseWidget) bindManager(manager *UIManager) {
	b.manager = manager
}

func (b *BaseWidget) GetVisible() (visible bool) {
	return b.visible
}

func (b *BaseWidget) SetVisible(visible bool) {
	b.visible = visible
}
