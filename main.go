package main

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
    player := NewMusicPlayer()
    SetupUI(player)
}

type MusicPlayer struct {
    streamer beep.StreamSeekCloser
    format   beep.Format
    ctrl     *beep.Ctrl
    paused   bool
}

func NewMusicPlayer() *MusicPlayer {
    f, err := os.Open("music.wav")
    if err != nil {
        panic(err)
    }

    streamer, format, err := wav.Decode(f)
    if err != nil {
        panic(err)
    }

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: true}

    speaker.Play(ctrl)

    return &MusicPlayer{
        streamer: streamer,
        format:   format,
        ctrl:     ctrl,
        paused:   true,
    }
}

func (p *MusicPlayer) Play() {
    if p.paused {
        p.ctrl.Paused = false
        p.paused = false
    }
}

func (p *MusicPlayer) Pause() {
    if !p.paused {
        p.ctrl.Paused = true
        p.paused = true
    }
}