package main

import (
    "fmt"
    "os"
)

func performCheck(cond bool, err string) {
    if(!cond) {
        fmt.Fprintf(os.Stderr, err + "\n")
        os.Exit(3)
    }
}

//NOTE: All failures halt the program.
func isArrayType(t interface{}) {
    _, ok := t.(*ArrayType)
    performCheck(ok, "Not an array type.")
}

func isAssignable(exp interface{}, expType interface{}) {
    _, expIsNilType := exp.(*Nil)
    _, typeisRecordType := expType.(*RecordType)

    fmt.Println("Vaules were ", expIsNilType, typeisRecordType)
    performCheck(expIsNilType && typeisRecordType,
                 fmt.Sprintf("Expression of type %T not compatible with type %T.\n", exp, expType))
}
