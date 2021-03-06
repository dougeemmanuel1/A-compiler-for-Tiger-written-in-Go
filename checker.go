package main

import (
    "fmt"
    "os"
)

func performCheck(cond bool, err string) {
    // fmt.Printf("Cond = %t\n", cond)
    if(!cond) {
        fmt.Fprintf(os.Stderr, err)
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
    case *StringPrimitive:
        expressionType = v
    // case *StringPrimitive:
        // expressionType = v
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
    case *ForExpression:
        expressionType = &UnitType{}
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
    case *Assignment:
        expressionType = getType(c, c.lookup(getIdForLValue(v.lValue.Exp)))
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
    } else if _, isStringL := t.(*StringPrimitive); isStringL {
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
        case *StringPrimitive: lineno = t.getLineno()
        // case *StringPrimitive: lineno = t.getLineno()
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

func expressionsHaveSameType(c *Context, e1 interface{}, e2 interface{}) bool {
    e1TypeAsStr := typeToString(getType(c, e1))
    e2TypeAsStr := typeToString(getType(c, e2))

    // fmt.Printf("Excpressions have %s %s.\n", e1TypeAsStr, e2TypeAsStr)
    performCheck(e1TypeAsStr == e2TypeAsStr || (e2TypeAsStr == "nil" && e1TypeAsStr == "recType") || (e2TypeAsStr == "recType" && e1TypeAsStr == "nil"),
                 fmt.Sprintf("ERROR: %d: Semantic: Expressions must have same type, %s not compitable with %s.", resolveLineNumber(e1), e1TypeAsStr, e2TypeAsStr))

    return (e1TypeAsStr == e2TypeAsStr || (e2TypeAsStr == "nil" && e1TypeAsStr == "recType") || (e2TypeAsStr == "recType" && e1TypeAsStr == "nil"))
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

//Function theoretically should return either a string or an int or a bool maybe or void
//by recursively calling it self and evaluating its expressions
func evaluateExpression(c *Context, exp interface{}) interface{} {
    var val interface{}

    // fmt.Printf("Evaluating type %T, \n", exp)
    switch t := exp.(type) {
    case int:
        // fmt.Printf("Returning int %d\n", t)
        val = t
    case *Integer:
        val = t.Number
    case *StringPrimitive:
        val = t.str
    case string:
        val = t
    case []interface{}:
        val = t
    case *InfixExpression:
        // fmt.Printf("Evaluating infix expresison %T l: %T, r:%T\n", t,t.leftNode.Exp, t.rightNode.Exp)
        //Get the value of the left node and right
        lVal := evaluateExpression(c, t.leftNode.Exp)
        rVal :=  evaluateExpression(c, t.rightNode.Exp)


        //then apply the correct infix operator to the values
        val = evaluateInfixExpression(c, t.opType, lVal, rVal)

    case *Identifier:
        // fmt.Printf("Was identifier returning %v\n", t.id)
        if(t.id == "nil") {
            val = NewNil(0)
        } else {
            // fmt.Printf("Posssible indeitEifer %T\n", c.values[t.id])
            val = evaluateExpression(c, c.values[t.id])
        }
    case *CallExpression:
        // fmt.Printf("Evaluation call to %s\n",t.callee)

        //get declaration of call
        callDec := c.lookup(t.callee).(*FuncDeclaration)
        if(callDec.returnType == "") { //no rreturn type = unit
            val = &UnitType{}
        } else { //Evaluate the expression of the body for a value
            //This should produce a value and will execute all
            //of the code of the body as well
            for i, _ := range callDec.paramNodes {
                paramDec  := callDec.paramNodes[i].Exp.(*Param)
                paramValue := t.paramNodes[i].Exp
                c.values[paramDec.id] = evaluateExpression(c, paramValue)
                val = c.values[paramDec.id]
            }

        }

        if(t.callee == "printi") {
            invokePrintI(c, t)
        } else if(t.callee == "print") {
            invokePrint(c, t)
        } else if(t.callee == "not") {
            invokeNot(c, t)
        } else { //Regular other user defined function do your thing!
            //invoke the body
            //Get the declaration of the calling function so we can execute it
            callDec := c.lookup(t.callee).(*FuncDeclaration)
            // fmt.Printf("Invoking random func \n")
            evaluateExpression(c, callDec.body.Exp)
        }
    case *IfThenElseExpression:
        condition := evaluateExpression(c, t.condNode.Exp).(bool)
        // fmt.Printf("Cond was %v \n", condition)
        //If else is nil then its an IfThen Exp
        if(t.elseNode == nil) {
            val = &UnitType{}
            if(condition) { //if the condition is true evaluatie the code inside
                evaluateExpression(c, t.thenNode.Exp)
            }
        } else { //otherwise its and ifThenElse
            if(condition) {
                val = evaluateExpression(c, t.thenNode.Exp)
            } else {
                val = evaluateExpression(c, t.elseNode.Exp)
            }
        }
    case *SeqExpression:
        // Value is equivalent to the last node of the seqence expression
        if(len(t.nodes) == 0) {
            val = &UnitType{}
        } else {
            // fmt.Printf("Seq type was %T\n", t.nodes[len(t.nodes)-1].Exp)
            val = evaluateExpression(c, t.nodes[len(t.nodes)-1].Exp)
        }
    case *Nil:
        val = NewNil(0)
    case *ArrayExp:
        arrType := getType(c, c.lookup(t.typeId)).(*Identifier)
        val = c.lookup(arrType.id)
    case *ForExpression:
        val = &UnitType{}
    case *LetExpression:
        if(len(t.exps) == 0) {
            val = &UnitType{}
        } else {
            // fmt.Printf("%T is last exp type\n", t.exps[len(t.exps)-1].Exp)
            // val = getType(c, t.exps[len(t.exps)-1].Exp)
        }
    case *Assignment:
        val = &UnitType{}
    case *RecordExp:
        var slc []interface{}
        for _, fcNode := range t.fieldCreateNodes {
            if b, isABinding := fcNode.Exp.(*Binding); isABinding {
                slc = append(slc, evaluateExpression(c, b.exp.Exp))
            }
        }
        val = slc
    default:
        fmt.Fprintf(os.Stderr, "Could not evaluate exp %T\n", t)
        os.Exit(4)
    }

    return val
}

func evaluateInfixExpression(c *Context, opType Op, l interface{}, r interface{}) interface{} {
    var result interface{}

    //Must resolve to int for mathemtical operations
    //in other ops we dont use these two
    lValue, _ := l.(int)
    rValue, _ := r.(int)

    switch(opType) {
    case Op_PLUS:
        result = lValue + rValue
        // fmt.Printf("infix was plus l:%d r:%d %d\n", lValue, rValue, result)
    case Op_MINUS:
        result = lValue - rValue
    case Op_MUL:
        result = lValue * rValue
    case Op_DIV:
        result = lValue / rValue
    case Op_LT:
        result = lValue < rValue
    case Op_LTE:
        result = lValue <= rValue
    case Op_GT:
        result = lValue > rValue
    case Op_GTE:
        result = lValue <= rValue
    case Op_EQUALS:
        result = expressionsHaveSameType(c, l, r)
        // fmt.Printf("EQUALS WAS %v\n", result)
    case Op_NEQ:
        result = !expressionsHaveSameType(c, l, r)
        // fmt.Printf("NEQ WAS %v\n", result)
    }
    return result
}
