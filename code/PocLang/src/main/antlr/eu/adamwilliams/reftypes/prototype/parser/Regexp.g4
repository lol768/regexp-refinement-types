// BNF: http://www.cs.sfu.ca/~cameron/Teaching/384/99-3/regexp-plg.html
// converted to (limited) ANTLR Adam Williams Nov 2018

grammar Regexp;
program        : re EOF ;
re             : simple_re union_prime ;
union_prime    : ALTERNATION re | /* ε */ ;

simple_re      : basic_re concat_prime ;
concat_prime   : simple_re | /* ε */ ;

basic_re       : kleene_star | plus | elementary_re ;
kleene_star    : elementary_re STAR ;
plus           : elementary_re PLUS;
elementary_re  : group | DOT | character | range ;
group          : '(' re ')' ;
range          : positive_range | negative_range ;

positive_range : '[' range_items ']' ;
negative_range : '[^' range_items ']' ;

range_items    : range_item | range_item range_items ;
range_item     : character '-' character | character;
character      : CHARACTER ;

// ANTLR does not support rule refs in sets, so no DRY here
DOT            : '.' ;
STAR           : '*' ;
PLUS           : '+' ;
ALTERNATION    : '|' ;
CHARACTER      : ~('\n'|'\r'|'.'|'('|')'|'['|']'|'*'|'+') | '\\' META_CHAR;
META_CHAR      : (DOT|PLUS|STAR|ALTERNATION|'['|']') ;
