package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func sort() {
	arr := [10]int{9, 3, 5, 2, 8, 0, 1, 4, 7, 6}
	fmt.Println("\nИсходный массив")
	fmt.Println(arr)

	for i := 1; i < 10; i++ {
		key := arr[i]                //фиксируем текущий элемент
		j := i - 1                   //сравниваем с предыдущим
		for j >= 0 && arr[j] > key { //просматриваем подмассив до предыдущего элемента, если тек эл-т меньше пред-го
			arr[j+1] = arr[j] //меняем местами
			j--               //уменьшаем счётчик
		}
		arr[j+1] = key //тек-й элемент на корректное место
	}

	fmt.Println("Отсортированный массив")
	fmt.Println(arr)
}

func storage_empl() {

	const size = 512
	empls := [size]*Employee{}
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scan(&empl.Name)

			fmt.Println("Возраст:")
			fmt.Scan(&empl.Age)

			fmt.Println("Позиция:")
			fmt.Scan(&empl.Position)

			fmt.Println("Зарплата:")
			fmt.Scan(&empl.Salary)

			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					break
				}
			}
		case 2:
			var name string
			fmt.Println("Удаляем сотрудника")
			fmt.Println("\nВведите имя сотрудника:")
			fmt.Scan(&name)
			flag := false
			for i := 0; i < size; i++ {
				if empls[i] != nil && empls[i].Name == name {
					empls[i] = nil
					flag = true
					fmt.Println("Сотрудник удалён")
					break
				}
			}
			if !flag {
				fmt.Println("Сотрудник не найден")
			}

		case 3:
			fmt.Println("Вывод сотрудников")
			count := 0
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf("%d) Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n", i+1, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
					count++
				}
			}
			if count == 0 {
				fmt.Println("\nСотрудников нет")
			}
		case 4:
			return

		default:
			fmt.Println("Неверная команда!")
		}
	}

}

func main() {
	fmt.Println("1 задание - сортировка вставками")
	sort()

	fmt.Println("\n2 задание")
	storage_empl()
}
