package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training — структура с данными о тренировке.
// Встроенная personaldata.Personal даёт доступ к весу и росту пользователя.
type Training struct {
	Steps        int                 // Количество шагов
	TrainingType string              // Тип тренировки (Ходьба или Бег)
	Duration     time.Duration        // Длительность тренировки
	personaldata.Personal // Встроенная структура с персональными данными (безымянное поле — promoted fields)
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в структуру.
func (t *Training) Parse(datastring string) (err error) {
	// Разделяем строку по запятой — должно быть ровно 3 части
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных: ожидается 3 параметра (шаги, тип, длительность)")
	}

	stepsStr := parts[0]
	trainingType := parts[1]
	durationStr := parts[2]

	// Проверяем что строка шагов не содержит пробелов
	if stepsStr != strings.TrimSpace(stepsStr) || strings.Contains(stepsStr, " ") {
		return errors.New("неверный формат шагов: обнаружены пробелы")
	}

	// Парсим количество шагов
	steps, err := strconv.ParseInt(stepsStr, 10, 64)
	if err != nil {
		return errors.New("неверное количество шагов: не является числом")
	}
	if steps <= 0 {
		return errors.New("неверное количество шагов: должно быть больше 0")
	}
	t.Steps = int(steps)

	// Сохраняем тип тренировки
	t.TrainingType = trainingType

	// Парсим длительность
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("неверный формат длительности")
	}
	// Проверяем что длительность больше 0
	if duration <= 0 {
		return errors.New("длительность должна быть больше 0")
	}
	// Проверяем что в строке есть хотя бы одна единица измерения "m" или "h"
	if !strings.Contains(durationStr, "m") && !strings.Contains(durationStr, "h") {
		return errors.New("неверный формат длительности: отсутствуют единицы измерения (m или h)")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// Проверяем что данные корректны
	if t.Steps <= 0 {
		return "", errors.New("некорректное количество шагов")
	}
	if t.Duration <= 0 {
		return "", errors.New("некорректная продолжительность")
	}
	if t.Weight <= 0 {
		return "", errors.New("некорректный вес")
	}
	if t.Height <= 0 {
		return "", errors.New("некорректный рост")
	}

	// Рассчитываем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)

	// Рассчитываем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// Рассчитываем калории в зависимости от типа тренировки
	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	// Формируем строка с информацией
	// Длительность в часах (float64)
	durationInHours := t.Duration.Hours()

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationInHours, distance, speed, calories)

	return info, nil
}
