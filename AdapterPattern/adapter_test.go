package AdapterPattern

import "testing"

func Test(t *testing.T) {
	t.Run("audio Test:", AudioPlayerTest)
}

func AudioPlayerTest(t *testing.T) {
	audio := NewAudioPlayer()
	audio.Play("mp4", "IU.mp4")
	audio.Play("rmvb", "mamamu.rmvb")
	audio.Play("mp3", "CNBLUE.mp3")
	audio.Play("avi", "viva la Vida.avi")
}
