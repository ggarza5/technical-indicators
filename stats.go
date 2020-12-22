package indicators

import "math"
import _ "fmt"
import _ "strconv"

type mfloat []float64

// Avg returns 'data' average.
func Avg(data []float64) float64 {

	return Sum(data) / float64(len(data))
}

// Sum returns the sum of all elements of 'data'.
func Sum(data []float64) float64 {

	var sum float64

	for _, x := range data {
		sum += x
	}

	return sum
}

// Std returns standard deviation of a slice.
func Std(slice []float64) float64 {

	var result []float64

	mean := Avg(slice)

	for i := 0; i < len(slice); i++ {
		result = append(result, math.Pow(slice[i]-mean, 2))
	}

	return math.Sqrt(Sum(result) / float64(len(result)))
}

// AddToAll adds a value to all slice elements.
func (slice mfloat) AddToAll(val float64) []float64 {

	var addedSlice []float64

	for i := 0; i < len(slice); i++ {
		addedSlice = append(addedSlice, slice[i] + val)
	}

	return addedSlice
}

// SubSlices subtracts two slices.
func SubSlices(slice1, slice2 []float64) []float64 {

	var result []float64

	for i := 0; i < len(slice1); i++ {
		result = append(result, slice1[i]-slice2[i])
	}

	return result	
}

// AddSlices adds two slices.
func AddSlices(slice1, slice2 []float64) []float64 {

	var result []float64
	// println(len(slice1))
	// println(len(slice2))
	for i, _ := range(slice2) {
		result = append(result, slice1[i]+slice2[i])
	}
	// println("done adding")
	return result	
}

func AddSlicesFromReverse(slice1, slice2 []float64) []float64 {

	var result []float64
	// println(len(slice1))
	// println(len(slice2))
	if len(slice1) > len(slice2) {
		for i, _ := range(slice2) {
			result = append(result, slice1[len(slice1) - i - 1]+slice2[len(slice2) - i - 1])
		}
	} else {
		for i, _ := range(slice1) {
			result = append(result, slice1[len(slice1) - i]+slice2[len(slice2) - i])
		}
	}
	var reversedResults []float64
	for i, _ := range result {
		reversedResults = append(reversedResults, result[len(result) - i - 1])
	}
	// println("done adding")
	return reversedResults	
} 

// DivSlice divides a slice by a float.
func DivSlice(slice []float64, n float64) []float64 {

	var result []float64

	for i := 0; i < len(slice); i++ {
		result = append(result, slice[i]/n)
	}

	return result
}

func sliceMax(slice []float64) float64 {
	var m float64 = 0
	for i, e := range slice {
	    if i == 0 || e > m {
	        m = e
	    }
	}
	// println(m)
	return m
}

func sliceMin(slice []float64) float64 {
	var m float64 = math.MaxFloat64
	for i, e := range slice {
	    if i == 0 || e < m {
	        m = e
	    }
	}
	return m
}