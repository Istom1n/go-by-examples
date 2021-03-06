////
// В Go, _array (массив)_ представляет собой пронумерованную
// последовательность элементов определенной длины.
////

package main

import "fmt"

func main() {

	// Здесь мы создаем массив `a`, который будет содержать
	// ровно 5 значений типа `int`. Тип элементов и длина являются
	// частью типа массива. По умолчанию массив является нулевым,
	// что для `int` равно значению `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Мы может указывать значение элемента массива по индексу
	// при помощи данного синтаксиса `array[index] = value`,
	// и получать это значение соответственно так `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Встроенная функция `len` возвращает количество элементов
	// в массиве.
	fmt.Println("len:", len(a))

	// Используйте данный синтаксис для обхявления и инициализации
	// массива одной строкой.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Типы массивов являются одномерными, но вы можете создавать
	// типы и для построения многомерных структур данных.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

// --- --- --- --- ---

// Обратите внимание, что массивы отображаются в форме `[v1 v2 v3 ...]`
// при печати с помощью функции `fmt.Println`.
//
// $ go run arrays.go
// emp: [0 0 0 0 0]
// set: [0 0 0 0 100]
// get: 100
// len: 5
// dcl: [1 2 3 4 5]
// 2d:  [[0 1 2] [1 2 3]]

// Вы увидите _слайсы_ гораздо чаще, чем массивы в типичном Go.
// Далее мы рассмотрим слайсы.
