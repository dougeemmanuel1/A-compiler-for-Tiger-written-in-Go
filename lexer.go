//      Author:    Emmanuel Douge
//      Due Date:  September 26, 2019
//      Course:    CSC425
//      Professor Name: Dr. Schwesinger
//      Assignment: 1
package main

//Import standard usage stuff
import (
    "fmt"
    "strings"
    "os"
)

//Import third party libraries
import (
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)

type golex struct {
	*lexmachine.Scanner
	stmts []*Node
}

var Literals []struct{ t string ; s string }   //The tokens representing literal strings
var Keywords []string   //Keywords from the tiger language
var OtherTokens []string     //All tokens including the above
var TokenIds map[string]int   //A map from the token name to their int ids

func initTokens() {
    Literals = []struct {
            t   string   // token eg: LPAREN
            s   string   // string eg  (

        } {{ t: "LPAREN" , s: "(" }, { t: "RPAREN" , s: ")" },
           { t: "LBRACKET" , s: "[" }, { t: "RBRACKET" , s: "]"  },
           { t: "LCURLY" , s: "{" }, { t: "RCURLY" , s: "}"  },
           { t: "COLON" , s: ":" }, { t: "COLONEQUALS" , s: ":="  },
           { t: "DOT" , s: "." }, { t: "COMMA" , s: ","  },
           { t: "SEMICOLON" , s: ";" }, { t: "STAR" , s: "*"  },
           { t: "FORWARDSLASH" , s: "/" }, { t: "PLUS" , s: "+"  },
           { t: "MINUS" , s: "-" }, { t: "EQUALS" , s: "="  },
           { t: "DOUBLEARROW" , s: "<>" }, { t: "RARROW" , s: ">"  },
           { t: "LARROW" , s: "<" }, { t: "GREATERTHANEQ" , s: ">="  },
           { t: "LESSTHANEQ" , s: "<=" }, { t: "AND" , s: "&"  },
           { t: "BAR" , s: "|"  },
    }

    Keywords = []string{
        "array",  "break",  "do",  "else",
        "end",  "for",  "function",  "if",
        "in",  "let",  "nil",  "of",
        "then",  "to",  "type",  "var",
        "while",
    }

    //Seperate tokens that dont exactly fit into the above categories
    OtherTokens = []string {
        "id",
        "intLit",
        "stringLit",
    }

    //Map that will hold tokens identifirs to type
    TokenIds = make(map[string]int)
}

func newLexer() (*lexmachine.Lexer, error) {
    initTokens()
    lexer := lexmachine.NewLexer()

	for id, name := range yyToknames {
		TokenIds[name] = id
        // fmt.Printf("yyTok:%d - %v\n", id, name)
	}

    //Populate the lexer with all the tokens and their action functions
    for _, lit := range Literals {
        //For the literals we are escaping all of the characters with a
        //a back slash to make sure they are interpreted as themselves in
        //regex.

        regex := "\\" + strings.Join(strings.Split(lit.s, ""), "\\")
        lexer.Add([]byte(regex), token(TokenIds[lit.t]))
        // fmt.Printf("reg: %s tok: %s id: %d\n", regex, lit.t, TokenIds[lit.t])
    }

    //For the keyword tokens, the regular expression for them will just be
    //the keyword itself.
    for _, name := range Keywords {
        // fmt.Printf("Keyword - reg: %s gettok: %s\n", name, strings.ToUpper(name))
        lexer.Add([]byte(name), token(TokenIds[strings.ToUpper(name)]))
    }

    //Add the tokens to the lexer with more complex patterns
    lexer.Add([]byte(`[a-zA-Z][a-zA-Z0-9_]*`), token(TokenIds["ID"]))

    //Catch integer literals
    lexer.Add([]byte(`[0-9]+`), token(TokenIds["INTLIT"]))

    //This regexp catches all none "\ characters and manually accepts
    //all other valid forms including a backslash
    //I tried to shorthang the uni definition \ddd with \\[0-9]{3}
    //but it refused to work, rip :(
    lexer.Add([]byte(`"([^"\\]|\\n|\\t|\\\^c|\\[0-9][0-9][0-9]|\\\\|\\")*"`), token(TokenIds["STRINGLIT"]))

    //Catch comments and turn them into a comment TokenIds
    //In a future version I will switch it for a skip action
    //instead as there is no reason to lex it.
    lexer.Add([]byte(`\/\*([^*\n]|\*+[^\*\/])*\*+\/`), skip)
    lexer.Add([]byte(`\/\*([^*]|\*+[^\*\/])*\*+\/`), newline)


    //Catch whitespace and skip it
    lexer.Add([]byte("( |\t)+"), skip)

    //Catch newlines and increment yylineno
    lexer.Add([]byte("(\r|\n)"), newline)

    //Compile the NFA
    err := lexer.Compile()
    if err != nil {
        // return nil, err
        fmt.Printf("Error compiling NFA: %v", err)
        panic(err)
    }
    return lexer, nil
}

