package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0.0, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0.0, errors.New("Failed to parse float from file.")
	}

	return value, nil
}
