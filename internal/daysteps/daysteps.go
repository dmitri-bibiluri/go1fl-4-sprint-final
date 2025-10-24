package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataToStrings := strings.Split(data, ",")
	if len(dataToStrings) != 2 {
		return 0, 0, errors.New("not enough data to parse a training")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(dataToStrings[0]))
	if err != nil{
		return 0, 0, errors.New("failed to convert steps to int: " + err.Error())
	}
	if steps < 1{
		return 0, 0, errors.New("can't execute program if there are less than 1 steps")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(dataToStrings[1]))
	if err != nil{
		return 0, 0, errors.New("failed to parse duration: " + err.Error())
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
