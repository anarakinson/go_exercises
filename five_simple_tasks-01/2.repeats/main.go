package main

/*

### 2. **Поиск повторяющихся строк в файле**
**Задача**: Программа читает текстовый файл (например, `text.txt`) и выводит строки, которые повторяются >1 раза, с указанием количества.
**Пример файла**:
```
apple
banana
apple
orange
```
**Вывод**:
```
apple: 2
```
**Упрощение**: Можно сначала загрузить весь файл в память.

*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	word_counter := make(map[string]int, 10)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // read until newline
		if err != nil {
			if err == io.EOF {
				break // end of file
			}
			log.Fatal(err)
		}
		word_counter[strings.TrimSpace(line)] += 1
	}

	for key, value := range word_counter {
		fmt.Printf("Word: %s - %d times\n", key, value)
	}

}
