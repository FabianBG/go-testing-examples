package main

import (
	"fmt"
	"os"
	"testing"

	// https://pkg.go.dev/gotest.tools?tab=overview
	"gotest.tools/assert"
)

func TestMain(t *testing.T) {

	// Init vars
	const (
		name    = "TEST"
		message = "Hola, TEST"
	)

	t.Run("printName", func(t *testing.T) {

		t.Run("SHOULD return custom message", func(t *testing.T) {
			os.Setenv("NAME", name)

			res := printName()
			assert.Assert(t, message == res, fmt.Sprintf("%s not equal %s", message, res))
		})

	})
}
