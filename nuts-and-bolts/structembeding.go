package nutsandbolts

import "time"

type base struct {
    yearOfBirth int
}

func (b base) age() int {
    year := time.Now().Year()
	return year - b.yearOfBirth
}

type occupation struct {
	name string
	yearOfFirstJob int
}

func (o occupation) worksForYears() int {
    year := time.Now().Year()
	return year - o.yearOfFirstJob
}

// A person embeds a base and occupation. An embedding looks like a field without a name.
type person struct {
    name string
	base
    occupation
	familyMembers []person
}

func (p person) noOfFamilyMembers() int {
	return len(p.familyMembers)
}