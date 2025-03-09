# Trimmed Mean

This package supports trimmed means for both float and integers in Go. 

This performs both symmetric and asymmetric trims on our data,
and has logic for handling both floating-point and integer inputs.

## What is a trimmed mean?

A trimmed mean is when we take a given proportion of the length of our data,
and then we remove elements from both ends.

This enables us to reduce outlier influence on our dataset,
perhaps for implementations in bootstrapping or other statistical methods.

## How we've implemented the logic
Our implementation follows an approach consistent with general data structures and loops.
We use arrays of floats/integers to first obtain a data slice.

After this, we perform the "trimming" by passing a proportion parameter, and then
cast the float as an int, and retrieve the elements within the lower bound and upper bound.

## Function Definitions

### Symmetric Trimming
```go
func TrimmedMean(data []float64, trim float64) (float64, error)
```
This function calculates a symmetric trimmed mean where the same proportion is trimmed from both ends.
- `data`: The input slice of float64 values
- `trim`: The proportion to trim from each end (must be between 0 and 0.5)
- Returns: The trimmed mean and any error that occurred

### Asymmetric Trimming
```go
func TrimmedMeanAsym(data []float64, trimLow, trimHigh float64) (float64, error)
```
This function allows different proportions to be trimmed from the lower and upper ends.
- `data`: The input slice of float64 values
- `trimLow`: The proportion to trim from the lower end (must be between 0 and 0.5)
- `trimHigh`: The proportion to trim from the upper end (must be between 0 and 0.5)

### Integer Input Functions
```go
func TrimmedMeanInt(data []int, trim float64) (float64, error)
```
Symmetric trimming for integer slices. Works the same as `TrimmedMean` but accepts integer input.

```go
func TrimmedMeanIntAsym(data []int, trimLow, trimHigh float64) (float64, error)
```
Asymmetric trimming for integer slices. Works the same as `TrimmedMeanAsym` but accepts integer input.


## Error Handling

The functions will return an error given the following cases:
- Empty input slice
- Invalid trim proportions (< 0 or > 0.5)
- Sum of trim proportions ≥ 1
- No elements remain after trimming

## Installation

To use this package in your Go project:

```bash
go get github.com/ro-mish/trimmed-mean-go
```
