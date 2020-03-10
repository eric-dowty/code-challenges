package main

import (
	"fmt"
)

type Location struct {
	Elevation	int
	MaxLeft	 	int
	MaxRight 	int
}

func (l *Location)MinWall() int {
	if l.MaxLeft < l.MaxRight {
		return l.MaxLeft
	} 
	
	return l.MaxRight
}

func (l *Location)TrappedWater() int {	
	if l.MinWall() > l.Elevation {
		return l.MinWall() - l.Elevation
	}
	
	return 0
}

func CalcTrappedWater(elevations []int) (trappedWater int) {
	locations := make([]Location, len(elevations))
	maxLeft, maxRight := 0, 0

	for i, elevation := range elevations {
		locations[i].Elevation = elevation

		if elevations[i] > maxLeft {
			maxLeft = elevations[i]
		}
		locations[i].MaxLeft = maxLeft

		revIndex := len(elevations) - 1 - i
		if elevations[revIndex] > maxRight {
			maxRight = elevations[revIndex]
		}
		locations[revIndex].MaxRight = maxRight
	}

	for _, location := range locations {
		trappedWater += location.TrappedWater()
	}

	return trappedWater
}

func main() {
	elevations := []int{1, 0, 2, 2, 4, 0, 1, 5, 2, 1, 6, 1}
	trappedWater := CalcTrappedWater(elevations)

	fmt.Println(trappedWater)
}
