lexer grammar PocLex;

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