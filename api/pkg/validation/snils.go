package validation

import (
	"fmt"
	"strconv"
)

func Snils(snils string) error {
	checkSum, err := strconv.Atoi(snils[len(snils)-2:])
	if err != nil {
		return fmt.Errorf("last 2 nums of snils must be numeric")
	}

	lengthWithoutCheckSum := len(snils) - 2
	sum := 0
	for i := 0; i < lengthWithoutCheckSum; i++ {
		currentNumber, err := strconv.Atoi(string(snils[i]))
		if err != nil {
			return fmt.Errorf("snils must be numbers")
		}

		sum += currentNumber * (lengthWithoutCheckSum - i)
	}

	if checkSum != sum {
		return fmt.Errorf("invalid check sum")
	}

	return nil
}
