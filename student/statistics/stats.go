package guessit2

func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Calculate the Linear Regression Line
func LinearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))
	sumX, sumY := 0.0, 0.0
	sumXY, sumX2 := 0.0, 0.0

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return slope, intercept
}

func Variance(x []float64) float64 {
	mean := Average(x)
	out := []float64{}

	for _, num := range x {
		dev := num - mean
		out = append(out, dev*dev)
	}
	return Average(out)
}
