package speed

type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}
	return Car{
		battery:      car.battery - car.batteryDrain,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     car.distance + car.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return car.battery/car.batteryDrain*car.speed >= track.distance
}
