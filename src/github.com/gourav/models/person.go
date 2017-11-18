package models

import (
	"fmt"
	"strings"

	"github.com/gourav/util"
)

type Person struct {
	First_name, Last_name util.Name
}

func (p Person) Full_name() string {
	name := fmt.Sprintf("%s %s", p.First_name.ToCapitalize(), p.Last_name.ToCapitalize())
	return name
}

func (p Person) IsValid() bool {
	if p == (Person{}) {
		fmt.Println("Person struct not present")
		return false
	} else if len(strings.Trim(string(p.First_name), " ")) <= 0 ||
		len(strings.Trim(string(p.Last_name), " ")) <= 0 {
		return false
	}
	return true
}
