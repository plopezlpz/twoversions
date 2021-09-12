package hello

import (
	"fmt"

	"github.com/plopezlpz/twoversions/v2/dep"
)

func SayHi() string {
	p := dep.Person{
		Name:    "Jonny",
		Surname: "Doe",
	}
	return fmt.Sprintf("I'm %v from v2", p.Name)
}
