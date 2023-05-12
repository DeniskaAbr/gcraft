package gc_input

import (
	"errors"
	"gcraft/gc_common/gc_interface"
)

var (
	// ErrHasReg shows the input system already has a registered handler
	ErrHasReg = errors.New("input system already has provided handler")
	// ErrNotReg shows the input system has no registered handler
	ErrNotReg = errors.New("input system does not have provided handler")
)

// Static checks to confirm struct conforms to interface
var _ gc_interface.InputEventHandler = &HandlerEvent{}
var _ gc_interface.KeyEvent = &KeyEvent{}
var _ gc_interface.KeyCharsEvent = &KeyCharsEvent{}
var _ gc_interface.MouseEvent = &MouseEvent{}
var _ gc_interface.MouseMoveEvent = &MouseMoveEvent{}
