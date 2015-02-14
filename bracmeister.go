package bracmeister

type Brac struct {
	Cup   string
	Under float32
}

var Cups = []string{
	"less than AAA", "AAA", "AA", "A",
	"B", "C", "D", "E",
	"F", "G", "H", "I",
	"J", "K", "L", "M",
	"N", "O", "P", "more than P",
}

func Calc(height, bust, waist int, correct bool) *Brac {
	var (
		num  int
		diff float32
	)

	h, b, w := float32(height), float32(bust), float32(waist)

	// Height correction
	if correct {
		diff = (b - (h * 0.54)) + (((h * 0.38) - w) * 0.73) + ((h - 158.8) * 0.3261)
	} else {
		diff = (b - (h * 0.54)) + (((h * 0.38) - w) * 0.73) + ((h - 158.8) * 0.1087)
	}

	if h <= 0.0 {
		num = 20
	} else if bust <= 0.0 {
		num = 21
	} else if waist <= 0.0 {
		num = 22
	} else if diff < -13.75 {
		num = 0
	} else if diff < -11.25 {
		num = 1
	} else if diff < -8.75 {
		num = 2
	} else if diff < -6.25 {
		num = 3
	} else if diff < -3.75 {
		num = 4
	} else if diff < -1.25 {
		num = 5
	} else if diff < 1.25 {
		num = 6
	} else if diff < 3.75 {
		num = 7
	} else if diff < 6.25 {
		num = 8
	} else if diff < 8.75 {
		num = 9
	} else if diff < 11.25 {
		num = 10
	} else if diff < 13.75 {
		num = 11
	} else if diff < 16.25 {
		num = 12
	} else if diff < 18.75 {
		num = 13
	} else if diff < 21.25 {
		num = 14
	} else if diff < 23.75 {
		num = 15
	} else if diff < 26.25 {
		num = 16
	} else if diff < 28.75 {
		num = 17
	} else if diff < 31.25 {
		num = 18
	} else {
		num = 19
	}

	// Calculate under bust from result value.
	brac := &Brac{Cup: Cups[num]}
	if num < 20 {
		brac.Under = b - (diff + 17.5)
	}
	return brac
}
