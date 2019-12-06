package main

import (
    "fmt"
    "os"
)

func performCheck(cond bool, err string) {
    // fmt.Printf("Cond = %t\n", cond)
    if(!cond) {
        fmt.Fprintf(os.Stderr, err + "\n")
        os.Exit(3)
    }
}

func getType(c *Context, e interface{}) interface{} {
    //Expression denoting the actual type usable within tiger
    var expressionType interface{}

    switch v := e.(type) {
    // case string:
        // expressionType = &StringLiteral{}
    case *ArrayType:
        expressionType = v.memberType
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
    case *UnitType:
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
            expressionType = &UnitType{}
        } else {
            expressionType = c.lookup(v.returnType)
        }
    case *Subscript:
        t := getType(c, c.lookup(v.getId()))
        // fmt.Printf("Returning type %T for sub\n", t)
        expressionType = t
    case *LetExpression:
        // If expressions has a body then take the type of the last element
        if(len(v.exps) > 0) {
            expressionType = v.exps[len(v.exps)-1].Exp
        } else {
            expressionType = &UnitType{}
            }
    case *SeqExpression:
        if(len(v.nodes) == 0) {
            expressionType = &UnitType{}
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
            if(getType(c, v.elseNode.Exp) != &UnitType{}) {
                expressionType = &UnitType{}
            }
        }
    default:
        fmt.Fprintf(os.Stderr, "ERROR: %d: Semantic: Type %T was unresolvable. %v \n", resolveLineNumber(e), e, e)
        os.Exit(3)
    }

    // fmt.Printf("Type %T %v resolved.\n", expressionType, expressionType)
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
    }  else if _, isVoid := t.(*UnitType); isVoid {
        typeStr = "void"
    } else if idType, isIdType := t.(*Identifier); isIdType {
        typeStr = idType.id
    } else if _, isRecordType := t.(*RecordExp); isRecordType {
        typeStr = "recType"
    } else if _, isRecordType := t.(*RecordType); isRecordType {
        typeStr = "recType"
    } else if _, isNilType := t.(*Nil); isNilType {
        typeStr = "nil"
    } else if at, isArrayType := t.(*ArrayType); isArrayType {
        typeStr = typeToString(at.memberType)
    } else {
        fmt.Fprintf(os.Stderr, "ERROR: %d: Semantic: Type %T not a valid tiger type.\n", resolveLineNumber(t), t)
        os.Exit(3)
    }

    return typeStr
}

func isInteger(c *Context, exp interface{}) {
    _, isInt := getType(c, exp).(*Integer)

    performCheck(isInt, fmt.Sprintf("ERROR: %d: Semantic: Type %T Not an integer", resolveLineNumber(exp), getType(c, exp)))
}

func isVoid(c *Context, exp interface{}) {
    _, isVoid := getType(c, exp).(*UnitType)

    performCheck(isVoid, fmt.Sprintf("ERROR: %d: Semantic: Type %T not UnitType", resolveLineNumber(exp), getType(c, exp)))
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
        case *Subscript: lineno = t.getLineno()
        case *UnitType: lineno = t.getLineno()
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
    _, isArrayType := exp.(*ArrayExp)
    performCheck(isArrayType, fmt.Sprintf("ERROR: %d: Semantic: Type %T not Array", resolveLineNumber(exp), exp))
}

func isArrayType(c *Context, exp interface{}) {
    _, isArrayType := exp.(*ArrayType)
    // fmt.Printf("Type was %T \n", exp)
    performCheck(isArrayType, fmt.Sprintf("ERROR: %d: Semantic: Type %T not ArrayType", resolveLineNumber(exp), getType(c, exp)))
}

func isIntegerOrString(c *Context, exp interface{}) {
    typeAsInt := typeToString(getType(c, exp))
    typeAsStr := typeToString(getType(c, exp))


    performCheck(typeAsInt == "int" || typeAsStr == "string", fmt.Sprintf("ERROR: %d: Semantic: Type %T Not an integer or string", resolveLineNumber(exp), getType(c, exp)))
}

//checks if the callee is an actual function type
func isFunction(c *Context, callee string) {
    possibleFunc := c.lookup(callee)
    _, isFuncDec := possibleFunc.(*FuncDeclaration)
    performCheck(isFuncDec, fmt.Sprintf("ERROR: %d: Semantic: ID: %s is not a function.", callee, possibleFunc))
}

func areLegalArguments(c *Context, decParams []Node, calleeParams []Node, callee string) {
    calleeExp := c.lookup(callee)
    //Check correct number of args to call
    performCheck(len(decParams) == len(calleeParams), fmt.Sprintf("ERROR: %d: Semantic: Expected %d args in call to %s, got %d.",  resolveLineNumber(calleeExp), len(decParams), callee, len(calleeParams)))

    //Check param types match
    for i, decNode := range decParams {
        dec, _ := decNode.Exp.(*Param)
        isAssignable(c, c.lookup(dec.fieldType), calleeParams[i].Exp)
    }
}

