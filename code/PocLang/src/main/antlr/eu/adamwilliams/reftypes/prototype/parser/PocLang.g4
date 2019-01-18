parser grammar PocLang;
options { tokenVocab=PocLex; }
body              : (
                    (WS? body_line? WS? NEWLINE)
                      | (WS? COMMENT_LINE NEWLINE)
                    )*;
function_body     : body ;
body_line         : (var_assignment | return_stmt | var_decl | function_call | if_stmt ) ;

return_stmt       : RETURN (SPACE expr )? ;
if_stmt           : IF SPACE BEGIN_GROUP expr END_GROUP SPACE BEGIN_BODY NEWLINE body END_BODY optional_else ;
optional_else     : SPACE ELSE SPACE BEGIN_BODY NEWLINE body END_BODY | SPACE ELSE SPACE if_stmt | /* ε */ ;
var_decl          : VAR SPACE IDENTIFIER COLON SPACE type_specifier ;
type_keyword      : UINT_T # UnsignedIntType |
                    STRING_T # StringType |
                    VOID_T # VoidType |
                    BOOL_T #BoolType ;
type_specifier    : (type_keyword)
                    | (type_keyword BEGIN_CONSTRAINT int_constraint END_CONSTRAINT)
                    | (type_keyword BEGIN_CONSTRAINT string_constraint END_CONSTRAINT) ;
int_constraint    : LT CONSTRAINT_SPACE CONSTRAINT_UINT # LessThanConstraint
                    | GT CONSTRAINT_SPACE CONSTRAINT_UINT # GreaterThanConstraint
                    | LE CONSTRAINT_SPACE CONSTRAINT_UINT # LessThanEqualsConstraint
                    | GE CONSTRAINT_SPACE CONSTRAINT_UINT # GreaterThanEqualsConstraint ;
string_constraint : RE_DELIMITER_OPEN re RE_DELIMITER_CLOSE ;
function_sig      : FUNCTION SPACE IDENTIFIER BEGIN_GROUP (argument_decl | argument_decl ARG_SEP)* END_GROUP COLON SPACE type_specifier SPACE BEGIN_BODY NEWLINE ;
argument_decl     : IDENTIFIER COLON SPACE type_specifier ;
function          : function_sig function_body END_BODY ;
program           : (function NEWLINE*)+ EOF;
expr              : expr (MULTIPLY|DIVIDE) expr
                    |	expr (ADD|SUBTRACT) expr
                    | value_ref
                    |   expr SPACE (LT_CONDITION|GT_CONDITION|LE_CONDITION|GE_CONDITION|EQ_CONDITION) SPACE expr
                    | function_call
                    |	BEGIN_GROUP expr END_GROUP ;
var_assignment    : IDENTIFIER WS? EQ_ASSIGNMENT expr ;
value_ref         : INT | STRING_LITERAL | TRUE_LIT | FALSE_LIT | identifier_ref ;
identifier_ref    : IDENTIFIER ;
function_call     : IDENTIFIER BEGIN_GROUP (expr (ARG_SEP expr)*)? END_GROUP ;

re             : simple_re union_prime ;
union_prime    : ALTERNATION re | /* ε */ ;

simple_re      : basic_re concat_prime ;
concat_prime   : simple_re | /* ε */ ;

basic_re       : kleene_star | plus | elementary_re ;
kleene_star    : elementary_re STAR ;
plus           : elementary_re PLUS;
elementary_re  : group | DOT | character | range_re ;
group          : BEGIN_RE_GROUP re END_RE_GROUP ;
range_re       : positive_range | negative_range ;

positive_range : BEGIN_RE_RANGE range_items RANGE_TERMINATE ;
negative_range : BEGIN_RE_RANGE RANGE_NEGATE range_items RANGE_TERMINATE ;
lax_character  : RANGE_CHARACTER ;
character      : CHARACTER|MINUS ; // ranges are a bit more lax w.r.t what characters are allowed

range_items    : range_item | range_item range_items ;
range_item     : lax_character RANGE_SEPARATOR lax_character | lax_character;
