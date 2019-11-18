package main

import (
    "fmt"
    _ "github.com/timtadh/lexmachine"
    "os"
)

//Define operation enum
type Op int
const (
    Op_MUL Op = 0
    Op_DIV Op = 1
    Op_PLUS Op = 2
    Op_MINUS Op = 3
    Op_EQUALS Op = 4
    Op_NEQ Op = 5
    Op_GT Op = 6
    Op_LT Op = 7
    Op_GTE Op = 8
    Op_LTE Op = 9
    Op_AND Op = 10
    Op_OR Op = 11
    Op_NEG Op = 12
)

type StringPrimitive struct {}
type VoidType struct {}
type Integer struct {
    expType    interface{}
    readOnly   bool
    Number     int
}

func NewInteger(number int) *Integer {
    return &Integer{
        Number: number,
    }
}

func NewIntegerPrimitive() *Integer {
    return &Integer{}
}


func (i *Integer) isReadOnly() bool { return i.readOnly }

func (i *Integer) visit() string {
    return fmt.Sprintf("(int %d)", i.Number)
}

func (i *Integer) analyze(c *Context)  {
}

type InfixExpression struct {
    expType    interface{}
    opType Op
    leftNode Node
    rightNode  Node
}

func NewInfixExpression(opType Op, leftNode Node , rightNode Node) *InfixExpression {
    return &InfixExpression{
        opType: opType,
        leftNode: leftNode,
        rightNode: rightNode,
    }
}

func (ie *InfixExpression) isReadOnly() bool { return false }

func (ie *InfixExpression) visit() string {
    return fmt.Sprintf("(%s %v %v)", resolveOp(ie.opType), string(ie.leftNode.Exp.visit()), ie.rightNode.Exp.visit())

}

func  (ie *InfixExpression) analyze(c *Context)  {
    ie.leftNode.Exp.analyze(c)
    ie.rightNode.Exp.analyze(c)

    if(isMathematicalOperator(ie.opType)) {
        fmt.Println("Is math op\n")
        isInteger(c, ie.leftNode.Exp)
        isInteger(c, ie.rightNode.Exp)
    } else if(isComparisonOperator(ie.opType)) {
        fmt.Println("Is comp op\n")
        expressionsHaveSameType(c, ie.leftNode.Exp, ie.rightNode.Exp)
        isIntegerOrString(c, ie.leftNode.Exp)
        isIntegerOrString(c, ie.rightNode.Exp)
    } else {
        expressionsHaveSameType(c, ie.leftNode.Exp, ie.rightNode.Exp)
    }

    ie.expType = &Integer{}
}

func resolveOp(opType Op) string {
    typeStr := ""
    switch opType {
        case 0:
            typeStr = "Op_MUL"
        case 1:
            typeStr = "Op_DIV"
        case 2:
            typeStr = "Op_PLUS"
        case 3:
            typeStr = "Op_MINUS"
        case 4:
            typeStr = "Op_EQUALS"
        case 5:
            typeStr = "Op_NEQ"
        case 6:
            typeStr = "Op_GT"
        case 7:
            typeStr = "Op_LT"
        case 8:
            typeStr = "Op_GTE"
        case 9:
            typeStr = "Op_LTE"
        case 10:
            typeStr = "Op_AND"
        case 11:
            typeStr = "Op_OR"
        case 12:
            typeStr = "Op_NEG"
    }
    return typeStr
}

type Negation struct {
    expType    interface{}
    n *Node
}

func NewNegation(exp *Node) *Negation {
    return &Negation{
        n: exp,
    }
}

func (ne *Negation) isReadOnly() bool { return false }

func (ne *Negation) visit() string {
    var str string
    if(ne.n == nil) {
        fmt.Println("n is null")
        str = string(ne.n.Token.Lexeme)
    } else {
        str = fmt.Sprintf("(NEG %v)", ne.n.visit())
    }
    return str
}

func (ne *Negation) analyze(c *Context)  {
}


type SeqExpression struct {
    expType    interface{}
    nodes []Node
}

func NewSeqExpression(expressions []Node) *SeqExpression {
    // var copiedExpressionContents []Node
    // for _, e := range expressions {
        // copiedExpressionContents = append(copiedExpressionContents, *e)
    // }
    return &SeqExpression{
        nodes: expressions,
    }
}

