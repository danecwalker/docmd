package utils

func RGB2HSL(r, g, b float64) (h, s, l float64) {
	max := max(r, g, b)
	min := min(r, g, b)

	l = (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}

		h /= 6
	}

	return
}

func HEX2HSL(hex string) (h, s, l float64) {
	r, g, b := HEX2RGB(hex)
	return RGB2HSL(r, g, b)
}

func HEX2RGB(hex string) (r, g, b float64) {
	r = float64(int(hex[0]) - 48)
	g = float64(int(hex[1]) - 48)
	b = float64(int(hex[2]) - 48)

	return
}
