package main

type Level struct {
	level int
	Spots []*ParkingSpot
}

func NewLevel(floor int, totalSpots int) *Level {
	carSpots := int(float64(totalSpots) * 0.4)
	bikeSpots := int(float64(totalSpots) * 0.4)

	l := &Level{level: floor}

	for i := 1; i <= carSpots; i++ {
		l.Spots = append(l.Spots, NewSpot(i, CAR))
	}
	for i := carSpots + 1; i <= carSpots+bikeSpots; i++ {
		l.Spots = append(l.Spots, NewSpot(i, MOTORCYCLE))
	}
	for i := carSpots + bikeSpots + 1; i <= totalSpots; i++ {
		l.Spots = append(l.Spots, NewSpot(i, TRUCK))
	}

	return l
}

func (l *Level) ParkVehicle(v Vehicle) bool {
	for _, spot := range l.Spots {
		if spot.IsAvailable() && spot.GetVehicleType() == v.getType() {
			spot.ParkVehicle(v)
			return true
		}
	}

	return false
}

func (l *Level) UnParkVehicle(v Vehicle) {
	for _, spot := range l.Spots {
		if !spot.IsAvailable() && spot.GetParkedVehicle().getNumberPlate() == v.getNumberPlate() {
			spot.UnParkVehicle(v)
		}
	}
}

func (l *Level) GetFloor() int {
	return l.level
}

func (l *Level) GetInfo() (occu int, unocc int) {
	for _, spot := range l.Spots {
		if spot.IsAvailable() {
			unocc++
		} else {
			occu++
		}
	}
	return
}
