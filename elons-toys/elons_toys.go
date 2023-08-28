package elon

import "fmt"

// Drive method drives the car one time.
// If there is not enough battery to drive one more time, the car will not move.
func (c *Car) Drive() {
	if c.battery > 0 && c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance method displays how long in meters Car has driven.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery method displays the percentage of Car battery.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (c *Car) CanFinish(trackDistance int) bool {
	return (c.speed*c.battery)/c.batteryDrain >= trackDistance
}
