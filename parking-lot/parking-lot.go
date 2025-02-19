package main

import "fmt"

type ParkingLot struct {
	Levels []*Level
}

func NewParkingLot(levels int) *ParkingLot {
	lot := &ParkingLot{}
	for i := 1; i <= levels; i++ {
		lot.Levels = append(lot.Levels, NewLevel(i, 100))
	}

	return lot
}

func (lot *ParkingLot) ParkVehicle(v Vehicle) bool {
	for _, level := range lot.Levels {
		if level.ParkVehicle(v) {
			return true
		}
	}

	return false
}

func (lot *ParkingLot) UnParkVehicle(v Vehicle) {
	for _, level := range lot.Levels {
		level.UnParkVehicle(v)
	}
}

func (lot *ParkingLot) GetInfo() {
	for _, level := range lot.Levels {
		occ, unocc := level.GetInfo()
		fmt.Printf("Level %d has %d occupied slots and %d unoccupied spots\n\n", level.GetFloor(), occ, unocc)
	}
}

func main() {
	ParkingLot := NewParkingLot(5)

	car := NewCar("ABC1234")
	bike := NewBike("XY00PS")
	truck := NewTruck("FU632")

	ParkingLot.ParkVehicle(car)
	ParkingLot.ParkVehicle(bike)
	ParkingLot.ParkVehicle(truck)

	ParkingLot.GetInfo()

	ParkingLot.UnParkVehicle(bike)

	ParkingLot.GetInfo()
}
