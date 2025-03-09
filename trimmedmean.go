package trimmedmean

import (
	"errors"
	"sort"
)

// Trim is the proportion (between 0 and 0.5) to trim from both ends.
func TrimmedMean(data []float64, trim float64) (float64, error) {
	return TrimmedMeanAsym(data, trim, trim)
}

// Each trim proportion must be between 0 and 0.5, and their sum must not exceed 1.
func TrimmedMeanAsym(data []float64, trimLow, trimHigh float64) (float64, error) {
	// Check if the data slice is empty
	if len(data) == 0 {
		return 0, errors.New("empty data slice")
	}

	// Check if the trim proportions are valid
	if trimLow < 0 || trimHigh < 0 || trimLow > 0.5 || trimHigh > 0.5 || (trimLow+trimHigh) >= 1 {
		return 0, errors.New("trim proportions must be between 0 and 0.5")
	}

	// create new copy to not modify original slice
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	// Calculate num elements to trim from each end
	lowTrim := int(float64(len(data)) * trimLow)
	highTrim := int(float64(len(data)) * trimHigh)

	// Calculate mean of remaining elements
	sum := 0.0
	count := 0

	// Sum the remaining elements
	for i := lowTrim; i < len(sortedData)-highTrim; i++ {
		sum += sortedData[i]
		count++
	}

	// Check if no elements remain after trimming
	if count == 0 {
		return 0, errors.New("no elements remain after trimming")
	}

	return sum / float64(count), nil
}

// TrimmedMeanInt calculates the trimmed mean of a slice of integers with symmetric trimming.
func TrimmedMeanInt(data []int, trim float64) (float64, error) {
	return TrimmedMeanIntAsym(data, trim, trim)
}

// TrimmedMeanIntAsym calculates the trimmed mean of integers with asymmetric trimming.
func TrimmedMeanIntAsym(data []int, trimLow, trimHigh float64) (float64, error) {
	// Convert integers to float64
	floatData := make([]float64, len(data))
	for i, v := range data {
		floatData[i] = float64(v)
	}
	return TrimmedMeanAsym(floatData, trimLow, trimHigh)
}
