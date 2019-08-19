package main

import (
	"math/rand"
	"time"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"strconv"
)

const (
	acurate = 10000
)
var (
	attack_sample = []int{6, 6, 6, 6}
	open_roll_weapon = 8
	strength = 4
	weapon = "Axe"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	degat_batch := [acurate]int{}
	degat_batch_plotter := make(plotter.Values, acurate)
	for i := 0; i < acurate; i++ {
		hits := resolve_attack_hit(attack_sample)
		degats := 0
		for j := 0; j < hits; j++ {
			degats += resolv_attack_degat(open_roll_weapon)
		}
		degat_batch_plotter[i] = float64(degats)
		degat_batch[i] = degats
	}

	avg := average(degat_batch)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = weapon + " sample 6, 6, 6, 6"
	p.Y.Label.Text = "Degats"

	q0, err := plotter.NewQuartPlot(0, degat_batch_plotter)
	if err != nil {
		panic(err)
	}
	p.Add(q0)

	p.NominalX(weapon + "\navg " + strconv.Itoa(avg))

	if err := p.Save(3*vg.Inch, 4*vg.Inch, weapon+"4.png"); err != nil {
		panic(err)
	}
}

func resolv_attack_degat(open_roll int) int {
	damage_amount := rand.Intn(10) + 1
	roll := damage_amount
	for roll >= open_roll {
		roll = rand.Intn(10) + 1
		damage_amount += roll
	}
	return damage_amount + strength
}

func resolve_attack_hit(sample_attack []int) int {
	hit_amount := 0
	for _, attack_SV := range sample_attack  {
		if rand.Intn(20) + 1 <= attack_SV {
			hit_amount += 1
		}
	}
	return hit_amount
}

func average(degats [acurate]int) int {
	sum_degats := 0
	for _, i := range degats {
		sum_degats += i
	}
	return sum_degats / len(degats)
}