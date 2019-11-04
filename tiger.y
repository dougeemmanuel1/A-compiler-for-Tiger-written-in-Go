%{

package main

import (
    "github.com/timtadh/lexmachine"
    _ "fmt"
)

var yylineno int = 1
var seqExps = []*Node{}
%}

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

exp_list    : /* episoln */         {  }
            | exp_list exp { seqExps = append(seqExps, $2.ast) }
            | exp_list SEMICOLON exp { seqExps = append(seqExps, $3.ast) }
            ;

expsComma   : exp COMMA expsComma
            | exp
            ;

exp         : lValue
            | NIL
            | INTLIT
            | STRINGLIT
            | seqExp            { $$.ast = $1.ast }
            | negation          { $$.ast = $1.ast }
            | callExp
            | infixExp
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

seqExp      : LPAREN exp_list RPAREN { $$.ast = NewNode("seqexp", nil, NewSeqExpression(seqExps)) }
            ;

negation    : MINUS exp     { $$.ast = NewNode("NEG", $1.token, NewNegation(NewNode("int", $2.token, nil))) }
            ;

callExp     : ID LPAREN expsComma RPAREN
            | ID LPAREN RPAREN
            ;

infixExp    : exp STAR exp              { $$.ast = NewNode("MUL", $2.token, NewInfixExpression(0, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp FORWARDSLASH exp      { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(1, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp PLUS exp              { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(2, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp MINUS exp             { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(3, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp EQUALS exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(4, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp DOUBLEARROW exp       { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(5, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp RARROW exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(6, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp LARROW exp            { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(7, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp GREATERTHANEQ exp     { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(8, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp LESSTHANEQ exp        { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(9, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp AND exp               { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(10, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
            | exp BAR exp               { $$.ast = NewNode("DIV", $2.token, NewInfixExpression(11, NewNode("int", $1.token, nil), NewNode("int", $3.token, nil))) }
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
