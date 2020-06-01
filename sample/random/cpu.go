package random

import "math/rand"

// CPUBrand returns random cpu brand
func CPUBrand() string {
	set := []string{"Intel", "AMD"}
	return randomStringFromSet(set...)
}

// CPUName returns random name for cpu
func CPUName(brand string) string {
	var set []string
	if brand == "Intel" {
		set = []string{
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		}
		return randomStringFromSet(set...)
	}
	set = []string{
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200U",
	}
	return randomStringFromSet(set...)
}

// Integer returns integer from min to max inclusively
func Integer(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// Float is float version of Integer
func Float(min, max float64) float64 {
	return min + rand.Float64()*(max-min+0.1)
}

func randomStringFromSet(set ...string) string {
	n := len(set)
	if n == 0 {
		return ""
	}
	return set[rand.Intn(n)]
}
