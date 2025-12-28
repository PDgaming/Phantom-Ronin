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
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel, UsePlatformModel: g.state.UsePlatformModel, IsDebug: g.state.isDebug})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	// Add checkbox for UseJiraiyaModel
	initialUseJiraiyaModel := g.state.UseJiraiyaModel
	g.state.UseJiraiyaModel = gui.CheckBox(rl.NewRectangle(float32(screenWidth)/2-50, 280, 20, 20), "Use Jiraiya Model", g.state.UseJiraiyaModel)

	if g.state.UseJiraiyaModel != initialUseJiraiyaModel {
		g.gameObjects.player.UseModel = g.state.UseJiraiyaModel
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel, UsePlatformModel: g.state.UsePlatformModel, IsDebug: g.state.isDebug})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	// Add checkbox for UsePlatformModel
	initialUsePlatformModel := g.state.UsePlatformModel
	g.state.UsePlatformModel = gui.CheckBox(rl.NewRectangle(float32(screenWidth)/2-50, 310, 20, 20), "Use Platform Model", g.state.UsePlatformModel)

	if g.state.UsePlatformModel != initialUsePlatformModel {
		g.currentLevel.resetLevel()
		g.currentLevel.loadLevel(fmt.Sprintf("./level-maps/level%d.csv", g.state.Level), g.gameObjects.platformModel, g.gameObjects.winPlatformModel, g.state.UsePlatformModel)
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel, UsePlatformModel: g.state.UsePlatformModel, IsDebug: g.state.isDebug})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	// Add checkbox for IsDebug
	initialIsDebug := g.state.isDebug
	g.state.isDebug = gui.CheckBox(rl.NewRectangle(float32(screenWidth)/2-50, 340, 20, 20), "Debug Mode", g.state.isDebug)
	if g.state.isDebug != initialIsDebug {
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel, UsePlatformModel: g.state.UsePlatformModel, IsDebug: g.state.isDebug})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
	}

	backButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 370, 100, 40), "Back") // Adjusted Y position
	if backButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		// Ensure settings are saved when leaving the settings menu
		err := SaveSettings(SettingsData{ShowIntro: g.state.ShowIntro, UseJiraiyaModel: g.state.UseJiraiyaModel, UsePlatformModel: g.state.UsePlatformModel, IsDebug: g.state.isDebug})
		if err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
		}
		g.state.menuState = "startMenu"
	}
}
