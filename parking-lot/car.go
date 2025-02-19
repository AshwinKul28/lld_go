package main

func NewCar(number string) *BaseV {
	return &BaseV{numberPlate: number, vType: CAR}
}
