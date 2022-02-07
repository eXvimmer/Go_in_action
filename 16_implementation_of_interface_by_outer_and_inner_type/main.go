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

func (a *admin) notify() {
	fmt.Printf("sending an admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "Mustafa Hayati",
			email: "mustafa@email.com",
		},
		level: "super",
	}

	ad.user.notify() // access to the inner type's method directly
	ad.notify()      // the inner type's method is not promoted
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
