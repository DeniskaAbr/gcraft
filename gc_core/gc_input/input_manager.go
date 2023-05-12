package gc_input

import (
	"gcraft/gc_common/gc_enum"
	"gcraft/gc_common/gc_interface"
	ebiten_input "gcraft/gc_core/gc_input/ebiten"
	"sort"
)

type inputManager struct {
	inputService gc_interface.InputService
	cursorX      int
	cursorY      int

	buttonMod gc_enum.MouseButtonMod
	keyMod    gc_enum.KeyMod

	entries handlerEntryList
}

func NewInputManager() gc_interface.InputManager {
	return &inputManager{
		inputService: ebiten_input.InputService{},
	}
}

func (im *inputManager) Advance(_, _ float64) error {
	im.updateKeyMod()
	im.updateButtonMod()

	cursorX, cursorY := im.inputService.CursorPosition()
	eventBase := HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	for key := gc_enum.KeyMin; key <= gc_enum.KeyMax; key++ {
		im.updateJustPressedKey(key, eventBase)
		im.updateJustReleasedKey(key, eventBase)
		im.updatePressedKey(key, eventBase)
	}

	im.updateInputChars(eventBase)

	for button := gc_enum.MouseButtonMin; button <= gc_enum.MouseButtonMax; button++ {
		im.updateJustPressedButton(button, eventBase)
		im.updateJustReleasedButton(button, eventBase)
		im.updatePressedButton(button, eventBase)
	}

	im.updateCursor(cursorX, cursorY, eventBase)

	return nil
}

func (im *inputManager) updateKeyMod() {
	im.keyMod = 0
	if im.inputService.IsKeyPressed(gc_enum.KeyAlt) {
		im.keyMod |= gc_enum.KeyModAlt
	}

	if im.inputService.IsKeyPressed(gc_enum.KeyControl) {
		im.keyMod |= gc_enum.KeyModControl
	}

	if im.inputService.IsKeyPressed(gc_enum.KeyShift) {
		im.keyMod |= gc_enum.KeyModShift
	}
}

func (im *inputManager) updateButtonMod() {
	im.buttonMod = 0
	if im.inputService.IsMouseButtonPressed(gc_enum.MouseButtonLeft) {
		im.buttonMod |= gc_enum.MouseButtonModLeft
	}

	if im.inputService.IsMouseButtonPressed(gc_enum.MouseButtonMiddle) {
		im.buttonMod |= gc_enum.MouseButtonModMiddle
	}

	if im.inputService.IsMouseButtonPressed(gc_enum.MouseButtonRight) {
		im.buttonMod |= gc_enum.MouseButtonModRight
	}
}

func (im *inputManager) updateJustPressedKey(k gc_enum.Key, e HandlerEvent) {
	if im.inputService.IsKeyJustPressed(k) {
		event := KeyEvent{HandlerEvent: e, key: k}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.KeyDownHandler); ok {
				return l.OnKeyDown(&event)
			}

			return false
		}

		im.propagate(fn)
	}
}

func (im *inputManager) updateJustReleasedKey(k gc_enum.Key, e HandlerEvent) {
	if im.inputService.IsKeyJustReleased(k) {
		event := KeyEvent{HandlerEvent: e, key: k}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.KeyUpHandler); ok {
				return l.OnKeyUp(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updatePressedKey(k gc_enum.Key, e HandlerEvent) {
	if im.inputService.IsKeyPressed(k) {
		event := KeyEvent{
			HandlerEvent: e,
			key:          k,
			duration:     im.inputService.KeyPressDuration(k),
		}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.KeyRepeatHandler); ok {
				return l.OnKeyRepeat(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateInputChars(eventBase HandlerEvent) {
	if chars := im.inputService.InputChars(); len(chars) > 0 {
		event := KeyCharsEvent{eventBase, chars}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.KeyCharsHandler); ok {
				l.OnKeyChars(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateJustPressedButton(b gc_enum.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonJustPressed(b) {
		event := MouseEvent{e, b}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.MouseButtonDownHandler); ok {
				return l.OnMouseButtonDown(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateJustReleasedButton(b gc_enum.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonJustReleased(b) {
		event := MouseEvent{e, b}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.MouseButtonUpHandler); ok {
				return l.OnMouseButtonUp(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updatePressedButton(b gc_enum.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonPressed(b) {
		event := MouseEvent{e, b}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.MouseButtonRepeatHandler); ok {
				return l.OnMouseButtonRepeat(&event)
			}

			return false
		}

		im.propagate(fn)
	}
}

func (im *inputManager) updateCursor(cursorX, cursorY int, e HandlerEvent) {
	if im.cursorX != cursorX || im.cursorY != cursorY {
		event := MouseMoveEvent{e}

		fn := func(handler gc_interface.InputEventHandler) bool {
			if l, ok := handler.(gc_interface.MouseMoveHandler); ok {
				return l.OnMouseMove(&event)
			}

			return false
		}
		im.propagate(fn)

		im.cursorX, im.cursorY = cursorX, cursorY
	}
}

func (im *inputManager) BindHandlerWithPriority(
	h gc_interface.InputEventHandler,
	p gc_enum.Priority) error {
	return im.bindHandler(h, p)
}

func (im *inputManager) BindHandler(h gc_interface.InputEventHandler) error {
	return im.bindHandler(h, gc_enum.PriorityDefault)
}

func (im *inputManager) bindHandler(h gc_interface.InputEventHandler, p gc_enum.Priority) error {
	for _, entry := range im.entries {
		if entry.handler == h {
			return ErrHasReg
		}
	}

	entry := handlerEntry{h, p}
	im.entries = append(im.entries, entry)
	sort.Sort(im.entries)

	return nil
}

func (im *inputManager) UnbindHandler(handler gc_interface.InputEventHandler) error {
	for i, entry := range im.entries {
		if entry.handler == handler {
			copy(im.entries[i:], im.entries[i+1:])
			im.entries = im.entries[:len(im.entries)-1]

			return nil
		}
	}

	return ErrNotReg
}

func (im *inputManager) propagate(callback func(gc_interface.InputEventHandler) bool) {
	var priority gc_enum.Priority

	var handler bool

	for _, entry := range im.entries {
		if priority > entry.priority && handler {
			break
		}

		if callback(entry.handler) {
			handler = true
		}

		priority = entry.priority
	}
}

type handlerEntry struct {
	handler  gc_interface.InputEventHandler
	priority gc_enum.Priority
}

type handlerEntryList []handlerEntry

func (lel handlerEntryList) Len() int {
	return len(lel)
}

func (lel handlerEntryList) Swap(i, j int) {
	lel[i], lel[j] = lel[j], lel[i]
}

func (lel handlerEntryList) Less(i, j int) bool {
	return lel[i].priority > lel[j].priority
}
