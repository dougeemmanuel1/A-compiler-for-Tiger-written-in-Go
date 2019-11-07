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
}

%token  LPAREN RPAREN LBRACKET RBRACKET LCURLY RCURLY COLON
%token  COLONEQUALS DOT COMMA SEMICOLON STAR FORWARDSLASH PLUS
%token  MINUS EQUALS DOUBLEARROW RARROW
%token  LARROW GREATERTHANEQ LESSTHANEQ AND BAR
%token  ARRAY BREAK DO ELSE END FOR FUNCTION IF IN LET
%token  NIL OF THEN TO TYPE VAR WHILE
%token  ID INTLIT STRINGLIT NEWLINE

%start Program
%%

Program     : exp   {  yylex.(*golex).stmts = append(yylex.(*golex).stmts, $1.ast) }
            ;

Dec         : tyDec
            | varDec
            | funDec
            ;

tyDec       : TYPE ID EQUALS ty
            ;

ty          : ID
            | arrTy
            | recTy
            ;

arrTy       : ARRAY OF ID
            ;

recTy       : LCURLY fieldDecs RCURLY
            | LCURLY RCURLY
            ;

fieldDecs   : fieldDec COMMA fieldDecs
            | fieldDec
            ;

fieldDec    : ID COLON ID
            ;

funDec      : FUNCTION ID LPAREN fieldDecs RPAREN EQUALS exp
            | FUNCTION ID LPAREN RPAREN EQUALS exp
            | FUNCTION ID LPAREN fieldDecs RPAREN COLON ID EQUALS exp
            | FUNCTION ID LPAREN RPAREN COLON ID EQUALS exp
            ;

varDec      : VAR ID COLONEQUALS exp
            | VAR ID COLON ID COLONEQUALS exp
            ;


subscript   : lValue LBRACKET exp RBRACKET
            | ID LBRACKET exp RBRACKET /*  verbose subscript to force reduce   */
            ;

arrCreate   : ID LBRACKET exp RBRACKET OF exp
            ;

lValue      : ID
            | subscript
            | fieldExp
            ;

fieldExp    : lValue DOT ID
            ;

exp_list    : /* episoln */         { $$ = []*Node{}  }
            | exp_list exp { seqExps = append(seqExps, $2.ast) }
            | exp_list SEMICOLON exp { seqExps = append(seqExps, $3.ast) }
            ;

exp_list_comma    : /* episoln */         {  }
                  | exp_list_comma exp { expsByComma = append(expsByComma, $2.ast) }
                  | exp_list_comma COMMA exp { expsByComma = append(expsByComma, $3.ast) }
                  ;

exp         : lValue
            | NIL               { $$.ast = NewNode("nil", $1.token, NewNil()) }
            | INTLIT            { $$.ast = NewNode("INTLIT", $1.token, NewInteger(toInt(string($1.token.Lexeme)))) }
            | STRINGLIT         { $$.ast = NewNode("STRINGLIT", $1.token, NewStringLiteral(string($1.token.Lexeme))) }
            | seqExp            { $$.ast = $1.ast }
            | negation          { $$.ast = $1.ast }
            | callExp           { $$.ast = $1.ast }
            | infixExp          { $$.ast = $1.ast}
            | arrCreate
            | recCreate
            | assignment
            | ifThenElse
            | ifThen
            | whileExp
            | forExp
            | BREAK
            | letExp
            ;

seqExp      : LPAREN exp_list RPAREN { $$.ast = NewNode("seqexp", $2.token, NewSeqExpression(seqExps)) }
            ;

negation    : MINUS exp  %prec UNARY   { $$.ast = NewNode("NEG", $1.token, NewNegation($2.ast)) }
            ;

callExp     :ID LPAREN exp_list_comma RPAREN              { $$.ast = NewNode("CALLEXP", $1.token, NewCallExpression(string($1.token.Lexeme), expsByComma)) }
            ;

infixExp    : exp STAR exp              { $$.ast = NewNode("MUL", $2.token, NewInfixExpression(Op_MUL,    $1.ast, $3.ast)) }
            | exp FORWARDSLASH exp      { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_DIV,    $1.ast, $3.ast)) }
            | exp PLUS exp              { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_PLUS,   $1.ast, $3.ast)) }
            | exp MINUS exp             { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_MINUS,  $1.ast, $3.ast)) }
            | exp EQUALS exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_EQUALS, $1.ast, $3.ast)) }
            | exp DOUBLEARROW exp       { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_NEQ,    $1.ast, $3.ast)) }
            | exp RARROW exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_GT,     $1.ast, $3.ast)) }
            | exp LARROW exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_LT,     $1.ast, $3.ast)) }
            | exp GREATERTHANEQ exp     { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_GTE,    $1.ast, $3.ast)) }
            | exp LESSTHANEQ exp        { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_LTE,    $1.ast, $3.ast)) }
            | exp AND exp               { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_AND,    $1.ast, $3.ast)) }
            | exp BAR exp               { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(Op_OR,     $1.ast, $3.ast)) }
            ;

recCreate   : ID LCURLY fieldCreates RCURLY
            | ID LCURLY RCURLY
            ;

fieldCreates: fieldCreate COMMA fieldCreates
            | fieldCreate
            ;

fieldCreate : ID EQUALS exp
            ;

assignment  : lValue COLONEQUALS exp
            ;

ifThenElse  : IF exp THEN exp ELSE exp
            ;

ifThen      : IF exp THEN exp
            ;

whileExp    : WHILE exp DO exp
            ;

forExp      : FOR ID COLONEQUALS exp TO exp DO exp
            ;

Decs        : Dec Decs
            | Dec
            ;

letExp      : LET Decs IN exp_list END
            ;

;
%%
