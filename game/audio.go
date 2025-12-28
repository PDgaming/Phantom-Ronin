package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) loadMusicAndSounds() {
	g.audioStreams.introMusic = rl.LoadMusicStream("./assets/audio/that-game-arcade.mp3")
	g.audioStreams.gameMusic = rl.LoadMusicStream("./assets/audio/356-8-bit-chiptune-game-music.mp3")
	g.audioStreams.jumpSound = rl.LoadSound("./assets/audio/pixel-jump.mp3")
	g.audioStreams.buttonSound = rl.LoadSound("./assets/audio/button-press.mp3")
	g.audioStreams.winSound = rl.LoadSound("./assets/audio/piglevelwin2.mp3")
}

func (g *Game) manageMusic() {
	var activeMusic rl.Music
	var inactiveMusic rl.Music

	if g.state.menuState == "startMenu" || g.state.menuState == "intro" {
		activeMusic = g.audioStreams.introMusic
		inactiveMusic = g.audioStreams.gameMusic
	} else if g.state.menuState == "inGame" {
		activeMusic = g.audioStreams.gameMusic
		inactiveMusic = g.audioStreams.introMusic
	}

	if rl.IsMusicStreamPlaying(inactiveMusic) {
		rl.StopMusicStream(inactiveMusic)
	}

	if !rl.IsMusicStreamPlaying(activeMusic) {
		rl.PlayMusicStream(activeMusic)
	}

	rl.UpdateMusicStream(activeMusic)

	if rl.GetMusicTimePlayed(activeMusic) >= rl.GetMusicTimeLength(activeMusic) {
		rl.StopMusicStream(activeMusic)
		rl.PlayMusicStream(activeMusic)
	}
}

func unloadAudio(g *Game) {
	rl.UnloadMusicStream(g.audioStreams.introMusic)
	rl.UnloadMusicStream(g.audioStreams.gameMusic)
	rl.UnloadSound(g.audioStreams.jumpSound)
	rl.UnloadSound(g.audioStreams.buttonSound)
	rl.UnloadSound(g.audioStreams.winSound)
	rl.CloseAudioDevice()
}