//The add function of the lexer takes actions. This is a predefined action which
//essentially performs as a no op. When nil is returned for the token (the interface here),
//the match is skipped which is useful for ignoring whitespace, comments, etc.
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

//Special action i created when you reach a newline character.
//It increments the yylineno var declared in the tiger.y file and ultimtealy
//generated in the y.go file. 
func newline(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
    yylineno = yylineno + 1
    // fmt.Println(yylineno)
	return nil, nil
}

// A function which returns Lexmachine.Action function which will construct a token
//of the given token type by the token types name
func token(tokenType int) lexmachine.Action {

	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
        //The token function defined by the scanner here creates a token
        //a token for you by taking the tokens type (retrieve through the map of token name to id)
        //and the lexeme followed by an error.
        // fmt.Printf("Returning token: %s id: %d\n", string(m.Bytes), tokenType)
		return s.Token(tokenType, string(m.Bytes), m), nil
	}
}

// Construct a new golex from a lexer object and the text to parse.
func newGoLex(lexer *lexmachine.Lexer, text []byte) (*golex, error) {
	scan, err := lexer.Scanner(text)
    if err != nil {
        fmt.Printf("Scanner error: %v", err)
    }

        // token := tok.(*lexmachine.Token)
        // fmt.Printf("\n%-7v | %-10v | %v:%v-%v:%v\n",
        //     TokenIds[string(token.Lexeme)],
        //     string(token.Lexeme),
        //     token.StartLine,
        //     token.StartColumn,
        //     token.EndLine,
        //     token.EndColumn)

	return &golex{Scanner: scan}, nil
}

// Lex implements yyLexer's interface for getting the next token. It returns the
// token type as an integer. The tokens should be defined in the $parser.y file.
// The actual number returned will be >= yyPrivate - 1 which is the range for
// custom token names.
func (g *golex) Lex(lval *yySymType) (tokenType int) {
	s := g.Scanner
    tok, err, eof := s.Next()

    //LExically analyze the input file.
    if err != nil {
            //Type assert lexer error to specific error type to inspect
            //error properties
            if inputError, ok := err.(*machines.UnconsumedInput); ok {
                fmt.Fprintf(os.Stderr, "ERROR: Line %d: Lexer: %v\n", inputError.StartLine, err)
            } else {
                fmt.Fprintf(os.Stderr, "ERROR: Lexer: %v\n", err)
            }
            os.Exit(1)
    } else if eof {
        return -1 // signals EOF to goyacc's yyParse
    }

	lval.token = tok.(*lexmachine.Token)

	// To return the correct number for goyacc you must add yyPrivate - 1 to
	// put the value into the correct range.
	return lval.token.Type + yyPrivate - 1
}

//This implements the error functionality of goyacc, theres not much you
//can do here given the context. The code generated through goyacc  only
//returns a string
func (l *golex) Error(message string) {
	fmt.Fprintf(os.Stderr, "ERROR: Line: %d Parser: %s\n", yylineno, message)
    os.Exit(2)
}
