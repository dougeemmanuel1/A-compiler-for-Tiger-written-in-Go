//      Author:    Emmanuel Douge
//      Due Date:  September 26, 2019
//      Course:    CSC425
//      Professor Name: Dr. Schwesinger
//      Assignment: 1
package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

import (
    "github.com/timtadh/lexmachine"
    _ "github.com/timtadh/lexmachine/machines"
)

func main () {
    if(len(os.Args) != 2) {
        fmt.Fprintf(os.Stderr, "Incorrect number of arguments.\n")
        fmt.Fprintf(os.Stderr, "Usage: go run main.go <filename>\n")
        os.Exit(1)
    }

    dat, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
    }

    //Creates and compiles lexer with rules
    lexer, _ := newLexer()

    //Lex the file contents first then parse them, but they are
    //done simulatneously within the parse call because of yacc.
    stmts, err1 := parse(lexer, []byte(dat))
    if err1 != nil {
		fmt.Fprintln(os.Stderr, err1)
		os.Exit(1)
	}

    // Print AST this is for later
	// for _, stmt := range stmts {
        // fmt.Println("%d statements.", len(stmts))
		fmt.Printf("Printing AST: %v\n", stmts)

    //Create context wehich is basically a symbol table
    c := NewContext(nil, nil, false)

    //Pre declare primitives
    c.predeclarePrimitives()
    
    //Begin semantic analysis
    stmts[0].analyze(c)
// }
    os.Exit(0)
}

func parse(lexer *lexmachine.Lexer, text []byte) (stmts []*Node, err error) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case error:
				err = e.(error)
				stmts = nil
			default:
				panic(e)
			}
		}
	}()
    //File is lexed within this function, returns scanner
	scanner, err := newGoLex(lexer, text)
	if err != nil {
		return nil, err
	}
	yyParse(scanner)
	return scanner.stmts, nil
}
