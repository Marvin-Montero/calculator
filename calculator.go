package main

import "fmt"

func main() {

  opc := 0

  var num1, num2, result int

  for{
      fmt.Println("Selecione una del menu:")
      fmt.Println("1-Suma \n2-Resta \n3-Multiplicar \n4-Division")
      fmt.Scanf("%d\n", &opc)

      fmt.Println("Ingrese el 1er numero: ")
      fmt.Scanf("%d\n", &num1)

      fmt.Println("SIngrese el 2do numero: ")
      fmt.Scanf("%d\n", &num2)

      if opc == 1{

          result = num1+num2
          fmt.Printf("El resultado de la suma es %d\n", result)

      }else if opc == 2{

          result = num1-num2
          fmt.Printf("El resultado de la resta es %d\n", result)

      }else if opc == 3{

          result = num1*num2
          fmt.Printf("El resultado de la multiplicacion es %d\n", result)

      }else if opc == 4{

          result = num1/num2
          fmt.Printf("El resultado de la division es %d\n", result)

    }

  }
}
