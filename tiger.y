%{

package main

import (
    "github.com/timtadh/lexmachine"
    _ "fmt"
    "strconv"
)

var yylineno int = 1
var seqExps = []*Node{}
var expsByComma = []*Node{}

func toInt(s string)  int {
    num, _ := strconv.ParseInt(s, 10, 64)
    return int(num)
}
%}

%left OR
%left AND
%nonassoc EQUALS NEQ GT LT GEQ LEQ
%left PLUS MINUS
%left STAR FORWARDSLASH
%right UNARY

%union{
    token *lexmachine.Token
    ast   *Node
    NodeList []Node
    FinalList []*Node
    strLit string
}

%token  LPAREN RPAREN LBRACKET RBRACKET LCURLY RCURLY COLON
%token  COLONEQUALS DOT COMMA SEMICOLON STAR FORWARDSLASH PLUS
%token  MINUS EQUALS DOUBLEARROW RARROW
%token  LARROW GREATERTHANEQ LESSTHANEQ AND BAR
%token  ARRAY BREAK DO ELSE END FOR FUNCTION IF IN LET
%token  NIL OF THEN TO TYPE VAR WHILE
%token  ID INTLIT STRINGLIT NEWLINE

%type <NodeList> exp_list
%type <ast> exp
%type <ast> NIL INTLIT STRINGLIT seqExp negation callExp infixExp arrCreate recCreate assignment ifThenElse ifThen whileExp forExp BREAK letExp lValue subscript fieldExp Decs Dec tyDec varDec funDec FUNCTION ty arrTy recTy ARRAY fieldDecs fieldDec
%type <strLit> ID

%start Program
%%

Program     : exp   {  yylex.(*golex).stmts = append(yylex.(*golex).stmts, $1) }
            ;

Dec         : tyDec    { $$ = $1 }
            | varDec    { $$ = $1 }
            | funDec { $$ = $1 }
            ;

tyDec       : TYPE ID EQUALS ty { $$ = $4 }
            ;

ty          : ID       { $$ = NewNode("sometyid", nil, nil) }
            | arrTy    { $$ = $1 }
            | recTy     { $$ = $1 }
            ;

arrTy       : ARRAY OF ID { $$ = $1 }
            ;

recTy       : LCURLY fieldDecs RCURLY  { $$ = $2 }
            | LCURLY RCURLY  { $$ = NewNode("recTy", nil, nil) }
            ;

fieldDecs   : fieldDec COMMA fieldDecs { $$ = $1 }
            | fieldDec  { $$ = $1 }
            ;

fieldDec    : ID COLON ID { $$ = NewNode("fieldDec", nil, nil) }
            ;

funDec      : FUNCTION ID LPAREN fieldDecs RPAREN EQUALS exp   { $$ = $1 }
            | FUNCTION ID LPAREN RPAREN EQUALS exp          { $$ = $1 }
            | FUNCTION ID LPAREN fieldDecs RPAREN COLON ID EQUALS exp { $$ = $1 }
            | FUNCTION ID LPAREN RPAREN COLON ID EQUALS exp { $$ = $1 }
            ;

varDec      : VAR ID COLONEQUALS exp { $$ = $4 }
            | VAR ID COLON ID COLONEQUALS exp { $$ = $6 }
            ;


subscript   : lValue LBRACKET exp RBRACKET { $$ = $1 }
            | ID LBRACKET exp RBRACKET  { $$ = $3 }/*  verbose subscript to force reduce   */
            ;

arrCreate   : ID LBRACKET exp RBRACKET OF exp { $$ = NewNode("arrCreate", nil, nil) }
            ;

lValue      : ID               { $$ = NewNode("someid", nil, nil) }
            | subscript         { $$ = $1 }
            | fieldExp          { $$ = $1 }
            ;

fieldExp    : lValue DOT ID
            ;

exp_list    : /* episoln */         {   }
            | exp_list exp { $$ = append($$, *$2) }
            | exp_list SEMICOLON exp { $$ = append($$, *$3) }
            ;

