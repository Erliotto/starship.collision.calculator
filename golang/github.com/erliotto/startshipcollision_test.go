package main

import "testing"

func Test_Collision(t *testing.T) {
	// arrange
	testData := []struct {
		kind     shipType
		weapon   weapon
		expected collistion
	}{
		{commander, laser, 5},
		{commander, blaster, 10},
		{patrol, laser, 15},
		{patrol, blaster, 25},
		{transport, laser, 20},
		{transport, blaster, 40},
	}

	// act
	for _, td := range testData {
		actual := collision(td.kind, td.weapon)

		// assert
		if actual != td.expected {
			t.Errorf("Collision on expected: %d, but actual: %d", td.expected, actual)
		}
	}
}
