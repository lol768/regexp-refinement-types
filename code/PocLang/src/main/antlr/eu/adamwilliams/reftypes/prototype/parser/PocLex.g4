lexer grammar PocLex ;

// ANTLR does not support rule refs in sets, so no DRY here
INT               : [0-9]+ ;


MULTIPLY       : '*' ;
ADD            : '+' ;
SUBTRACT       : '-' ;
DIVIDE         : '/' ;
RETURN         : 'return' ;
VAR            : 'var' ;
FUNCTION       : 'function' ;
IF             : 'if' ;
ELSE           : 'else' ;
COLON          : ':' ;
SPACE          : ' ' ;
UINT_T         : 'uint' ;
STRING_T       : 'string' ;
BOOL_T         : 'bool' ;
VOID_T         : 'void' ;
TRUE_LIT       : 'true' ;
FALSE_LIT      : 'false' ;

BEGIN_CONSTRAINT : '[' -> pushMode(CONSTRAINT);

BEGIN_GROUP: '(' ;
END_GROUP: ')' ;


BEGIN_BODY: '{' ;
END_BODY: '}' ;

IDENTIFIER     : [A-Za-z_] [A-Za-z_0-9]* ;


WS                : (' ' | '\t')+ -> channel(HIDDEN) ;
COMMENT_LINE      : '//'  ~[\n\r]* ;
NEWLINE           : [\n]+ ;
ARG_SEP           : [,] WS?;
// SO 19237249
fragment ESCAPED_QUOTE : '\\"';
STRING_LITERAL :   '"' ( ESCAPED_QUOTE | ~('\n'|'\r') )*? '"' ;

mode REGEX;

BEGIN_RE_RANGE : '[' -> pushMode(REGEX_RANGE) ;
END_RE_RANGE   : ']' ;
CHARACTER      : ~('\n'|'\r'|'.'|'('|')'|'['|']'|'*'|'+'|'/'|'|') | ESCAPED_META;
BEGIN_RE_GROUP : '(' ;
END_RE_GROUP   : ')' ;
MINUS          : '-' ;
DOT            : '.' ;
STAR           : '*' ;
PLUS           : '+' ;
ALTERNATION    : '|' ;
RANGE_NEGATE   : '^' ;
RE_DELIMITER_CLOSE: '/' -> popMode ;
META_CHAR      : (DOT|PLUS|STAR|ALTERNATION|'['|']'|'/') ;
fragment ESCAPED_FWD_SLASH: '\\/' ;
fragment ESCAPED_META: '\\' META_CHAR ;

mode REGEX_RANGE;
RANGE_SEPARATOR        : '-' ;
RANGE_CHARACTER        : ~('\n'|'\r'|']') | ESCAPED_RANGE;
RANGE_TERMINATE        : ']' -> popMode ;
fragment ESCAPED_RANGE : '\\' ']' ;

mode CONSTRAINT;
GT: '>' ;
GE: '>=' ;
LT: '<' ;
LE: '<=' ;
EQ: '=' ;
CONSTRAINT_UINT: [0-9]+ ;
CONSTRAINT_SPACE: ' ' ;
RE_DELIMITER_OPEN: '/' -> pushMode(REGEX) ;
END_CONSTRAINT   : ']' -> popMode ;
