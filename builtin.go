package main

import (
    // "fmt"
    // "os"
)

//Builtin function that prints an integer to standard out
func invokePrintI(c *Context, ce *CallExpression) {

    //Integer must be the first param since there is only 1
    // intVal := evaluateExpression(c, ce.paramNodes[0].Exp)

    // fmt.Printf("Invoking printi %T,\n", intVal)

    // fmt.Fprintf(os.Stdout, "%d\n", intVal)
}

func invokePrint(c *Context, ce *CallExpression) {
    // fmt.Println("Invoking print ")

    //String must be the first param since there is only 1
    // str := evaluateExpression(c, ce.paramNodes[0].Exp)
    // fmt.Fprintf(os.Stdout, "%s\n", str)
}

func invokeNot(c *Context, ce *CallExpression) int {
    // fmt.Println("Invoking not")
    intVal := evaluateExpression(c, ce.paramNodes[0].Exp)
    var result int
    if(intVal == 0) {
        result = 1
    } else {
        result = 0
    }
    return result
}
