# GoMouse

Library to generate human-like mouse mouvement.

# Usage

```go
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
```

# Credits

Thanks to [@BenLand100](https://github.com/BenLand100) for the [original WindMouse library in Java](https://github.com/BenLand100/SMART/blob/157e50691b4b63a0950fac06deccac26aae31f88/src/EventNazi.java#L201). All I did is porting it to Go.

# Visualizer

You can use [Mouse Data Visualizer](https://github.com/arevi/mouse-data-visualizer) made by [@arevi](https://github.com/arevi) to tune your mouse settings.