// Package entities contains support for types of
// people in the system.
package entities

// user defines a user in the program.
type user struct {
	Name  string
	Email string
}

// Admin defines an admin in the program.
type Admin struct {
	user   // The embedded type is unexported.
	Rights int
}
