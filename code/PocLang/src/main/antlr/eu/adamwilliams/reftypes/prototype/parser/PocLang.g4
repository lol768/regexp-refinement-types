grammar PocLang;
body              : (
                    (WS? body_line? WS? NEWLINE)
                      | (WS? COMMENT_LINE NEWLINE)
                    )*;
body_line         : (var_assignment | return_stmt | var_decl | function_call ) ;

return_stmt       : 'return ' expr ;
var_decl          : 'var ' IDENTIFIER ': ' type ;
type_keyword      : 'uint' # UnsignedIntType |
                    'string' # StringType |
                    'void' # VoidType ;
type              : (type_keyword)
                    | (type_keyword '[' int_constraint ']') ;
int_constraint    : '< ' INT # LessThanConstraint
                    | '> ' INT # GreaterThanConstraint
                    | '<= ' INT # LessThanEqualsConstraint
                    | '>= ' INT # GreaterThanEqualsConstraint ;
function_sig      : 'function ' IDENTIFIER '(' (argument_decl | argument_decl ARG_SEP)* '): ' type ' {' NEWLINE ;
argument_decl     : IDENTIFIER ': ' type ;
function          : function_sig body '}' ;
program           : (function NEWLINE*)+ EOF;
expr              : expr ('*'|'/') expr
                    |	expr ('+'|'-') expr
                    | value_ref
                    | function_call
                    |	'(' expr ')' ;
var_assignment    : IDENTIFIER WS? '=' expr;
value_ref         : INT | STRING_LITERAL | identifier_ref ;
identifier_ref    : IDENTIFIER ;
function_call     : IDENTIFIER '(' (expr (ARG_SEP expr)*)? ')' ;

IDENTIFIER        : [A-Za-z_][A-Za-z_0-9]* ;
WS                : (' ' | '\t')+ -> channel(HIDDEN);
COMMENT_LINE      : '//'  ~[\n\r]* -> channel(HIDDEN);
NEWLINE           : [\n]+ ;
ARG_SEP           : [,] WS?;
INT               : [0-9]+ ;
// SO 19237249
fragment ESCAPED_QUOTE : '\\"';
STRING_LITERAL :   '"' ( ESCAPED_QUOTE | ~('\n'|'\r') )*? '"';