package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://testovoe.kata.academy/go/step4
func main() {
	fmt.Println("Input:")
	reader := bufio.NewReader(os.Stdin)
	exprStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	expr := NewExpression(exprStr[:len(exprStr)-1])

	fmt.Println("Output")
	fmt.Println(expr.Result())
}
