package main

import "fmt"

func main() {
	var day int
	var month int

	var zodiac string

	fmt.Scanln(&day)
	fmt.Scanln(&month)

	if month == 1 {
		if day < 20 {
			zodiac = "capricornio"
		} else {
			zodiac = "acuario"
		}
	}

	if month == 2 {
		if day < 19 {
			zodiac = "acuario"
		} else {
			zodiac = "piscis"
		}
	}

	if month == 3 {
		if day < 21 {
			zodiac = "piscis"
		} else {
			zodiac = "aries"
		}
	}

	if month == 4 {
		if day < 20 {
			zodiac = "aries"
		} else {
			zodiac = "tauro"
		}
	}

	if month == 5 {
		if day < 21 {
			zodiac = "tauro"
		} else {
			zodiac = "geminis"
		}
	}

	if month == 6 {
		if day < 21 {
			zodiac = "geminis"
		} else {
			zodiac = "cancer"
		}
	}

	if month == 7 {
		if day < 23 {
			zodiac = "cancer"
		} else {
			zodiac = "leo"
		}
	}

	if month == 8 {
		if day < 23 {
			zodiac = "leo"
		} else {
			zodiac = "virgo"
		}
	}

	if month == 9 {
		if day < 23 {
			zodiac = "virgo"
		} else {
			zodiac = "libra"
		}
	}

	if month == 10 {
		if day < 23 {
			zodiac = "libra"
		} else {
			zodiac = "escorpio"
		}
	}

	if month == 11 {
		if day < 22 {
			zodiac = "escorpio"
		} else {
			zodiac = "sagitario"
		}
	}

	if month == 12 {
		if day < 22 {
			zodiac = "sagitario"
		} else {
			zodiac = "capricornio"
		}
	}

	fmt.Println(zodiac)
}
