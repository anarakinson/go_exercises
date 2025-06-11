package main

import "fmt"

/*

### 3. **Конвертер температуры**
**Задача**: Напишите функцию, которая конвертирует:
- Градусы Цельсия в Фаренгейты и наоборот.
**Интерфейс**:
```go
func convert(temp float64, unit string) (float64, string) {}
```
**Пример**:
```go
convert(0, "C") // возвращает (32, "F")
convert(32, "F") // возвращает (0, "C")

*/

func main() {

	fmt.Println(convert(0, "C")) // возвращает (32, "F")
	fmt.Println(convert(32, "F")) // возвращает (0, "C")

}


func convert(temp float64, unit string) (float64, string) {
	switch unit {
	case "F":
		return (temp - 32) * 5/9, "C"
	case "C":
		return temp * 9/5 + 32, "F"
	default:
		panic("Unrecognizable units")
	}

}