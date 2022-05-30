package main

import "log"

// interface
type Building interface {
	SetRoundWindow()
}

// one object
type HobbitHome struct {
	Radius float64
}

// second object
type HumanHome struct {
	Square float64
}

// third object
type Builder struct{}

// object adapter
type Adapter struct {
	home *HumanHome
}

// method available only for first object
func (h *HobbitHome) SetRoundWindow() {
	log.Println("Round window has been set")
}

// method available only for second object
func (hh *HumanHome) SetClassicWindow() {
	log.Println("Classic window has been set")
}

// adapt method available for first object for second object, using adapter struct
func (a *Adapter) SetRoundWindow() {
	a.home.SetClassicWindow()
}

// use adapted method
func (b *Builder) SetRoundWindowInBothBuildings(building Building) {
	building.SetRoundWindow()
}

func setRoundWinsowInHumanBuilding() {
	builder := &Builder{}
	humanBuilding := &HumanHome{}
	hobbitBuilding := &HobbitHome{}
	hobbitBuilding.SetRoundWindow()
	humanAdapter := &Adapter{home: humanBuilding}
	builder.SetRoundWindowInBothBuildings(humanAdapter)
}
