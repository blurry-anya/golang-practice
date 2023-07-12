package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	distance := car.distance
	battery := car.battery
	if battery-car.batteryDrain >= 0 {
		distance += car.speed
		battery -= car.batteryDrain
	}

	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      battery,
		distance:     distance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	possibleRacesAmount := int(car.battery / car.batteryDrain)
	possibleDistance := possibleRacesAmount * car.speed

	return possibleDistance >= track.distance
}
