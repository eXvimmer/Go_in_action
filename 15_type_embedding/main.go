package main

import (
	"fmt"
)

type notifier interface {
  notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending an user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // embedded type
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "Mustafa Hayati",
			email: "mustafa@email.com",
		},
		level: "super",
	}

	ad.user.notify() // direct access to the inner types
	ad.notify()      // The inner type's method is promoted

  sendNotification(&ad)
  // NOTE: thanks to inner type promotion, the implementation of the interface
  // by the inner type has been promoted up to the outer type. That means the
  // outer type now implements the interface, thanks to the inner type's
  // implementation.
}

func sendNotification(n notifier) {
  n.notify()
}
