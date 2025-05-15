package nutsandbolts

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestPerson(t *testing.T) {


    John := person{
		name: "John",
        base: base{
            yearOfBirth: 1990,
        },
		occupation: occupation {
			name: "DevOps",
			yearOfFirstJob: 2015,
		},
    }

	JohnsAge := time.Now().Year()-1990
	JohnWorksForYears := time.Now().Year()-2015

    // here the interesting bit is how person embeds base and occupation, 
	// the methods of base also become methods of a person. 
	assert.Equal(t, JohnsAge, John.age())
	assert.Equal(t, JohnWorksForYears, John.worksForYears())

	// even more interesting how can we compose it with interfaces

	type employee interface {
        worksForYears() int
    }

	var employeeJohn employee  = John

	assert.Equal(t, JohnWorksForYears, employeeJohn.worksForYears())

	Ann := person {
		name: "Ann",
	}

	
	John.familyMembers = append(John.familyMembers, Ann)

	assert.Equal(t, 1, John.noOfFamilyMembers())

	// we still can call the age function as defaults were initiated
	// this is quite error prone
	assert.Equal(t, time.Now().Year(),  Ann.age())
}
