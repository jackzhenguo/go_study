package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Get the user's input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the first number: ")
    num1, _ := reader.ReadString('\n')
    num1 = num1[:len(num1)-1] // remember: remove newline character

    fmt.Print("Enter the operator (+, -, *, /): ")
    operator, _ := reader.ReadString('\n')
    operator = operator[:len(operator)-1] // remove newline character

    fmt.Print("Enter the second number: ")
    num2, _ := reader.ReadString('\n')
    num2 = num2[:len(num2)-1] // remove newline character

    // Convert the input to float64
    num1Float, _ := strconv.ParseFloat(num1, 64)
    num2Float, _ := strconv.ParseFloat(num2, 64)

    // Perform the calculation
    var result float64
    switch operator {
    case "+":
        result = num1Float + num2Float
    case "-":
        result = num1Float - num2Float
    case "*":
        result = num1Float * num2Float
    case "/":
        if num2Float == 0 {
            fmt.Println("Cannot divide by zero")
            return
        }
        result = num1Float / num2Float
    default:
        fmt.Println("Invalid operator")
        return
    }

    // Display the result
    fmt.Printf("%v %v %v = %v\n", num1Float, operator, num2Float, result)
}
