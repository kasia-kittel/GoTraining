package patterns

// to support a huge number of objects that my not fit into available RAM

// intrinsic state - the constant data of an object that can be shared across many instances
// extrinsic state - the data of an object that changes

// a flyweight should initialize its state just once
//  more convenient access to various flyweights, you can create a
//  factory method that manages a pool of existing flyweight objects

// initialize an imaginary screen of 900 graphical objects of 3 types

const screenResolution = 3000

type Shape interface {
	Content() []byte
}

type RainDrop struct {
	graphics [1000000]byte //1MB
}

type SnowFlake struct {
	graphics [1000000]byte //1MB
}

type Cloud struct {
	graphics [1000000]byte //1MB
}
