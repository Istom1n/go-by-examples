////
// Ветвление с `if` и` else` в Go является
// довольно простым.
////

package main

import "fmt"

func main() {

	// Вот основной пример.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// У вас может быть инструкция `if` без `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Выражение может предшествовать условностям;
	// любые переменные, объявленные в этом выражениях,
	// доступны во всех ветвях.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Обратите внимание, что вам не нужны скобки вокруг
// условий в Go, но требуются фигурные скобки.

// --- --- --- --- ---

// $ go run if-else.go
// 7 is odd
// 8 is divisible by 4
// 9 has 1 digit

// В Go отсутствует тернарный оператор (https://bit.ly/1qXRzGF),
// поэтому вам нужно будет использовать полный оператор `if`
// даже для основных условий.
