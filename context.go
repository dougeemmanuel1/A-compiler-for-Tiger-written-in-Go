package main

import(
    "fmt"
    "os"
)

type Context struct {
    parent              *Context
    currentFunction     interface{}
    inLoop              bool
    locals              map[string]interface{}
}

func NewContext(parent *Context, currentFunction interface{}, inLoop bool) *Context {
    return &Context{
        parent:          parent,
        currentFunction: currentFunction,
        inLoop:          inLoop,
        locals:          make(map[string]interface{}),
    }
}

func (c *Context) createChildContextForBlock() *Context {
    //For a block we have to retain both the function and loop settings.
    return NewContext(c, c.currentFunction, c.inLoop)
}


func (c *Context) createChildContextForFunctionBody() *Context {
    //No longer in a loop when entering functions
    return NewContext(c, nil, false)
}

func (c *Context) createChildContextForLoop() *Context {
    return NewContext(c, c.currentFunction, true)
}

func (c *Context) predeclarePrimitives() {
     //Pre declare primitives
     c.locals["int"] = NewIntegerPrimitive()
     c.locals["string"] = &StringPrimitive{}
     c.locals["nil"] = NewNil()

     fmt.Printf("map conmtets %v\n", c.locals)
 }

//Adds a declaration to the current context
func (c *Context) add(identifier string, declaration interface{}) {
    // visitor := declaration.(*Visitor)
    id := ""
    if(identifier == "") {
        id = resolveDeclarationId(declaration)
    } else {
        id = identifier
    }
    _, hasKey := c.locals[id]

    _, isVarDec := declaration.(*Variable)

    if(hasKey && isVarDec) {
        //empty case
    } else if(hasKey) {
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

    fmt.Printf("Adding Dec: %s type: %T\n", id, declaration)
    c.locals[id] = declaration
}

func (c *Context) lookup(id string) interface{} {
    fmt.Printf("Looking up id:%s\n", id)
    for ; c != nil; c = c.parent {
        fmt.Printf("Checking for id: %s in %v\n", id, c.locals )
        e, hasKey := c.locals[id]
        if(hasKey) {
            return e
        }
    }

    fmt.Fprintf(os.Stderr, "%s was not declared. \n", id)
    os.Exit(3)

    //Empty return to satisfy condition code will never each this point
    return Nil{}
}


func resolveDeclarationId(declaration interface{}) string {
    var id string
    typeDec, isTypeDec := declaration.(*TypeDeclaration)
    varDec, isVarDec := declaration.(*Variable)
    funcDec, isFuncDec := declaration.(*FuncDeclaration)

    if(isTypeDec) {
        id = typeDec.id
    } else if(isVarDec) {
        id = varDec.id
    } else if(isFuncDec) {
        id = funcDec.id
    }
    // fmt.Printf("Resolved dec id: %s \n", id)
    return id
}
