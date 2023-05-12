package ebiten

import (
	"gcraft/gc_common/gc_enum"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	//nolint:gochecknoglobals
	keyToEbiten = map[gc_enum.Key]ebiten.Key{
		gc_enum.Key0:            ebiten.Key0,
		gc_enum.Key1:            ebiten.Key1,
		gc_enum.Key2:            ebiten.Key2,
		gc_enum.Key3:            ebiten.Key3,
		gc_enum.Key4:            ebiten.Key4,
		gc_enum.Key5:            ebiten.Key5,
		gc_enum.Key6:            ebiten.Key6,
		gc_enum.Key7:            ebiten.Key7,
		gc_enum.Key8:            ebiten.Key8,
		gc_enum.Key9:            ebiten.Key9,
		gc_enum.KeyA:            ebiten.KeyA,
		gc_enum.KeyB:            ebiten.KeyB,
		gc_enum.KeyC:            ebiten.KeyC,
		gc_enum.KeyD:            ebiten.KeyD,
		gc_enum.KeyE:            ebiten.KeyE,
		gc_enum.KeyF:            ebiten.KeyF,
		gc_enum.KeyG:            ebiten.KeyG,
		gc_enum.KeyH:            ebiten.KeyH,
		gc_enum.KeyI:            ebiten.KeyI,
		gc_enum.KeyJ:            ebiten.KeyJ,
		gc_enum.KeyK:            ebiten.KeyK,
		gc_enum.KeyL:            ebiten.KeyL,
		gc_enum.KeyM:            ebiten.KeyM,
		gc_enum.KeyN:            ebiten.KeyN,
		gc_enum.KeyO:            ebiten.KeyO,
		gc_enum.KeyP:            ebiten.KeyP,
		gc_enum.KeyQ:            ebiten.KeyQ,
		gc_enum.KeyR:            ebiten.KeyR,
		gc_enum.KeyS:            ebiten.KeyS,
		gc_enum.KeyT:            ebiten.KeyT,
		gc_enum.KeyU:            ebiten.KeyU,
		gc_enum.KeyV:            ebiten.KeyV,
		gc_enum.KeyW:            ebiten.KeyW,
		gc_enum.KeyX:            ebiten.KeyX,
		gc_enum.KeyY:            ebiten.KeyY,
		gc_enum.KeyZ:            ebiten.KeyZ,
		gc_enum.KeyApostrophe:   ebiten.KeyApostrophe,
		gc_enum.KeyBackslash:    ebiten.KeyBackslash,
		gc_enum.KeyBackspace:    ebiten.KeyBackspace,
		gc_enum.KeyCapsLock:     ebiten.KeyCapsLock,
		gc_enum.KeyComma:        ebiten.KeyComma,
		gc_enum.KeyDelete:       ebiten.KeyDelete,
		gc_enum.KeyDown:         ebiten.KeyDown,
		gc_enum.KeyEnd:          ebiten.KeyEnd,
		gc_enum.KeyEnter:        ebiten.KeyEnter,
		gc_enum.KeyEqual:        ebiten.KeyEqual,
		gc_enum.KeyEscape:       ebiten.KeyEscape,
		gc_enum.KeyF1:           ebiten.KeyF1,
		gc_enum.KeyF2:           ebiten.KeyF2,
		gc_enum.KeyF3:           ebiten.KeyF3,
		gc_enum.KeyF4:           ebiten.KeyF4,
		gc_enum.KeyF5:           ebiten.KeyF5,
		gc_enum.KeyF6:           ebiten.KeyF6,
		gc_enum.KeyF7:           ebiten.KeyF7,
		gc_enum.KeyF8:           ebiten.KeyF8,
		gc_enum.KeyF9:           ebiten.KeyF9,
		gc_enum.KeyF10:          ebiten.KeyF10,
		gc_enum.KeyF11:          ebiten.KeyF11,
		gc_enum.KeyF12:          ebiten.KeyF12,
		gc_enum.KeyGraveAccent:  ebiten.KeyGraveAccent,
		gc_enum.KeyHome:         ebiten.KeyHome,
		gc_enum.KeyInsert:       ebiten.KeyInsert,
		gc_enum.KeyKP0:          ebiten.KeyKP0,
		gc_enum.KeyKP1:          ebiten.KeyKP1,
		gc_enum.KeyKP2:          ebiten.KeyKP2,
		gc_enum.KeyKP3:          ebiten.KeyKP3,
		gc_enum.KeyKP4:          ebiten.KeyKP4,
		gc_enum.KeyKP5:          ebiten.KeyKP5,
		gc_enum.KeyKP6:          ebiten.KeyKP6,
		gc_enum.KeyKP7:          ebiten.KeyKP7,
		gc_enum.KeyKP8:          ebiten.KeyKP8,
		gc_enum.KeyKP9:          ebiten.KeyKP9,
		gc_enum.KeyKPAdd:        ebiten.KeyKPAdd,
		gc_enum.KeyKPDecimal:    ebiten.KeyKPDecimal,
		gc_enum.KeyKPDivide:     ebiten.KeyKPDivide,
		gc_enum.KeyKPEnter:      ebiten.KeyKPEnter,
		gc_enum.KeyKPEqual:      ebiten.KeyKPEqual,
		gc_enum.KeyKPMultiply:   ebiten.KeyKPMultiply,
		gc_enum.KeyKPSubtract:   ebiten.KeyKPSubtract,
		gc_enum.KeyLeft:         ebiten.KeyLeft,
		gc_enum.KeyLeftBracket:  ebiten.KeyLeftBracket,
		gc_enum.KeyMenu:         ebiten.KeyMenu,
		gc_enum.KeyMinus:        ebiten.KeyMinus,
		gc_enum.KeyNumLock:      ebiten.KeyNumLock,
		gc_enum.KeyPageDown:     ebiten.KeyPageDown,
		gc_enum.KeyPageUp:       ebiten.KeyPageUp,
		gc_enum.KeyPause:        ebiten.KeyPause,
		gc_enum.KeyPeriod:       ebiten.KeyPeriod,
		gc_enum.KeyPrintScreen:  ebiten.KeyPrintScreen,
		gc_enum.KeyRight:        ebiten.KeyRight,
		gc_enum.KeyRightBracket: ebiten.KeyRightBracket,
		gc_enum.KeyScrollLock:   ebiten.KeyScrollLock,
		gc_enum.KeySemicolon:    ebiten.KeySemicolon,
		gc_enum.KeySlash:        ebiten.KeySlash,
		gc_enum.KeySpace:        ebiten.KeySpace,
		gc_enum.KeyTab:          ebiten.KeyTab,
		gc_enum.KeyUp:           ebiten.KeyUp,
		gc_enum.KeyAlt:          ebiten.KeyAlt,
		gc_enum.KeyControl:      ebiten.KeyControl,
		gc_enum.KeyShift:        ebiten.KeyShift,
	}
	//nolint:gochecknoglobals
	mouseButtonToEbiten = map[gc_enum.MouseButton]ebiten.MouseButton{
		gc_enum.MouseButtonLeft:   ebiten.MouseButtonLeft,
		gc_enum.MouseButtonMiddle: ebiten.MouseButtonMiddle,
		gc_enum.MouseButtonRight:  ebiten.MouseButtonRight,
	}
)

