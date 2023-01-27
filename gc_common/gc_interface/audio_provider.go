package gc_interface

type AudioProvider interface {
	PlayBGM(song string)
	LoadSound(sfx string, loop bool, bgm bool) (SoundEffect, error)
	SetVolumes()
}
