// BNF: http://www.cs.sfu.ca/~cameron/Teaching/384/99-3/regexp-plg.html
// converted to (limited) ANTLR Adam Williams Nov 2018

grammar Regexp;
re_program     : re EOF ;
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
lax_character  : CHARACTER | (DOT|PLUS|STAR|ALTERNATION|MINUS) ;
character      : CHARACTER|MINUS ; // ranges are a bit more lax w.r.t what characters are allowed

range_items    : range_item | range_item range_items ;
range_item     : lax_character MINUS lax_character | lax_character;

// ANTLR does not support rule refs in sets, so no DRY here
MINUS          : '-' ;
DOT            : '.' ;
STAR           : '*' ;
PLUS           : '+' ;
ALTERNATION    : '|' ;
CHARACTER      : ~('\n'|'\r'|'.'|'('|')'|'['|']'|'*'|'+'|'/') | ESCAPED_META;
META_CHAR      : (DOT|PLUS|STAR|ALTERNATION|'['|']'|'/') ;
fragment ESCAPED_FWD_SLASH: '\\/' ;
fragment ESCAPED_META: '\\' META_CHAR ;
