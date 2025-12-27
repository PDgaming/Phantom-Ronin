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

	header := []string{"posX", "posY", "posZ", "width", "height", "length", "final"}
	writer.Write(header)

	posX := 3.0
	initialPosY := -2.0
	maxPosY := 10.0

	xIncrement := (28.0 - posX) / float64(numPlatforms)

	previousPosZ := 0.0

	for i := 0; i < numPlatforms; i++ {
		isFinal := (i == numPlatforms-1)

		width := 0.8 + rand.Float64()*(1.0-0.8)
		length := 0.4 + rand.Float64()*(0.8-0.4)
		height := 0.3

		minZ := -0.9 + length/2
		maxZ := 1.1 - length/2

		var posZ float64
		if i == 0 {
			posZ = -0.3 + rand.Float64()*(-0.3 - -0.3)
		} else {
			var deltaZ float64
			if rand.Float64() < 0.2 { // 20% chance of a trick
				sign := float64(rand.Intn(2)*2 - 1) // -1 or 1
				deltaZ = sign * (0.4 + rand.Float64()*(0.6-0.4))
			} else {
				deltaZ = -0.1 + rand.Float64()*(0.1 - -0.1)
			}
			posZ = previousPosZ + deltaZ

			if posZ < minZ {
				posZ = minZ
			} else if posZ > maxZ {
				posZ = maxZ
			}
		}
		previousPosZ = posZ

		var posY float64
		if i == 0 { // First platform
			posY = -3.0 + rand.Float64()*(-2.0 - -3.0) // Ensure the first platform is low
		} else {
			var basePosY float64
			if numPlatforms > 1 {
				basePosY = initialPosY + ((maxPosY-initialPosY)/float64(numPlatforms-1))*float64(i)
			} else {
				basePosY = initialPosY
			}
			posY = basePosY + (-1.0 + rand.Float64()*(2.0))

			if posY > 10 {
				posY = 10
			} else if posY < -2 {
				posY = -2
			}
		}

		record := []string{
			strconv.FormatFloat(posX, 'f', 6, 64),
			strconv.FormatFloat(posY, 'f', 6, 64),
			strconv.FormatFloat(posZ, 'f', 6, 64),
			strconv.FormatFloat(width, 'f', 6, 64),
			strconv.FormatFloat(height, 'f', -1, 64),
			strconv.FormatFloat(length, 'f', 6, 64),
			strconv.FormatBool(isFinal),
		}
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}

		posX += xIncrement + (-0.2 + rand.Float64()*(0.4))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if _, err := os.Stat("level-maps"); os.IsNotExist(err) {
		os.Mkdir("level-maps", os.ModePerm)
	}

	for i := 1; i <= 10; i++ {
		numPlatforms := rand.Intn(11) + 20 // 20 to 30
		generateLevel(i, numPlatforms)
		fmt.Printf("Generated level-maps/level%d.csv\n", i)
	}
}
