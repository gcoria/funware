package main

import (
	"fmt"

	"math/rand"
)

type FulboPlayer struct {
	name    string
	age     int
	stamina int
	power   int
}

type ElDiego struct {
	name             string
	age              int
	stamina          int
	power            int
	barrileteCosmico int
}

type Player interface {
	KickBall()
	Name() string
}

func (d *ElDiego) KickBall() {
	shoot := (d.stamina + d.power) * d.barrileteCosmico
	fmt.Println("Ta ta ta: ", shoot)
}

func (d *ElDiego) Name() string {
	return d.name
}

func (f *FulboPlayer) Name() string {
	return f.name
}

func (f *FulboPlayer) KickBall() {
	shoot := f.stamina + f.power
	fmt.Println("Kick: ", shoot)
}

func main() {
	galactic_team := make([]Player, 11)
	for i := 0; i < len(galactic_team)-1; i++ {
		galactic_team[i] = &FulboPlayer{
			name:    fmt.Sprintf("galactic-%v", i),
			age:     rand.Intn(10),
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	galactic_team[10] = &ElDiego{
		name:             "El Diego",
		age:              30,
		stamina:          20,
		power:            20,
		barrileteCosmico: 100,
	}

	for i := 0; i < len(galactic_team); i++ {
		fmt.Printf("Player: %s --- ", galactic_team[i].Name())
		galactic_team[i].KickBall()
	}
}
