package main

import (
	"fmt"
	"log"
)

//TowerOfHanoi struct for Tower of Hanoi
type TowerOfHanoi struct {
	//source is source rod
	source string
	//destination is destination rod
	destination string
	//extra is extra rod
	extra string
	//discs is number of discs
	discs int
}

func newTowerOfHanoi(source, destination, extra string, discs int) *TowerOfHanoi {
	return &TowerOfHanoi{
		source:      source,
		destination: destination,
		extra:       extra,
		discs:       discs,
	}
}

//SolveTowerOfHanoi solves tower of hanoi problem recursively
func (t *TowerOfHanoi) SolveTowerOfHanoi() {
	RecursiveTowerOfHanoi(t.source, t.destination, t.extra, t.discs)
}

//RecursiveTowerOfHanoi runs recursively
func RecursiveTowerOfHanoi(source, destination, extra string, discs int) {
	if discs <= 0 {
		return
	}
	RecursiveTowerOfHanoi(source, extra, destination, discs-1)
	fmt.Printf("Move Disk-%d from %s to %s\n", discs, source, destination)
	RecursiveTowerOfHanoi(extra, destination, source, discs-1)
}

func main() {
	log.Print("Implementation of Tower of Hanoi")
	TowerOfHanoiInstance := newTowerOfHanoi("Source", "Destination", "Extra", 3)
	TowerOfHanoiInstance.SolveTowerOfHanoi()
}
