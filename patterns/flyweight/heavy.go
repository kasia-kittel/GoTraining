package patterns

import (
	"math/rand/v2"
)

func (r RainDrop) Content() []byte {
	return r.graphics[:]
}

func (r SnowFlake) Content() []byte {
	return r.graphics[:]
}

func (r Cloud) Content() []byte {
	return r.graphics[:]
}

type Object struct {
	shape Shape
	x, y  int
}

type Screen struct {
	objects []Object
}

func InitVeryHeavy() Screen {

	screen := &Screen{}

	for range 300 {
		rd := RainDrop{}
		screen.objects = append(screen.objects, Object{
			shape: rd,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		sf := SnowFlake{}
		screen.objects = append(screen.objects, Object{
			shape: sf,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		c := Cloud{}
		screen.objects = append(screen.objects, Object{
			shape: c,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	// we have 1800 objects ~ 1,6GB
	return *screen
}


func InitHeavy() Screen {

	rainDrop := RainDrop{}
	snowFlake := SnowFlake{}
	cloud := Cloud{}

	screen := &Screen{}

	for range 300 {
		screen.objects = append(screen.objects, Object{
			shape: rainDrop,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, Object{
			shape: snowFlake,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, Object{
			shape: cloud,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	// we have 900 objects ~ 900GB
	return *screen
}
