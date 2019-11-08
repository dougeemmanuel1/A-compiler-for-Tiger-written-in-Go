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
    Left *Node
    Right *Node
}

func NewInfixExpression(opType Op, left *Node , right *Node) *InfixExpression {
    return &InfixExpression{
        OpType: opType,
        Left: left,
        Right: right,
    }
}

func (ie *InfixExpression) visit() string {
    return fmt.Sprintf("(%s %v %v)", resolveOp(ie.OpType), string(ie.Left.visit()), ie.Right.visit())

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
    Exps []*Node
}

func NewSeqExpression(expressions []*Node) *SeqExpression {
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

type Nil struct {}

func NewNil() *Nil {
    return &Nil{}
}

func (ni *Nil) visit() string {
    return fmt.Sprintf("(nil)")
}

type CallExpression struct {
    name    string
    exps    []*Node
}

func NewCallExpression(name string, exps []*Node) *CallExpression {
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
