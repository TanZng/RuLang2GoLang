grammar Ru;

programa
 : bloque EOF
 ;

bloque
 : (sentencia)*
 ;

sentencia
 : asignacion 
 | sentenciaIf
 | sentenciaWhile
 | log
 | imprimir
 | OTRO {fmt.Println("caracter desconocido: ", $OTRO.text);}
 ;

asignacion
 : ID ASIGNA expr PTOCOMA
 ;

sentenciaIf
 : IF bloqueCondicional (ELSE IF bloqueCondicional)* (ELSE bloqueDeSentencia)?
 ;

bloqueCondicional
 : APAR expr CPAR bloqueDeSentencia
 ;

bloqueDeSentencia
 : ALLAVE bloque CLLAVE
 ;

sentenciaWhile
 : WHILE expr bloqueDeSentencia
 ;

log
 : LOG expr PTOCOMA
 ;
 
imprimir
: IMPRIMIR expr PTOCOMA
;


expr
 : MENOS right=expr                                           #MenosUnarioExpr
 | NOT right=expr                                             #notExpr
 | left=expr op=(MULT | DIV | MOD) right=expr                 #multiplicacionExpr
 | left=expr op=(MAS | MENOS) right=expr                      #aditivaExpr
 | left=expr op=(MAYIG | MENIG | MENORQ | MAYORQ) right=expr  #relacionalExpr
 | left=expr op=(IGUAL | DIFQ) right=expr                     #igualdadExpr
 | left=expr AND right=expr                                   #andExpr
 | left=expr OR right=expr                                    #orExpr
 | atomo                                                      #atomExpr
 ;

atomo
 : APAR expr CPAR #parExpr
 | (INT | FLOAT)  #numberAtom
 | (TRUE | FALSE) #booleanAtom
 | ID             #idAtom
 | STRING         #stringAtom
 | NIL            #nilAtom
 ;

OR : '||';
AND : '&&';
IGUAL : '==';
DIFQ : '!=';
MAYORQ : '>';
MENORQ : '<';
MENIG : '<=';
MAYIG : '>=';
MAS : '+';
MENOS : '-';
MULT : '*';
DIV : '/';
MOD : '%';
POW : '^';
NOT : '!';

PTOCOMA : ';';
ASIGNA : '=';
APAR : '(';
CPAR : ')';
ALLAVE : '{';
CLLAVE : '}';

TRUE : 'true';
FALSE : 'false';
NIL : 'nil';
IF : 'if';
ELSE : 'else';
WHILE : 'while';
LOG : 'log';

IMPRIMIR : 'imprime';

ID
 : [a-zA-Z_] [a-zA-Z_0-9]*
 ;

INT
 : [0-9]+
 ;

FLOAT
 : [0-9]+ '.' [0-9]* 
 | '.' [0-9]+
 ;

STRING
 : '"' (~["\r\n] | '""')* '"'
 ;
COMENTARIO
 : '#' ~[\r\n]* -> skip
 ;
ESPACIO
 : [ \t\r\n] -> skip
 ;
OTRO
 : . 
 ;
