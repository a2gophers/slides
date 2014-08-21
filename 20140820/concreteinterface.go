// +build OMIT
package main

// START OMIT
type Concrete interface {
	Name() string
	Size() int64

	canPeopleImplementMe() bool
}

// STOP OMIT
