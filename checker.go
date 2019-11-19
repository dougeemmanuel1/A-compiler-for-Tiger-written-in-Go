package main

import (
    "fmt"
    "os"
)

func performCheck(cond bool, err string) {
    fmt.Printf("Cond = %t\n", cond)
    if(!cond) {
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
    case *ArrayExp:
        expressionType = getType(c, c.lookup(v.typeId))
    case *Identifier:
        expressionType = getType(c, c.lookup(v.id))
    case *Integer:
        expressionType = v
    case *StringLiteral:
        expressionType = v
    case *StringPrimitive:
        expressionType = v
    case *Nil:
        expressionType = &Nil{}
    case *VoidType:
        expressionType = v
    case *RecordExp:
        expressionType = getType(c, c.lookup(v.id))
    case *RecordType:
        expressionType = v
    case *Param:
        expressionType = c.lookup(v.fieldType)
    case *MemberExp:
        //Get record type this expresion is a part of
        // so we can distinguish the field expressions type
        recordType := getType(c, v.record.Exp).(*RecordType)
        expressionType = recordType.getTypeOfRecordMember(c, v.id)
    case *Variable:
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
    case *SeqExpression:
        if(len(v.nodes) == 0) {
            expressionType = &VoidType{}
        } else {
            expressionType = getType(c, v.nodes[len(v.nodes)-1].Exp)
        }
    case *CallExpression:
        expressionType = getType(c, c.lookup(v.callee))
    case *InfixExpression:
        expressionType = &Integer{}
    case *IfThenElseExpression:
        isInteger(c, v.condNode.Exp) //cond typem ust be int
        if(v.elseNode != nil) { // if then else
            expressionsHaveSameType(c, v.thenNode.Exp, v.elseNode.Exp)
            expressionType = v.thenNode.Exp
        } else { //if then
            if(getType(c, v.elseNode.Exp) != &VoidType{}) {
                expressionType = &VoidType{}
            }
        }
    default:
        fmt.Fprintf(os.Stderr, "ERROR: <%d> : Semantic: Type %T was unresolvable.\n", resolveLineNumber(e), e)
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
    } else if _, isRecordType := t.(*RecordExp); isRecordType {
        typeStr = "recType"
    } else if _, isRecordType := t.(*RecordType); isRecordType {
        typeStr = "recType"
    } else if _, isNilType := t.(*Nil); isNilType {
        typeStr = "nil"
    }  else {
        fmt.Fprintf(os.Stderr, "ERROR: <%d> : Semantic: Type %T not a valid tiger type.\n", resolveLineNumber(t), t)
        os.Exit(3)
    }

    return typeStr
}

func isInteger(c *Context, exp interface{}) {
    _, isInt := getType(c, exp).(*Integer)

    performCheck(isInt, fmt.Sprintf("ERROR: <%d> : Semantic: Type %T Not an integer", resolveLineNumber(exp), getType(c, exp)))
}

func isVoid(c *Context, exp interface{}) {
    _, isVoid := getType(c, exp).(*VoidType)

    performCheck(isVoid, fmt.Sprintf("ERROR: <%d> : Semantic: Type %T not VoidType", resolveLineNumber(exp), getType(c, exp)))
}

func resolveLineNumber(exp interface{}) int {
    var lineno int
    switch t := exp.(type) {
        case *ArrayType: lineno = t.getLineno()
        case *ArrayExp: lineno = t.getLineno()
        case *Identifier: lineno = t.getLineno()
        case *Integer: lineno = t.getLineno()
        case *StringLiteral: lineno = t.getLineno()
        case *StringPrimitive: lineno = t.getLineno()
        case *Nil: lineno = t.getLineno()
        case *VoidType: lineno = t.getLineno()
        case *RecordExp: lineno = t.getLineno()
        case *RecordType: lineno = t.getLineno()
        case *Param: lineno = t.getLineno()
        case *MemberExp: lineno = t.getLineno()
        case *Variable: lineno = t.getLineno()
        case *FuncDeclaration: lineno = t.getLineno()
        case *LetExpression: lineno = t.getLineno()
        case *SeqExpression: lineno = t.getLineno()
        case *CallExpression: lineno = t.getLineno()
        case *InfixExpression: lineno = t.getLineno()
        case *IfThenElseExpression: lineno = t.getLineno()

    }
    return lineno
}
func isArray(c *Context, exp interface{}) {
    _, isArrayType := getType(c, exp).(*ArrayType)
    performCheck(isArrayType, fmt.Sprintf("ERROR: <%d> : Semantic: Type %T not Array", resolveLineNumber(exp), getType(c, exp)))
}

func isArrayType(c *Context, exp interface{}) {
    _, isArrayType := exp.(*ArrayType)
    fmt.Printf("Type was %T \n", exp)
    performCheck(isArrayType, fmt.Sprintf("ERROR: <%d> : Semantic: Type %T not ArrayType", resolveLineNumber(exp), getType(c, exp)))
}

