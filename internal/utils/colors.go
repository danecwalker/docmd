package utils

import "strconv"

func HEX2HSL(hex string) (h, s, l float64) {
	r, g, b := HEX2RGB(hex)
	return RGB2HSL(r, g, b)
}

func HEX2RGB(hexStr string) (r, g, b int) {
	hexStr = hexStr[1:] // Remove the '#' prefix
	if len(hexStr) == 6 {
		rx, _ := strconv.ParseInt(hexStr[0:2], 16, 32)
		gx, _ := strconv.ParseInt(hexStr[2:4], 16, 32)
		bx, _ := strconv.ParseInt(hexStr[4:6], 16, 32)

		r = int(rx)
		g = int(gx)
		b = int(bx)
	} else if len(hexStr) == 3 {
		rx, _ := strconv.ParseInt(hexStr[0:1], 16, 32)
		gx, _ := strconv.ParseInt(hexStr[1:2], 16, 32)
		bx, _ := strconv.ParseInt(hexStr[2:3], 16, 32)

		r = int(rx * 17)
		g = int(gx * 17)
		b = int(bx * 17)
	}
	return
}

func RGB2HSL(r, g, b int) (h, s, l float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0
	max := max(rf, gf, bf)
	min := min(rf, gf, bf)
	l = (max + min) / 2.0

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case rf:
			h = (gf - bf) / d
			if g < b {
				h += 6
			}
		case gf:
			h = (bf-rf)/d + 2
		case bf:
			h = (rf-gf)/d + 4
		}
		h *= 60
		if h < 0 {
			h += 360
		}
	}
	return
}

// max returns the maximum of three float64 values.
func max(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}

// min returns the minimum of three float64 values.
func min(x, y, z float64) float64 {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}
