/*
https://exercism.org/tracks/go/exercises/cars-assemble/edit
*/
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) / 60 * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var carFullGroups, remainingCars uint

	if carsCount >= 10 {
		carFullGroups = uint(carsCount) / 10
		remainingCars = uint(carsCount) - carFullGroups*10
	} else {
		carFullGroups = uint(carsCount) / 10
		remainingCars = uint(carsCount)
	}

	return carFullGroups*95000 + remainingCars*10000
}
