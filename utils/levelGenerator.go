package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateLevel(levelNumber int, numPlatforms int) {
	filepath := fmt.Sprintf("level-maps/level%d.csv", levelNumber)
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"posX", "posY", "posZ", "width", "height", "length", "final"})

	// --- WORLD BOUNDARIES ---
	minX := 1.0  // Start slightly in from 0
	maxX := 27.0 // End slightly before 28 to ensure the goal is visible

	minZBound := 1.0
	maxZBound := 3.0

	// --- PLAYER LIMITS ---
	maxVerticalJump := 1.0
	minVerticalStep := 0.3
	maxZStep := 0.7 // Safe Z-distance for 0.8 player reach

	// --- PLATFORM DIMENSIONS ---
	pWidth := 1.0
	pLength := 0.6
	pHeight := 0.3

	// --- MATH FOR X-SPACING ---
	// Total space taken by platforms = numPlatforms * pWidth
	// Remaining space for gaps = (maxX - minX) - (Total Platform Width)
	totalAvailableSpace := (maxX - minX) - (float64(numPlatforms) * pWidth)

	// If numPlatforms is 20, there are 19 gaps between them.
	gapSize := totalAvailableSpace / float64(numPlatforms-1)

	// --- INITIAL STATE ---
	currentX := minX
	posY := -2.0
	currentZ := 2.0 // Start center

	for i := 0; i < numPlatforms; i++ {
		isFinal := (i == numPlatforms-1)

		// 1. Z-Axis Walk (Step-constrained)
		if i > 0 {
			zShift := (rand.Float64()*2 - 1) * maxZStep
			currentZ += zShift

			// Clamp Z
			if currentZ < minZBound {
				currentZ = minZBound + 0.1
			} else if currentZ > maxZBound {
				currentZ = maxZBound - 0.1
			}
		}

		// 2. Y-Axis Walk (Step-constrained)
		if i > 0 {
			posY += minVerticalStep + rand.Float64()*(maxVerticalJump-minVerticalStep)
		}

		// 3. X-Axis Placement (calculated to fit the 1.0-27.0 range)
		// We use currentX for the platform center or edge depending on your engine.
		// Here, currentX represents the start of the platform.

		record := []string{
			strconv.FormatFloat(currentX, 'f', 2, 64),
			strconv.FormatFloat(posY, 'f', 2, 64),
			strconv.FormatFloat(currentZ, 'f', 2, 64),
			strconv.FormatFloat(pWidth, 'f', 2, 64),
			strconv.FormatFloat(pHeight, 'f', 2, 64),
			strconv.FormatFloat(pLength, 'f', 2, 64),
			strconv.FormatBool(isFinal),
		}
		writer.Write(record)

		// Advance X for the next platform
		currentX += pWidth + gapSize
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if _, err := os.Stat("level-maps"); os.IsNotExist(err) {
		os.Mkdir("level-maps", os.ModePerm)
	}

	for i := 1; i <= 10; i++ {
		// Choosing 15-22 platforms keeps the gaps reasonable for a 28-unit world
		generateLevel(i, rand.Intn(8)+15)
		fmt.Printf("Generated level %d\n", i)
	}
}
