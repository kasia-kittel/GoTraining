package patterns

import (
	"math/rand/v2"
)

type LightShape interface {
	LightContent() []byte
}

func (r *RainDrop) LightContent() []byte {
	return r.graphics[:]
}

func (r *SnowFlake) LightContent() []byte {
	return r.graphics[:]
}

func (r *Cloud) LightContent() []byte {
	return r.graphics[:]
}

// factory
// todo enum
type LightShapes int

const (
	RainDropShape LightShapes = iota
	SnowFlakeShape
	CloudShape
)

type LightShapeFactory struct {
	shapes map[LightShapes]LightShape
}

func (f *LightShapeFactory) getShapeByType(shapeType LightShapes) LightShape {

	if f.shapes[shapeType] != nil {
		return f.shapes[shapeType]
	}

	switch shapeType {
	case RainDropShape:
		f.shapes[RainDropShape] = &RainDrop{}
		return f.shapes[RainDropShape]
	case SnowFlakeShape:
		f.shapes[SnowFlakeShape] = &RainDrop{}
		return f.shapes[SnowFlakeShape]
	case CloudShape:
		f.shapes[CloudShape] = &RainDrop{}
		return f.shapes[CloudShape]
	}

	return nil

}

var shapeFactory = &LightShapeFactory{
	shapes: make(map[LightShapes]LightShape),
}

func getShapeFactory() *LightShapeFactory {
	return shapeFactory
}

type LightObject struct {
	shape LightShape
	x, y  int
}

type LightScreen struct {
	objects []LightObject
}

func InitLight() LightScreen {

	// we can turn LightShape creation to a factory patter so
	// we are sure in more complex cases only one instance of
	// each Shape is created
	rainDrop := &RainDrop{}
	snowFlake := &SnowFlake{}
	cloud := &Cloud{}

	screen := &LightScreen{}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: rainDrop,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: snowFlake,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: cloud,
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	// we have 900 objects ~ 900 x 32B
	return *screen
}

func InitEvenLighter() LightScreen {

	screen := &LightScreen{}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: getShapeFactory().getShapeByType(RainDropShape),
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: getShapeFactory().getShapeByType(SnowFlakeShape),
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	for range 300 {
		screen.objects = append(screen.objects, LightObject{
			shape: getShapeFactory().getShapeByType(CloudShape),
			x:     rand.IntN(screenResolution),
			y:     rand.IntN(screenResolution),
		})
	}

	// we have 900 objects ~ 900 x 32B
	return *screen
}
