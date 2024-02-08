package eho

func GetDirection(year int16, language string) (degree int, name string) {
	onesPlace5Rem := (year % 10) % 5

	if onesPlace5Rem == 0 {
		degree = 225
		name = "西南西"
	} else if onesPlace5Rem == 2 {
		degree = 345
		name = "北北西"
	} else if onesPlace5Rem == 4 {
		degree = 75
		name = "東北東"
	} else {
		degree = 165
		name = "南南東"
	}

	return degree, name
}
