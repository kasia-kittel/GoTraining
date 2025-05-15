package nutsandbolts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeParameters(t *testing.T) {

	b:= Being{}
	assert.Equal(t, Identity(b), b)

	duck := &Bird{}

	// this doesn't check if the function is really implemented
	assert.Implements(t, (*living)(nil), duck)

	// but this will not pass:
	// dinosaur := &Being{}
	// assert.Implements(t, (*living)(nil), dinosaur)
}

func TestSchool(t *testing.T) {

	fish1 := &Fish{ Being: Being{ name: "Fish1"} }
	fish2 := &Fish{}

	// because of embedding we can also do this
	fish2.name = "Fish2"

	fish3 := &Fish{}

	school := School[Fish]{}
    school.Add(*fish1)
	school.Add(*fish2)
	school.Add(*fish3)

	assert.Contains(t, school.AllAquatics(), *fish1)
	assert.Contains(t, school.AllAquatics(), *fish2)
	assert.Contains(t, school.AllAquatics(), *fish3)

	assert.Equal(t, FindIndex(school.AllAquatics(), *fish1), 0)
}