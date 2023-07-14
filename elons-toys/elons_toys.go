package elon

import "fmt"

// Drive updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage
func (c *Car) Drive() {
	distance := c.distance
	battery := c.battery
	if battery-c.batteryDrain >= 0 {
		distance += c.speed
		battery -= c.batteryDrain
	}
	c.battery = battery
	c.distance = distance
}

// DisplayDistance returns a current state of the driven distance
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a current state of the battery
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if a remote control car can finish a race
func (c *Car) CanFinish(trackDistance int) bool {
	possibleRacesAmount := int(c.battery / c.batteryDrain)
	possibleDistance := possibleRacesAmount * c.speed

	return possibleDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
