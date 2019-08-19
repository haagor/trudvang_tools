package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	da := resolv_degat(8)
	fmt.Println(da)
}

func resolv_degat(open_roll int) int {
	damage_amount := rand.Intn(10) + 1
	roll := damage_amount
	for roll >= open_roll {
		roll = rand.Intn(10) + 1
		damage_amount += roll
	}
	return damage_amount
}
