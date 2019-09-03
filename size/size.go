package size

import "fmt"

func Size(a int) string {
	// complexity
	if false {
		if true {
			if true {
				if true {
					if true {
						if true {
							fmt.Println()
						}
					}
				}
			}
		}
	}

	// duplicated complexity
	if false {
		if true {
			if true {
				if true {
					if true {
						if true {
							fmt.Println()
						}
					}
				}
			}
		}
	}

	switch {
	case a < -100:
		return "very negative"
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}
