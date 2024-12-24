
grammar LangMak;

prog:   stat* EOF;

stat:   ID '=' expr ';'   
    |   'if' '(' expr ')' '{' stat* '}' ('else' '{' stat* '}')?   
    |   'while' '(' expr ')' '{' stat* '}'       
    |   'print' '(' expr ')' ';'          
    ;

expr:   expr op=('<' | '==' | '>') expr
    |   expr op=('*' | '/') expr
    |   expr op=('+' | '-') expr
    |   INT
    |   STRING
    |   ID
    |   '(' expr ')'
    ;


LPAREN: '(' ;
RPAREN: ')' ;

MULT  : '*' ;
DIV   : '/' ;

PLUS  : '+' ;
MINUS : '-' ;

ASSIGN: '=' ;
GREAT : '>' ;
LESS  : '<' ;
EQUAL : '==';

ID  : [a-zA-Z_][a-zA-Z0-9_]* ;
INT : [0-9]+ ;
STRING: '"' ('\\' . | ~["\\])* '"' ;
WS  : [ \t\r\n]+ -> skip ;
COMMENT : '/*' .*? '*/' -> skip ;
LINE_COMMENT : '//' ~[\r\n]* -> skip ;