func (se *SeqExpression) isReadOnly() bool { return false }

func (se *SeqExpression) visit() string {
    str := "(seqexp "
    for _, n := range se.nodes {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\n%v ", n.Exp.visit())
        }
    }

    str += "\n)"
    return str
}

func (se *SeqExpression) analyze(c *Context)  {
    for _, node := range se.nodes {
        node.Exp.analyze(c)
    }
}


type StringLiteral struct {
    expType    interface{}
    str string
}

func NewStringLiteral(s string) *StringLiteral {
    return &StringLiteral{
        str: s,
    }
}

func (sl *StringLiteral) isReadOnly() bool { return false }

func (sl *StringLiteral) visit() string {
    str := fmt.Sprintf("(strlit %s)", sl.str)
    return str
}

func (sl *StringLiteral) analyze(c *Context)  {
}



type Assignment struct {
    expType    interface{}
    lValue  Node
    exp     Node
}

func NewAssignment(lValue Node, exp Node) *Assignment {
    return &Assignment{
        lValue: lValue,
        exp:    exp,
    }
}

func (as *Assignment) isReadOnly() bool { return false }

func (as *Assignment) visit() string {
    return fmt.Sprintf("(assignment lValue:%v exp:%v)", as.lValue.Exp.visit(), as.exp.Exp.visit())
}

func (as *Assignment) analyze(c *Context)  {
    as.lValue.Exp.analyze(c)
    as.exp.Exp.analyze(c)
    fmt.Printf("Is exp:%T assignables to expType:%T (BEFORE expansion)\n", as.lValue.Exp, as.exp.Exp)
    isAssignable(c, as.lValue.Exp, as.exp.Exp)

    //Checking for assignment to read only variables here.
    if identifier, isIdentifier := as.lValue.Exp.(*Identifier); isIdentifier {
        if _, isVariable := c.lookup(identifier.id).(*Variable); isVariable {
            performCheck(true, fmt.Sprintf("Assignment to read only variable"))
        }
    }
}



type Nil struct {
    expType interface{}
}

func NewNil() *Nil {
    return &Nil{}
}

func (ni *Nil) isReadOnly() bool { return false }

func (ni *Nil) visit() string {
    return fmt.Sprintf("(nil)")
}

func (ni *Nil) analyze(c *Context)  {
}


type CallExpression struct {
    expType    interface{}
    callee    string
    paramNodes    []Node
}

func NewCallExpression(callee string, paramNodes []Node) *CallExpression {
    return &CallExpression{
        callee: callee,
        paramNodes: paramNodes,
    }
}

func (ce *CallExpression) isReadOnly() bool { return false }

func (ce *CallExpression) visit() string {
    str := fmt.Sprintf("(callExp: %s", ce.callee)
    for i, n := range ce.paramNodes {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\nparam %d: %v ", i+1, n.Exp.visit())
        }
    }
    str += "\n)"
    return str
}

func (ce *CallExpression) analyze(c *Context)  {
    //Check that the id they are calling a function with is actually a function
    isFunction(c, ce.callee)

    //We know this will work cause we checked above
    funcDec, _ := c.lookup(ce.callee).(*FuncDeclaration)

    areLegalArguments(c, funcDec.paramNodes, ce.paramNodes, ce.callee)

    //DO semantic analysis for the param node
    for _, paramNode := range ce.paramNodes {
        paramNode.Exp.analyze(c)
    }
}


type TypeDeclaration struct {
    expType    interface{}
    id    string
    n    Node
}

func NewTypeDeclaration(identifier string, n *Node) *TypeDeclaration {
    return &TypeDeclaration{
        id: identifier,
        n: *n,
    }
}

func (td *TypeDeclaration) isReadOnly() bool { return false }

func (td *TypeDeclaration) visit() string {
    return fmt.Sprintf("(tyDec: type:%s %s)", td.id, td.n.visit())
}

func (td *TypeDeclaration) analyze(c *Context)  {
    if rt, isRecordType := td.n.Exp.(*RecordType); isRecordType {
        rt.typeId = td.id
    }

    td.n.Exp.analyze(c)
}


