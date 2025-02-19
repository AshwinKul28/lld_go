package main

type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

type Vehicle interface {
	getNumberPlate() string
	getType() VehicleType
}

type BaseV struct {
	numberPlate string
	vType       VehicleType
}

func (b *BaseV) getNumberPlate() string {
	return b.numberPlate
}

func (b *BaseV) getType() VehicleType {
	return b.vType
}
