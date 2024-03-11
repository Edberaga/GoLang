package main

import "fmt"

type Celcius float64

func (c Celcius) toFahrenhiet() float64 {
	//return c *9 / 5 + 32 Note: error for not able to use Celcius data even if its from float64
	return float64(c * 9 / 5 + 32)
}

func main() {
	var derajat Celcius = 20.0
	fmt.Println("Derajat Celcius: ", derajat)
	fmt.Println("Derajat Fahrenhit: ", derajat.toFahrenhiet())
}
