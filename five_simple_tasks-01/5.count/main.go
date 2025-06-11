package main

/*

### 5. **Подсчет слов в строке**
**Задача**: Напишите функцию, которая принимает строку и возвращает количество слов (разделенных пробелами).
**Пример**:
```go
countWords("Hello world!") // 2
countWords("  Extra   spaces   ") // 2 (игнорируем лишние пробелы)
```
**Усложнение**: Удалите знаки препинания (например, `!?,.)` перед подсчетом.

*/

import "fmt"

func main() {

	fmt.Println(countWords("  Extra   spaces   "))
	fmt.Println(countWords("  Extra   spaces"))
	fmt.Println(countWords("Extra spaces"))
	fmt.Println(countWords("Extra  spaces      "))

}

func countWords(str string) int {

	var count int
	var temp rune = 0
	for i, ch := range str {
		// если не начало строки, текущая руна равна " " и предыдущая не равна " "
		if ch == ' ' && temp != ' ' && temp != 0 {
			// значит слово только закончилось
			count += 1
			// если конец строки и руна не равна " "
		} else if ch != ' ' && i == len(str)-1 {
			// значит слово только закончилось
			count += 1
		}
		temp = ch
	}
	return count

}
