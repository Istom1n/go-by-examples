////
// Наша первая программа выведет классическое
// сообщение "hello world". Ниже представлен исходный код.
////

package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

// --- --- --- --- ---

// Чтобы запустить программу, введите код в `hello-world.go` и
// используйте команду `go run`.
//
// $ go run hello-world.go
// hello world

// Иногда мы хотим собрать нашу программу в
// виде двоичного кода. Мы можем сделать это
// с помощью команды `go build`.
//
// $ go build hello-world.go
// $ ls
// hello-world	hello-world.go

// Также, мы может запустить наш собранный код,
// сохраненного в двоичном виде.
//
// $ ./hello-world
// hello world

// Также, мы может запустить наш собранный код,
// сохраненного в двоичном виде.