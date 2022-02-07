package entities

type user struct {
	Name  string // exported
	Email string // exported
}

type Admin struct {
	user   // embedded and unexported
	Rights int
}
