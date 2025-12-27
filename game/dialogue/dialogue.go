package dialogue

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Manager struct {
	dialogues          []string
	currentDialogueIdx int
	boxX               int32
	boxY               int32
	nextBtnRect        rl.Rectangle
	IsActive           bool
}

func NewManager(dialogues []string, screenWidth, screenHeight int32) *Manager {
	boxX := (screenWidth / 2) - (600 / 2)
	boxY := screenHeight - 180
	return &Manager{
		dialogues:          dialogues,
		currentDialogueIdx: 0,
		boxX:               boxX,
		boxY:               boxY,
		nextBtnRect:        rl.NewRectangle(float32(boxX+480), float32(boxY+100), 100, 40),
		IsActive:           true,
	}
}

func (m *Manager) Update(buttonSound rl.Sound) (dialogueFinished bool) {
	if !m.IsActive {
		return false
	}

	buttonLabel := "Next"
	if m.currentDialogueIdx == len(m.dialogues)-1 {
		buttonLabel = "Begin"
	}

	if gui.Button(m.nextBtnRect, buttonLabel) || rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(buttonSound)
		if m.currentDialogueIdx < len(m.dialogues)-1 {
			m.currentDialogueIdx++
		} else {
			m.IsActive = false
			return true // Dialogue finished
		}
	}
	return false
}

func (m *Manager) Draw() {
	if !m.IsActive {
		return
	}
	// Draw dialogue box
	rl.DrawRectangle(m.boxX, m.boxY, 600, 150, rl.Black)
	rl.DrawRectangle(m.boxX+5, m.boxY+5, 590, 140, rl.Gray)

	// Draw dialogue text
	rl.DrawText(m.dialogues[m.currentDialogueIdx], m.boxX+10, m.boxY+20, 20, rl.Black)
}

func (m *Manager) Reset() {
	m.currentDialogueIdx = 0
	m.IsActive = true
}
