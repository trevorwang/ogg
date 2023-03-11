package ogg

import "testing"

func TestOggToWavByPath(t *testing.T) {
	oggfile := "ogg.ogg"
	wavfile := "output.wav"

	OggToWavByPath(oggfile, wavfile)

}
