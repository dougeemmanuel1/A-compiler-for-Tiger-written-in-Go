%{

package main

import (
    "github.com/timtadh/lexmachine"
    _ "fmt"
    "strconv"
)

var yylineno int = 1

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
    strLit string
}

%token  LPAREN RPAREN LBRACKET RBRACKET LCURLY RCURLY COLON
%token  COLONEQUALS DOT COMMA SEMICOLON STAR FORWARDSLASH PLUS
%token  MINUS EQUALS DOUBLEARROW RARROW
%token  LARROW GREATERTHANEQ LESSTHANEQ AND BAR
%token  ARRAY BREAK DO ELSE END FOR FUNCTION IF IN LET
%token  NIL OF THEN TO TYPE VAR WHILE
%token  ID INTLIT STRINGLIT NEWLINE

%type <NodeList> exp_list_semi exp_list_comma Decs fieldDecs fieldCreates
%type <ast> exp seqExp negation callExp infixExp arrCreate recCreate assignment
%type <ast> funDec FUNCTION ty arrTy recTy ARRAY fieldDec
%type <ast> fieldCreate ifThenElse ifThen whileExp forExp BREAK letExp lValue
%type <ast> NIL  subscript fieldExp Dec tyDec varDec
%type <token> ID INTLIT STRINGLIT
%start Program
%%

Program     : exp   {  yylex.(*golex).stmts = append(yylex.(*golex).stmts, $1) }
            ;

Dec         : tyDec    { $$ = $1 }
            | varDec    { $$ = $1 }
            | funDec { $$ = $1 }
            ;

tyDec       : TYPE ID EQUALS ty { $$ = NewNode("typeDec", nil, NewTypeDeclaration(string($2.Lexeme), $4)) }
            ;

ty          : ID       { $$ = NewNode("ID", $1, NewIdentifier(string($1.Lexeme))) }
            | arrTy    { $$ = $1 }
            | recTy     { $$ = $1 }
            ;

arrTy       : ARRAY OF ID { $$ = NewNode("arrTy", nil, NewArrayType(string($3.Lexeme))) }
            ;

recTy       : LCURLY fieldDecs RCURLY  { $$ = NewNode("sad",nil, NewRecordType($2)) }
            | LCURLY RCURLY  { $$ = NewNode("recTy", nil, NewRecordType([]Node{})) }
            ;

fieldDecs   : fieldDec COMMA fieldDecs { $$ = append($$, *$1) }
            | fieldDec  { $$ = append($$, *$1) }
            ;

fieldDec    : ID COLON ID { $$ = NewNode("fieldDec", nil, NewFieldDeclaration(string($1.Lexeme), string($3.Lexeme))) }
            ;

funDec      : FUNCTION ID LPAREN fieldDecs RPAREN EQUALS exp   { $$ = NewNode("funDec", nil, NewFuncDeclaration(string($2.Lexeme), "", $4,  *$7)) }
            | FUNCTION ID LPAREN RPAREN EQUALS exp             { $$ = NewNode("funDec", nil, NewFuncDeclaration(string($2.Lexeme), "", []Node{}, *$6)) }
            | FUNCTION ID LPAREN fieldDecs RPAREN COLON ID EQUALS exp  { $$ = NewNode("funDec", nil, NewFuncDeclaration(string($2.Lexeme), string($7.Lexeme), $4, *$9)) }
            | FUNCTION ID LPAREN RPAREN COLON ID EQUALS exp  { $$ = NewNode("funDec", nil, NewFuncDeclaration(string($2.Lexeme), string($6.Lexeme), []Node{}, *$8)) }
            ;

varDec      : VAR ID COLONEQUALS exp { $$ = NewNode("varDec", nil, NewVarDeclaration(string($2.Lexeme), "", $4)) }
            | VAR ID COLON ID COLONEQUALS exp { $$ = NewNode("varDec", nil, NewVarDeclaration(string($2.Lexeme), string($4.Lexeme), $6)) }
            ;


subscript   : lValue LBRACKET exp RBRACKET { $$ = $1 }
            | ID LBRACKET exp RBRACKET  { $$ = $3 }/*  verbose subscript to force reduce   */
            ;

arrCreate   : ID LBRACKET exp RBRACKET OF exp { $$ = NewNode("arrCreate", nil, NewArrayCreate(string($1.Lexeme), *$3, *$6)) }
            ;

lValue      : ID                { $$ = NewNode("ID", $1, NewIdentifier(string($1.Lexeme))) }
            | subscript         { $$ = $1 }
            | fieldExp          { $$ = $1 }
            ;

fieldExp    : lValue DOT ID     { $$ = NewNode("fieldExp", nil, NewFieldExpression(*$1, string($3.Lexeme)))}
            ;

exp_list_semi : /* episoln */         {   }
              | exp_list_semi exp { $$ = append($$, *$2) }
              | exp_list_semi SEMICOLON exp { $$ = append($$, *$3) }
              ;

