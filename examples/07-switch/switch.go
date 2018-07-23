////
// _Оператор switch_ выстро обрабатвает все "условные
// ветви".
////

package main

import "fmt"
import "time"

func main() {

	// Это основа оператора `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Вы можете использовать запятые для разделения нескольких
	// выражений в одном и том же `case`. В этом примере мы
	// также используем необязательный случай `default`.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Оператор `switch` без условия еще один пусть
	// повторить механику if/else. Здесь же, мы покажем
	// вам, как выражение `case` может быть не константным.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Тип `switch` сравнивает типы вместо значений. Вы можете
	// использовать это, чтобы узнать тип интерфейса. В этом
	// примере переменная `t` будет иметь тип, соответствующий
	// его значению при вызове функции.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// --- --- --- --- ---

// $ go run switch.go
// Write 2 as two
// It's a weekday
// It's after noon
// I'm a bool
// I'm an int
// Don't know type string
