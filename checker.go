package main

import (
    "fmt"
    "os"
)

func performCheck(cond bool, err string) {
    if(cond) {
        fmt.Fprintf(os.Stderr, err + "\n")
        os.Exit(3)
    }
}

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
    case *StringLiteral:
        expressionType = v
    case *StringPrimitive:
        expressionType = v
    case *RecordCreate:
        expressionType = v
    case *RecordType:
        expressionType = v
    case *FieldDeclaration:
        expressionType = c.lookup(v.id)
    case *FieldExpression:
        //Get record type this expresion is a part of
        // so we can distinguish the field expressions type
        recordType := getType(c, v.lValue.Exp).(*RecordType)
        expressionType = recordType.getTypeOfRecordMember(c, v.id)
    case *VarDeclaration:
        if(v.typeId != "") {
            expressionType = c.lookup(v.typeId)
        } else {
            expressionType = getType(c, v.Exp.Exp)
        }
    case *FuncDeclaration:
        if(v.returnType == "") {
            expressionType = &VoidType{}
        } else {
            expressionType = c.lookup(v.returnType)
        }
    case *LetExpression:
        //If expressions has a body then take the type of the last element
        // if(len(le.exps) > 0) {
            // expressionType = le.exps[len(le.exps)-1].Exp
        // } else {
            // expressionType = &VoidType{}
        // }
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
    } else if _, isStringP := t.(*StringPrimitive); isStringP {
        typeStr = "string"
    } else if _, isStringL := t.(*StringLiteral); isStringL {
        typeStr = "string"
    }  else if _, isVoid := t.(*VoidType); isVoid {
        typeStr = "void"
    } else if idType, isIdType := t.(*Identifier); isIdType {
        typeStr = idType.id
    } else if _, isRecordType := t.(*RecordCreate); isRecordType {
        typeStr = "recType"
    } else if _, isRecordType := t.(*RecordType); isRecordType {
        typeStr = "recType"
    }  else {
        fmt.Fprintf(os.Stderr, "Type %T not a valid tiger type.\n", t)
        os.Exit(3)
    }

    return typeStr
}

//Function checks if any fields have been redclared for an id
func fieldHasNotBeenUsed (id string , declaredFields []string) {
    for _, fieldDecId := range declaredFields {
        performCheck(id == fieldDecId, fmt.Sprintf("%s already declared!\n", id))
    }
}

//NOTE: All failures halt the program.
func isArrayType(t interface{}) {
    _, ok := t.(*ArrayType)
    performCheck(!ok, "Not an array type.")
}

func isRecordType(t interface{}) {
    _, ok := t.(*RecordType)
    performCheck(!ok, fmt.Sprintf("%T is not a record type\n", t))
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
    performCheck((!expIsNilType || !typeisRecordType) && !(expTypeStr == expStr) ,
                 fmt.Sprintf("Expression of type %s not compatible with type %s.\n", expStr, expTypeStr))
    fmt.Println("Was assingable.")
}
