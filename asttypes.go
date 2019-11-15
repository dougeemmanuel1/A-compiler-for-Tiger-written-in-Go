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
    Number int
}

func NewInteger(number int) *Integer {
    return &Integer{
        Number: number,
    }
}

func NewIntegerPrimitive() *Integer {
    return &Integer{}
}


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

func (ie *InfixExpression) visit() string {
    return fmt.Sprintf("(%s %v %v)", resolveOp(ie.opType), string(ie.leftNode.Exp.visit()), ie.rightNode.Exp.visit())

}

func  (ie *InfixExpression) analyze(c *Context)  {
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

func (as *Assignment) visit() string {
    return fmt.Sprintf("(assignment lValue:%v exp:%v)", as.lValue.Exp.visit(), as.exp.Exp.visit())
}

func (as *Assignment) analyze(c *Context)  {
    as.lValue.Exp.analyze(c)
    as.exp.Exp.analyze(c)
    fmt.Printf("Is exp:%T assignables to expType:%T (BEFORE expansion)\n", as.lValue.Exp, as.exp.Exp)
    isAssignable(c, as.lValue.Exp, as.exp.Exp)
}



type Nil struct {
    expType interface{}
}

func NewNil() *Nil {
    return &Nil{}
}

func (ni *Nil) visit() string {
    return fmt.Sprintf("(nil)")
}

func (ni *Nil) analyze(c *Context)  {
}


type CallExpression struct {
    expType    interface{}
    name    string
    paramNodes    []Node
}

func NewCallExpression(name string, paramNodes []Node) *CallExpression {
    return &CallExpression{
        name: name,
        paramNodes: paramNodes,
    }
}

func (ce *CallExpression) visit() string {
    str := fmt.Sprintf("(callExp: %s", ce.name)
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

func (td *TypeDeclaration) visit() string {
    return fmt.Sprintf("(tyDec: type:%s %s)", td.id, td.n.visit())
}

func (td *TypeDeclaration) analyze(c *Context)  {
    td.n.Exp.analyze(c)
}


type FuncDeclaration struct {
    expType            interface{}
    id                 string
    returnType         string
    body                Node
    bodyContext        *Context
    declarationNodes   []Node

}

func NewFuncDeclaration(id string, returnType string, declarations []Node, n Node) *FuncDeclaration {
    return &FuncDeclaration{
        id: id,
        returnType: returnType,
        declarationNodes: declarations,
        body: n,
    }
}

func (fd *FuncDeclaration) visit() string {
    str := fmt.Sprintf("(funDec: id:%s returnType:%s declarationNodes:", fd.id, fd.returnType)
    for _, n := range fd.declarationNodes {
        str += fmt.Sprintf("(%v)\n", n.Exp.visit())
    }
    str += fmt.Sprintf("body:%s)", fd.body.Exp.visit())
    return str
}

func (fd *FuncDeclaration) analyze(c *Context)  {
    for _, decNode := range fd.declarationNodes {
        decNode.Exp.analyze(c)
    }

    isAssignable(c, fd.body.Exp, getType(c, fd))
}

func (fd *FuncDeclaration) analyzeSignature(c *Context) {
    fd.bodyContext = c.createChildContextForFunctionBody(fd)

    for _, decNode := range fd.declarationNodes {
        decNode.Exp.analyze(fd.bodyContext)
    }

    //Remove body context, useless now
    fd.bodyContext = nil
}

type FieldDeclaration struct {
    expType    interface{}
    id    string
    fieldType   string
}

func NewFieldDeclaration(identifier1 string, fieldType string) *FieldDeclaration {
    return &FieldDeclaration{
        id: identifier1,
        fieldType: fieldType,
    }
}

func (fid *FieldDeclaration) visit() string {
    return fmt.Sprintf("fieldDec: (id:%s) (fieldType:%s)", fid.id, fid.fieldType)
}

func (fid *FieldDeclaration) analyze(c *Context)  {
    c.add(fid.id, fid)
}


type FieldExpression struct {
    expType    interface{}
    id        string
    lValue    Node
}

func NewFieldExpression(lValue Node, id string) *FieldExpression {
    return &FieldExpression{
        lValue: lValue,
        id: id,
    }
}

func (fe *FieldExpression) visit() string {
    return fmt.Sprintf("(fieldExp: (lValue:%v) (id:%s))", fe.lValue.Exp.visit(), fe.id)
}

func (fe *FieldExpression) analyze(c *Context)  {
    fe.lValue.Exp.analyze(c)
    isRecordType(getType(c, fe.lValue.Exp))

    //Check if this record type even has a member named <id>
    //+ Dont have to type check since we confirmed its a record type earlier

    rt := getType(c, fe.lValue.Exp).(*RecordType)
    if(!rt.definesId(fe.id)) {
        fmt.Fprintf(os.Stderr, "Record does not define %s.\n", fe.id)
        os.Exit(3)
    }

}


type FieldCreate struct {
    expType    interface{}
    id    string
    exp   Node
}

func NewFieldCreate(identifier string, exp Node) *FieldCreate {
    return &FieldCreate{
        id: identifier,
        exp: exp,
    }
}

func (fc *FieldCreate) visit() string {
    return fmt.Sprintf("fieldCreate: id:%s exp:(%v)", fc.id, fc.exp.Exp.visit())
}

func (fc *FieldCreate) analyze(c *Context)  {
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
    } else { // Inference type from init experssion:O
        vd.expType = vd.Exp.Exp
    }

    //add type to context
    c.add(vd.id, vd)
}


type Identifier struct {
    expType    interface{}
    id    string
}

func NewIdentifier(identifier string) *Identifier {
    return &Identifier{
        id: identifier,
    }
}

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
}