exp_list_comma    : /* episoln */         {   }
            | exp_list_comma exp { $$ = append($$, *$2) }
            | exp_list_comma COMMA exp { $$ = append($$, *$3) }
            ;


exp         : NIL               { $$ = NewNode("nil", nil, NewNil()) }
            | INTLIT            { $$ = NewNode("INTLIT", nil, NewInteger(toInt(string($1.Lexeme)))) }
            | STRINGLIT         { $$ = NewNode("STRINGLIT", nil, NewStringLiteral(string($1.Lexeme))) }
            | seqExp            { $$ = $1 }
            | negation          { $$ = $1 }
            | callExp           { $$ = $1 }
            | infixExp          { $$ = $1 }
            | arrCreate         { $$ = $1 }
            | recCreate         { $$ = $1 }
            | assignment        { $$ = $1 }
            | ifThenElse        { $$ = $1 }
            | ifThen            { $$ = $1 }
            | whileExp          { $$ = $1 }
            | forExp            { $$ = $1 }
            | BREAK             { $$ = $1 }
            | letExp            { $$ = $1 }
            | lValue            { $$ = $1 }
            ;

seqExp      : LPAREN exp_list_semi RPAREN { $$ = NewNode("seqexp", nil, NewSeqExpression($2)) }
            ;

negation    : MINUS exp  %prec UNARY   { $$ = NewNode("NEG", nil, NewNegation($2)) }
            ;

callExp     :ID LPAREN exp_list_comma RPAREN    { $$ = NewNode("CALLEXP", nil, NewCallExpression(string($1.Lexeme), $3)) }
            ;

infixExp    : exp STAR exp              { $$ = NewNode("MUL", nil, NewInfixExpression(Op_MUL,    *$1, *$3)) }
            | exp FORWARDSLASH exp      { $$ = NewNode("DIV", nil, NewInfixExpression(Op_DIV,    *$1, *$3)) }
            | exp PLUS exp              { $$ = NewNode("DIV", nil, NewInfixExpression(Op_PLUS,   *$1, *$3)) }
            | exp MINUS exp             { $$ = NewNode("DIV", nil, NewInfixExpression(Op_MINUS,  *$1, *$3)) }
            | exp EQUALS exp            { $$ = NewNode("DIV", nil, NewInfixExpression(Op_EQUALS, *$1, *$3)) }
            | exp DOUBLEARROW exp       { $$ = NewNode("DIV", nil, NewInfixExpression(Op_NEQ,    *$1, *$3)) }
            | exp RARROW exp            { $$ = NewNode("DIV", nil, NewInfixExpression(Op_GT,     *$1, *$3)) }
            | exp LARROW exp            { $$ = NewNode("DIV", nil, NewInfixExpression(Op_LT,     *$1, *$3)) }
            | exp GREATERTHANEQ exp     { $$ = NewNode("DIV", nil, NewInfixExpression(Op_GTE,    *$1, *$3)) }
            | exp LESSTHANEQ exp        { $$ = NewNode("DIV", nil, NewInfixExpression(Op_LTE,    *$1, *$3)) }
            | exp AND exp               { $$ = NewNode("DIV", nil, NewInfixExpression(Op_AND,    *$1, *$3)) }
            | exp BAR exp               { $$ = NewNode("DIV", nil, NewInfixExpression(Op_OR,     *$1, *$3)) }
            ;

recCreate   : ID LCURLY fieldCreates RCURLY { $$ = NewNode("recCreate", nil, NewRecordCreate(string($1.Lexeme), $3)) }
            | ID LCURLY RCURLY  { $$ = NewNode("recCreate", nil, NewRecordCreate(string($1.Lexeme), []Node{})) }
            ;

fieldCreates: fieldCreates COMMA fieldCreate   { $$ = append($$, *$3) }
            | fieldCreate   { $$ = append($$, *$1) }
            ;

fieldCreate : ID EQUALS exp { $$ = NewNode("fieldCreate", nil, NewFieldCreate(string($1.Lexeme), *$3)) }
            ;

assignment  : lValue COLONEQUALS exp    { $$ = NewNode("assignment", nil, NewAssignment(*$1, *$3))}
            ;

ifThenElse  : IF exp THEN exp ELSE exp  { $$ = NewNode("Iftheneelse", nil, NewIfThenElseExpression(*$2, *$4, *$6)) }
            ;

ifThen      : IF exp THEN exp { $$ = $2 }
            ;

whileExp    : WHILE exp DO exp { $$ = $2 }
            ;

forExp      : FOR ID COLONEQUALS exp TO exp DO exp { $$ = $4 }
            ;

Decs        : Decs Dec { $$ = append($$, *$2) }
            | Dec { $$ = append($$, *$1) }
            ;

letExp      : LET Decs IN exp_list_semi END { $$ = NewNode("letexp", nil, NewLetExpression($2, $4)) }
            ;

;
%%
