package gc_interface

type SoundEffect interface {
	Play()
	Stop()
	SetPan(pan float64)
	IsPlaying() bool
	SetVolume(volume float64)
}