func isIntegerOrString(c *Context, exp interface{}) {
    typeAsInt := typeToString(getType(c, exp))
    typeAsStr := typeToString(getType(c, exp))


    performCheck(typeAsInt == "int" || typeAsStr == "string", fmt.Sprintf("ERROR: <%d> : Semantic: Type %T Not an integer or string", resolveLineNumber(exp), getType(c, exp)))
}

//checks if the callee is an actual function type
func isFunction(c *Context, callee string) {
    possibleFunc := c.lookup(callee)
    _, isFuncDec := possibleFunc.(*FuncDeclaration)
    performCheck(isFuncDec, fmt.Sprintf("ID: %s is not a function.", callee, possibleFunc))
}

func areLegalArguments(c *Context, decParams []Node, calleeParams []Node, callee string) {
    calleeExp := c.lookup(callee)
    //Check correct number of args to call
    performCheck(len(decParams) != len(calleeParams), fmt.Sprintf("ERROR: <%d> : Semantic: Expected %d args in call to %s, got %d.",  resolveLineNumber(calleeExp), len(decParams), callee, len(calleeParams)))

    //Check param types match
    for i, decNode := range decParams {
        dec, _ := decNode.Exp.(*Param)
        isAssignable(c, c.lookup(dec.fieldType), calleeParams[i].Exp)
    }
}

//Function checks if any fields have been redclared for an id
func fieldHasNotBeenUsed (id string , declaredFields []string) {
    for fd, fieldDecId := range declaredFields {
        performCheck(id != fieldDecId, fmt.Sprintf("ERROR: <%d> : Semantic : %s already declared!", resolveLineNumber(fd), id))
    }
}

//NOTE: All failures halt the program.
func isRecordType(t interface{}) {
    _, ok := t.(*RecordType)
    performCheck(ok, fmt.Sprintf("ERROR: <%d>: Semantic: %T is not a record type\n", resolveLineNumber(t), t))
}

func isNotNil(t interface{}) {
    _, isNil := t.(*Nil)
    performCheck(!isNil, fmt.Sprintf("ERROR: <%d>: Semantic: %T cannot be nil.", resolveLineNumber(t), t))
}

func isAssignable(c *Context, exp interface{}, expType interface{}) {
    //These two are used to check if they are assigning nil to a record type declaration
    _, expIsNilType := getType(c, exp).(*Nil)
    _, typeisRecordType := getType(c, expType).(*RecordType)

    _, expIsNilType2 := getType(c, expType).(*RecordType)
    _, typeisRecordType2 := getType(c, exp).(*Nil)

    //Check any other valid type combinations with these two
    fmt.Printf("Is exp:%T assignable to expType:%T (BEFORE expansion)\n", exp, expType)
    expStr := typeToString(getType(c,exp))
    expTypeStr := typeToString(getType(c, expType))
    fmt.Printf("Expstr: %s, expTypeStr:%s\n", expStr, expTypeStr)

    fmt.Printf("Is exp:%s assignable to expType:%s (POST expansion)\n", expStr, expTypeStr)
    recNilCheck := (expIsNilType && typeisRecordType || expIsNilType2 && typeisRecordType2)
    typesAreEqual := (expTypeStr == expStr)
    fmt.Printf("RecNil:%t typesareequals:%t\n", recNilCheck, typesAreEqual)
    performCheck(recNilCheck || typesAreEqual,
                 fmt.Sprintf("ERROR: <%d>: Semantic: Expression of typ %s not compatible with type %s.", resolveLineNumber(exp), expStr, expTypeStr))
    fmt.Println("Was assingable.")
}

func expressionsHaveSameType(c *Context, e1 interface{}, e2 interface{}) {
    e1TypeAsStr := typeToString(getType(c, e1))
    e2TypeAsStr := typeToString(getType(c, e2))

    fmt.Printf("Excpressions have %s %s.\n", e1TypeAsStr, e2TypeAsStr)
    performCheck(e1TypeAsStr == e2TypeAsStr || (e2TypeAsStr == "nil" && e1TypeAsStr == "recType") || (e2TypeAsStr == "recType" && e1TypeAsStr == "nil"),
                 fmt.Sprintf("Expressions must have same type, %s not compitable with %s.", e1TypeAsStr, e2TypeAsStr))
}

func isMathematicalOperator(op Op) bool {
    isMathOp := false
    switch(op) {
    case Op_MINUS:
        isMathOp = true
    case Op_PLUS:
        isMathOp = true
    case Op_MUL:
        isMathOp = true
    case Op_DIV:
        isMathOp = true
    case Op_AND:
        isMathOp = true
    case Op_OR:
        isMathOp = true
    }
    return isMathOp
}

func isComparisonOperator(op Op) bool {
    isCompOp := false
    switch(op) {
    case Op_LTE:
        isCompOp = true
    case Op_LT:
        isCompOp = true
    case Op_GTE:
        isCompOp = true
    case Op_GT:
        isCompOp = true
    }
    return isCompOp
}

func noRecursiveTypeCyclesWithoutRecordTypes() {

}
