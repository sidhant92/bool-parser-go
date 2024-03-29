grammar BooleanExpression;

parse
 : expression EOF
 ;

expression
 : LPAREN expression RPAREN                                          #parentExpression
 | NOT expression                                                    #notExpression
 | left=expression op=comparator right=expression                    #comparatorExpression
 | left=expression op=binary right=expression                        #binaryExpression
 | types                                                             #typesExpression
 | (field=WORD)? lower=numericTypes TO upper=numericTypes            #toExpression
 | (field=WORD)? (not=NOT)? IN data=wordlist                         #inExpression
 | (field=WORD)? op=arrayOperators data=wordlist                     #arrayExpression
 ;

comparator
 : GT | GE | LT | LE | EQ | NE
 ;


 wordlist
 : LPAREN WS* first=types WS* (',' WS* rest=types WS*)* RPAREN
 ;

 arrayOperators
  : CONTAINS_ANY | CONTAINS_ALL
  ;


 numericTypes
 : INTEGER | DECIMAL
 ;

 types
 : INTEGER | DECIMAL | APP_VERSION | bool | WORD |
 ;


binary
 : AND | OR
 ;

bool
 : TRUE | FALSE
 ;


IN           : 'IN' | 'in';
TO           : 'TO' | 'to';
AND          : 'AND' | 'and' | '&&';
OR           : 'OR' | 'or' | '||';
NOT          : 'NOT' | 'not';
TRUE         : 'TRUE' | 'true';
FALSE        : 'FALSE' | 'false';
CONTAINS_ALL : 'CONTAINS_ALL' | 'contains_all';
CONTAINS_ANY : 'CONTAINS_ANY' | 'contains_any';
NE           : '!=';
GT           : '>' ;
GE           : '>=' ;
LT           : '<' ;
LE           : '<=' ;
EQ           : '=' ;
LPAREN       : '(' ;
RPAREN       : ')' ;
DECIMAL      : [0-9]+ '.' [0-9]+;
APP_VERSION  : [0-9] ('.' INTEGER)+;
INTEGER      : [0-9]+;
WS           : [ \r\t\u000C\n]+ -> skip;
WORD         : (ALPHANUMERIC | '_' | '-' | '.' | SQ | DQ)+;
ALPHANUMERIC : [a-zA-Z0-9];
SQ           : '\''.*? '\'';
DQ           : '"'.*? '"';
