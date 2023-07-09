package validators

import (
	"errors"
	"strconv"
)

func NumValidator(num string) error {
	if len(num) > 5 || len(num) < 0 {
		return errors.New("number must be between 1 and 5 digits")
	}

	if convFloat, err := strconv.ParseFloat(num, 64); err == nil {
		if convFloat < 0 {
			return errors.New("number must be positive and greater than 0")
		}
	} else {
		return errors.New("invalid number")
	}

	return nil
}
