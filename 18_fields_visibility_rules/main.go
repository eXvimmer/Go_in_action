package main

import (
	"fmt"
	"github.com/eXvimmer/fields_visibility_rules/entities"
)

func main() {
	// u := entities.User{
	//   Name:  "Mustafa",
	//   email: "mustafa@email.com", // error
	// }
	// error: cannot refer to unexported field 'email' in struct literal of type
	// entities.User

	u := entities.User{
		Name:  "Mustafa",
		Email: "mustafa@email.com",
	}

	fmt.Printf("User: %v\n", u)
}
