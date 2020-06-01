package random

// GPUBrand returns random gpu brand
func GPUBrand() string {
	set := []string{"NVIDIA", "AMD"}
	return randomStringFromSet(set...)
}

// GPUName returns random gpu name for the brand
func GPUName(brand string) string {
	var set []string
	if brand == "NVIDIA" {
		set = []string{
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		}
		return randomStringFromSet(set...)
	}
	set = []string{
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	}
	return randomStringFromSet(set...)
}
