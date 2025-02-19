package main

func NewBike(number string) *BaseV {
	return &BaseV{numberPlate: number, vType: MOTORCYCLE}
}
