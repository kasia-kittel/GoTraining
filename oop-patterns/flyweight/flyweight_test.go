package patterns

import (
	"runtime"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestHeavyInit(t *testing.T) {

	screen := InitHeavy()
	screenLight := InitLight()
	var m runtime.MemStats
    runtime.ReadMemStats(&m)

	assert.NotNil(t, screen)
	assert.NotNil(t, screenLight)

	m1 := unsafe.Sizeof(screen)
	m2 := unsafe.Sizeof(screenLight)

	assert.True(t, m1 == m2, "The structure size should be the same")
}

func BenchmarkInitVeryHeavy(b *testing.B) {
    for i := 0; i < b.N; i++ {
        InitVeryHeavy()
    }
	// BenchmarkInitVeryHeavy-8   	       5	 204385153 ns/op	1813779424 B/op	    1824 allocs/op
}

func BenchmarkInitHeavy(b *testing.B) {
    for i := 0; i < b.N; i++ {
        InitHeavy()
    }
	// BenchmarkInitHeavy-8   	       8	 137476527 ns/op	909947626 B/op	     919 allocs/op
}

func BenchmarkInitLight(b *testing.B) {
    for i := 0; i < b.N; i++ {
		InitLight()
    }
	// BenchmarkInitLight-8   	    2127	    537162 ns/op	 3093153 B/op	      16 allocs/op
}

func BenchmarkInitLighter(b *testing.B) {
    for i := 0; i < b.N; i++ {
		InitEvenLighter()
    }
	// BenchmarkInitLigher-8   	   30294	     39248 ns/op	   70241 B/op	      11 allocs/op
}
