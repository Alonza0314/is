package is

// PowerOf reports whether input is a power of base. It returns false when
// input is less than 1 or base is less than 2.
func PowerOf(input int, base int) bool {
	if input < 1 || base < 2 {
		return false
	}

	for input%base == 0 {
		input /= base
	}

	return input == 1
}

// PowerOfTwo reports whether input is a power of 2.
func PowerOfTwo(input int) bool {
	return input > 0 && input&(input-1) == 0
}
