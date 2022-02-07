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
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("sending an admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	u := user{"Mustafa", "mustafa@email.com"}
	sendNotification(&u)
	a := admin{"admin", "admin@email.com"}
	sendNotification(&a)
}

func sendNotification(u notifier) {
	u.notify()
}
