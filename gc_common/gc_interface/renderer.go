package gc_interface

type renderCallback = func(Surface) error

type updateCallback = func() error

type Renderer interface {
	ShowPanicScreen(message string)
	Run(r renderCallback, u updateCallback, width, height int, title string) error
}
