package main

import "fmt"

func main() {

  opc := 0

  var num1, num2, result int

  for{
      fmt.Println("Select one option from the menu:")
      fmt.Println("1-Addition \n2-Subtraction \n3-Multiplication \n4-Division")
      fmt.Scanf("%d\n", &opc)

      fmt.Println("Type the first number: ")
      fmt.Scanf("%d\n", &num1)

      fmt.Println("Type the second number: ")
      fmt.Scanf("%d\n", &num2)

      if opc == 1{

          result = num1+num2
          fmt.Printf("The result of the addition is %d\n", result)

      }else if opc == 2{

          result = num1-num2
          fmt.Printf("The result of the subtraction is %d\n", result)

      }else if opc == 3{

          result = num1*num2
          fmt.Printf("The result of the multiplication is %d\n", result)

      }else if opc == 4{

          result = num1/num2
          fmt.Printf("The result of the division is %d\n", result)

    }

  }
}
