In this blog, let's add more features to previous calculator by go, such as 

- input exception checking
- support more input operator, etc.


#### input exception support
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
    num1, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    num1 = num1[:len(num1)-1] // remember: remove newline character

    fmt.Print("Enter the operator (+, -, *, /): ")
    operator, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    operator = operator[:len(operator)-1] // remove newline character

    fmt.Print("Enter the second number: ")
    num2, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    num2 = num2[:len(num2)-1] // remove newline character

    // Convert the input to float64
    num1Float, err := strconv.ParseFloat(num1, 64)
    if err != nil {
        fmt.Println("Error converting input to number:", err)
        return
    }
    num2Float, err := strconv.ParseFloat(num2, 64)
    if err != nil {
        fmt.Println("Error converting input to number:", err)
        return
    }

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




```

(base) zhenguo@zhenguo src % go run calculator_v1.1.go
```

```
Enter the first number: d
Enter the operator (+, -, *, /): e
Enter the second number: s
Error converting input to number: strconv.ParseFloat: parsing "d": invalid syntax
```

