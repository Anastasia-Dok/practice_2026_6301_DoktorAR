package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stack - структура стека
type stack struct {
	s    []any // слайс для хранения значений
	head int   // индекс головы стека (вершины)
}

// newStack - создание нового стека
func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1, //  стек пуст
	}
}

// push - добавление значения в стек
func push(s *stack, v any) bool {
	// Проверяем, есть ли место в стеке
	if s.head >= len(s.s)-1 {
		return false // стек заполнен
	}
	s.head++
	s.s[s.head] = v
	return true
}

// pop - получение значения из стека и его удаление из вершины
func pop(s *stack) (any, bool) {
	// Проверяем, пуст ли стек
	if s.head < 0 {
		return nil, false // стек пуст
	}
	val := s.s[s.head]
	s.s[s.head] = nil // очищаем ячейку
	s.head--
	return val, true
}

// peek - просмотр значения на вершине стека без удаления
func peek(s *stack) (any, bool) {
	if s.head < 0 {
		return nil, false // стек пуст
	}
	return s.s[s.head], true
}

// isEmpty - проверка, пуст ли стек
func isEmpty(s *stack) bool {
	return s.head < 0
}

///////////////////////////////////////

type queue struct {
	s         []any // слайс в котором хранятся значения
	low, high int   // индексы верхней и нижней границы очереди
	size      int   // размер очереди
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  0,
		high: 0,
	}
}

// push - добавление в очередь значения
func push1(q *queue, v any) bool {
	// Проверяем, заполнена ли очередь
	if q.high >= q.size {
		return false // очередь заполнена
	}
	q.s[q.high] = v
	q.high++
	return true
}

// pop - получения значения из очереди и его удаление
func pop1(q *queue) (any, bool) {
	// Проверяем, пуста ли очередь
	if q.low >= q.high {
		return nil, false // очередь пуста
	}
	val := q.s[q.low]
	q.s[q.low] = nil // очищаем ячейку
	q.low++
	return val, true
}

// isEmpty - проверка, пуста ли очередь
func isEmpty1(q *queue) bool {
	return q.low >= q.high
}

///////////////////////////////////////////////////

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// add - добавление значения в связный список
func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// addFirst - добавление значения в начало списка
func addFirst(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		newItem.next = l.first
		l.first = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) (any, bool) {
	if idx < 0 || idx >= l.size {
		return nil, false // индекс вне диапазона
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v, true
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) bool {
	if idx < 0 || idx >= l.size {
		return false // индекс вне диапазона
	}
	if idx == 0 {
		// Удаляем первый элемент
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		// Ищем предыдущий элемент
		current := l.first
		for i := 0; i < idx-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		if idx == l.size-1 {
			l.last = current // обновляем last
		}
	}
	l.size--
	return true
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	result := make([]any, 0, l.size)
	current := l.first
	for current != nil {
		result = append(result, current.v)
		current = current.next
	}
	return result
}

/////////////////////////////////////////////////////////////

func rim_to_arab(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	// Идём с конца
	for i := len(s) - 1; i >= 0; i-- {
		curr := romanMap[s[i]]
		if curr < prev {
			result -= curr
		} else {
			result += curr
		}
		prev = curr
	}

	return result
}

////////////////////////////////////////////////////////

func fillUnique(rows, cols int) ([][]int, bool) {
	totalCells := rows * cols
	minVal := 1
	maxVal := totalCells * 2 // диапазон в 2 раза больше ячеек
	rangeSize := maxVal - minVal + 1

	// Создаём двумерный массив
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}

	// Map для проверки уникальности
	used := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())

	// Заполняем массив
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for {
				num := rand.Intn(rangeSize) + minVal
				// Проверяем, есть ли число в map
				if _, ok := used[num]; !ok {
					used[num] = true
					arr[i][j] = num
					break
				}
			}
		}
	}

	return arr, true
}

func printMatrix(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%4d ", arr[i][j])
		}
		fmt.Println()
	}
}

////////////////////////////////////////////////////////////////////

func task_1_1() {
	fmt.Println("\nСозданиче стека")
	// Создаём стек на 5 элементов
	st := newStack(5)

	// Добавляем элементы
	push(st, 10)
	push(st, 20)
	push(st, 30)

	// Смотрим на вершину
	val, ok := peek(st)
	if ok {
		fmt.Println("Вершина:", val) // 30
	}

	// Извлекаем элементы
	for !isEmpty(st) {
		val, _ := pop(st)
		fmt.Println("Извлечено:", val)
	}

	fmt.Println("Стек пуст:", isEmpty(st))
}

func task_1_2() {
	fmt.Println("\nСозданиче очереди")
	// Создаём очередь на 5 элементов
	q := newQueue(5)

	// Добавляем элементы
	push1(q, 10)
	push1(q, 20)
	push1(q, 30)

	// Извлекаем элементы
	for !isEmpty1(q) {
		val, _ := pop1(q)
		fmt.Println("Извлечено:", val)
	}

	fmt.Println("Очередь пуста:", isEmpty1(q))
}

func task_1_3() {
	fmt.Println("\nСодаём односвязный список")
	list := newSinglyLinkedList()

	// Добавляем элементы
	add(list, 10)
	add(list, 20)
	add(list, 30)
	addFirst(list, 5)

	// Получаем все значения
	fmt.Println("Все элементы:", values(list))

	// Получаем элемент по индексу
	val, ok := get(list, 2)
	if ok {
		fmt.Println("Элемент [2]:", val) // 20
	}

	// Удаляем элемент
	remove(list, 2) // удаляем 20
	fmt.Println("После удаления:", values(list))

}

func task_2() {
	testCases := []string{
		"III",     // 3
		"IV",      // 4
		"IX",      // 9
		"LVIII",   // 58
		"MMXXIV",  // 2024
		"D",       // 500
		"CM",      // 900
		"MCMXCIV", // 1994
	}

	fmt.Println("\n Римские → Арабские")
	for _, roman := range testCases {
		arabic := rim_to_arab(roman)
		fmt.Printf("%-10s → %d\n", roman, arabic)
	}
}

func task_3() {
	var rows, cols int

	// Ввод размеров с консоли
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)

	fmt.Print("Введите количество столбцов: ")
	fmt.Scan(&cols)

	if rows <= 0 || cols <= 0 {
		fmt.Println("Ошибка: количество строк и столбцов должно быть больше 0")
		return
	}

	totalCells := rows * cols
	maxVal := totalCells * 2
	fmt.Printf("\nСоздание массива %dx%d с уникальными числами от 1 до %d\n", rows, cols, maxVal)

	// Заполняем массив
	arr, ok := fillUnique(rows, cols)
	if !ok {
		return
	}

	// Выводим массив
	fmt.Println("\nСгенерированный массив:")
	printMatrix(arr)
}

func main() {
	task_1_1()
	task_1_2()
	task_1_3()
	task_2()
	task_3()
}
