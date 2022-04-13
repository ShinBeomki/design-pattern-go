package AdapterPattern

import "fmt"

type MediaPlayer interface {
	Play(audioType, fileName string)
}

type AdvanceMediaPlayer interface {
	PlayRmvb(fileName string)
	PlayMP4(fileName string)
}

type RmvbPlayer struct{}

type MP4Player struct{}

func NewRmvbPlayer() *RmvbPlayer {
	return &RmvbPlayer{}
}

func (rmvb *RmvbPlayer) PlayRmvb(fileName string) {
	fmt.Println("Playing rmvb file. Name: ", fileName)
}

func (rmvb *RmvbPlayer) PlayMP4(fileName string) {
	//do nothing
}

func NewMp4Player() *MP4Player {
	return &MP4Player{}
}

func (mp4 *MP4Player) PlayRmvb(fileName string) {
	//do nothing
}

func (mp4 *MP4Player) PlayMP4(fileName string) {
	fmt.Println("Playing mp4 file. Name: ", fileName)
}

type MediaAdapter struct {
	MusicPlayer AdvanceMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "rmvb":
		return &MediaAdapter{MusicPlayer: NewRmvbPlayer()}
	case "mp4":
		return &MediaAdapter{MusicPlayer: NewMp4Player()}
	}
	return nil
}

func (mAdapter *MediaAdapter) Play(audioType, fileName string) {
	switch audioType {
	case "rmvb":
		mAdapter.MusicPlayer.PlayRmvb(fileName)
		break
	case "mp4":
		mAdapter.MusicPlayer.PlayMP4(fileName)
		break
	}
}

type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

func (auPlayer *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: ", fileName)
	} else if audioType == "rmvb" || audioType == "mp4" {
		auPlayer.mediaAdapter = NewMediaAdapter(audioType)
		if auPlayer.mediaAdapter == nil {
			fmt.Println("mediaAdapter create fail.")
		}
		auPlayer.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media type. ", audioType, "format not supported")
	}
}
