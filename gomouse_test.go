package gomouse

import (
	"log"
	"math"
	"testing"
)

func TestGeneratePoints(t *testing.T) {
	settings := MouseSettings{
		StartX:     math.Ceil(RandomNumberFloat() * 1920),
		StartY:     math.Ceil(RandomNumberFloat() * 1080),
		EndX:       math.Ceil(RandomNumberFloat() * 1920),
		EndY:       math.Ceil(RandomNumberFloat() * 1080),
		Gravity:    math.Ceil(RandomNumberFloat() * 10),
		Wind:       math.Ceil(RandomNumberFloat() * 10),
		MinWait:    2.0,
		MaxWait:    math.Ceil(RandomNumberFloat() * 5),
		MaxStep:    math.Ceil(RandomNumberFloat() * 3),
		TargetArea: math.Ceil(RandomNumberFloat() * 10),
	}

	points := GeneratePoints(settings)

	log.Print(points)
}
