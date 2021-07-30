package ratingparser

import "fmt"

func FormatSnils(snils string) string {
	return fmt.Sprintf("%s-%s-%s %s", snils[:3], snils[3:6], snils[6:9], snils[9:11])
}