type RecordType struct {
    expType    interface{}
    fieldDecNodes    []Node
}

func NewRecordType(fieldDecNodes []Node) *RecordType {
    return &RecordType{
        fieldDecNodes: fieldDecNodes,
    }
}

func (rt *RecordType) nodesAsFieldDecs() []*FieldDeclaration {
    var interfaceList []interface{}
    var fdList        []*FieldDeclaration

    for _, fdNode := range rt.fieldDecNodes {
        interfaceList = append(interfaceList, fdNode.Exp)
    }

    //Cast all to field decs and return that
    for _, fdInterface := range interfaceList {
        fd := fdInterface.(*FieldDeclaration)
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


type RecordCreate struct {
    expType    interface{}
    id string
    fieldCreateNodes    []Node
}

func NewRecordCreate(id string, fieldCreateNodes []Node) *RecordCreate {
    return &RecordCreate{
        id: id,
        fieldCreateNodes: fieldCreateNodes,
    }
}

func (rc *RecordCreate) visit() string {
    str := fmt.Sprintf("(recCreate: id:%s fieldCreateNodes:(", rc.id)
    for _, n := range rc.fieldCreateNodes {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")\n"
    return str
}

func (rc *RecordCreate) analyze(c *Context)  {
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

func (at *ArrayType) visit() string {
    return fmt.Sprintf("(arrType: %s)", at.id)
}


func (at *ArrayType) analyze(c *Context)  {
     at.expType = c.lookup(at.id)
     fmt.Printf("Assigning type %T to arraytype\n", at.expType)
}


type ArrayCreate struct {
    expType  interface{}
    typeId    string
    subscriptNode  Node
    expNode  Node
}

func NewArrayCreate(typeIdentifier string, subscriptNode Node, expNode Node) *ArrayCreate {
    return &ArrayCreate{
        typeId: typeIdentifier,
        subscriptNode: subscriptNode,
        expNode: expNode,
    }
}

func (ac *ArrayCreate) visit() string {
    return fmt.Sprintf("(arrCreate: typeId:%s subscriptNode:%v expNode:%v)", ac.typeId, ac.subscriptNode.visit(), ac.expNode.visit())
}

func (ac *ArrayCreate) analyze(c *Context)  {}

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
        td, isTypeDec := d.Exp.(*TypeDeclaration)
        if(isTypeDec) { //If its a type declaration, add it to the new context
            newContext.add(td.id, td.n.Exp)
        }
    }

    for _, d := range le.declarationNodes {
        d.analyze(newContext)
    }

    for _, d := range le.exps {
        d.analyze(newContext)
    }

    //Check for no recursive type cycles with out record types in decs..
}



type IfThenElseExpression struct {
    expType    interface{}
    condNode  Node
    thenNode  Node
    elseNode  Node
}

func NewIfThenElseExpression(condNode Node, thenNode Node, elseNode Node) *IfThenElseExpression {
    return &IfThenElseExpression{
        condNode:condNode,
        thenNode:thenNode,
        elseNode:elseNode,
    }
}

func (itee *IfThenElseExpression) visit() string {
    return fmt.Sprintf("(ifThenElse if:%v then:%v else:%v)", itee.condNode.Exp.visit(), itee.thenNode.Exp.visit(), itee.elseNode.Exp.visit())
}

func (itee *IfThenElseExpression) analyze(c *Context)  {
}


type IfThenExpression struct {
    expType    interface{}
    condNode  Node
    thenNode  Node
}

func NewIfThenExpression(condNode Node, thenNode Node) *IfThenExpression {
    return &IfThenExpression{
        condNode:condNode,
        thenNode:thenNode,
    }
}

func (ite *IfThenExpression) visit() string {
    return fmt.Sprintf("(ifThen if:%v then:%v)", ite.condNode.Exp.visit(), ite.thenNode.Exp.visit())
}

func (ite *IfThenExpression) analyze(c *Context)  {
}


type WhileExpression struct {
    expType    interface{}
    exp1  Node
    exp2  Node
}

func NewWhileExpression(exp1 Node, exp2 Node) *WhileExpression {
    return &WhileExpression{
        exp1:exp1,
        exp2:exp2,
    }
}

func (we *WhileExpression) visit() string {
    return fmt.Sprintf("(whileExp cond:%v do:%v)", we.exp1.Exp.visit(), we.exp2.Exp.visit())
}


func (we *WhileExpression) analyze(c *Context)  {
}


type ForExpression struct {
    expType    interface{}
    id    string
    exp1  Node
    exp2  Node
    exp3  Node
}

func NewForExpression(id string, exp1 Node, exp2 Node, exp3 Node) *ForExpression {
    return &ForExpression{
        id:id,
        exp1:exp1,
        exp2:exp2,
        exp3:exp3,
    }
}

func (fe *ForExpression) visit() string {
    return fmt.Sprintf("(forexp id:%s exp1:%v exp2:%v exp3:%v)", fe.id, fe.exp1.Exp.visit(), fe.exp2.Exp.visit(), fe.exp3.Exp.visit())
}

func (fe *ForExpression) analyze(c *Context)  {
}