//Function checks if any fields have been redclared for an id
func fieldHasNotBeenUsed (id string , declaredFields []string) {
    for fd, fieldDecId := range declaredFields {
        performCheck(id != fieldDecId, fmt.Sprintf("ERROR: %d: Semantic: %s already declared!", resolveLineNumber(fd), id))
    }
}

//NOTE: All failures halt the program.
func isRecordType(t interface{}) {
    _, ok := t.(*RecordType)
    performCheck(ok, fmt.Sprintf("ERROR: %d: Semantic: %T is not a record type.", resolveLineNumber(t), t))
}

func isNotNil(t interface{}) {
    _, isNil := t.(*Nil)
    performCheck(!isNil, fmt.Sprintf("ERROR: %d: Semantic: %T cannot be nil.", resolveLineNumber(t), t))
}

func isAssignable(c *Context, exp interface{}, expType interface{}) {
    //These two are used to check if they are assigning nil to a record type declaration
    _, expIsNilType := getType(c, exp).(*Nil)
    _, typeisRecordType := getType(c, expType).(*RecordType)

    _, expIsNilType2 := getType(c, expType).(*RecordType)
    _, typeisRecordType2 := getType(c, exp).(*Nil)

    //Check any other valid type combinations with these two
    // fmt.Printf("Is exp:%T assignable to expType:%T (BEFORE expansion)\n", exp, expType)
    expStr := typeToString(getType(c,exp))
    expTypeStr := typeToString(getType(c, expType))
    // fmt.Printf("Expstr: %s, expTypeStr:%s\n", expStr, expTypeStr)

    // fmt.Printf("Is exp:%s assignable to expType:%s (POST expansion)\n", expStr, expTypeStr)
    recNilCheck := (expIsNilType && typeisRecordType || expIsNilType2 && typeisRecordType2)
    typesAreEqual := (expTypeStr == expStr)
    // fmt.Printf("RecNil:%t typesareequals:%t\n", recNilCheck, typesAreEqual)
    performCheck(recNilCheck || typesAreEqual,
                 fmt.Sprintf("ERROR: %d: Semantic: Expression of typ %s not compatible with type %s.",  resolveLineNumber(exp), expStr, expTypeStr))
    // fmt.Println("Was assingable.")
}

func expressionsHaveSameType(c *Context, e1 interface{}, e2 interface{}) {
    e1TypeAsStr := typeToString(getType(c, e1))
    e2TypeAsStr := typeToString(getType(c, e2))

    // fmt.Printf("Excpressions have %s %s.\n", e1TypeAsStr, e2TypeAsStr)
    performCheck(e1TypeAsStr == e2TypeAsStr || (e2TypeAsStr == "nil" && e1TypeAsStr == "recType") || (e2TypeAsStr == "recType" && e1TypeAsStr == "nil"),
                 fmt.Sprintf("ERROR: %d: Semantic: Expressions must have same type, %s not compitable with %s.", resolveLineNumber(e1), e1TypeAsStr, e2TypeAsStr))
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

func getIdForLValue(exp interface{}) string {
    str := ""
    switch t := exp.(type) {
    case *Identifier:
        str = t.id
    case *Subscript:
        str = t.getId()
    case *MemberExp:
        str = getIdForLValue(t.record.Exp)
    }
    return str
}

//Function theoretically should return either a string or an int or a bool maybe?
//by recursively calling it self and evaluating its expressions
func evaluateExpression(c *Context, exp interface{}) interface{} {
    var val interface{}

    switch t := exp.(type) {
    case *InfixExpression:
        fmt.Printf("Evaluating infix expresison %T l: %T, r:%T\n", t,t.leftNode.Exp, t.rightNode.Exp)
        lVal := evaluateExpression(c, t.leftNode.Exp)
        rVal :=  evaluateExpression(c, t.rightNode.Exp)
        val = evaluateInfixExpression(t.opType, lVal, rVal)
    case *Integer:
        val = t.Number
    case *Identifier:
        // fmt.Printf("Was identifier returning %v\n", c.values[t.id])
        val = evaluateExpression(c, c.values[t.id])
    case int:
        val = t
    }

    return val
}

func evaluateInfixExpression(opType Op, l interface{}, r interface{}) interface{} {
    var result interface{}
    lValue := l.(int)
    rValue := r.(int)
    switch(opType) {
    case Op_PLUS:
        result = lValue + rValue
        fmt.Printf("infix was plus l:%d r:%d %d\n", lValue, rValue, result)
    }
    return result
}
