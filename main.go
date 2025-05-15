package main

import (
	"fmt"
	"os"

	"github.com/kasia-kittel/GoTraining/coinflipper"
	"github.com/kasia-kittel/GoTraining/diceroller"
	"github.com/kasia-kittel/GoTraining/url"
)

func main() {
	fmt.Println("Hello World!")

	r1 := coinflipper.Simulate()
	fmt.Printf("Probability: %.4f\n", r1)

	r2 := diceroller.Simulate()
	fmt.Printf("Probability: %.4f\n", r2)

	os.Exit(url.Run(os.Stdout))
}
