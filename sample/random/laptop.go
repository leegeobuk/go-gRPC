package random

// LaptopBrand returns random laptop brand
func LaptopBrand() string {
	set := []string{"Apple", "Dell", "Lenovo"}
	return randomStringFromSet(set...)
}

// LaptopName returns random laptop name of a brand
func LaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
