package utils

import (
	"os"
)

func GetData(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
