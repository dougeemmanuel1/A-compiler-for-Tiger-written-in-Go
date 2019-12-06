package main

import (
    "fmt"
)

func invokePrintI(ce *CallExpression) {
    fmt.Println("Invoking printi")
}

func invokePrint(ce *CallExpression) {
    fmt.Println("Invoking printi")
}

func invokeNot(ce *CallExpression) {
    fmt.Println("Invoking printi")
}
