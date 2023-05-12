package gc_interface

import "gcraft/gc_common/gc_enum"

type InputService interface {
	CursorPosition() (x int, y int)
	InputChars() []rune
	IsKeyPressed(key gc_enum.Key) bool
	IsKeyJustPressed(key gc_enum.Key) bool
	IsKeyJustReleased(key gc_enum.Key) bool
	IsMouseButtonPressed(button gc_enum.MouseButton) bool
	IsMouseButtonJustPressed(button gc_enum.MouseButton) bool
	IsMouseButtonJustReleased(button gc_enum.MouseButton) bool
	KeyPressDuration(key gc_enum.Key) int
}