type FuncDeclaration struct {
    expType            interface{}
    id                 string
    returnType         string
    body                Node
    paramNodes   []Node
    bodyContext        *Context

}

func NewFuncDeclaration(id string, params []Node, returnType string, body Node ) *FuncDeclaration {
    return &FuncDeclaration{
        id: id,
        paramNodes: params,
        returnType: returnType,
        body: body,
    }
}

func (fd *FuncDeclaration) isReadOnly() bool { return false }

func (fd *FuncDeclaration) visit() string {
    str := fmt.Sprintf("(funDec: id:%s returnType:%s paramNodes:", fd.id, fd.returnType)
    for _, n := range fd.paramNodes {
        str += fmt.Sprintf("(%v)\n", n.Exp.visit())
    }
    str += fmt.Sprintf("body:%s)", fd.body.Exp.visit())
    return str
}

func (fd *FuncDeclaration) analyze(c *Context)  {
    fd.body.Exp.analyze(fd.bodyContext)


    isAssignable(c, fd.body.Exp, getType(c, fd))

    //Remove body context, useless now
    fd.bodyContext = nil
}

func (fd *FuncDeclaration) analyzeSignature(c *Context) {
    fd.bodyContext = c.createChildContextForFunctionBody()

    for _, param := range fd.paramNodes {
        param.Exp.analyze(fd.bodyContext)
    }
}

type Param struct {
    expType    interface{}
    id         string
    fieldType  string
}

func NewParam(identifier1 string, fieldType string) *Param {
    return &Param{
        id: identifier1,
        fieldType: fieldType,
    }
}

func (p *Param) isReadOnly() bool { return false }

func (p *Param) visit() string {
    return fmt.Sprintf("fieldDec: (id:%s) (fieldType:%s)", p.id, p.fieldType)
}

func (p *Param) analyze(c *Context)  {
    // c.add(p.id, p)
}


type MemberExp struct {
    expType    interface{}
    id        string
    record    Node
}

func NewMemberExp(record Node, id string) *MemberExp {
    return &MemberExp{
        record: record,
        id: id,
    }
}

func (me *MemberExp) isReadOnly() bool { return false }

func (me *MemberExp) visit() string {
    return fmt.Sprintf("(fieldExp: (record:%v) (id:%s))", me.record.Exp.visit(), me.id)
}

func (me *MemberExp) analyze(c *Context)  {
    me.record.Exp.analyze(c)
    isRecordType(getType(c, me.record.Exp))

    //Check if this record type even has a member named <id>
    //+ Dont have to type check since we confirmed its a record type earlier

    rt := getType(c, me.record.Exp).(*RecordType)
    if(!rt.definesId(me.id)) {
        fmt.Fprintf(os.Stderr, "Record does not define %s.\n", me.id)
        os.Exit(3)
    }
}


type Binding struct {
    expType    interface{}
    id    string
    exp   Node
}

func NewBinding(identifier string, exp Node) *Binding {
    return &Binding{
        id: identifier,
        exp: exp,
    }
}

func (fc *Binding) isReadOnly() bool { return false }

func (fc *Binding) visit() string {
    return fmt.Sprintf("fieldCreate: id:%s exp:(%v)", fc.id, fc.exp.Exp.visit())
}

func (fc *Binding) analyze(c *Context)  {
}

type Variable struct {
    id          string
    expType     interface{}
    readOnly    bool
}

func (v *Variable) isReadOnly() bool {
    fmt.Printf("Returning %v \n", v.readOnly)
    return v.readOnly
}

type VarDeclaration struct {
    expType    interface{}
    id         string
    typeId     string
    Exp        Node
}

func NewVarDeclaration(identifier1 string, typeId string, n *Node) *VarDeclaration {
    return &VarDeclaration{
        id: identifier1,
        typeId: typeId,
        Exp: *n,
    }
}

func (vd *VarDeclaration) isReadOnly() bool { return false }

func (vd *VarDeclaration) visit() string {
    return fmt.Sprintf("(varDec: id:%s typeId:%s exp:%s)", vd.id, vd.typeId, vd.Exp.visit())
}

