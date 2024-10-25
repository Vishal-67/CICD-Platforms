package main

import (
    "fmt"
    "os"
    "strconv"
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
    return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b float64) float64 {
    return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
    return a * b
}

// Divide returns the quotient of two numbers
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run calculator.go <add|subtract|multiply|divide> <num1> <num2>")
        return
    }

    operation := os.Args[1]
    num1, err1 := strconv.ParseFloat(os.Args[2], 64)
    num2, err2 := strconv.ParseFloat(os.Args[3], 64)

    if err1 != nil || err2 != nil {
        fmt.Println("Error: Please provide valid numbers.")
        return
    }

    switch operation {
    case "add":
        fmt.Printf("Result: %f\n", Add(num1, num2))
    case "subtract":
        fmt.Printf("Result: %f\n", Subtract(num1, num2))
    case "multiply":
        fmt.Printf("Result: %f\n", Multiply(num1, num2))
    case "divide":
        result, err := Divide(num1, num2)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("Result: %f\n", result)
        }
    default:
        fmt.Println("Error: Invalid operation. Use add, subtract, multiply, or divide.")
    }
}
