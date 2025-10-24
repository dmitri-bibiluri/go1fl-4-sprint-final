package spentcalories

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
 	dataToStrings := strings.Split(data, ",")
    if len(dataToStrings) != 2 {
        return 0, "", 0, errors.New("Not enough data to parse a training")
    }

	steps, err := strconv.Atoi(dataToStrings[0])
	if err != nil{
		return 0, "", 0, errors.New("Failed to convert steps to int: " + err.Error())
	}else if steps < 1{
		return 0, "", 0, errors.New("Can't execute programm if there are less than 1 steps")
	}

	duration, err := time.ParseDuration(dataToStrings[2])
		if err != nil{
			return 0, "", 0, errors.New("Failed to parse duration: " + err.Error())
		}
	return steps, dataToStrings[1], duration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