exp_list_comma    : /* episoln */         {  }
                  | exp_list_comma exp { expsByComma = append(expsByComma, $2.ast) }
                  | exp_list_comma COMMA exp { expsByComma = append(expsByComma, $3.ast) }
                  ;

exp         : NIL               { $$.ast = NewNode("nil", $1.token, NewNil()) }
            | INTLIT            { $$.ast = NewNode("INTLIT", $1.token, NewInteger(toInt(string($1.token.Lexeme)))) }
            | STRINGLIT         { $$.ast = NewNode("STRINGLIT", $1.token, NewStringLiteral(string($1.token.Lexeme))) }
            | seqExp            { $$.ast = $1.ast }
            | negation          { $$.ast = $1.ast }
            | callExp           { $$.ast = $1.ast }
            | infixExp          { $$.ast = $1.ast }
            | arrCreate         { $$.ast = $1.ast }
            | recCreate         { $$.ast = $1.ast }
            | assignment        { $$.ast = $1.ast }
            | ifThenElse        { $$.ast = $1.ast }
            | ifThen            { $$.ast = $1.ast }
            | whileExp          { $$.ast = $1.ast }
            | forExp            { $$.ast = $1.ast }
            | BREAK             { $$.ast = $1.ast }
            | letExp            { $$.ast = $1.ast }
            | lValue            { $$.ast = $1.ast }
            ;

seqExp      : LPAREN exp_list RPAREN { $$.ast = NewNode("seqexp", $2.token, NewSeqExpression(seqExps)) }
            ;

negation    : MINUS exp  %prec UNARY   { $$.ast = NewNode("NEG", $2.token, NewNegation($2.ast)) }
            ;

callExp     :ID LPAREN exp_list_comma RPAREN              { $$.ast = NewNode("CALLEXP", $1.token, NewCallExpression(string($1.token.Lexeme), expsByComma)) }
            ;

infixExp    : exp STAR exp              { $$.ast = NewNode("MUL", $1.token, NewInfixExpression(Op_MUL,    $1.ast, $3.ast)) }
            | exp FORWARDSLASH exp      { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_DIV,    $1.ast, $3.ast)) }
            | exp PLUS exp              { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_PLUS,   $1.ast, $3.ast)) }
            | exp MINUS exp             { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_MINUS,  $1.ast, $3.ast)) }
            | exp EQUALS exp            { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_EQUALS, $1.ast, $3.ast)) }
            | exp DOUBLEARROW exp       { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_NEQ,    $1.ast, $3.ast)) }
            | exp RARROW exp            { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_GT,     $1.ast, $3.ast)) }
            | exp LARROW exp            { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_LT,     $1.ast, $3.ast)) }
            | exp GREATERTHANEQ exp     { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_GTE,    $1.ast, $3.ast)) }
            | exp LESSTHANEQ exp        { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_LTE,    $1.ast, $3.ast)) }
            | exp AND exp               { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_AND,    $1.ast, $3.ast)) }
            | exp BAR exp               { $$.ast = NewNode("DIV", $1.token, NewInfixExpression(Op_OR,     $1.ast, $3.ast)) }
            ;

recCreate   : ID LCURLY fieldCreates RCURLY { $$.ast = $1.ast }
            | ID LCURLY RCURLY  { $$.ast = $1.ast }
            ;

fieldCreates: fieldCreate COMMA fieldCreates
            | fieldCreate
            ;

fieldCreate : ID EQUALS exp
            ;

assignment  : lValue COLONEQUALS exp
            ;

ifThenElse  : IF exp THEN exp ELSE exp  { $$.ast = $2.ast }
            ;

ifThen      : IF exp THEN exp { $$.ast = $2.ast }
            ;

whileExp    : WHILE exp DO exp { $$.ast = $2.ast }
            ;

forExp      : FOR ID COLONEQUALS exp TO exp DO exp { $$.ast = $4.ast }
            ;

Decs        : Dec Decs { $$.ast = $1.ast }
            | Dec { $$.ast = $1.ast }
            ;

letExp      : LET Decs IN exp_list END { $$.ast = $2.ast }
            ;

;
%%
