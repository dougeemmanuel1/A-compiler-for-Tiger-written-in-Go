package main

import (
    "fmt"
    "os"
)

func getType(c *Context, e interface{}) interface{} {
    //Expression denoting the actual type usable within tiger
    var expressionType interface{}

    switch v := e.(type) {
    case *ArrayType:
        expressionType = c.lookup(v.id)
    case *ArrayCreate:
        expressionType = getType(c, c.lookup(v.typeId))
    case *Identifier:
        expressionType = getType(c, c.lookup(v.id))
    case *Integer:
        expressionType = v
    case *StringPrimitive:
        expressionType = v
    default:
        fmt.Fprintf(os.Stderr, "Type %T was unresolvable.\n", e)
        os.Exit(3)
    }

    fmt.Printf("Type %T %v resolved.\n", expressionType, expressionType)
    return expressionType
}

func typeToString(t interface{}) (string) {
    typeStr := ""

    if _, isInt := t.(*Integer); isInt {
        typeStr = "int"
    } else if _, isString := t.(*StringPrimitive); isString {
        typeStr = "string"
    } else if _, isVoid := t.(*VoidType); isVoid {
        typeStr = "void"
    } else if idType, isIdType := t.(*Identifier); isIdType {
        typeStr = idType.id
    } else if _, isStringPrimitive := t.(*StringPrimitive); isStringPrimitive {
        typeStr = "string"
    } else {
        fmt.Fprintf(os.Stderr, "Type %T not a valid tiger type.\n", t)
        os.Exit(3)
    }

    return typeStr
}

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

func isAssignable(c *Context, exp interface{}, expType interface{}) {
    //These two are used to check if they are assigning nil to a record type declaration
    _, expIsNilType := exp.(*Nil)
    _, typeisRecordType := expType.(*RecordType)

    //Check any other valid type combinations with these two
    fmt.Printf("Is exp:%T assignable to expType:%T (BEFORE expansion)\n", exp, expType)
    expStr := typeToString(getType(c,exp))
    expTypeStr := typeToString(getType(c, expType))
    fmt.Printf("Expstr: %s, expTypeStr:%s\n", expStr, expTypeStr)

    fmt.Println("Vaules were ", expIsNilType && typeisRecordType, expTypeStr == expStr)
    performCheck((expIsNilType && typeisRecordType) || (expTypeStr == expStr) ,
                 fmt.Sprintf("Expression of type %s not compatible with type %s.\n", expStr, expTypeStr))
}
