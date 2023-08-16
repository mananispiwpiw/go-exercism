package elon

import "strconv"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.batteryDrain > c.battery {
		c.distance = 0
	} else {
		c.battery -= c.batteryDrain
		c.distance = c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	remainingDistance := (c.battery / c.batteryDrain) * c.speed
	return remainingDistance >= trackDistance
}

// name: "Car has 40% battery. Car can finish the race",
// 			car: Car{
// 				speed:        5,
// 				batteryDrain: 2,
// 				battery:      40,
// 			},
// 			trackDistance: 100,
// 			expected:      true,

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
