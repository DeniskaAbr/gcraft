package ebiten

import "gcraft/gc_common/gc_interface"

const sampleRate = 44100

const logPrefix = "Ebien Audio Provider"

var _ gc_interface.AudioProvider = &AudioProvider{}

type AudioProvider struct{}

func CreateAudio() *AudioProvider {
	result := &AudioProvider{}

	return result
}

func (a *AudioProvider) PlayBGM(song string) {
	// TODO implement me
	panic("implement me")
}

func (a *AudioProvider) LoadSound(sfx string, loop bool, bgm bool) (gc_interface.SoundEffect, error) {
	// TODO implement me
	panic("implement me")
}

func (a *AudioProvider) SetVolumes() {
	// TODO implement me
	panic("implement me")
}
