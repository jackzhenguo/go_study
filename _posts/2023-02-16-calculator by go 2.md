In this blog, let's add one of more features to previous calculator by go, such as exception handling.


#### support input exception 

To support handling of invalid user input, we can add error handling to the existing code. 

Here's an updated version of the program with error handling:

case 1: 
It will prompts us `nil` exception if we don't input valid data

```
num1, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
```

case 2：

It will prompts us error convering if we perform parseFloat convertion.

```    
    num1Float, err := strconv.ParseFloat(num1, 64)
    if err != nil {
        fmt.Println("Error converting input to number:", err)
        return
    }

```

Put it all together, we will get the v1.1 of calculator after we add exception handling:

[calculator_v1.1.go](https://github.com/jackzhenguo/go_study/blob/my-pages/src/calculator_v1.1.go)

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


#### let's run it after adding exception handling

```

(base) zhenguo@zhenguo src % go run calculator_v1.1.go
```

```
Enter the first number: d
Enter the operator (+, -, *, /): e
Enter the second number: s
Error converting input to number: strconv.ParseFloat: parsing "d": invalid syntax
```

Yes. It shows `Error converting input to number`, and that's the exact exception handling we add in v1.1 of code.

```
num1Float, err := strconv.ParseFloat(num1, 64)
    if err != nil {
        fmt.Println("Error converting input to number:", err)
        return
    }
```

#### Challenges

Adding exceptions in Go can be challenging compared to other programming languages that have built-in support for exceptions, such as Java or Python. Instead, Go relies on error handling through the use of return values, which can make the code more verbose.

Since Go relies on returning errors as values, it can lead to cluttered code with many if statements to check for errors. This can make the code harder to read and maintain.

Sometimes, I forget to handle an error or ignore an error when it is returned, which can cause bugs or unexpected behavior.






