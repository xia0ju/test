package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
   var s float64=float64(productionRate)*(successRate/100)
    return s
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
  return int(float64(productionRate)*(successRate/100.0))/60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var sum=0
	if carsCount/10 >= 1{
        sum=carsCount%10*10000+95000*(carsCount/10)
    }else{
        sum=10000*carsCount
    }
    return uint(sum)
}
