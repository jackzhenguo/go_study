Today blog I will talk about a small project using GO to implement a calculator.

### Caculator of GO

Here's the coding of a simple calculator program in Go.

```
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
    num1 = num1[:len(num1)-1] // remove newline character

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

```



In this program, the user is prompted to enter two numbers and an operator (+, -, *, or /) to perform a calculation. The program then uses a `switch` statement to determine which calculation to perform and displays the result.


In this simple implementation example, we don't support the funcitions, such as more operators, handle errors and invalid input more gracefully, or create a graphical user interface (GUI) for the calculator.

Next blogs, let's add these more features to this programming.


### Challenages

I was familiar with Python, java, C++ and other languages before, but this is my first small project writing in GO language. During this period, I encountered problems such as unskilled grammar and subtle differences from other languages, but I overcame them one by one. up.

