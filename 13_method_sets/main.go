// Sample program to show how to use interfaces in Go.
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
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"Mustafa", "mustafa@email.com"}

	// sendNotification(u) // compiler error

	// error: cannot use u (type user) as type notifier in argument to
	// sendNotification: user doesn't implement notifier (notify method has
	// pointer receiver)

	// NOTE: method sets define the set of methods that are associated with
	// values or pointers of a given type. The type of receiver used will
	// determine whether a method is associated with a value, pointer or both.

	// NOTE: method sets as described by the specification:
	//    values                                         method receivers
	// -------------------------------------------------------------------------
	//    T                                              (t T)
	//    *T                                             (t T) and (t *T)
	// -------------------------------------------------------------------------
	// A value of type T only has methods declared that have a value receiver, as
	// part of it's method set. But pointers of type T have methods declared with
	// both value and pointer receivers, as part of it's method set.

	// NOET: method sets from the perspective of the receiver type:
	//    method receivers                               values
	// -------------------------------------------------------------------------
	//    (t T)                                          T and *T
	//    (t *T)                                         *T
	// -------------------------------------------------------------------------
	// If you implement an interface using a pointer receiver, then only pointers
	// of that type implement the interface. If you implement an interfaces using
	// a vlaue receiver, then both values and pointers of that type implement the
	// interfaces.

	// We get an error, because we implemented the interface using a pointer
	// receiver and attempted to pass a value of type user to the
	// sendNotification function.

	sendNotification(&u) // Now, this works fine
}

func sendNotification(u notifier) {
	u.notify()
}
