grammar Delphi;


// ================== BEGIN - TOKENS ==================

TokenIdentifier         : (Letter | '_') (Letter | Digit | '_')*;
TokenInteger            : DigitSequence;
TokenReal               : DigitSequence '.' DigitSequence;


// ================== BEGIN - SECTION ==================

StringLiteral           : '\'' ('\'\'\'' | ~('\''))* '\'';
ControlLiteral          : Control+;
DigitSequence           : Digit+;
HexDigitSequence        : HexDigit+;


// ================== BEGIN - FRAGMENTS ==================

fragment Digit          : '0' ..'9';
fragment HexDigit       : Digit | 'A' ..'F' | 'a' ..'f';
fragment Letter         : 'A' ..'Z' | 'a' ..'z' | '\u0080' ..'\uFFFE' ~('\uFEFF');
fragment Control        : '#' HexDigitSequence | '#' '$' HexDigitSequence;


// ================== BEGIN - SKIP ==================

WS                      : [ \t\r\n\f]+ -> skip;
UNICODE_BOM             : '\uFEFF' -> skip;
COMMENT                 : ( '//' ~('\r' | '\n')* '\r' ? '\n' 
                        | '(*' .*? '*)' | '{' .*? '}' ) -> skip;

