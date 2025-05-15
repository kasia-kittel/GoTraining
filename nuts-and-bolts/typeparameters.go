package nutsandbolts


// very general type of some being
type Being struct {
	name string
}

// something that has ability to be alive in the current times
type living interface {
	alive() bool
}

// an animal is a living being
type Animal struct {
	Being
	// this is dangerous as there is no compile time checks if functions from the interface are implemented
	// and may end-up with "panic: runtime error: invalid memory address or nil pointer dereference"
	living 
}

// but a dinosaur is extinct
type Dinosaur Being


// we can implement the alive function
func (a *Animal) alive() bool {
	return true
}

// and create some more Animals types
type Bird Animal
type Dog Animal
type Fish Animal
type Whale Animal
type Porpoise Animal

// accept all types and just return the same return type as the parameter
func Identity[T any](v T) T {
    return v
}

// let's create a data structure that only accepts animals that gather in schools: 
// Fish, Whales, Porpoises, etc
// First the type constraint:
// "In Go, type constraints are described through a mixture of method and type requirements 
// which together define a type set: this is the set of all the types that satisfy all the 
// requirements. Go uses a generalized form of interfaces for this purpose. An interface 
// enumerates a set of methods and types, and the type set described by such an interface 
// consists of all the types that implement those methods and that are included in the 
// enumerated types.""
// (https://go.dev/blog/coretypes)
type Aquatic interface {
    Fish | Whale | Porpoise
}

type element[T any] struct {
    next *element[T]
    val  T
}

// linked list
type School[T Aquatic] struct {
    head, tail *element[T]
}

func (lst *School[T]) Add(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *School[T]) AllAquatics() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

// this function will accept array types and the 
// type E must be comparable, where comparable is a build in type
func FindIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

