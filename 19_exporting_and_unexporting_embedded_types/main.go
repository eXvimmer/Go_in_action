package main

import (
	"fmt"

	"github.com/eXvimmer/exporting_and_unexporting_embedded_types/entities"
)

func main() {
	a := entities.Admin{
		Rights: 13,
	}

  // NOTE: the fields declared within the inner type (user) are exported.
  // Since the identifiers from the inner type are promoted to the outer type,
  // those exported fields are known through a value of the outer type. There's
  // no access to the inner type directly, since the user type is unexported. 
	a.Name = "Mustafa"
	a.Email = "mustafa@email.com"

	fmt.Printf("User: %v\n", a)
}
