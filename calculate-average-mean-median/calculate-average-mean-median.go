/*
Simple example of calculations
involving the concepts of
MEAN and MEDIAN

Author: Guilherme Ferreira (guilhermeaferreira_t at yahoo dot com dot br)
https://guilherme-ferreira.me
*/

package calculate_average_mean_median

import (
	"fmt"
	"math"
	"sort"
)

// type
type MeanMedian struct {
	numbers []int
}

func main() {
	// range of int numbers
	mmType := MeanMedian{
		numbers: []int{
			//1671, 1817, 1763, 1733, 1722, 1745, 1773, 1778, 1725, 1696, 1689, 1718, 1735, 1705, 1734,
			1, 2, 3, 4,
		},
	}
	
	sort.Ints(mmType.numbers) // sort

	fmt.Printf("Sorted numbers:\t\t%v\n", mmType.numbers)
	fmt.Println("------------------")
	fmt.Printf("Minimum value:\t\t%v\n", mmType.GetMinValue())
	fmt.Printf("Maximum value:\t\t%v\n", mmType.GetMaxValue())
	fmt.Println("------------------")
	fmt.Printf("Range of values:\t%v\n", mmType.CalcRangeValues())
	fmt.Println("------------------")
	fmt.Printf("Mean value:\t\t%v\n", mmType.CalcMean())
	fmt.Println("------------------")
	fmt.Printf("Median value:\t\t%v\n", mmType.CalcMedian())
}

// return the minimum value
func (mm *MeanMedian) GetMinValue() float64 {
	sort.Ints(mm.numbers) // sort the numbers

	return float64(mm.numbers[0])
}

// return the maximum value
func (mm *MeanMedian) GetMaxValue() float64 {
	sort.Ints(mm.numbers) // sort the numbers

	return float64(mm.numbers[len(mm.numbers)-1])
}

// calculate the range values
// last value - first value
func (mm *MeanMedian) CalcRangeValues() float64 {
	sort.Ints(mm.numbers) // sort the numbers

	return float64(mm.numbers[len(mm.numbers)-1]) - float64(mm.numbers[0])
}

// calculate the "mean" value
// sum of all the int values
// divided by its quantity
func (mm *MeanMedian) CalcMean() float64 {
	total := 0

	for _, v := range mm.numbers {
		total += v
	}

	// IMPORTANT: return was rounded!
	return math.Round(float64(total) / float64(len(mm.numbers)))
}

// calculate the "median" value
// if the total of numbers is odd
// 	takes the middle value
//
// if the total of numbers is even
//	calculate the "mean" of the middle two values
func (mm *MeanMedian) CalcMedian(n ...int) float64 {
	sort.Ints(mm.numbers) // sort the numbers

	mNumber := len(mm.numbers) / 2

	if mm.IsOdd() {
		return float64(mm.numbers[mNumber])
	}

	return float64(mm.numbers[mNumber-1]+mm.numbers[mNumber]) / float64(2)
}

// check if the total of numbers is
// odd or even
func (mm *MeanMedian) IsOdd() bool {
	if len(mm.numbers)%2 == 0 {
		return false
	}

	return true
}
