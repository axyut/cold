package main

import (
	"github.com/axyut/cold/cmd"
)

func main() {
	cmd.Execute()
}

// todo:
// write tests

// known bugs:
// when installe dfrom `go install`, cold -v does not work, because ikd how it builds and how to intert ldflags there