func (vd *VarDeclaration) analyze(c *Context)  {
    vd.Exp.analyze(c)
    if(vd.typeId != "") {//If type id is declared then we know the type from a lookup!
        vd.expType = c.lookup(vd.typeId)

        fmt.Printf("Type in lookup was %T\n", vd.expType)

        //Check assignable to ?
        isAssignable(c, vd.Exp.Exp, vd.expType)

        //When type is declared for record, it must match type declared in reccreate
        if rc, isRecordExp := vd.Exp.Exp.(*RecordExp); isRecordExp {
            if(rc.id != vd.typeId) {
                fmt.Fprintf(os.Stderr, fmt.Sprintf("Record type %s not compatible with %s.\n", rc.id, vd.typeId))
                os.Exit(3)
            }
        }

        //When type is declared for array, it must match type declared in ArrayExp
        // if ac, isArrayExp := vd.Exp.Exp.(*ArrayExp); isArrayExp {
        //     if(ac.typeId != vd.typeId) {
        //         fmt.Fprintf(os.Stderr, fmt.Sprintf("Array type %s not compatible with %s.\n", ac.typeId, vd.typeId))
        //         os.Exit(3)
        //     }
        // }
    } else { // Inference type from init experssion:O
        vd.expType = vd.Exp.Exp
    }

    //add type to context
    c.add(vd.id, vd)
}


type Identifier struct {
    expType    interface{}
    id         string
    readOnly   bool
}

func NewIdentifier(identifier string) *Identifier {
    return &Identifier{
        id: identifier,
    }
}

func (id *Identifier) isReadOnly() bool { return id.readOnly }

func (id *Identifier) visit() string {
    return fmt.Sprintf("(ID: %s)", id.id)
}


func (id *Identifier) analyze(c *Context)  {
}


type Subscript struct {
    expType    interface{}
    id          string
    expId        *Node
    subscriptExp Node
}

func NewSubscriptExpression(id string, expId *Node, subscriptExp Node) *Subscript {
    return &Subscript{
        id: id,
        expId: expId,
        subscriptExp: subscriptExp,
    }
}

func (se *Subscript) isReadOnly() bool { return false }

func (se *Subscript) visit() string {
    var str string
    if(se.id != "") {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", se.id, se.subscriptExp.Exp.visit())
    } else {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", se.expId.Exp.visit(), se.subscriptExp.Exp.visit())
    }
    return str
}

func (se *Subscript) analyze(c *Context)  {
    var arr interface{}
    if(se.id == "") { //name is defined by expId then
        if id, isId := se.expId.Exp.(*Identifier); isId {
            arr = c.lookup(id.id)
        }

        // sub, isSubscript := se.expId.(*Subscript); isSubscript {
        //     arr = c.lookup(id.id)
        // }

        if field, isFieldExp := se.expId.Exp.(*MemberExp); isFieldExp {
            arr = c.lookup(field.id)
        }
    } else {
        arr = c.lookup(se.id)
    }
    isArray(c, arr)
    se.subscriptExp.analyze(c)
    isInteger(c, se.subscriptExp.Exp)
}


type RecordType struct {
    expType          interface{}
    typeId           string
    fieldDecNodes    []Node
}

func NewRecordType(fieldDecNodes []Node) *RecordType {
    return &RecordType{
        fieldDecNodes: fieldDecNodes,
    }
}

func (rt *RecordType) nodesAsFieldDecs() []*Param {
    var interfaceList []interface{}
    var fdList        []*Param

    for _, fdNode := range rt.fieldDecNodes {
        interfaceList = append(interfaceList, fdNode.Exp)
    }

    //Cast all to field decs and return that
    for _, fdInterface := range interfaceList {
        fd := fdInterface.(*Param)
        fdList = append(fdList, fd)
    }
    return fdList
}

func (rt *RecordType) definesId(id string) bool {
    doesDeclareId := false
    for _, fd := range rt.nodesAsFieldDecs() {
        if(fd.id == id) {
            doesDeclareId = true
            continue
        }
    }
    return doesDeclareId
}

func (rt *RecordType) getTypeOfRecordMember(c *Context, id string) interface{} {
    var t interface{}
    for _, fd := range rt.nodesAsFieldDecs() {
        if(fd.id == id) {
            t = c.lookup(fd.fieldType)
            continue
        }
    }
    return t
}

