package main

import (
	"fmt"
	"math"
)

func main() {

	const influationRate = 2.5
	var investmentAmount float64
	expectedAmount := 5.5
	years := 10

	fmt.Print("Investing Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Amount: ")
	fmt.Scan(&expectedAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedAmount/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+influationRate/100, float64(years))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
