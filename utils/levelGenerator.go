package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateLevel(levelNumber int, numPlatforms int) {
	filepath := fmt.Sprintf("level-maps/level%d.csv", levelNumber)
	file, _ := os.Create(filepath)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"posX", "posY", "posZ", "width", "height", "length", "final"})

	// --- CONFIG ---
	minX, maxX := 1.0, 27.0
	minZ, maxZ := 0.0, 2.0
	maxZStep := 0.7

	totalDist := maxX - minX
	pWidth := 1.0

	// --- Y-AXIS S-CURVE CONFIG ---
	// We'll use: Y = Sine(progress) * Amplitude + UpwardTrend
	amplitude := 2.5 + rand.Float64()*2.0     // How high the "S" peaks
	verticalTrend := 6.0 + rand.Float64()*4.0 // Total height gain

	currentZ := 1.5 + rand.Float64() // Start somewhere random in Z

	for i := 0; i < numPlatforms; i++ {
		isFinal := (i == numPlatforms-1)
		progress := float64(i) / float64(numPlatforms-1) // 0.0 to 1.0

		// 1. X-AXIS WITH JITTER
		// Base linear position
		baseX := minX + (progress * totalDist)
		// Add jitter so platforms aren't perfectly spaced (except first and last)
		jitterX := 0.0
		if i > 0 && !isFinal {
			jitterX = (rand.Float64() - 0.5) * 0.8
		}
		actualX := baseX + jitterX

		// 2. Y-AXIS "S" PATTERN
		// Sine wave + linear climb
		sCurve := math.Sin(progress*math.Pi*1.5) * amplitude
		climb := progress * verticalTrend
		actualY := -2.0 + sCurve + climb + (rand.Float64() * 0.5) // add bit of noise

		// 3. Z-AXIS WALK
		if i > 0 {
			zShift := (rand.Float64()*2 - 1) * maxZStep
			currentZ += zShift
			// Clamp with a small bounce
			if currentZ < minZ {
				currentZ = minZ + 0.1
			}
			if currentZ > maxZ {
				currentZ = maxZ - 0.1
			}
		}

		record := []string{
			strconv.FormatFloat(actualX, 'f', 2, 64),
			strconv.FormatFloat(actualY, 'f', 2, 64),
			strconv.FormatFloat(currentZ, 'f', 2, 64),
			strconv.FormatFloat(pWidth, 'f', 2, 64),
			strconv.FormatFloat(0.3, 'f', 2, 64),
			strconv.FormatFloat(0.6, 'f', 2, 64),
			strconv.FormatBool(isFinal),
		}
		writer.Write(record)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if _, err := os.Stat("level-maps"); os.IsNotExist(err) {
		os.Mkdir("level-maps", os.ModePerm)
	}

	for i := 1; i <= 10; i++ {
		generateLevel(i, rand.Intn(6)+16) // 16 to 22 platforms
		fmt.Printf("Generated level %d\n", i)
	}
}
