package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе.
// Возвращает количество калорий и ошибку, если входные данные некорректны.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем корректность входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные для расчета калорий")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Формула: (weight * speed * durationInMinutes) / minInH
	calories := (weight * speed * durationInMinutes) / minInH

	// Умножаем на корректирующий коэффициент для ходьбы
	calories = calories * walkingCaloriesCoefficient

	return calories, nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
// Возвращает количество калорий и ошибку, если входные данные некорректны.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем корректность входных данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные для расчета калорий")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Формула: (weight * speed * durationInMinutes) / minInH
	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
// Если продолжительность <= 0 или шаги <= 0, возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Если продолжительность нулевая или отрицательная, возвращаем 0
	if duration <= 0 {
		return 0
	}

	// Если шаги отрицательные или нулевые, возвращаем 0
	if steps <= 0 {
		return 0
	}

	// Рассчитываем дистанцию
	distance := Distance(steps, height)

	// Скорость = дистанция (км) / время (часы)
	// duration.Hours() возвращает продолжительность в часах (float64)
	speed := distance / duration.Hours()

	return speed
}

// Distance рассчитывает пройденную дистанцию в километрах.
// Длина шага = рост * stepLengthCoefficient
// Дистанция = шаги * длина шага / mInKm
func Distance(steps int, height float64) float64 {
	// Переводим steps в float64 для умножения
	stepLength := float64(steps) * height * stepLengthCoefficient
	distance := stepLength / mInKm

	return distance
}
