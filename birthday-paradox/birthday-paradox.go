package birthdayparadox


import (
	"math/rand/v2"
)

func success(getBDay func() int) bool{

	var results = make(map[int]int)

	for v := range 23 {
		bDay := getBDay()
		_, prs := results[bDay]

		if prs {
			return true
		}
		results[bDay] = v
	}
	return false
}

func Simulate() float32{
	count := 1000
	successCount := 0

	randomBDay := func() int { 
		return rand.IntN(355) 
	}

	for i := 0; i < count; i++ {
		res := success(randomBDay)
		if(res) {
			successCount ++
		}	
	}

	return float32(successCount) / float32(count)
}