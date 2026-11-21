package personaldata

import "fmt"

// Personal — структура с персональными данными пользователя.
type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес пользователя в кг
	Height float64 // Рост пользователя в метрах
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
