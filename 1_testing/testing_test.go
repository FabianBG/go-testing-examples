package main

import (
	"testing"
)

// Testing functions always has prefix Test
func TestMain(t *testing.T) {

	// Init vars
	const (
		normalNumberA  = 4
		normlaNumberB  = 2
		divisionResult = 2
		divisionZero   = 0
		name           = "TEST"
	)

	t.Run("Division", func(t *testing.T) {

		t.Run("WHEN all params are ok SHOULD return an integer", func(t *testing.T) {
			res := division(normalNumberA, normlaNumberB)
			if res != divisionResult {
				t.Errorf(
					"division (%d, %d) = %d; want %d",
					normalNumberA,
					normlaNumberB,
					res,
					divisionResult,
				)
			}

		})

		t.Run("WHEN b is 0 SHOULD panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Te division should panic")
				}
			}()
			division(normalNumberA, divisionZero)
		})

	})
}
