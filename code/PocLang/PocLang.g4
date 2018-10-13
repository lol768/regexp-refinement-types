grammar PocLang;
body            : (WS? expr NEWLINE)* ;
expr            : expr ('*'|'/') expr
                  |	expr ('+'|'-') expr
                  |	INT
                  |	'(' expr ')' ;
type_keyword      : 'uint' # UnsignedIntType ;
type              : type_keyword
                    | type_keyword '[' int_refinement ']' ;
int_refinement    : int_constraint ' ' INT ;
int_constraint    : '<' # LessThanConstraint
                    | '>' # GreaterThanConstraint
                    | '<=' # LessThanEqualsConstraint
                    | '>=' # GreaterThanEqualsConstraint ;                 
function_sig      : 'function ' IDENTIFIER '(' (argument | argument ARG_SEP)* '): ' type ' {' ;
argument          : IDENTIFIER ': ' type ;
function          : function_sig (NEWLINE | ' ') body '}' ;
NEWLINE           : [\n]+ ;
ARG_SEP           : [,] WS;
WS                : (' ' | '\t')+ -> channel(HIDDEN);
INT               : [0-9]+ ;
IDENTIFIER        : [A-Za-z_][A-Za-z_0-9]* ;
