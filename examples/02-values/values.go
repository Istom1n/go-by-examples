////
// Go имеет различные типы данных, включая strings
// (строки), integers (вещественные числа), floats
// (числа с плавающей точкой), booleans, и т.д.
// Вот несколько основных примеров.
////

package main

import "fmt"

func main() {

	// Строки, которые могут быть соединены знаком `+`.
	fmt.Println("go" + "lang")

	// Вещественные числа и с плавающей точкой.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Булево значение, возвращающие булевый оператор
	// как и следовало ожидать.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// --- --- --- --- ---

// $ go run values.go
// golang
// 1+1 = 2
// 7.0/3.0 = 2.3333333333333335
// false
// true
// false
