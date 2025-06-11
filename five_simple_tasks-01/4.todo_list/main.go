package main

/*

### 4. **Упрощенный "Todo List"**
**Задача**: Создайте программу для управления списком дел в памяти (без БД):
- Добавление задачи (`add "Buy milk"`).
- Удаление по индексу (`remove 1`).
- Вывод списка (`list`).
**Пример работы**:
```bash
> add "Learn Go"
> add "Read a book"
> list
1. Learn Go
2. Read a book
> remove 1
> list
1. Read a book

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TodoList struct {
	tasks []string
}

func (tl *TodoList) Add(task string) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *TodoList) Remove(idx int) {
	if idx < 1 || idx > len(tl.tasks) {
		fmt.Println("Wrong index:", idx)
		return
	}
	removed_task := tl.tasks[idx-1]
	tl.tasks = append(tl.tasks[:idx-1], tl.tasks[idx:]...)
	fmt.Println("Task removed:", removed_task)
}

func (tl *TodoList) List() {
	if len(tl.tasks) < 0 {
		fmt.Println("List is empty")
	}
	for i, val := range tl.tasks {
		fmt.Printf("%d. %s\n", i+1, val)
	}
}

func main() {
	tdlist := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Available commands: \n\t1. add <task>\n\t2. remove <index>\n\t3. list\n\t4. exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}
		input_parts := strings.SplitN(input, " ", 2)
		command := input_parts[0]

		switch command {
		case "add":
			if len(input_parts) <= 1 {
				fmt.Println("Missing task's text")
				continue
			}
			tdlist.Add(input_parts[1])
		case "remove":
			if len(input_parts) <= 1 {
				fmt.Println("Missing task's index")
				continue
			}
			var idx int
			// сканируем вторую часть ввода и пытаемся считать число
			// число будет сохранено по ссылке в переменную
			_, err := fmt.Sscanf(input_parts[1], "%d", &idx)
			if err != nil {
				fmt.Println("Index must be integer")
				continue
			}
			tdlist.Remove(idx)
		case "list":
			tdlist.List()
		case "exit":
			fmt.Println("Exiting program")
			return
		default:
			fmt.Println("Unrecognized command.")
			fmt.Println("Available commands: \n\t1. add <task>\n\t2. remove <index>\n\t3. list\n\t4. exit")
		}

	}

}
