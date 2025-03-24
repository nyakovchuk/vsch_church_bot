package service

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const coordPattern = `^(-?\d+\.\d+),?\s+(-?\d+\.\d+)$`

var (
	ErrInvalidCoordinatesFormat         = errors.New("некорректный формат координат")
	ErrInvalidCoordinatesTransformation = errors.New("ошибка преобразования координат")
)

func ParseCoordinates(text string) (float64, float64, error) {
	latStr, lonStr, err := splitCoordinates(text)
	if err != nil {
		return 0, 0, err
	}

	lat, lon, err := parseFloats(latStr, lonStr)
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}

func splitCoordinates(input string) (string, string, error) {
	re := regexp.MustCompile(coordPattern)
	matches := re.FindStringSubmatch(strings.TrimSpace(input))

	if len(matches) != 3 {
		return "", "", ErrInvalidCoordinatesFormat
	}

	return matches[1], matches[2], nil
}

func parseFloats(latStr, lonStr string) (float64, float64, error) {
	lat, err1 := strconv.ParseFloat(latStr, 64)
	lon, err2 := strconv.ParseFloat(lonStr, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, ErrInvalidCoordinatesTransformation
	}

	return lat, lon, nil
}
