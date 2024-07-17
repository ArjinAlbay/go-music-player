package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupUI(player *MusicPlayer) {
    a := app.New()
    w := a.NewWindow("Arjinamp")

    playButton := widget.NewButton("Play", func() {
        player.Play()
    })
    pauseButton := widget.NewButton("Pause", func() {
        player.Pause()
    })

    w.SetContent(container.NewVBox(
        widget.NewLabel("Music Player"),
        playButton,
        pauseButton,
    ))

    w.ShowAndRun()
}