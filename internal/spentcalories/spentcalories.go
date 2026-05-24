package spentcalories

import (
	"fmt"
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
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format: expected 3 values, got %d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid data format: %w", err)
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid data format: steps must be positive")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid data format: %w", err)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid data format: duration must be positive")
	}
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLen := height * stepLengthCoefficient
	distM := float64(steps) * stepLen
	return distM / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distKm := distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distKm / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	// Вычисляем общие метрики
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	var calories float64
	switch activity {
	case "Бег":
		cal, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		calories = cal
	case "Ходьба":
		cal, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		calories = cal
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	// Форматируем вывод согласно тестам
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity,
		durationHours,
		dist,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive, got %v", duration)
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * averageSpeed * durationMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive, got %v", duration)
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * averageSpeed * durationMinutes) / minInH
	calories = calories * walkingCaloriesCoefficient
	return calories, nil
}
