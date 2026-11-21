package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для парсинга данных и формирования информации.
// Реализуется структурами DaySteps и Training из соответствующих пакетов.
type DataParser interface {
	// Parse парсит строку с данными и записывает их в структуру
	Parse(data string) error
	// ActionInfo формирует строку с информацией о активности
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными и экземпляр структуры, реализующей DataParser.
// Для каждой строки выполняется парсинг и вывод информации.
func Info(dataset []string, dp DataParser) {
	// Перебираем все строки с данными
	for _, data := range dataset {
		// Парсим строку
		err := dp.Parse(data)
		if err != nil {
			// Если ошибка парсинга — логируем и переходим к следующей строке
			log.Printf("ошибка парсинга: %v", err)
			continue
		}

		// Формируем строку с информацией
		info, err := dp.ActionInfo()
		if err != nil {
			// Если ошибка формирования — логируем и пропускаем
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		// Выводим информацию в stdout
		fmt.Print(info)
	}
}
