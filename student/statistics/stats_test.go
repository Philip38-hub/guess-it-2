package guessit2

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive numbers",
			args: args{
				data: []float64{1, 2, 3, 4, 5},
			},
			want: 3.0, // Expected mean: (1+2+3+4+5)/5 = 15/5 = 3.0
		},
		{
			name: "Negative numbers",
			args: args{
				data: []float64{-1, -2, -3, -4, -5},
			},
			want: -3.0, // Expected mean: (-1+-2+-3+-4+-5)/5 = -15/5 = -3.0
		},
		{
			name: "Mixed numbers",
			args: args{
				data: []float64{-1, 2, -3, 4, -5},
			},
			want: -0.6, // Expected mean: (-1+2+-3+4+-5)/5 = -3/5 = -0.6
		},
		{
			name: "Single value",
			args: args{
				data: []float64{10},
			},
			want: 10.0, // Expected mean: 10/1 = 10.0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.data); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "Basic linear relationship",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{2, 4, 6, 8, 10},
			},
			want:  2.0,    // Expected slope (m)
			want1: 0.0,    // Expected intercept (c)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LinearRegression(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("LinearRegression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LinearRegression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test with positive integers",
			args: args{x: []float64{10, 20, 30, 40, 50}},
			want: 200,
		},
		{
			name: "Test with a single element",
			args: args{x: []float64{42}},
			want: 0,
		},
		{
			name: "Test with identical elements",
			args: args{x: []float64{5, 5, 5, 5}},
			want: 0,
		},
		{
			name: "Test with a mix of positive and negative numbers",
			args: args{x: []float64{-10, -5, 0, 5, 10}},
			want: 50,
		},
		{
			name: "Test with decimals",
			args: args{x: []float64{1.5, 2.5, 3.5, 4.5}},
			want: 1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.x); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
