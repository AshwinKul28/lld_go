package main

type ParkingSpot struct {
	id            int
	vehicleType   VehicleType
	available     bool
	parkedVehicle Vehicle
}

func NewSpot(id int, tp VehicleType) *ParkingSpot {
	return &ParkingSpot{id: id, vehicleType: tp, available: true}
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.available
}

func (ps *ParkingSpot) ParkVehicle(v Vehicle) {
	ps.parkedVehicle = v
	ps.available = false
}

func (ps *ParkingSpot) UnParkVehicle(v Vehicle) {
	ps.parkedVehicle = nil
	ps.available = true
}
func (ps *ParkingSpot) GetSpotNumber() int {
	return ps.id
}

func (ps *ParkingSpot) GetVehicleType() VehicleType {
	return ps.vehicleType
}

func (ps *ParkingSpot) GetParkedVehicle() Vehicle {
	return ps.parkedVehicle
}
