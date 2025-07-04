package utils

import "math"

func CalculateGoldPrice(Purity int64, dryWeight float64, weightReduction float64, goldPrice float64) float64 {
	netWeight := RoundTo(dryWeight-weightReduction, 2)
	purityFactor := float64(Purity) / 1000
	price := netWeight * purityFactor * goldPrice
	return price
}

func RoundTo(val float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(val*factor) / factor
}
