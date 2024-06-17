package utils

import (
	"errors"
	"strings"
)

// Gets the file extension of a given string, when the format is file.type
func GetFileExtension(name string) (string, error) {
	p := strings.Split(name, ".")
	if len(p) > 1 {
		return p[len(p)-1], nil
	}

	return "", errors.New("bad file name or file")
}
