package main

func NewTruck(number string) *BaseV {
	return &BaseV{numberPlate: number, vType: TRUCK}
}
