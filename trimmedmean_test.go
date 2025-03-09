package trimmedmean

import (
	"math"
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		trim     float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "basic symmetric trim",
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			trim:     0.2,
			expected: 5.5, // removes 1,2 from low end and 9,10 from high end
			wantErr:  false,
		},
		{
			name:     "empty slice",
			data:     []float64{},
			trim:     0.2,
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "invalid trim proportion",
			data:     []float64{1, 2, 3},
			trim:     0.6,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrimmedMean(tt.data, tt.trim)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrimmedMean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("TrimmedMean() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTrimmedMeanAsym(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		trimLow  float64
		trimHigh float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "asymmetric trim",
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			trimLow:  0.1,
			trimHigh: 0.3,
			expected: 4.5,
			wantErr:  false,
		},
		{
			name:     "invalid total trim",
			data:     []float64{1, 2, 3},
			trimLow:  0.5,
			trimHigh: 0.5,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrimmedMeanAsym(tt.data, tt.trimLow, tt.trimHigh)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrimmedMeanAsym() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("TrimmedMeanAsym() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTrimmedMeanInt(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		trim     float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "basic integer trim",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			trim:     0.2,
			expected: 5.5,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrimmedMeanInt(tt.data, tt.trim)
			if (err != nil) != tt.wantErr {
				t.Errorf("error, trimmed mean int = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("error, trimmed mean int = %v, want %v", got, tt.expected)
			}
		})
	}
}
