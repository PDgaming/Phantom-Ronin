package main

import (
	"encoding/csv"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Platforms []Platform
}

func (l *Level) loadLevel(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to open level file.")
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		rl.TraceLog(rl.LogError, "Failed to read CSV records.")
		panic(err)
	}

	for i, row := range records {
		if i == 0 && row[0] == "posX" {
			continue
		}

		posX, err := strconv.ParseFloat(row[0], 32)
		if err != nil {
			continue
		}
		posY, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			continue
		}
		posZ, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			continue
		}
		width, err := strconv.ParseFloat(row[3], 32)
		if err != nil {
			continue
		}
		height, err := strconv.ParseFloat(row[4], 32)
		if err != nil {
			continue
		}
		length, err := strconv.ParseFloat(row[5], 32)
		if err != nil {
			continue
		}
		final, err := strconv.ParseBool(row[6])
		if err != nil {
			continue
		}

		platformTopTexture := rl.LoadTexture("./assets/grass.jpg")
		platformSideTexture := rl.LoadTexture("./assets/dirt.png")

		newPlatform := Platform{
			Position: rl.NewVector3(float32(posX), float32(posY), float32(posZ)),
			Width:    float32(width),
			Height:   float32(height),
			Length:   float32(length),
			Color:    rl.Brown,

			TextureProvided: true,
			TopTexture:      platformTopTexture,
			SideTexture:     platformSideTexture,

			final: final,
		}

		l.Platforms = append(l.Platforms, newPlatform)
	}
}

func (l *Level) resetLevel() {
	l.Platforms = []Platform{}
}
