package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux if its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows if its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {

	opc := 0

	var num1, num2, result int

	for {
		fmt.Println("Select one option from the menu:")
		fmt.Println("1-Addition \n2-Subtraction \n3-Multiplication \n4-Division \n5-Exit")
		fmt.Scanf("%d\n", &opc)

		if opc == 5 {

			CallClear()
			fmt.Println("See ya!") //Nice
			break
		}

		fmt.Println("Type the first number: ")
		fmt.Scanf("%d\n", &num1)

		fmt.Println("Type the second number: ")
		fmt.Scanf("%d\n", &num2)

		if opc == 1 {

			result = num1 + num2
			CallClear()
			fmt.Printf("The result of the addition is %d\n", result)

		} else if opc == 2 {

			result = num1 - num2
			CallClear()
			fmt.Printf("The result of the subtraction is %d\n", result)

		} else if opc == 3 {

			result = num1 * num2
			CallClear()
			fmt.Printf("The result of the multiplication is %d\n", result)

		} else if opc == 4 {

			result = num1 / num2
			CallClear()
			fmt.Printf("The result of the division is %d\n", result)

		}

	}
}
