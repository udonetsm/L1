package main

import "log"

type Building interface {
	SetRoundWindow()
}

type HobbitHome struct {
	Radius float64
}
type HumanHome struct {
	Square float64
}
type Builder struct{}
type Adapter struct {
	home *HumanHome
}
type Test struct {
}

func (h *HobbitHome) SetRoundWindow() {
	log.Println("Round window has been set")
}

func (hh *HumanHome) SetClassicWindow() {
	log.Println("Classic window has been set")
}

func (a *Adapter) SetRoundWindow() {
	a.home.SetClassicWindow()
}

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
