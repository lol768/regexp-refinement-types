grammar PocLang;
body            : (
                    (WS? expr? WS? NEWLINE)
                    | (WS? COMMENT_LINE NEWLINE)
                    )*;
expr            : expr ('*'|'/') expr
                  |	expr ('+'|'-') expr
                  |	INT
                  | function_call
                  | value_ref
                  |	'(' expr ')' ;
type_keyword      : 'uint' # UnsignedIntType ;
type              : type_keyword
                    | type_keyword '[' int_refinement ']' ;
int_refinement    : int_constraint ' ' INT ;
int_constraint    : '<' # LessThanConstraint
                    | '>' # GreaterThanConstraint
                    | '<=' # LessThanEqualsConstraint
                    | '>=' # GreaterThanEqualsConstraint ;
value_ref         : INT | identifier_ref ;
function_sig      : 'function ' IDENTIFIER '(' (argument_decl | argument_decl ARG_SEP)* '): ' type ' {' ;
function_call     : IDENTIFIER '(' (expr | expr ARG_SEP)* ')' ;
argument_decl     : IDENTIFIER ': ' type ;
identifier_ref      : IDENTIFIER ;
function          : function_sig (NEWLINE | ' ') body '}' ;
program           : (function NEWLINE*)+ ;
COMMENT_LINE      : '//'  ~[\n\r]* -> channel(HIDDEN);
NEWLINE           : [\n]+ ;
ARG_SEP           : [,] WS;
WS                : (' ' | '\t')+ -> channel(HIDDEN);
INT               : [0-9]+ ;
IDENTIFIER        : [A-Za-z_][A-Za-z_0-9]* ;