package main

/*

### 1. **Калькулятор для терминала**
**Задача**: Напишите программу, которая:
- Принимает два числа и операцию (`+`, `-`, `*`, `/`) как аргументы командной строки.
- Выводит результат вычисления.
**Пример**:
```bash
go run calculator.go 5 3 +
> 8
```
**Дополнительно**: Обработайте деление на ноль и неверные аргументы.

*/

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	// fmt.Println("Program path:", os.Args[0])
	// fmt.Println("Arguments:", os.Args[1:])

	operators := []string{"+", "-", "*", "/"}

	if len(os.Args) <= 1 || len(os.Args) > 4 {
		fmt.Println("\n[!] Usage: \n\tprogram x y operator")
		return
	}

	val_x, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("[!] Error parsing x: ", err)
		return
	}
	val_y, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("[!] Error parsing y: ", err)
		return
	}
	if !slices.Contains(operators, os.Args[3]) {
		fmt.Println("[!] Operator must be one of: + - * /")
		return
	}

	switch os.Args[3] {
	case "+":
		fmt.Println(val_x, " + ", val_y, " = ", val_x+val_y)
	case "-":
		fmt.Println(val_x, " - ", val_y, " = ", val_x-val_y)
	case "*":
		fmt.Println(val_x, " * ", val_y, " = ", val_x*val_y)
	case "/":
		if val_y == 0 {
			fmt.Println("Zero division!")
			return
		}
		fmt.Println(val_x, " / ", val_y, " = ", val_x/val_y)
	}

}