func (rt *RecordType) isReadOnly() bool { return false }

func (rt *RecordType) visit() string {
    str := fmt.Sprintf("(recTy: fieldDecNodes:(")
    for _, n := range rt.fieldDecNodes {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")"
    return str
}

func (rt *RecordType) analyze(c *Context)  {
    //Keep track of fields declared so we can err, when
    //they try to redeclare a field
    var declaredFields []string

    for _, fieldDec := range rt.nodesAsFieldDecs() {
        //get field list as list of interfaces so we can type assert and access id
        //if it hasnt been used add it to the list of used identifiers for fields
        fieldHasNotBeenUsed(fieldDec.id, declaredFields)

        declaredFields = append(declaredFields, fieldDec.id)

        fieldDec.analyze(c)
    }
}


type RecordExp struct {
    expType    interface{}
    id string
    fieldCreateNodes    []Node
}

func NewRecordExp(id string, fieldCreateNodes []Node) *RecordExp {
    return &RecordExp{
        id: id,
        fieldCreateNodes: fieldCreateNodes,
    }
}

func (rc *RecordExp) isReadOnly() bool { return false }

func (rc *RecordExp) visit() string {
    str := fmt.Sprintf("(recCreate: id:%s fieldCreateNodes:(", rc.id)
    for _, n := range rc.fieldCreateNodes {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")\n"
    return str
}

func (rc *RecordExp) analyze(c *Context)  {
    recCreate := c.lookup(rc.id)
    isRecordType(recCreate)
    // for _, fieldCreate := range rc.fieldCreateNodes {

    // }
}


type ArrayType struct {
    expType    interface{}
    id    string
}

func NewArrayType(identifier string) *ArrayType {
    return &ArrayType{
        id: identifier,
    }
}

func (at *ArrayType) isReadOnly() bool { return false }

func (at *ArrayType) visit() string {
    return fmt.Sprintf("(arrType: %s)", at.id)
}


func (at *ArrayType) analyze(c *Context)  {
     at.expType = c.lookup(at.id)
     fmt.Printf("Assigning type %T to arraytype\n", at.expType)
}


type ArrayExp struct {
    expType  interface{}
    typeId    string
    subscriptNode  Node
    expNode  Node
}

func NewArrayExp(typeIdentifier string, subscriptNode Node, expNode Node) *ArrayExp {
    return &ArrayExp{
        typeId: typeIdentifier,
        subscriptNode: subscriptNode,
        expNode: expNode,
    }
}

func (ae *ArrayExp) isReadOnly() bool { return false }

func (ae *ArrayExp) visit() string {
    return fmt.Sprintf("(arrCreate: typeId:%s subscriptNode:%v expNode:%v)", ae.typeId, ae.subscriptNode.visit(), ae.expNode.visit())
}

func (ae *ArrayExp) analyze(c *Context)  {
    arr := c.lookup(ae.typeId)
    fmt.Printf("IN AC type iss %T\n",arr)
    isArrayType(c, arr)
    ae.subscriptNode.Exp.analyze(c)
    isInteger(c, ae.subscriptNode.Exp)

}

type LetExpression struct {
    expType    interface{}
    declarationNodes    []Node
    exps    []Node
}

func NewLetExpression(declarations []Node, expressions []Node) *LetExpression {
    return &LetExpression{
        declarationNodes: declarations,
        exps: expressions,
    }
}

func (le *LetExpression) isReadOnly() bool { return false }

func (le *LetExpression) visit() string {
    str := fmt.Sprintf("(letExp: declarationNodes:(")
    for _, n := range le.declarationNodes {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }
    str += fmt.Sprintf(")\n(exps: ")
    for _, n := range le.exps {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }

    str += "))"
    return str
}

func (le *LetExpression) analyze(c *Context)  {
    newContext := c.createChildContextForBlock()
    for _, d := range le.declarationNodes {
        td, isTypeDec := d.Exp.(*TypeDeclaration)
        if(isTypeDec) { //If its a type declaration, add it to the new context
            newContext.add(td.id, td.n.Exp)
        }
    }

    for _, d := range le.declarationNodes {
        fd, isFuncDec := d.Exp.(*FuncDeclaration)
        if(isFuncDec) { //If its a fimc declaration, add it to the new context
            fmt.Println("IS A FUNC")
            fd.analyzeSignature(newContext)
        }
    }

    for _, d := range le.declarationNodes {
        fd, isFuncDec := d.Exp.(*FuncDeclaration)
        if(isFuncDec) { //If its a fimc declaration, add it to the new context
            newContext.add(fd.id, fd)
        }
    }

    for _, d := range le.declarationNodes {
        d.Exp.analyze(newContext)
    }

    //Check for no recursive type cycles with out record types in decs..
    noRecursiveTypeCyclesWithoutRecordTypes()
    for _, d := range le.exps {
        d.Exp.analyze(newContext)
    }
}



type IfThenElseExpression struct {
    expType    interface{}
    condNode  Node
    thenNode  Node
    elseNode  *Node
}

func NewIfThenElseExpression(condNode Node, thenNode Node, elseNode *Node) *IfThenElseExpression {
    return &IfThenElseExpression{
        condNode: condNode,
        thenNode: thenNode,
        elseNode: elseNode,
    }
}

func (itee *IfThenElseExpression) isReadOnly() bool { return false }

func (itee *IfThenElseExpression) visit() string {
    str := ""
    if(itee.elseNode == nil) {
        str = fmt.Sprintf("(ifThenElse if:%v then:%v)", itee.condNode.Exp.visit(),
                                                        itee.thenNode.Exp.visit())
    } else {
        str = fmt.Sprintf("(ifThenElse if:%v then:%v else:%v)", itee.condNode.Exp.visit(),
                                                                itee.thenNode.Exp.visit(),
                                                                itee.elseNode.Exp.visit())
    }
    return str
}

func (itee *IfThenElseExpression) analyze(c *Context)  {
    itee.condNode.Exp.analyze(c)
    isInteger(c, itee.condNode.Exp)
    itee.thenNode.Exp.analyze(c)

    if(itee.elseNode != nil) { // else
        itee.elseNode.analyze(c)
        if(getType(c, itee.elseNode.Exp) != &VoidType{}) {
            expressionsHaveSameType(c, itee.thenNode.Exp, itee.elseNode.Exp)
        } else {
            isVoid(c, itee.elseNode.Exp)
        }
    } else {
        isVoid(c, itee.thenNode.Exp)
    }
}


type WhileExpression struct {
    expType    interface{}
    cond       Node
    body       Node
}

func NewWhileExpression(cond Node, body Node) *WhileExpression {
    return &WhileExpression{
        cond:cond,
        body:body,
    }
}

func (we *WhileExpression) isReadOnly() bool { return false }

func (we *WhileExpression) visit() string {
    return fmt.Sprintf("(whileExp cond:%v do:%v)", we.cond.Exp.visit(), we.body.Exp.visit())
}


func (we *WhileExpression) analyze(c *Context)  {
    we.cond.Exp.analyze(c)
    isInteger(c, we.cond.Exp)
    we.body.Exp.analyze(c.createChildContextForLoop())
    isVoid(c, we.body.Exp)
}


type ForExpression struct {
    expType    interface{}
    id    string
    low   Node
    high  Node
    body  Node
}

func NewForExpression(id string, low Node, high Node, body Node) *ForExpression {
    return &ForExpression{
        id:id,
        low:low,
        high:high,
        body:body,
    }
}

func (fe *ForExpression) isReadOnly() bool { return false }

func (fe *ForExpression) visit() string {
    return fmt.Sprintf("(forexp id:%s low:%v high:%v body:%v)", fe.id, fe.low.Exp.visit(), fe.high.Exp.visit(), fe.body.Exp.visit())
}

func (fe *ForExpression) analyze(c *Context)  {
    fe.low.Exp.analyze(c)
    isInteger(c, fe.low.Exp)
    fe.high.Exp.analyze(c)
    isInteger(c, fe.high.Exp)
    bodyContext := c.createChildContextForLoop()

    index := &Variable{
        id:         fe.id,
        expType:    fe.low.Exp,
        readOnly:   true,
    }

    fmt.Printf("index is %v \n", index)
    bodyContext.add(fe.id, index)
    // bodyContext.lookup(fe.id).(*Identifier).readOnly = true
    fe.body.analyze(bodyContext)

}
