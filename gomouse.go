package gomouse

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math"
	"math/rand"
)

// MouseSettings initiate the mouse settings
type MouseSettings struct {
	StartX     float64
	StartY     float64
	EndX       float64
	EndY       float64
	Gravity    float64
	Wind       float64
	MinWait    float64
	MaxWait    float64
	MaxStep    float64
	TargetArea float64
}

func RandomNumberFloat() float64 {
	// avoid pitfalls of clock based seed value
	var b [8]byte
	_, err := cryptorand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r.Float64()
}

func hypot(dx, dy float64) float64 {
	return math.Sqrt(dx*dx + dy*dy)
}

func GeneratePoints(settings MouseSettings) [][]float64 {
	if settings.Gravity < 1 {
		settings.Gravity = 1
	}

	if settings.MaxStep == 0 {
		settings.MaxStep = 0.01
	}

	windX := math.Floor(RandomNumberFloat() * 10)
	windY := math.Floor(RandomNumberFloat() * 10)

	var oldX float64
	var oldY float64
	newX := math.Floor(settings.StartX)
	newY := math.Floor(settings.StartY)

	waitDiff := settings.MaxWait - settings.MinWait

	// Hardcore instead of doing math.sqrt, maybe saving us some computiong time
	sqrt2 := 1.4142135623730951
	sqrt3 := 1.7320508075688772
	sqrt5 := 2.23606797749979

	var randomDist float64
	var velocityX float64 = 0
	var velocityY float64 = 0
	var dist float64
	var veloMag float64
	var step float64

	var points [][]float64
	var currentWait float64 = 0

	dist = hypot(settings.EndX-settings.StartX, settings.EndY-settings.StartY)

	for dist > 1.0 {
		settings.Wind = math.Min(settings.Wind, dist)

		if dist >= settings.TargetArea {
			w := math.Floor(RandomNumberFloat()*math.Round(settings.Wind)*2 + 1)

			windX = windX/sqrt3 + (w-settings.Wind)/sqrt5
			windY = windY/sqrt3 + (w-settings.Wind)/sqrt5
		} else {
			windX = windX / sqrt2
			windY = windY / sqrt2

			if settings.MaxStep < 3 {
				settings.MaxStep = math.Floor(RandomNumberFloat()*3) + 3.0
			} else {
				settings.MaxStep = settings.MaxStep / sqrt5
			}
		}

		velocityX += windX
		velocityY += windY
		velocityX = velocityX + (settings.Gravity*(settings.EndX-settings.StartX))/dist
		velocityY = velocityY + (settings.Gravity*(settings.EndY-settings.StartY))/dist

		if hypot(velocityX, velocityY) > settings.MaxStep {
			randomDist = settings.MaxStep/2.0 + math.Floor((RandomNumberFloat()*math.Round(settings.MaxStep))/2)
			veloMag = hypot(velocityX, velocityY)
			velocityX = (velocityX / veloMag) * randomDist
			velocityY = (velocityY / veloMag) * randomDist
		}

		oldX = math.Round(settings.StartX)
		oldY = math.Round(settings.StartY)

		settings.StartX += velocityX
		settings.StartY += velocityY

		dist = hypot(settings.EndX-settings.StartX, settings.EndY-settings.StartY)

		newX = math.Round(settings.StartX)
		newY = math.Round(settings.StartY)

		step = hypot(settings.StartX-oldX, settings.StartY-oldY)
		wait := math.Round(waitDiff*(step/settings.MaxStep) + settings.MinWait)

		currentWait += wait

		if oldX != newY || oldY != newY {
			points = append(points, []float64{
				newX,
				newY,
				currentWait,
			})
		}
	}

	return points
}