// InputService ...
type InputService struct{}

// CursorPosition ...
func (is InputService) CursorPosition() (x, y int) {
	return ebiten.CursorPosition()
}

// InputChars ...
func (is InputService) InputChars() []rune {
	return ebiten.InputChars()
}

// IsKeyPressed ...
func (is InputService) IsKeyPressed(key gc_enum.Key) bool {
	return ebiten.IsKeyPressed(keyToEbiten[key])
}

// IsKeyJustPressed ...
func (is InputService) IsKeyJustPressed(key gc_enum.Key) bool {
	return inpututil.IsKeyJustPressed(keyToEbiten[key])
}

// IsKeyJustReleased ...
func (is InputService) IsKeyJustReleased(key gc_enum.Key) bool {
	return inpututil.IsKeyJustReleased(keyToEbiten[key])
}

// IsMouseButtonPressed ...
func (is InputService) IsMouseButtonPressed(button gc_enum.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustPressed ...
func (is InputService) IsMouseButtonJustPressed(button gc_enum.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustReleased ...
func (is InputService) IsMouseButtonJustReleased(button gc_enum.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(mouseButtonToEbiten[button])
}

// KeyPressDuration ...
func (is InputService) KeyPressDuration(key gc_enum.Key) int {
	return inpututil.KeyPressDuration(keyToEbiten[key])
}
