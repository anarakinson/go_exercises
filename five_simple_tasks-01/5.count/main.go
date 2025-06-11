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

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(countWords("  Extra   spaces   "))
	fmt.Println(countWords("  Extra   spaces"))
	fmt.Println(countWords("Extra spaces"))
	fmt.Println(countWords("Extra  spaces      "))
	fmt.Println(countWords("... Extra,   spaces !  "))

}

func countWords(str string) int {

	var count int
	var temp rune = 0
	for i, r := range str {
		// если не начало строки, текущая руна пробел или пункт. и предыдущая не пробел или пункт.
		if space_or_punct(r) && !space_or_punct(temp) && temp != 0 {
			// значит слово только закончилось
			count += 1
			// если конец строки и руна не пробел или пункт.
		} else if !space_or_punct(r) && i == len(str)-1 {
			// значит слово только закончилось
			count += 1
		}
		temp = r
	}
	return count

}

func space_or_punct(r rune) bool {
	return r == ' ' || unicode.IsPunct(r)
}
