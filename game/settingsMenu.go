package game

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawSettingsMenu(g *Game) {
	rl.DrawText("Settings", 80, 150, 80, rl.Red)
	// Store the initial state to check for changes
	initialShowIntro := g.state.ShowIntro
	g.state.ShowIntro = gui.CheckBox(rl.NewRectangle(float32(screenWidth)/2-50, 250, 20, 20), "Show Intro", g.state.ShowIntro)

	if g.state.ShowIntro != initialShowIntro {
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	// Add checkbox for UseJiraiyaModel
	initialUseJiraiyaModel := g.state.UseJiraiyaModel
	g.state.UseJiraiyaModel = gui.CheckBox(rl.NewRectangle(float32(screenWidth)/2-50, 280, 20, 20), "Use Jiraiya Model", g.state.UseJiraiyaModel)

	if g.state.UseJiraiyaModel != initialUseJiraiyaModel {
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	backButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 320, 100, 40), "Back") // Adjusted Y position
	if backButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		// Ensure settings are saved when leaving the settings menu
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
		g.state.menuState = "startMenu"
	}
}
