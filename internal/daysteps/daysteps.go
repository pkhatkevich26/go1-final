package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps — структура с данными о дневных прогулках.
// Встроенная personaldata.Personal даёт доступ к весу и росту пользователя.
type DaySteps struct {
	Steps    int                 // Количество шагов
	Duration time.Duration       // Длительность прогулки
	personaldata.Personal // Встроенная структура с персональными данными (безымянное поле — promoted fields)
}

// Parse парсит строку с данными формата "678,0h50m" и записывает данные в структуру.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем строку по запятой — должно быть ровно 2 части
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных: ожидается 2 параметра (шаги, длительность)")
	}

	stepsStr := parts[0]
	durationStr := parts[1]

	// Проверяем что строка шагов не содержит пробелов
	// (пробелы в начале, в конце или внутри — всё ошибка)
	if stepsStr != strings.TrimSpace(stepsStr) || strings.Contains(stepsStr, " ") {
		return errors.New("неверный формат шагов: обнаружены пробелы")
	}

	// Парсим количество шагов
	// strconv.ParseInt обрабатывает "+" в начале, но не пробелы
	steps, err := strconv.ParseInt(stepsStr, 10, 64)
	if err != nil {
		return errors.New("неверное количество шагов: не является числом")
	}
	if steps <= 0 {
		return errors.New("неверное количество шагов: должно быть больше 0")
	}
	ds.Steps = int(steps)

	// Парсим длительность
	// time.ParseDuration парсит форматы: "1h30m", "30m", "2h", "1.5h", "30.5m"
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("неверный формат длительности")
	}
	// Проверяем что длительность больше 0
	if duration <= 0 {
		return errors.New("длительность должна быть больше 0")
	}
	// Проверяем что в строке есть хотя бы одна единица измерения "m" или "h"
	// Это отсекает форматы вроде "1.5d" (дни)
	if !strings.Contains(durationStr, "m") && !strings.Contains(durationStr, "h") {
		return errors.New("неверный формат длительности: отсутствуют единицы измерения (m или h)")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Проверяем что данные корректны
	if ds.Steps <= 0 {
		return "", errors.New("некорректное количество шагов")
	}
	if ds.Duration <= 0 {
		return "", errors.New("некорректная продолжительность")
	}
	if ds.Weight <= 0 {
		return "", errors.New("некорректный вес")
	}
	if ds.Height <= 0 {
		return "", errors.New("некорректный рост")
	}

	// Рассчитываем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Рассчитываем сожжённые калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем строку с информацией
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return info, nil
}
