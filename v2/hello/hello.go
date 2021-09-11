package hello

import (
	"fmt"

	"github.com/plopezlpz/twoversions/v2/dep"
)

func SayHi() string {
	p := dep.Person{
		Name:    "John",
		Surname: "Doe",
	}
	return fmt.Sprintf("I'm %v from V0", p.Name)
}
