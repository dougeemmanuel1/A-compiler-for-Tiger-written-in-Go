package main

import (
    "fmt"
    _ "github.com/timtadh/lexmachine"
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

type Integer struct {
    Number int
}

func NewInteger(number int) *Integer {
    return &Integer{
        Number: number,
    }
}

func (i *Integer) visit() string {
    return fmt.Sprintf("(int %d)", i.Number)
}



type InfixExpression struct {
    OpType Op
    Left Node
    Right  Node
}

func NewInfixExpression(opType Op, left Node , right Node) *InfixExpression {
    return &InfixExpression{
        OpType: opType,
        Left: left,
        Right: right,
    }
}

func (ie *InfixExpression) visit() string {
    return fmt.Sprintf("(%s %v %v)", resolveOp(ie.OpType), string(ie.Left.Exp.visit()), ie.Right.Exp.visit())

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
    Exp *Node
}

func NewNegation(exp *Node) *Negation {
    return &Negation{
        Exp: exp,
    }
}

func (ne *Negation) visit() string {
    var str string
    if(ne.Exp == nil) {
        fmt.Println("Exp is null")
        str = string(ne.Exp.Token.Lexeme)
    } else {
        str = fmt.Sprintf("(NEG %v)", ne.Exp.visit())
    }
    return str
}

type SeqExpression struct {
    Exps []Node
}

func NewSeqExpression(expressions []Node) *SeqExpression {
    // var copiedExpressionContents []Node
    // for _, e := range expressions {
        // copiedExpressionContents = append(copiedExpressionContents, *e)
    // }
    return &SeqExpression{
        Exps: expressions,
    }
}

func (se *SeqExpression) visit() string {
    str := "(seqexp "
    for _, n := range se.Exps {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\n%v ", n.Exp.visit())
        }
    }

    str += "\n)"
    return str
}

type StringLiteral struct {
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

type Assignment struct {
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


type Nil struct {}

func NewNil() *Nil {
    return &Nil{}
}

func (ni *Nil) visit() string {
    return fmt.Sprintf("(nil)")
}

type CallExpression struct {
    name    string
    exps    []Node
}

func NewCallExpression(name string, exps []Node) *CallExpression {
    return &CallExpression{
        name: name,
        exps: exps,
    }
}

func (ce *CallExpression) visit() string {
    str := fmt.Sprintf("(callExp: %s", ce.name)
    for i, n := range ce.exps {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\nparam %d: %v ", i+1, n.Exp.visit())
        }
    }
    str += "\n)"
    return str
}

type TypeDeclaration struct {
    id    string
    Exp    Node
}

func NewTypeDeclaration(identifier string, n *Node) *TypeDeclaration {
    return &TypeDeclaration{
        id: identifier,
        Exp: *n,
    }
}

func (td *TypeDeclaration) visit() string {
    return fmt.Sprintf("(tyDec: type:%s %s)", td.id, td.Exp.visit())
}

type FuncDeclaration struct {
    id1     string
    id2     string
    decs   []Node
    exp    Node
}

func NewFuncDeclaration(id1 string, id2 string, declarations []Node, n Node) *FuncDeclaration {
    return &FuncDeclaration{
        id1: id1,
        id2: id2,
        decs: declarations,
        exp: n,
    }
}

func (fd *FuncDeclaration) visit() string {
    str := fmt.Sprintf("(funDec: id1:%s id2:%s decs:", fd.id1, fd.id2)
    for _, n := range fd.decs {
        str += fmt.Sprintf("(%v)\n", n.Exp.visit())
    }
    str += fmt.Sprintf("exp:%s)", fd.exp.Exp.visit())
    return str
}

type FieldDeclaration struct {
    id1    string
    id2    string
}

func NewFieldDeclaration(identifier1 string, identifier2 string) *FieldDeclaration {
    return &FieldDeclaration{
        id1: identifier1,
        id2: identifier2,
    }
}

func (fid *FieldDeclaration) visit() string {
    return fmt.Sprintf("fieldDec: (id1:%s) (id2:%s)", fid.id1, fid.id2)
}

type FieldExpression struct {
    lValue    Node
    id        string
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

type FieldCreate struct {
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


type VarDeclaration struct {
    id1    string
    id2    string
    Exp    Node
}

func NewVarDeclaration(identifier1 string, identifier2 string, n *Node) *VarDeclaration {
    return &VarDeclaration{
        id1: identifier1,
        id2: identifier2,
        Exp: *n,
    }
}

func (vd *VarDeclaration) visit() string {
    return fmt.Sprintf("(varDec: id1:%s id2:%s exp:%s)", vd.id1, vd.id2, vd.Exp.visit())
}


type Identifier struct {
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

type Subscript struct {
    id          string
    expId        Node
    subscriptExp Node
}

func NewSubscript(id string, expId Node, subscriptExp Node) *Subscript {
    return &Subscript{
        id: id,
        expId: expId,
        subscriptExp: subscriptExp,
    }
}

func (s *Subscript) visit() string {
    var str string
    if(s.id != "") {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", s.id, s.subscriptExp.Exp.visit())
    } else {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", s.expId.Exp.visit(), s.subscriptExp.Exp.visit())
    }
    return str
}

type RecordType struct {
    decs    []Node
}

func NewRecordType(decs []Node) *RecordType {
    return &RecordType{
        decs: decs,
    }
}

func (rt *RecordType) visit() string {
    str := fmt.Sprintf("(recTy: decs:(")
    for _, n := range rt.decs {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")"
    return str
}

type RecordCreate struct {
    id string
    decs    []Node
}

func NewRecordCreate(id string, decs []Node) *RecordCreate {
    return &RecordCreate{
        id: id,
        decs: decs,
    }
}

func (rc *RecordCreate) visit() string {
    str := fmt.Sprintf("(recCreate: decs:(")
    for _, n := range rc.decs {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")\n"
    return str
}



type ArrayType struct {
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


type ArrayCreate struct {
    id    string
    exp1  Node
    exp2  Node
}

func NewArrayCreate(identifier string, exp1 Node, exp2 Node) *ArrayCreate {
    return &ArrayCreate{
        id: identifier,
        exp1: exp1,
        exp2: exp2,
    }
}

func (ac *ArrayCreate) visit() string {
    return fmt.Sprintf("(arrCreate: id:%s exp1:%v exp2:%v)", ac.id, ac.exp1.visit(), ac.exp2.visit())
}


type LetExpression struct {
    decs    []Node
    exps    []Node
}

func NewLetExpression(declarations []Node, expressions []Node) *LetExpression {
    return &LetExpression{
        decs: declarations,
        exps: expressions,
    }
}

func (le *LetExpression) visit() string {
    str := fmt.Sprintf("(letExp: decs:(")
    for _, n := range le.decs {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }
    str += fmt.Sprintf(")\n(exps: ")
    for _, n := range le.exps {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }

    str += "))"
    return str
}

type IfThenElseExpression struct {
    exp1  Node
    exp2  Node
    exp3  Node
}

func NewIfThenElseExpression(exp1 Node, exp2 Node, exp3 Node) *IfThenElseExpression {
    return &IfThenElseExpression{
        exp1:exp1,
        exp2:exp2,
        exp3:exp3,
    }
}

func (itee *IfThenElseExpression) visit() string {
    return fmt.Sprintf("(ifThenElse if:%v then:%v else:%v)", itee.exp1.Exp.visit(), itee.exp2.Exp.visit(), itee.exp3.Exp.visit())
}
