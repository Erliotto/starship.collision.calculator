package main

/*
 Task description: https://github.com/Erliotto/starship.collision.calculator/blob/master/StarShipCollision-v1.txt
*/

import (
	"fmt"
)

type shipType int

const (
	commander shipType = iota
	patrol
	transport
)

type weapon int

const (
	laser weapon = iota
	blaster
)

type ship struct {
	kind   shipType
	weapon weapon
}

type collistion int

func collision(kind shipType, weapon weapon) collistion {
	switch kind {
	case commander:
		switch weapon {
		case laser:
			return 5
		case blaster:
			return 10
		}

	case patrol:
		switch weapon {
		case laser:
			return 15
		case blaster:
			return 25
		}

	case transport:
		switch weapon {
		case laser:
			return 20
		case blaster:
			return 40
		}
	}
	return 0
}

func main() {
	fmt.Println("[starship.collision .v1]")
}
