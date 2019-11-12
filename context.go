package main

import(
    "fmt"
    "os"
)

type Context struct {
    parent              *Context
    currentFunction     *Visitor
    inLoop              bool
    locals              map[string]*Visitor
}

func NewContext(parent *Context, currentFunction *Visitor, inLoop bool) *Context {
    return &Context{
        parent:          parent,
        currentFunction: currentFunction,
        inLoop:          inLoop,
        locals:          make(map[string]*Visitor),
    }
}

func (c *Context) createChildContextForBlock() *Context {
    //For a block we have to retain both the function and loop settings.
    return NewContext(c, c.currentFunction, c.inLoop)
}

//Adds a declaration to the current context
func (c *Context) add(declaration interface{}) {
    visitor := declaration.(*Visitor)
    declarationId := resolveDeclarationId(visitor)
    id, hasKey := c.locals[declarationId]
    if(hasKey) {
        fmt.Printf("%s already declared in this scope.\n", id)
        os.Exit(3)
    }

    // typeDec, ok := visitor.(*TypeDeclaration)
    // var entity *Visitor
    // if(ok) { //If we can ensure its a type declaration give them the type expression
    //     entity = typeDec.Exp
    // } else { //If not give them the basic entity and well figure it out later
    //     fmt.Println("Type cast to typedec failed...")
    //     entity = declaration
    // }

    c.locals[declarationId] = visitor
}

func resolveDeclarationId(declaration *Visitor) string {
    var id string
    typeDec, isTypeDec := declaration.(*TypeDeclaration)
    varDec, isVarDec := declaration.(*VarDeclaration)
    funcDec, isFuncDec := declaration.(*FuncDeclaration)

    if(isTypeDec) {
        id = typeDec.id
    } else if(isVarDec) {
        id = varDec.id
    } else if(isFuncDec) {
        id = funcDec.id
    }
    return id
}
