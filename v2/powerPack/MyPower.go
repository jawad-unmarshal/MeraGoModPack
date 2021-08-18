package powerPack

func FindPow(base float64, raise int) float64 {
	var pow float64 = 1.
	for i := 1; i <= raise; i++ {
		pow *= base
	}
	return pow
}
