grammar ECMAScript;

@lexer::members {
    var strict bool
}

// 11: Lexical Grammar

fragment SourceCharacter
    : [\u0000-\u{10FFFF}]
    ;

inputElementDiv
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | DivPunctuator
    | RightBracePunctuator
    ;

inputElementRegExp
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | RightBracePunctuator
    | regularExpressionLiteral
    ;

inputElementRegExpOrTemplateTail
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | regularExpressionLiteral
    | TemplateSubstitutionTail
    ;

inputElementTemplateTail
    : WhiteSpace
    | LineTerminator
    | Comment
    | CommonToken
    | DivPunctuator
    | TemplateSubstitutionTail
    ;

// 11.1: Unicode Format-Control Characters

fragment ZWNJ   : '\u200C'; // Zero Width Non-Joiner
fragment ZWJ    : '\u200D'; // Zero Width Joiner
fragment ZWNBSP : '\uFEFF'; // Zero Width No-Break Space

// 11.2: White Space

fragment TAB  : '\u0009';
fragment VT   : '\u000B';
fragment FF   : '\u000C';
fragment SP   : '\u0020';
fragment NBSP : '\u00A0';
fragment USP
    : '\u1680'
    | '\u2001'
    | '\u2002'
    | '\u2003'
    | '\u2004'
    | '\u2005'
    | '\u2006'
    | '\u2007'
    | '\u2008'
    | '\u2009'
    | '\u200A'
    | '\u202F'
    | '\u205F'
    | '\u3000'
    ;

WhiteSpace
    : TAB
    | VT
    | FF
    | SP
    | NBSP
    | ZWNBSP
    | USP
    ;

// 11.3: Line Terminators

fragment CRLF : '\u000D\u000A'; // does this work?
fragment LF   : '\u000A';
fragment CR   : '\u000D';
fragment LS   : '\u2028';
fragment PS   : '\u2029';

LineTerminator
    : LF
    | CR
    | LS
    | PS
    ;

LineTerminatorSequence
    : LF
    | CRLF
    | CR
    | LS
    | PS
    ;

// 11.4: Comments

Comment
    : MultiLineComment
    | SingleLineComment
    ;

MultiLineComment
    : '/*' MultiLineCommentChars? '*/'
    ;

fragment MultiLineCommentChars
    : MultiLineNotAsteriskChar MultiLineCommentChars?
    | Multiply PostAsteriskCommentChars?
    ;

fragment PostAsteriskCommentChars
    : MultilineNotForwardSlashOrAsteriskChar MultiLineCommentChars?
    | Multiply PostAsteriskCommentChars?
    ;

fragment MultiLineNotAsteriskChar
    : ~[*] // SourceCharacter but not *
    ;

fragment MultilineNotForwardSlashOrAsteriskChar
    : ~[*/] // SourceCharacter but not one of / or *
    ;

SingleLineComment
    : '//' SingleLineCommentChars?
    ;

fragment SingleLineCommentChars
    : SingleLineCommentChar SingleLineCommentChars?
    ;

fragment SingleLineCommentChar
    : ~[\r\n\u2028\u2029]
    ;

// 11.5: Tokens

CommonToken
    : IdentifierName
    | Punctuator
    | NumericLiteral
    | StringLiteral
    | Template
    ;

// 11.6: Names and Keywords

// IdentifierName
//     : IdentifierStart
//     | IdentifierName IdentifierPart
//     ;

IdentifierName
    : IdentifierStart IdentifierPart?
    ;

fragment IdentifierStart
    : UnicodeIDStart
    | '$'
    | '_'
    | '\\' UnicodeEscapeSequence
    ;

fragment IdentifierPart
    : UnicodeIDContinue
    | '$'
    | '\\' UnicodeEscapeSequence
    | ZWNJ
    | ZWJ
    ;

fragment UnicodeIDStart
    : [A-Za-z]
    | '\u00AA'
    | '\u00B5'
    | '\u00BA'
    | '\u00C0'..'\u00D6'
    | '\u00D8'..'\u00F6'
    | '\u00F8'..'\u02C1'
    | '\u02C6'..'\u02D1'
    | '\u02E0'..'\u02E4'
    | '\u02EC'
    | '\u02EE'
    | '\u0370'..'\u0374'
    | '\u0376'
    | '\u0377'
    | '\u037A'..'\u037D'
    | '\u037F'
    | '\u0386'
    | '\u0388'..'\u038A'
    | '\u038C'
    | '\u038E'..'\u03A1'
    | '\u03A3'..'\u03F5'
    | '\u03F7'..'\u0481'
    | '\u048A'..'\u052F'
    | '\u0531'..'\u0556'
    | '\u0559'
    | '\u0560'..'\u0588'
    | '\u05D0'..'\u05EA'
    | '\u05EF'..'\u05F2'
    | '\u0620'..'\u064A'
    | '\u066E'
    | '\u066F'
    | '\u0671'..'\u06D3'
    | '\u06D5'
    | '\u06E5'
    | '\u06E6'
    | '\u06EE'
    | '\u06EF'
    | '\u06FA'..'\u06FC'
    | '\u06FF'
    | '\u0710'
    | '\u0712'..'\u072F'
    | '\u074D'..'\u07A5'
    | '\u07B1'
    | '\u07CA'..'\u07EA'
    | '\u07F4'
    | '\u07F5'
    | '\u07FA'
    | '\u0800'..'\u0815'
    | '\u081A'
    | '\u0824'
    | '\u0828'
    | '\u0840'..'\u0858'
    | '\u0860'..'\u086A'
    | '\u08A0'..'\u08B4'
    | '\u08B6'..'\u08BD'
    | '\u0904'..'\u0939'
    | '\u093D'
    | '\u0950'
    | '\u0958'..'\u0961'
    | '\u0971'..'\u0980'
    | '\u0985'..'\u098C'
    | '\u098F'
    | '\u0990'
    | '\u0993'..'\u09A8'
    | '\u09AA'..'\u09B0'
    | '\u09B2'
    | '\u09B6'..'\u09B9'
    | '\u09BD'
    | '\u09CE'
    | '\u09DC'
    | '\u09DD'
    | '\u09DF'..'\u09E1'
    | '\u09F0'
    | '\u09F1'
    | '\u09FC'
    | '\u0A05'..'\u0A0A'
    | '\u0A0F'
    | '\u0A10'
    | '\u0A13'..'\u0A28'
    | '\u0A2A'..'\u0A30'
    | '\u0A32'
    | '\u0A33'
    | '\u0A35'
    | '\u0A36'
    | '\u0A38'
    | '\u0A39'
    | '\u0A59'..'\u0A5C'
    | '\u0A5E'
    | '\u0A72'..'\u0A74'
    | '\u0A85'..'\u0A8D'
    | '\u0A8F'..'\u0A91'
    | '\u0A93'..'\u0AA8'
    | '\u0AAA'..'\u0AB0'
    | '\u0AB2'
    | '\u0AB3'
    | '\u0AB5'..'\u0AB9'
    | '\u0ABD'
    | '\u0AD0'
    | '\u0AE0'
    | '\u0AE1'
    | '\u0AF9'
    | '\u0B05'..'\u0B0C'
    | '\u0B0F'
    | '\u0B10'
    | '\u0B13'..'\u0B28'
    | '\u0B2A'..'\u0B30'
    | '\u0B32'
    | '\u0B33'
    | '\u0B35'..'\u0B39'
    | '\u0B3D'
    | '\u0B5C'
    | '\u0B5D'
    | '\u0B5F'..'\u0B61'
    | '\u0B71'
    | '\u0B83'
    | '\u0B85'..'\u0B8A'
    | '\u0B8E'..'\u0B90'
    | '\u0B92'..'\u0B95'
    | '\u0B99'
    | '\u0B9A'
    | '\u0B9C'
    | '\u0B9E'
    | '\u0B9F'
    | '\u0BA3'
    | '\u0BA4'
    | '\u0BA8'..'\u0BAA'
    | '\u0BAE'..'\u0BB9'
    | '\u0BD0'
    | '\u0C05'..'\u0C0C'
    | '\u0C0E'..'\u0C10'
    | '\u0C12'..'\u0C28'
    | '\u0C2A'..'\u0C39'
    | '\u0C3D'
    | '\u0C58'..'\u0C5A'
    | '\u0C60'
    | '\u0C61'
    | '\u0C80'
    | '\u0C85'..'\u0C8C'
    | '\u0C8E'..'\u0C90'
    | '\u0C92'..'\u0CA8'
    | '\u0CAA'..'\u0CB3'
    | '\u0CB5'..'\u0CB9'
    | '\u0CBD'
    | '\u0CDE'
    | '\u0CE0'
    | '\u0CE1'
    | '\u0CF1'
    | '\u0CF2'
    | '\u0D05'..'\u0D0C'
    | '\u0D0E'..'\u0D10'
    | '\u0D12'..'\u0D3A'
    | '\u0D3D'
    | '\u0D4E'
    | '\u0D54'..'\u0D56'
    | '\u0D5F'..'\u0D61'
    | '\u0D7A'..'\u0D7F'
    | '\u0D85'..'\u0D96'
    | '\u0D9A'..'\u0DB1'
    | '\u0DB3'..'\u0DBB'
    | '\u0DBD'
    | '\u0DC0'..'\u0DC6'
    | '\u0E01'..'\u0E30'
    | '\u0E32'
    | '\u0E33'
    | '\u0E40'..'\u0E46'
    | '\u0E81'
    | '\u0E82'
    | '\u0E84'
    | '\u0E86'..'\u0E8A'
    | '\u0E8C'..'\u0EA3'
    | '\u0EA5'
    | '\u0EA7'..'\u0EB0'
    | '\u0EB2'
    | '\u0EB3'
    | '\u0EBD'
    | '\u0EC0'..'\u0EC4'
    | '\u0EC6'
    | '\u0EDC'..'\u0EDF'
    | '\u0F00'
    | '\u0F40'..'\u0F47'
    | '\u0F49'..'\u0F6C'
    | '\u0F88'..'\u0F8C'
    | '\u1000'..'\u102A'
    | '\u103F'
    | '\u1050'..'\u1055'
    | '\u105A'..'\u105D'
    | '\u1061'
    | '\u1065'
    | '\u1066'
    | '\u106E'..'\u1070'
    | '\u1075'..'\u1081'
    | '\u108E'
    | '\u10A0'..'\u10C5'
    | '\u10C7'
    | '\u10CD'
    | '\u10D0'..'\u10FA'
    | '\u10FC'..'\u1248'
    | '\u124A'..'\u124D'
    | '\u1250'..'\u1256'
    | '\u1258'
    | '\u125A'..'\u125D'
    | '\u1260'..'\u1288'
    | '\u128A'..'\u128D'
    | '\u1290'..'\u12B0'
    | '\u12B2'..'\u12B5'
    | '\u12B8'..'\u12BE'
    | '\u12C0'
    | '\u12C2'..'\u12C5'
    | '\u12C8'..'\u12D6'
    | '\u12D8'..'\u1310'
    | '\u1312'..'\u1315'
    | '\u1318'..'\u135A'
    | '\u1380'..'\u138F'
    | '\u13A0'..'\u13F5'
    | '\u13F8'..'\u13FD'
    | '\u1401'..'\u166C'
    | '\u166F'..'\u167F'
    | '\u1681'..'\u169A'
    | '\u16A0'..'\u16EA'
    | '\u16EE'..'\u16F8'
    | '\u1700'..'\u170C'
    | '\u170E'..'\u1711'
    | '\u1720'..'\u1731'
    | '\u1740'..'\u1751'
    | '\u1760'..'\u176C'
    | '\u176E'..'\u1770'
    | '\u1780'..'\u17B3'
    | '\u17D7'
    | '\u17DC'
    | '\u1820'..'\u1878'
    | '\u1880'..'\u18A8'
    | '\u18AA'
    | '\u18B0'..'\u18F5'
    | '\u1900'..'\u191E'
    | '\u1950'..'\u196D'
    | '\u1970'..'\u1974'
    | '\u1980'..'\u19AB'
    | '\u19B0'..'\u19C9'
    | '\u1A00'..'\u1A16'
    | '\u1A20'..'\u1A54'
    | '\u1AA7'
    | '\u1B05'..'\u1B33'
    | '\u1B45'..'\u1B4B'
    | '\u1B83'..'\u1BA0'
    | '\u1BAE'
    | '\u1BAF'
    | '\u1BBA'..'\u1BE5'
    | '\u1C00'..'\u1C23'
    | '\u1C4D'..'\u1C4F'
    | '\u1C5A'..'\u1C7D'
    | '\u1C80'..'\u1C88'
    | '\u1C90'..'\u1CBA'
    | '\u1CBD'..'\u1CBF'
    | '\u1CE9'..'\u1CEC'
    | '\u1CEE'..'\u1CF3'
    | '\u1CF5'
    | '\u1CF6'
    | '\u1CFA'
    | '\u1D00'..'\u1DBF'
    | '\u1E00'..'\u1F15'
    | '\u1F18'..'\u1F1D'
    | '\u1F20'..'\u1F45'
    | '\u1F48'..'\u1F4D'
    | '\u1F50'..'\u1F57'
    | '\u1F59'
    | '\u1F5B'
    | '\u1F5D'
    | '\u1F5F'..'\u1F7D'
    | '\u1F80'..'\u1FB4'
    | '\u1FB6'..'\u1FBC'
    | '\u1FBE'
    | '\u1FC2'..'\u1FC4'
    | '\u1FC6'..'\u1FCC'
    | '\u1FD0'..'\u1FD3'
    | '\u1FD6'..'\u1FDB'
    | '\u1FE0'..'\u1FEC'
    | '\u1FF2'..'\u1FF4'
    | '\u1FF6'..'\u1FFC'
    | '\u2071'
    | '\u207F'
    | '\u2090'..'\u209C'
    | '\u2102'
    | '\u2107'
    | '\u210A'..'\u2113'
    | '\u2115'
    | '\u2118'..'\u211D'
    | '\u2124'
    | '\u2126'
    | '\u2128'
    | '\u212A'..'\u2139'
    | '\u213C'..'\u213F'
    | '\u2145'..'\u2149'
    | '\u214E'
    | '\u2160'..'\u2188'
    | '\u2C00'..'\u2C2E'
    | '\u2C30'..'\u2C5E'
    | '\u2C60'..'\u2CE4'
    | '\u2CEB'..'\u2CEE'
    | '\u2CF2'
    | '\u2CF3'
    | '\u2D00'..'\u2D25'
    | '\u2D27'
    | '\u2D2D'
    | '\u2D30'..'\u2D67'
    | '\u2D6F'
    | '\u2D80'..'\u2D96'
    | '\u2DA0'..'\u2DA6'
    | '\u2DA8'..'\u2DAE'
    | '\u2DB0'..'\u2DB6'
    | '\u2DB8'..'\u2DBE'
    | '\u2DC0'..'\u2DC6'
    | '\u2DC8'..'\u2DCE'
    | '\u2DD0'..'\u2DD6'
    | '\u2DD8'..'\u2DDE'
    | '\u3005'..'\u3007'
    | '\u3021'..'\u3029'
    | '\u3031'..'\u3035'
    | '\u3038'..'\u303C'
    | '\u3041'..'\u3096'
    | '\u309B'..'\u309F'
    | '\u30A1'..'\u30FA'
    | '\u30FC'..'\u30FF'
    | '\u3105'..'\u312F'
    | '\u3131'..'\u318E'
    | '\u31A0'..'\u31BA'
    | '\u31F0'..'\u31FF'
    | '\u3400'..'\u4DB5'
    | '\u4E00'..'\u9FEF'
    | '\uA000'..'\uA48C'
    | '\uA4D0'..'\uA4FD'
    | '\uA500'..'\uA60C'
    | '\uA610'..'\uA61F'
    | '\uA62A'
    | '\uA62B'
    | '\uA640'..'\uA66E'
    | '\uA67F'..'\uA69D'
    | '\uA6A0'..'\uA6EF'
    | '\uA717'..'\uA71F'
    | '\uA722'..'\uA788'
    | '\uA78B'..'\uA7BF'
    | '\uA7C2'..'\uA7C6'
    | '\uA7F7'..'\uA801'
    | '\uA803'..'\uA805'
    | '\uA807'..'\uA80A'
    | '\uA80C'..'\uA822'
    | '\uA840'..'\uA873'
    | '\uA882'..'\uA8B3'
    | '\uA8F2'..'\uA8F7'
    | '\uA8FB'
    | '\uA8FD'
    | '\uA8FE'
    | '\uA90A'..'\uA925'
    | '\uA930'..'\uA946'
    | '\uA960'..'\uA97C'
    | '\uA984'..'\uA9B2'
    | '\uA9CF'
    | '\uA9E0'..'\uA9E4'
    | '\uA9E6'..'\uA9EF'
    | '\uA9FA'..'\uA9FE'
    | '\uAA00'..'\uAA28'
    | '\uAA40'..'\uAA42'
    | '\uAA44'..'\uAA4B'
    | '\uAA60'..'\uAA76'
    | '\uAA7A'
    | '\uAA7E'..'\uAAAF'
    | '\uAAB1'
    | '\uAAB5'
    | '\uAAB6'
    | '\uAAB9'..'\uAABD'
    | '\uAAC0'
    | '\uAAC2'
    | '\uAADB'..'\uAADD'
    | '\uAAE0'..'\uAAEA'
    | '\uAAF2'..'\uAAF4'
    | '\uAB01'..'\uAB06'
    | '\uAB09'..'\uAB0E'
    | '\uAB11'..'\uAB16'
    | '\uAB20'..'\uAB26'
    | '\uAB28'..'\uAB2E'
    | '\uAB30'..'\uAB5A'
    | '\uAB5C'..'\uAB67'
    | '\uAB70'..'\uABE2'
    | '\uAC00'..'\uD7A3'
    | '\uD7B0'..'\uD7C6'
    | '\uD7CB'..'\uD7FB'
    | '\uF900'..'\uFA6D'
    | '\uFA70'..'\uFAD9'
    | '\uFB00'..'\uFB06'
    | '\uFB13'..'\uFB17'
    | '\uFB1D'
    | '\uFB1F'..'\uFB28'
    | '\uFB2A'..'\uFB36'
    | '\uFB38'..'\uFB3C'
    | '\uFB3E'
    | '\uFB40'
    | '\uFB41'
    | '\uFB43'
    | '\uFB44'
    | '\uFB46'..'\uFBB1'
    | '\uFBD3'..'\uFD3D'
    | '\uFD50'..'\uFD8F'
    | '\uFD92'..'\uFDC7'
    | '\uFDF0'..'\uFDFB'
    | '\uFE70'..'\uFE74'
    | '\uFE76'..'\uFEFC'
    | '\uFF21'..'\uFF3A'
    | '\uFF41'..'\uFF5A'
    | '\uFF66'..'\uFFBE'
    | '\uFFC2'..'\uFFC7'
    | '\uFFCA'..'\uFFCF'
    | '\uFFD2'..'\uFFD7'
    | '\uFFDA'..'\uFFDC'
    ;

fragment UnicodeIDContinue
    : [0-9A-Z_a-z]
    | '\u00AA'
    | '\u00B5'
    | '\u00B7'
    | '\u00BA'
    | '\u00C0'..'\u00D6'
    | '\u00D8'..'\u00F6'
    | '\u00F8'..'\u02C1'
    | '\u02C6'..'\u02D1'
    | '\u02E0'..'\u02E4'
    | '\u02EC'
    | '\u02EE'
    | '\u0300'..'\u0374'
    | '\u0376'
    | '\u0377'
    | '\u037A'..'\u037D'
    | '\u037F'
    | '\u0386'..'\u038A'
    | '\u038C'
    | '\u038E'..'\u03A1'
    | '\u03A3'..'\u03F5'
    | '\u03F7'..'\u0481'
    | '\u0483'..'\u0487'
    | '\u048A'..'\u052F'
    | '\u0531'..'\u0556'
    | '\u0559'
    | '\u0560'..'\u0588'
    | '\u0591'..'\u05BD'
    | '\u05BF'
    | '\u05C1'
    | '\u05C2'
    | '\u05C4'
    | '\u05C5'
    | '\u05C7'
    | '\u05D0'..'\u05EA'
    | '\u05EF'..'\u05F2'
    | '\u0610'..'\u061A'
    | '\u0620'..'\u0669'
    | '\u066E'..'\u06D3'
    | '\u06D5'..'\u06DC'
    | '\u06DF'..'\u06E8'
    | '\u06EA'..'\u06FC'
    | '\u06FF'
    | '\u0710'..'\u074A'
    | '\u074D'..'\u07B1'
    | '\u07C0'..'\u07F5'
    | '\u07FA'
    | '\u07FD'
    | '\u0800'..'\u082D'
    | '\u0840'..'\u085B'
    | '\u0860'..'\u086A'
    | '\u08A0'..'\u08B4'
    | '\u08B6'..'\u08BD'
    | '\u08D3'..'\u08E1'
    | '\u08E3'..'\u0963'
    | '\u0966'..'\u096F'
    | '\u0971'..'\u0983'
    | '\u0985'..'\u098C'
    | '\u098F'
    | '\u0990'
    | '\u0993'..'\u09A8'
    | '\u09AA'..'\u09B0'
    | '\u09B2'
    | '\u09B6'..'\u09B9'
    | '\u09BC'..'\u09C4'
    | '\u09C7'
    | '\u09C8'
    | '\u09CB'..'\u09CE'
    | '\u09D7'
    | '\u09DC'
    | '\u09DD'
    | '\u09DF'..'\u09E3'
    | '\u09E6'..'\u09F1'
    | '\u09FC'
    | '\u09FE'
    | '\u0A01'..'\u0A03'
    | '\u0A05'..'\u0A0A'
    | '\u0A0F'
    | '\u0A10'
    | '\u0A13'..'\u0A28'
    | '\u0A2A'..'\u0A30'
    | '\u0A32'
    | '\u0A33'
    | '\u0A35'
    | '\u0A36'
    | '\u0A38'
    | '\u0A39'
    | '\u0A3C'
    | '\u0A3E'..'\u0A42'
    | '\u0A47'
    | '\u0A48'
    | '\u0A4B'..'\u0A4D'
    | '\u0A51'
    | '\u0A59'..'\u0A5C'
    | '\u0A5E'
    | '\u0A66'..'\u0A75'
    | '\u0A81'..'\u0A83'
    | '\u0A85'..'\u0A8D'
    | '\u0A8F'..'\u0A91'
    | '\u0A93'..'\u0AA8'
    | '\u0AAA'..'\u0AB0'
    | '\u0AB2'
    | '\u0AB3'
    | '\u0AB5'..'\u0AB9'
    | '\u0ABC'..'\u0AC5'
    | '\u0AC7'..'\u0AC9'
    | '\u0ACB'..'\u0ACD'
    | '\u0AD0'
    | '\u0AE0'..'\u0AE3'
    | '\u0AE6'..'\u0AEF'
    | '\u0AF9'..'\u0AFF'
    | '\u0B01'..'\u0B03'
    | '\u0B05'..'\u0B0C'
    | '\u0B0F'
    | '\u0B10'
    | '\u0B13'..'\u0B28'
    | '\u0B2A'..'\u0B30'
    | '\u0B32'
    | '\u0B33'
    | '\u0B35'..'\u0B39'
    | '\u0B3C'..'\u0B44'
    | '\u0B47'
    | '\u0B48'
    | '\u0B4B'..'\u0B4D'
    | '\u0B56'
    | '\u0B57'
    | '\u0B5C'
    | '\u0B5D'
    | '\u0B5F'..'\u0B63'
    | '\u0B66'..'\u0B6F'
    | '\u0B71'
    | '\u0B82'
    | '\u0B83'
    | '\u0B85'..'\u0B8A'
    | '\u0B8E'..'\u0B90'
    | '\u0B92'..'\u0B95'
    | '\u0B99'
    | '\u0B9A'
    | '\u0B9C'
    | '\u0B9E'
    | '\u0B9F'
    | '\u0BA3'
    | '\u0BA4'
    | '\u0BA8'..'\u0BAA'
    | '\u0BAE'..'\u0BB9'
    | '\u0BBE'..'\u0BC2'
    | '\u0BC6'..'\u0BC8'
    | '\u0BCA'..'\u0BCD'
    | '\u0BD0'
    | '\u0BD7'
    | '\u0BE6'..'\u0BEF'
    | '\u0C00'..'\u0C0C'
    | '\u0C0E'..'\u0C10'
    | '\u0C12'..'\u0C28'
    | '\u0C2A'..'\u0C39'
    | '\u0C3D'..'\u0C44'
    | '\u0C46'..'\u0C48'
    | '\u0C4A'..'\u0C4D'
    | '\u0C55'
    | '\u0C56'
    | '\u0C58'..'\u0C5A'
    | '\u0C60'..'\u0C63'
    | '\u0C66'..'\u0C6F'
    | '\u0C80'..'\u0C83'
    | '\u0C85'..'\u0C8C'
    | '\u0C8E'..'\u0C90'
    | '\u0C92'..'\u0CA8'
    | '\u0CAA'..'\u0CB3'
    | '\u0CB5'..'\u0CB9'
    | '\u0CBC'..'\u0CC4'
    | '\u0CC6'..'\u0CC8'
    | '\u0CCA'..'\u0CCD'
    | '\u0CD5'
    | '\u0CD6'
    | '\u0CDE'
    | '\u0CE0'..'\u0CE3'
    | '\u0CE6'..'\u0CEF'
    | '\u0CF1'
    | '\u0CF2'
    | '\u0D00'..'\u0D03'
    | '\u0D05'..'\u0D0C'
    | '\u0D0E'..'\u0D10'
    | '\u0D12'..'\u0D44'
    | '\u0D46'..'\u0D48'
    | '\u0D4A'..'\u0D4E'
    | '\u0D54'..'\u0D57'
    | '\u0D5F'..'\u0D63'
    | '\u0D66'..'\u0D6F'
    | '\u0D7A'..'\u0D7F'
    | '\u0D82'
    | '\u0D83'
    | '\u0D85'..'\u0D96'
    | '\u0D9A'..'\u0DB1'
    | '\u0DB3'..'\u0DBB'
    | '\u0DBD'
    | '\u0DC0'..'\u0DC6'
    | '\u0DCA'
    | '\u0DCF'..'\u0DD4'
    | '\u0DD6'
    | '\u0DD8'..'\u0DDF'
    | '\u0DE6'..'\u0DEF'
    | '\u0DF2'
    | '\u0DF3'
    | '\u0E01'..'\u0E3A'
    | '\u0E40'..'\u0E4E'
    | '\u0E50'..'\u0E59'
    | '\u0E81'
    | '\u0E82'
    | '\u0E84'
    | '\u0E86'..'\u0E8A'
    | '\u0E8C'..'\u0EA3'
    | '\u0EA5'
    | '\u0EA7'..'\u0EBD'
    | '\u0EC0'..'\u0EC4'
    | '\u0EC6'
    | '\u0EC8'..'\u0ECD'
    | '\u0ED0'..'\u0ED9'
    | '\u0EDC'..'\u0EDF'
    | '\u0F00'
    | '\u0F18'
    | '\u0F19'
    | '\u0F20'..'\u0F29'
    | '\u0F35'
    | '\u0F37'
    | '\u0F39'
    | '\u0F3E'..'\u0F47'
    | '\u0F49'..'\u0F6C'
    | '\u0F71'..'\u0F84'
    | '\u0F86'..'\u0F97'
    | '\u0F99'..'\u0FBC'
    | '\u0FC6'
    | '\u1000'..'\u1049'
    | '\u1050'..'\u109D'
    | '\u10A0'..'\u10C5'
    | '\u10C7'
    | '\u10CD'
    | '\u10D0'..'\u10FA'
    | '\u10FC'..'\u1248'
    | '\u124A'..'\u124D'
    | '\u1250'..'\u1256'
    | '\u1258'
    | '\u125A'..'\u125D'
    | '\u1260'..'\u1288'
    | '\u128A'..'\u128D'
    | '\u1290'..'\u12B0'
    | '\u12B2'..'\u12B5'
    | '\u12B8'..'\u12BE'
    | '\u12C0'
    | '\u12C2'..'\u12C5'
    | '\u12C8'..'\u12D6'
    | '\u12D8'..'\u1310'
    | '\u1312'..'\u1315'
    | '\u1318'..'\u135A'
    | '\u135D'..'\u135F'
    | '\u1369'..'\u1371'
    | '\u1380'..'\u138F'
    | '\u13A0'..'\u13F5'
    | '\u13F8'..'\u13FD'
    | '\u1401'..'\u166C'
    | '\u166F'..'\u167F'
    | '\u1681'..'\u169A'
    | '\u16A0'..'\u16EA'
    | '\u16EE'..'\u16F8'
    | '\u1700'..'\u170C'
    | '\u170E'..'\u1714'
    | '\u1720'..'\u1734'
    | '\u1740'..'\u1753'
    | '\u1760'..'\u176C'
    | '\u176E'..'\u1770'
    | '\u1772'
    | '\u1773'
    | '\u1780'..'\u17D3'
    | '\u17D7'
    | '\u17DC'
    | '\u17DD'
    | '\u17E0'..'\u17E9'
    | '\u180B'..'\u180D'
    | '\u1810'..'\u1819'
    | '\u1820'..'\u1878'
    | '\u1880'..'\u18AA'
    | '\u18B0'..'\u18F5'
    | '\u1900'..'\u191E'
    | '\u1920'..'\u192B'
    | '\u1930'..'\u193B'
    | '\u1946'..'\u196D'
    | '\u1970'..'\u1974'
    | '\u1980'..'\u19AB'
    | '\u19B0'..'\u19C9'
    | '\u19D0'..'\u19DA'
    | '\u1A00'..'\u1A1B'
    | '\u1A20'..'\u1A5E'
    | '\u1A60'..'\u1A7C'
    | '\u1A7F'..'\u1A89'
    | '\u1A90'..'\u1A99'
    | '\u1AA7'
    | '\u1AB0'..'\u1ABD'
    | '\u1B00'..'\u1B4B'
    | '\u1B50'..'\u1B59'
    | '\u1B6B'..'\u1B73'
    | '\u1B80'..'\u1BF3'
    | '\u1C00'..'\u1C37'
    | '\u1C40'..'\u1C49'
    | '\u1C4D'..'\u1C7D'
    | '\u1C80'..'\u1C88'
    | '\u1C90'..'\u1CBA'
    | '\u1CBD'..'\u1CBF'
    | '\u1CD0'..'\u1CD2'
    | '\u1CD4'..'\u1CFA'
    | '\u1D00'..'\u1DF9'
    | '\u1DFB'..'\u1F15'
    | '\u1F18'..'\u1F1D'
    | '\u1F20'..'\u1F45'
    | '\u1F48'..'\u1F4D'
    | '\u1F50'..'\u1F57'
    | '\u1F59'
    | '\u1F5B'
    | '\u1F5D'
    | '\u1F5F'..'\u1F7D'
    | '\u1F80'..'\u1FB4'
    | '\u1FB6'..'\u1FBC'
    | '\u1FBE'
    | '\u1FC2'..'\u1FC4'
    | '\u1FC6'..'\u1FCC'
    | '\u1FD0'..'\u1FD3'
    | '\u1FD6'..'\u1FDB'
    | '\u1FE0'..'\u1FEC'
    | '\u1FF2'..'\u1FF4'
    | '\u1FF6'..'\u1FFC'
    | '\u203F'
    | '\u2040'
    | '\u2054'
    | '\u2071'
    | '\u207F'
    | '\u2090'..'\u209C'
    | '\u20D0'..'\u20DC'
    | '\u20E1'
    | '\u20E5'..'\u20F0'
    | '\u2102'
    | '\u2107'
    | '\u210A'..'\u2113'
    | '\u2115'
    | '\u2118'..'\u211D'
    | '\u2124'
    | '\u2126'
    | '\u2128'
    | '\u212A'..'\u2139'
    | '\u213C'..'\u213F'
    | '\u2145'..'\u2149'
    | '\u214E'
    | '\u2160'..'\u2188'
    | '\u2C00'..'\u2C2E'
    | '\u2C30'..'\u2C5E'
    | '\u2C60'..'\u2CE4'
    | '\u2CEB'..'\u2CF3'
    | '\u2D00'..'\u2D25'
    | '\u2D27'
    | '\u2D2D'
    | '\u2D30'..'\u2D67'
    | '\u2D6F'
    | '\u2D7F'..'\u2D96'
    | '\u2DA0'..'\u2DA6'
    | '\u2DA8'..'\u2DAE'
    | '\u2DB0'..'\u2DB6'
    | '\u2DB8'..'\u2DBE'
    | '\u2DC0'..'\u2DC6'
    | '\u2DC8'..'\u2DCE'
    | '\u2DD0'..'\u2DD6'
    | '\u2DD8'..'\u2DDE'
    | '\u2DE0'..'\u2DFF'
    | '\u3005'..'\u3007'
    | '\u3021'..'\u302F'
    | '\u3031'..'\u3035'
    | '\u3038'..'\u303C'
    | '\u3041'..'\u3096'
    | '\u3099'..'\u309F'
    | '\u30A1'..'\u30FA'
    | '\u30FC'..'\u30FF'
    | '\u3105'..'\u312F'
    | '\u3131'..'\u318E'
    | '\u31A0'..'\u31BA'
    | '\u31F0'..'\u31FF'
    | '\u3400'..'\u4DB5'
    | '\u4E00'..'\u9FEF'
    | '\uA000'..'\uA48C'
    | '\uA4D0'..'\uA4FD'
    | '\uA500'..'\uA60C'
    | '\uA610'..'\uA62B'
    | '\uA640'..'\uA66F'
    | '\uA674'..'\uA67D'
    | '\uA67F'..'\uA6F1'
    | '\uA717'..'\uA71F'
    | '\uA722'..'\uA788'
    | '\uA78B'..'\uA7BF'
    | '\uA7C2'..'\uA7C6'
    | '\uA7F7'..'\uA827'
    | '\uA840'..'\uA873'
    | '\uA880'..'\uA8C5'
    | '\uA8D0'..'\uA8D9'
    | '\uA8E0'..'\uA8F7'
    | '\uA8FB'
    | '\uA8FD'..'\uA92D'
    | '\uA930'..'\uA953'
    | '\uA960'..'\uA97C'
    | '\uA980'..'\uA9C0'
    | '\uA9CF'..'\uA9D9'
    | '\uA9E0'..'\uA9FE'
    | '\uAA00'..'\uAA36'
    | '\uAA40'..'\uAA4D'
    | '\uAA50'..'\uAA59'
    | '\uAA60'..'\uAA76'
    | '\uAA7A'..'\uAAC2'
    | '\uAADB'..'\uAADD'
    | '\uAAE0'..'\uAAEF'
    | '\uAAF2'..'\uAAF6'
    | '\uAB01'..'\uAB06'
    | '\uAB09'..'\uAB0E'
    | '\uAB11'..'\uAB16'
    | '\uAB20'..'\uAB26'
    | '\uAB28'..'\uAB2E'
    | '\uAB30'..'\uAB5A'
    | '\uAB5C'..'\uAB67'
    | '\uAB70'..'\uABEA'
    | '\uABEC'
    | '\uABED'
    | '\uABF0'..'\uABF9'
    | '\uAC00'..'\uD7A3'
    | '\uD7B0'..'\uD7C6'
    | '\uD7CB'..'\uD7FB'
    | '\uF900'..'\uFA6D'
    | '\uFA70'..'\uFAD9'
    | '\uFB00'..'\uFB06'
    | '\uFB13'..'\uFB17'
    | '\uFB1D'..'\uFB28'
    | '\uFB2A'..'\uFB36'
    | '\uFB38'..'\uFB3C'
    | '\uFB3E'
    | '\uFB40'
    | '\uFB41'
    | '\uFB43'
    | '\uFB44'
    | '\uFB46'..'\uFBB1'
    | '\uFBD3'..'\uFD3D'
    | '\uFD50'..'\uFD8F'
    | '\uFD92'..'\uFDC7'
    | '\uFDF0'..'\uFDFB'
    | '\uFE00'..'\uFE0F'
    | '\uFE20'..'\uFE2F'
    | '\uFE33'
    | '\uFE34'
    | '\uFE4D'..'\uFE4F'
    | '\uFE70'..'\uFE74'
    | '\uFE76'..'\uFEFC'
    | '\uFF10'..'\uFF19'
    | '\uFF21'..'\uFF3A'
    | '\uFF3F'
    | '\uFF41'..'\uFF5A'
    | '\uFF66'..'\uFFBE'
    | '\uFFC2'..'\uFFC7'
    | '\uFFCA'..'\uFFCF'
    | '\uFFD2'..'\uFFD7'
    | '\uFFDA'..'\uFFDC'
    ;

// 11.6.2: Reserved Words

ReservedWord
    : Keyword
    | FutureReservedWord
    | NullLiteral
    | BooleanLiteral
    ;

// 11.6.2.1: Keywords

Keyword
    : Await
    | Break
    | Case
    | Catch
    | Class
    | Const
    | Continue
    | Debugger
    | Default
    | Delete
    | Do
    | Else
    | Export
    | Extends
    | Finally
    | For
    | Function
    | If
    | Import
    | In
    | Instanceof
    | { strict }? Let
    | New
    | Return
    | { strict }? Static
    | Super
    | Switch
    | This
    | Throw
    | Try
    | Typeof
    | Var
    | Void
    | While
    | With
    | Yield
    ;

// 11.6.2.2: Future Reserved Words

FutureReservedWord
    : 'enum'
    | { strict }? 'implements'
    | { strict }? 'interface'
    | { strict }? 'package'
    | { strict }? 'private'
    | { strict }? 'protected'
    | { strict }? 'public'
    ;

// Symbols

OpenBracket                : '[';
CloseBracket               : ']';
OpenParen                  : '(';
CloseParen                 : ')';
OpenBrace                  : '{';
CloseBrace                 : '}';
SemiColon                  : ';';
Comma                      : ',';
Assign                     : '=';
QuestionMark               : '?';
Colon                      : ':';
Dot                        : '.';
PlusPlus                   : '++';
MinusMinus                 : '--';
Plus                       : '+';
Minus                      : '-';
BitNot                     : '~';
Not                        : '!';
Multiply                   : '*';
Power                      : '**';
Divide                     : '/';
Modulus                    : '%';
RightShiftArithmetic       : '>>';
LeftShiftArithmetic        : '<<';
RightShiftLogical          : '>>>';
LessThan                   : '<';
MoreThan                   : '>';
LessThanEquals             : '<=';
MoreThanEquals             : '>=';
Equals                     : '==';
NotEquals                  : '!=';
IdentityEquals             : '===';
IdentityNotEquals          : '!==';
BitAnd                     : '&';
BitXor                     : '^';
BitOr                      : '|';
And                        : '&&';
Or                         : '||';
MultiplyAssign             : '*=';
PowerAssign                : '**=';
DivideAssign               : '/='; 
ModulusAssign              : '%='; 
PlusAssign                 : '+='; 
MinusAssign                : '-='; 
LeftShiftArithmeticAssign  : '<<='; 
RightShiftArithmeticAssign : '>>='; 
RightShiftLogicalAssign    : '>>>='; 
BitAndAssign               : '&='; 
BitXorAssign               : '^='; 
BitOrAssign                : '|=';
Arrow                      : '=>';
Ellipsis                   : '...';

Let                        : 'let';
LetBracket                 : 'let [';
For                        : 'for';
Async                      : 'async';
Function                   : 'function';
AsyncFunction              : 'async function';
Class                      : 'class';
Export                     : 'export';
Default                    : 'default';
From                       : 'from';
In                         : 'in';
As                         : 'as';
Import                     : 'import';
Await                      : 'await';
Break                      : 'break';
Case                       : 'case';
Catch                      : 'catch';
Const                      : 'const';
Continue                   : 'continue';
Debugger                   : 'debugger';
Delete                     : 'delete';
Do                         : 'do';
Else                       : 'else';
Extends                    : 'extends';
Finally                    : 'finally';
Get                        : 'get';
If                         : 'if';
Instanceof                 : 'instanceof';
New                        : 'new';
Of                         : 'of';
Return                     : 'return';
Set                        : 'set';
Static                     : 'static';
Super                      : 'super';
Switch                     : 'switch';
Target                     : 'target';
This                       : 'this';
Throw                      : 'throw';
Try                        : 'try';
Typeof                     : 'typeof';
Var                        : 'var';
Void                       : 'void';
While                      : 'while';
With                       : 'with';
Yield                      : 'yield';

// 11.7: Punctuators

Punctuator
    : OpenBrace 
    | OpenParen 
    | CloseParen 
    | OpenBracket 
    | CloseBracket 
    | Dot 
    | Ellipsis 
    | SemiColon 
    | Comma 
    | LessThan 
    | MoreThan 
    | LessThanEquals 
    | MoreThanEquals 
    | Equals 
    | NotEquals 
    | IdentityEquals 
    | IdentityNotEquals 
    | Plus 
    | Minus 
    | Multiply 
    | Modulus 
    | Power 
    | PlusPlus 
    | MinusMinus 
    | LeftShiftArithmetic 
    | RightShiftArithmetic 
    | RightShiftLogical 
    | BitAnd 
    | BitOr 
    | BitXor 
    | Not 
    | BitNot 
    | And 
    | Or 
    | QuestionMark 
    | Colon 
    | Assign 
    | PlusAssign 
    | MinusAssign 
    | MultiplyAssign 
    | ModulusAssign 
    | PowerAssign 
    | LeftShiftArithmeticAssign 
    | RightShiftArithmeticAssign 
    | RightShiftLogicalAssign 
    | BitAndAssign 
    | BitOrAssign 
    | BitXorAssign 
    | Arrow
    ;

DivPunctuator
    : Divide
    | DivideAssign
    ;

RightBracePunctuator
    : CloseBrace
    ;

// 11.8: Literals

// 11.8.1: Null Literals

NullLiteral
    : 'null'
    ;

// 11.8.2: Boolean Literals

BooleanLiteral
    : 'true'
    | 'false'
    ;

// 11.8.3 Numeric Literals

NumericLiteral
    : DecimalLiteral
    | BinaryIntegerLiteral
    | OctalIntegerLiteral
    | HexIntegerLiteral
    ;

DecimalLiteral
    : DecimalIntegerLiteral Dot DecimalDigits? ExponentPart?
    | Dot DecimalDigits ExponentPart?
    | DecimalIntegerLiteral ExponentPart?
    ;

fragment DecimalIntegerLiteral
    : '0'
    | NonZeroDigit DecimalDigits?
    ;

fragment DecimalDigits
    : DecimalDigit+
    ;

DecimalDigit
    : [0-9]
    ;

fragment NonZeroDigit
    : [1-9]
    ;

fragment ExponentPart
    : ExponentIndicator SignedInteger
    ;

fragment ExponentIndicator
    : 'e'
    | 'E'
    ;

fragment SignedInteger
    : DecimalDigits
    | Plus DecimalDigits
    | Minus DecimalDigits
    ;

BinaryIntegerLiteral
    : '0b' BinaryDigits
    | '0B' BinaryDigits
    ;

fragment BinaryDigits
    : BinaryDigit+
    ;

fragment BinaryDigit
    : '0'
    | '1'
    ;

OctalIntegerLiteral
    : '0o' OctalDigits
    | '0O' OctalDigits
    ;

fragment OctalDigits
    : OctalDigit+
    ;

fragment OctalDigit
    : [0-7]
    ;

HexIntegerLiteral
    : '0x' HexDigits
    | '0X' HexDigits
    ;

fragment HexDigits
    : HexDigit+
    ;

HexDigit
    : [0-9a-zA-Z]
    ;

// 11.8.4: String Literals

StringLiteral
    : '"' DoubleStringCharacters '"'
    | '\'' SingleStringCharacters '\''
    ;

fragment DoubleStringCharacters
    : DoubleStringCharacter DoubleStringCharacters?
    ;

fragment SingleStringCharacters
    : SingleStringCharacter SingleStringCharacters?
    ;

fragment DoubleStringCharacter
    : ~["\\\r\n\u2028\u2029]
    | '\\' EscapeSequence
    | LineContinuation
    ;

fragment SingleStringCharacter
    : ~['\\\r\n\u2028\u2029]
    | '\\' EscapeSequence
    | LineContinuation
    ;

fragment LineContinuation
    : '\\' LineTerminatorSequence
    ;

fragment EscapeSequence
    : CharacterEscapeSequence
    | '0' { p.negativeLookahead(ECMAScriptLexerDecimalDigit) }?
    | HexEscapeSequence
    | UnicodeEscapeSequence
    ;

fragment CharacterEscapeSequence
    : SingleEscapeCharacter
    | NonEscapeCharacter
    ;

fragment SingleEscapeCharacter
    : '\''
    | '"'
    | '\\'
    | 'b'
    | 'f'
    | 'n'
    | 'r'
    | 't'
    | 'v'
    ;

fragment NonEscapeCharacter
    : ~['"\bfnrtv0-9xu\r\n\u2028\u2029] // SourceCharacter but not one of EscapeCharacter or LineTerminator
    ;

fragment EscapeCharacter
    : SingleEscapeCharacter
    | DecimalDigit
    | 'x'
    | 'u'
    ;

fragment HexEscapeSequence
    : 'x' HexDigit HexDigit
    ;

fragment UnicodeEscapeSequence
    : 'u' Hex4Digits
    | 'u{' CodePoint CloseBrace
    ;

fragment Hex4Digits
    : HexDigit HexDigit HexDigit HexDigit
    ;

// 11.8.5: Regular Expression Literals

regularExpressionLiteral
    : Divide RegularExpressionBody Divide RegularExpressionFlags*
    ;

RegularExpressionBody
    : RegularExpressionFirstChar RegularExpressionChars*
    ;

RegularExpressionChars
    : RegularExpressionChar
    ;

RegularExpressionFirstChar
    : ~[*/[\r\n\u2028\u2029] // RegularExpressionNonTerminator but not one of * or \ or / or [
    | RegularExpressionBackslashSequence
    | RegularExpressionClass
    ;

RegularExpressionChar
    : ~[/[\r\n\u2028\u2029] // RegularExpressionNonTerminator but not one of \ or / or [
    | RegularExpressionBackslashSequence
    | RegularExpressionClass
    ;

RegularExpressionBackslashSequence
    : '\\' RegularExpressionNonTerminator
    ;

RegularExpressionNonTerminator
    : ~[\r\n\u2028\u2029]
    ;

RegularExpressionClass
    : OpenBracket RegularExpressionClassChars* CloseBracket
    ;

RegularExpressionClassChars
    : RegularExpressionClassChar
    ;

RegularExpressionClassChar
    : ~[\r\n\u2028\u2029\]\\]
    | RegularExpressionBackslashSequence
    ;

RegularExpressionFlags
    : IdentifierPart
    ;

// 11.8.6: Template Literal Lexical Components

Template
    : NoSubstitutionTemplate
    | TemplateHead
    ;

NoSubstitutionTemplate
    : '`' TemplateCharacters? '`'
    ;

TemplateHead
    : '`' TemplateCharacters? '${'
    ;

TemplateSubstitutionTail
    : TemplateMiddle
    | TemplateTail
    ;

TemplateMiddle
    : CloseBrace TemplateCharacters? '${'
    ;

TemplateTail
    : CloseBrace TemplateCharacters? '`'
    ;

TemplateCharacters
    : TemplateCharacter TemplateCharacters?
    ;

TemplateCharacter
    : '$' { p.negativeLookahead(ECMAScriptLexerOpenBrace) }?
    | '\\' EscapeSequence
    | '\\' NotEscapeSequence
    | LineContinuation
    | LineTerminatorSequence
    | ~[`\\$\r\n\u2028\u2029]
    ;

NotEscapeSequence
    : '0' DecimalDigit
    | [1-9]
    | 'x' { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'x' HexDigit { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' { p.negativeLookahead(ECMAScriptLexerHexDigit) && p.negativeLookahead(ECMAScriptLexerOpenBrace) }?
    | 'u' HexDigit { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' HexDigit HexDigit { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' HexDigit HexDigit HexDigit { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' OpenBrace { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' OpenBrace NotCodePoint { p.negativeLookahead(ECMAScriptLexerHexDigit) }?
    | 'u' OpenBrace CodePoint { p.negativeLookahead(ECMAScriptLexerHexDigit) && p.negativeLookahead(ECMAScriptLexerOpenBrace) }?
    ;

NotCodePoint
    : HexDigits { p.notCodePoint(p.GetText()) }?
    ;

CodePoint
    : HexDigits { p.codePoint(p.GetText()) }?
    ;

// ********************************************
// *** 12: ECMAScript Language: Expressions ***
// ********************************************

// 12.1: Identifiers

identifierReference
    : Identifier
    | Yield
    | Await
    ;

identifierReference_Yield
    : Identifier
    | Await
    ;

identifierReference_Await
    : Identifier
    | Yield
    ;

identifierReference_Yield_Await
    : Identifier
    ;

bindingIdentifier
    : Identifier
    | Yield
    | Await
    ;

bindingIdentifier_Yield
    : Identifier
    | Yield
    | Await
    ;

bindingIdentifier_Await
    : Identifier
    | Yield
    | Await
    ;

bindingIdentifier_Yield_Await
    : Identifier
    | Yield
    | Await
    ;

labelIdentifier
    : Identifier
    | Yield
    | Await
    ;

labelIdentifier_Yield
    : Identifier
    | Await
    ;

labelIdentifier_Await
    : Identifier
    | Yield
    ;

labelIdentifier_Yield_Await
    : Identifier
    ;


Identifier
    : IdentifierName { p.not(ECMAScriptLexerReservedWord) }?
    ;

// 12.2: Primary Expression

primaryExpression
    : This
    | identifierReference
    | literal
    | arrayLiteral
    | objectLiteral
    | functionExpression
    | classExpression
    | generatorExpression
    | asyncFunctionExpression
    | asyncGeneratorExpression
    | regularExpressionLiteral
    | templateLiteral
    | coverParenthesizedExpressionAndArrowParameterList
    ;

primaryExpression_Yield
    : This
    | identifierReference_Yield
    | literal
    | arrayLiteral_Yield
    | objectLiteral_Yield
    | functionExpression
    | classExpression_Yield
    | generatorExpression
    | asyncFunctionExpression
    | asyncGeneratorExpression
    | regularExpressionLiteral
    | templateLiteral_Yield
    | coverParenthesizedExpressionAndArrowParameterList_Yield
    ;

primaryExpression_Await
    : This
    | identifierReference_Await
    | literal
    | arrayLiteral_Await
    | objectLiteral_Await
    | functionExpression
    | classExpression_Await
    | generatorExpression
    | asyncFunctionExpression
    | asyncGeneratorExpression
    | regularExpressionLiteral
    | templateLiteral_Await
    | coverParenthesizedExpressionAndArrowParameterList_Await
    ;

primaryExpression_Yield_Await
    : This
    | identifierReference_Yield_Await
    | literal
    | arrayLiteral_Yield_Await
    | objectLiteral_Yield_Await
    | functionExpression
    | classExpression_Yield_Await
    | generatorExpression
    | asyncFunctionExpression
    | asyncGeneratorExpression
    | regularExpressionLiteral
    | templateLiteral_Yield_Await
    | coverParenthesizedExpressionAndArrowParameterList_Yield_Await
    ;

coverParenthesizedExpressionAndArrowParameterList
    : OpenParen expression_In CloseParen
    | OpenParen expression_In CloseParen
    | OpenParen CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen expression_In Comma Ellipsis bindingIdentifier CloseParen
    | OpenParen expression_In Comma Ellipsis bindingPattern CloseParen
    ;

coverParenthesizedExpressionAndArrowParameterList_Yield
    : OpenParen expression_In_Yield CloseParen
    | OpenParen expression_In_Yield CloseParen
    | OpenParen CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen expression_In_Yield Comma Ellipsis bindingIdentifier_Yield CloseParen
    | OpenParen expression_In_Yield Comma Ellipsis bindingPattern_Yield CloseParen
    ;

coverParenthesizedExpressionAndArrowParameterList_Await
    : OpenParen expression_In_Await CloseParen
    | OpenParen expression_In_Await CloseParen
    | OpenParen CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen expression_In_Await Comma Ellipsis bindingIdentifier_Await CloseParen
    | OpenParen expression_In_Await Comma Ellipsis bindingPattern_Await CloseParen
    ;

coverParenthesizedExpressionAndArrowParameterList_Yield_Await
    : OpenParen expression_In_Yield_Await CloseParen
    | OpenParen expression_In_Yield_Await CloseParen
    | OpenParen CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen Ellipsis bindingIdentifier CloseParen
    | OpenParen expression_In_Yield_Await Comma Ellipsis bindingIdentifier_Yield_Await CloseParen
    | OpenParen expression_In_Yield_Await Comma Ellipsis bindingPattern_Yield_Await CloseParen
    ;

// 12.2.4: Literals

literal
    : NullLiteral
    | BooleanLiteral
    | NumericLiteral
    | StringLiteral
    ;

// 12.2.5: Array Initializer

arrayLiteral
    : OpenBracket elision? CloseBracket
    | OpenBracket elementList CloseBracket
    | OpenBracket elementList Comma elision? CloseBracket
    ;

arrayLiteral_Yield
    : OpenBracket elision? CloseBracket
    | OpenBracket elementList_Yield CloseBracket
    | OpenBracket elementList_Yield Comma elision? CloseBracket
    ;

arrayLiteral_Await
    : OpenBracket elision? CloseBracket
    | OpenBracket elementList_Await CloseBracket
    | OpenBracket elementList_Await Comma elision? CloseBracket
    ;

arrayLiteral_Yield_Await
    : OpenBracket elision? CloseBracket
    | OpenBracket elementList_Yield_Await CloseBracket
    | OpenBracket elementList_Yield_Await Comma elision? CloseBracket
    ;

elementList
    : elision? assignmentExpression_In
    | elision? spreadElement
    | elementList Comma elision? assignmentExpression_In
    | elementList Comma elision? spreadElement
    ;

elementList_Yield
    : elision? assignmentExpression_In_Yield
    | elision? spreadElement_Yield
    | elementList_Yield Comma elision? assignmentExpression_In_Yield
    | elementList_Yield Comma elision? spreadElement_Yield
    ;

elementList_Await
    : elision? assignmentExpression_In_Await
    | elision? spreadElement_Await
    | elementList_Await Comma elision? assignmentExpression_In_Await
    | elementList_Await Comma elision? spreadElement_Await
    ;

elementList_Yield_Await
    : elision? assignmentExpression_In_Yield_Await
    | elision? spreadElement_Yield_Await
    | elementList_Yield_Await Comma elision? assignmentExpression_In_Yield_Await
    | elementList_Yield_Await Comma elision? spreadElement_Yield_Await
    ;

elision
    : Comma
    | elision Comma
    ;

spreadElement
    : Ellipsis assignmentExpression_In
    ;

spreadElement_Yield
    : Ellipsis assignmentExpression_In_Yield
    ;

spreadElement_Await
    : Ellipsis assignmentExpression_In_Await
    ;

spreadElement_Yield_Await
    : Ellipsis assignmentExpression_In_Yield_Await
    ;

// 12.2.6: Object Initializer

objectLiteral
    : OpenBrace CloseBrace
    | OpenBrace propertyDefinitionList CloseBrace
    | OpenBrace propertyDefinitionList Comma CloseBrace
    ;

objectLiteral_Yield
    : OpenBrace CloseBrace
    | OpenBrace propertyDefinitionList_Yield CloseBrace
    | OpenBrace propertyDefinitionList_Yield Comma CloseBrace
    ;

objectLiteral_Await
    : OpenBrace CloseBrace
    | OpenBrace propertyDefinitionList_Await CloseBrace
    | OpenBrace propertyDefinitionList_Await Comma CloseBrace
    ;

objectLiteral_Yield_Await
    : OpenBrace CloseBrace
    | OpenBrace propertyDefinitionList_Yield_Await CloseBrace
    | OpenBrace propertyDefinitionList_Yield_Await Comma CloseBrace
    ;

propertyDefinitionList
    : propertyDefinition*
    ;

propertyDefinitionList_Yield
    : propertyDefinition_Yield*
    ;

propertyDefinitionList_Await
    : propertyDefinition_Await*
    ;

propertyDefinitionList_Yield_Await
    : propertyDefinition_Yield_Await*
    ;

propertyDefinition
    : identifierReference
    | coverInitializedName
    | propertyName Colon assignmentExpression_In
    | methodDefinition
    | Ellipsis assignmentExpression_In
    ;

propertyDefinition_Yield
    : identifierReference_Yield
    | coverInitializedName_Yield
    | propertyName_Yield Colon assignmentExpression_In_Yield
    | methodDefinition_Yield
    | Ellipsis assignmentExpression_In_Yield
    ;

propertyDefinition_Await
    : identifierReference_Await
    | coverInitializedName_Await
    | propertyName_Await Colon assignmentExpression_In_Await
    | methodDefinition_Await
    | Ellipsis assignmentExpression_In_Await
    ;

propertyDefinition_Yield_Await
    : identifierReference_Yield_Await
    | coverInitializedName_Yield_Await
    | propertyName_Yield_Await Colon assignmentExpression_In_Yield_Await
    | methodDefinition_Yield_Await
    | Ellipsis assignmentExpression_In_Yield_Await
    ;

propertyName
    : literalPropertyName
    | computedPropertyName
    ;

propertyName_Yield
    : literalPropertyName
    | computedPropertyName_Yield
    ;

propertyName_Await
    : literalPropertyName
    | computedPropertyName_Await
    ;

propertyName_Yield_Await
    : literalPropertyName
    | computedPropertyName_Yield_Await
    ;

literalPropertyName
    : IdentifierName
    | StringLiteral
    | NumericLiteral
    ;

computedPropertyName
    : OpenBracket assignmentExpression_In CloseBracket
    ;

computedPropertyName_Yield
    : OpenBracket assignmentExpression_In_Yield CloseBracket
    ;

computedPropertyName_Await
    : OpenBracket assignmentExpression_In_Await CloseBracket
    ;

computedPropertyName_Yield_Await
    : OpenBracket assignmentExpression_In_Yield_Await CloseBracket
    ;

coverInitializedName
    : identifierReference initializer_In
    ;

coverInitializedName_Yield
    : identifierReference_Yield initializer_In_Yield
    ;

coverInitializedName_Await
    : identifierReference_Await initializer_In_Await
    ;

coverInitializedName_Yield_Await
    : identifierReference_Yield_Await initializer_In_Yield_Await
    ;

initializer
    : Assign assignmentExpression
    ;

initializer_In
    : Assign assignmentExpression_In
    ;

initializer_Yield
    : Assign assignmentExpression_Yield
    ;

initializer_In_Yield
    : Assign assignmentExpression_In_Yield
    ;

initializer_Await
    : Assign assignmentExpression_Await
    ;

initializer_In_Await
    : Assign assignmentExpression_In_Await
    ;

initializer_Yield_Await
    : Assign assignmentExpression_Yield_Await
    ;

initializer_In_Yield_Await
    : Assign assignmentExpression_In_Yield_Await
    ;

// 12.2.9: Template Literals

templateLiteral
    : NoSubstitutionTemplate
    | substitutionTemplate
    ;

templateLiteral_Yield
    : NoSubstitutionTemplate
    | substitutionTemplate_Yield
    ;

templateLiteral_Await
    : NoSubstitutionTemplate
    | substitutionTemplate_Await
    ;

templateLiteral_Yield_Await
    : NoSubstitutionTemplate
    | substitutionTemplate_Yield_Await
    ;

templateLiteral_Tagged
    : NoSubstitutionTemplate
    | substitutionTemplate_Tagged
    ;

templateLiteral_Yield_Tagged
    : NoSubstitutionTemplate
    | substitutionTemplate_Yield_Tagged
    ;

templateLiteral_Await_Tagged
    : NoSubstitutionTemplate
    | substitutionTemplate_Await_Tagged
    ;

templateLiteral_Yield_Await_Tagged
    : NoSubstitutionTemplate
    | substitutionTemplate_Yield_Await_Tagged
    ;

substitutionTemplate
    : TemplateHead expression_In templateSpans
    ;

substitutionTemplate_Yield
    : TemplateHead expression_In_Yield templateSpans_Yield
    ;

substitutionTemplate_Await
    : TemplateHead expression_In_Await templateSpans_Await
    ;

substitutionTemplate_Yield_Await
    : TemplateHead expression_In_Yield_Await templateSpans_Yield_Await
    ;

substitutionTemplate_Tagged
    : TemplateHead expression_In templateSpans_Tagged
    ;

substitutionTemplate_Yield_Tagged
    : TemplateHead expression_In_Yield templateSpans_Yield_Tagged
    ;

substitutionTemplate_Await_Tagged
    : TemplateHead expression_In_Await templateSpans_Await_Tagged
    ;

substitutionTemplate_Yield_Await_Tagged
    : TemplateHead expression_In_Yield_Await templateSpans_Yield_Await_Tagged
    ;

templateSpans
    : templateMiddleList* TemplateTail
    ;

templateSpans_Yield
    : templateMiddleList_Yield* TemplateTail
    ;

templateSpans_Await
    : templateMiddleList_Await* TemplateTail
    ;

templateSpans_Yield_Await
    : templateMiddleList_Yield_Await* TemplateTail
    ;

templateSpans_Tagged
    : templateMiddleList_Tagged* TemplateTail
    ;

templateSpans_Yield_Tagged
    : templateMiddleList_Yield_Tagged* TemplateTail
    ;

templateSpans_Await_Tagged
    : templateMiddleList_Await_Tagged* TemplateTail
    ;

templateSpans_Yield_Await_Tagged
    : templateMiddleList_Yield_Await_Tagged* TemplateTail
    ;

templateMiddleList
    : (TemplateTail expression_In)+
    ;

templateMiddleList_Yield
    : (TemplateTail expression_In_Yield)+
    ;

templateMiddleList_Await
    : (TemplateTail expression_In_Await)+
    ;

templateMiddleList_Yield_Await
    : (TemplateTail expression_In_Yield_Await)+
    ;

templateMiddleList_Tagged
    : (TemplateTail expression_In)+
    ;

templateMiddleList_Yield_Tagged
    : (TemplateTail expression_In_Yield)+
    ;

templateMiddleList_Await_Tagged
    : (TemplateTail expression_In_Await)+
    ;

templateMiddleList_Yield_Await_Tagged
    : (TemplateTail expression_In_Yield_Await)+
    ;

// 12.3: Left-Hand-Side Expressions

memberExpression
    : primaryExpression
    | memberExpression OpenBracket expression_In CloseBracket
    | memberExpression Dot IdentifierName
    | memberExpression templateLiteral_Tagged
    | superProperty
    | metaProperty
    | New memberExpression arguments
    ;

memberExpression_Yield
    : primaryExpression_Yield
    | memberExpression_Yield OpenBracket expression_In_Yield CloseBracket
    | memberExpression_Yield Dot IdentifierName
    | memberExpression_Yield templateLiteral_Yield_Tagged
    | superProperty_Yield
    | metaProperty
    | New memberExpression_Yield arguments_Yield
    ;

memberExpression_Await
    : primaryExpression_Await
    | memberExpression_Await OpenBracket expression_In_Await CloseBracket
    | memberExpression_Await Dot IdentifierName
    | memberExpression_Await templateLiteral_Await_Tagged
    | superProperty_Await
    | metaProperty
    | New memberExpression_Await arguments_Await
    ;

memberExpression_Yield_Await
    : primaryExpression_Yield_Await
    | memberExpression_Yield_Await OpenBracket expression_In_Yield_Await CloseBracket
    | memberExpression_Yield_Await Dot IdentifierName
    | memberExpression_Yield_Await templateLiteral_Yield_Await_Tagged
    | superProperty_Yield_Await
    | metaProperty
    | New memberExpression_Yield_Await arguments_Yield_Await
    ;

superProperty
    : Super OpenBracket expression_In CloseBracket
    | Super Dot IdentifierName
    ;

superProperty_Yield
    : Super OpenBracket expression_In_Yield CloseBracket
    | Super Dot IdentifierName
    ;

superProperty_Await
    : Super OpenBracket expression_In_Await CloseBracket
    | Super Dot IdentifierName
    ;

superProperty_Yield_Await
    : Super OpenBracket expression_In_Yield_Await CloseBracket
    | Super Dot IdentifierName
    ;

metaProperty
    : newTarget
    ;

newTarget
    : New Dot Target
    ;

theNewExpression
    : memberExpression
    | New theNewExpression
    ;

theNewExpression_Yield
    : memberExpression_Yield
    | New theNewExpression_Yield
    ;

theNewExpression_Await
    : memberExpression_Await
    | New theNewExpression_Await
    ;

theNewExpression_Yield_Await
    : memberExpression_Yield_Await
    | New theNewExpression_Yield_Await
    ;

callExpression
    : coverCallExpressionAndAsyncArrowHead
    | superCall
    | callExpression arguments
    | callExpression OpenBracket expression_In CloseBracket
    | callExpression Dot IdentifierName
    | callExpression templateLiteral_Tagged
    ;

callExpression_Yield
    : coverCallExpressionAndAsyncArrowHead_Yield
    | superCall_Yield
    | callExpression_Yield arguments_Yield
    | callExpression_Yield OpenBracket expression_In_Yield CloseBracket
    | callExpression_Yield Dot IdentifierName
    | callExpression_Yield templateLiteral_Yield_Tagged
    ;

callExpression_Await
    : coverCallExpressionAndAsyncArrowHead_Await
    | superCall_Await
    | callExpression_Await arguments_Await
    | callExpression_Await OpenBracket expression_In_Await CloseBracket
    | callExpression_Await Dot IdentifierName
    | callExpression_Await templateLiteral_Await_Tagged
    ;

callExpression_Yield_Await
    : coverCallExpressionAndAsyncArrowHead_Yield_Await
    | superCall_Yield_Await
    | callExpression_Yield_Await arguments_Yield_Await
    | callExpression_Yield_Await OpenBracket expression_In_Yield_Await CloseBracket
    | callExpression_Yield_Await Dot IdentifierName
    | callExpression_Yield_Await templateLiteral_Yield_Await_Tagged
    ;

superCall
    : Super arguments
    ;

superCall_Yield
    : Super arguments_Yield
    ;

superCall_Await
    : Super arguments_Await
    ;

superCall_Yield_Await
    : Super arguments_Yield_Await
    ;

arguments
    : OpenParen CloseParen
    | OpenParen argumentList? Comma CloseParen
    ;

arguments_Yield
    : OpenParen CloseParen
    | OpenParen argumentList_Yield? Comma CloseParen
    ;

arguments_Await
    : OpenParen CloseParen
    | OpenParen argumentList_Await? Comma CloseParen
    ;

arguments_Yield_Await
    : OpenParen CloseParen
    | OpenParen argumentList_Yield_Await? Comma CloseParen
    ;

argumentList
    : Ellipsis? assignmentExpression_In
    | argumentList Comma Ellipsis? assignmentExpression_In
    ;

argumentList_Yield
    : Ellipsis? assignmentExpression_In_Yield
    | argumentList_Yield Comma Ellipsis? assignmentExpression_In_Yield
    ;

argumentList_Await
    : Ellipsis? assignmentExpression_In_Await
    | argumentList_Await Comma Ellipsis? assignmentExpression_In_Await
    ;

argumentList_Yield_Await
    : Ellipsis? assignmentExpression_In_Yield_Await
    | argumentList_Yield_Await Comma Ellipsis? assignmentExpression_In_Yield_Await
    ;

leftHandSideExpression
    : theNewExpression
    | callExpression
    ;

leftHandSideExpression_Yield
    : theNewExpression_Yield
    | callExpression_Yield
    ;

leftHandSideExpression_Await
    : theNewExpression_Await
    | callExpression_Await
    ;

leftHandSideExpression_Yield_Await
    : theNewExpression_Yield_Await
    | callExpression_Yield_Await
    ;

// 12.4: Update Expressions

updateExpression
    : leftHandSideExpression
    | leftHandSideExpression PlusPlus
    | leftHandSideExpression MinusMinus
    | PlusPlus unaryExpression
    | MinusMinus unaryExpression
    ;

updateExpression_Yield
    : leftHandSideExpression_Yield
    | leftHandSideExpression_Yield PlusPlus
    | leftHandSideExpression_Yield MinusMinus
    | PlusPlus unaryExpression_Yield
    | MinusMinus unaryExpression_Yield
    ;

updateExpression_Await
    : leftHandSideExpression_Await
    | leftHandSideExpression_Await PlusPlus
    | leftHandSideExpression_Await MinusMinus
    | PlusPlus unaryExpression_Await
    | MinusMinus unaryExpression_Await
    ;

updateExpression_Yield_Await
    : leftHandSideExpression_Yield_Await
    | leftHandSideExpression_Yield_Await PlusPlus
    | leftHandSideExpression_Yield_Await MinusMinus
    | PlusPlus unaryExpression_Yield_Await
    | MinusMinus unaryExpression_Yield_Await
    ;

// 12.5: Unary Operators

unaryExpression
    : updateExpression
    | Delete unaryExpression
    | Void unaryExpression
    | Typeof unaryExpression
    | Plus unaryExpression
    | Minus unaryExpression
    | BitNot unaryExpression
    | Not unaryExpression
    ;

unaryExpression_Yield
    : updateExpression_Yield
    | Delete unaryExpression_Yield
    | Void unaryExpression_Yield
    | Typeof unaryExpression_Yield
    | Plus unaryExpression_Yield
    | Minus unaryExpression_Yield
    | BitNot unaryExpression_Yield
    | Not unaryExpression_Yield
    ;

unaryExpression_Await
    : updateExpression_Await
    | Delete unaryExpression_Await
    | Void unaryExpression_Await
    | Typeof unaryExpression_Await
    | Plus unaryExpression_Await
    | Minus unaryExpression_Await
    | BitNot unaryExpression_Await
    | Not unaryExpression_Await
    | awaitExpression
    ;

unaryExpression_Yield_Await
    : updateExpression_Yield_Await
    | Delete unaryExpression_Yield_Await
    | Void unaryExpression_Yield_Await
    | Typeof unaryExpression_Yield_Await
    | Plus unaryExpression_Yield_Await
    | Minus unaryExpression_Yield_Await
    | BitNot unaryExpression_Yield_Await
    | Not unaryExpression_Yield_Await
    | awaitExpression_Yield
    ;

// 12.6: Exponentiation Operator

exponentationExpression
    : unaryExpression
    | updateExpression Power exponentationExpression
    ;

exponentationExpression_Yield
    : unaryExpression_Yield
    | updateExpression_Yield Power exponentationExpression_Yield
    ;

exponentationExpression_Await
    : unaryExpression_Await
    | updateExpression_Await Power exponentationExpression_Await
    ;

exponentationExpression_Yield_Await
    : unaryExpression_Yield_Await
    | updateExpression_Yield_Await Power exponentationExpression_Yield_Await
    ;

// 12.7: Multiplicative Operators

multiplicativeExpression
    : exponentationExpression
    | multiplicativeExpression MultiplicativeOperator exponentationExpression
    ;

multiplicativeExpression_Yield
    : exponentationExpression_Yield
    | multiplicativeExpression_Yield MultiplicativeOperator exponentationExpression_Yield
    ;

multiplicativeExpression_Await
    : exponentationExpression_Await
    | multiplicativeExpression_Await MultiplicativeOperator exponentationExpression_Await
    ;

multiplicativeExpression_Yield_Await
    : exponentationExpression_Yield_Await
    | multiplicativeExpression_Yield_Await MultiplicativeOperator exponentationExpression_Yield_Await
    ;

MultiplicativeOperator
    : [*/%]
    ;

// 12.8: Additive Operators

additiveExpression
    : multiplicativeExpression
    | additiveExpression Plus multiplicativeExpression
    | additiveExpression Minus multiplicativeExpression
    ;

additiveExpression_Yield
    : multiplicativeExpression_Yield
    | additiveExpression_Yield Plus multiplicativeExpression_Yield
    | additiveExpression_Yield Minus multiplicativeExpression_Yield
    ;

additiveExpression_Await
    : multiplicativeExpression_Await
    | additiveExpression_Await Plus multiplicativeExpression_Await
    | additiveExpression_Await Minus multiplicativeExpression_Await
    ;

additiveExpression_Yield_Await
    : multiplicativeExpression_Yield_Await
    | additiveExpression_Yield_Await Plus multiplicativeExpression_Yield_Await
    | additiveExpression_Yield_Await Minus multiplicativeExpression_Yield_Await
    ;

// 12.9: Bitwise Shift Operators

shiftExpression
    : additiveExpression
    | shiftExpression LeftShiftArithmetic additiveExpression
    | shiftExpression RightShiftArithmetic additiveExpression
    | shiftExpression RightShiftLogical additiveExpression
    ;

shiftExpression_Yield
    : additiveExpression_Yield
    | shiftExpression_Yield LeftShiftArithmetic additiveExpression_Yield
    | shiftExpression_Yield RightShiftArithmetic additiveExpression_Yield
    | shiftExpression_Yield RightShiftLogical additiveExpression_Yield
    ;

shiftExpression_Await
    : additiveExpression_Await
    | shiftExpression_Await LeftShiftArithmetic additiveExpression_Await
    | shiftExpression_Await RightShiftArithmetic additiveExpression_Await
    | shiftExpression_Await RightShiftLogical additiveExpression_Await
    ;

shiftExpression_Yield_Await
    : additiveExpression_Yield_Await
    | shiftExpression_Yield_Await LeftShiftArithmetic additiveExpression_Yield_Await
    | shiftExpression_Yield_Await RightShiftArithmetic additiveExpression_Yield_Await
    | shiftExpression_Yield_Await RightShiftLogical additiveExpression_Yield_Await
    ;

// 12.10: Relational Operators

relationalExpression
    : shiftExpression
    | relationalExpression LessThan shiftExpression
    | relationalExpression MoreThan shiftExpression
    | relationalExpression LessThanEquals shiftExpression
    | relationalExpression MoreThanEquals shiftExpression
    | relationalExpression Instanceof shiftExpression
    ;

relationalExpression_In
    : shiftExpression
    | relationalExpression_In LessThan shiftExpression
    | relationalExpression_In MoreThan shiftExpression
    | relationalExpression_In LessThanEquals shiftExpression
    | relationalExpression_In MoreThanEquals shiftExpression
    | relationalExpression_In Instanceof shiftExpression
    | relationalExpression In shiftExpression
    ;

relationalExpression_Yield
    : shiftExpression
    | relationalExpression_Yield LessThan shiftExpression
    | relationalExpression_Yield MoreThan shiftExpression
    | relationalExpression_Yield LessThanEquals shiftExpression
    | relationalExpression_Yield MoreThanEquals shiftExpression
    | relationalExpression_Yield Instanceof shiftExpression
    ;

relationalExpression_In_Yield
    : shiftExpression
    | relationalExpression_In_Yield LessThan shiftExpression
    | relationalExpression_In_Yield MoreThan shiftExpression
    | relationalExpression_In_Yield LessThanEquals shiftExpression
    | relationalExpression_In_Yield MoreThanEquals shiftExpression
    | relationalExpression_In_Yield Instanceof shiftExpression
    | relationalExpression_Yield In shiftExpression_Yield
    ;

relationalExpression_Await
    : shiftExpression
    | relationalExpression_Await LessThan shiftExpression
    | relationalExpression_Await MoreThan shiftExpression
    | relationalExpression_Await LessThanEquals shiftExpression
    | relationalExpression_Await MoreThanEquals shiftExpression
    | relationalExpression_Await Instanceof shiftExpression
    ;

relationalExpression_In_Await
    : shiftExpression
    | relationalExpression_In_Await LessThan shiftExpression
    | relationalExpression_In_Await MoreThan shiftExpression
    | relationalExpression_In_Await LessThanEquals shiftExpression
    | relationalExpression_In_Await MoreThanEquals shiftExpression
    | relationalExpression_In_Await Instanceof shiftExpression
    | relationalExpression_Await In shiftExpression_Await
    ;

relationalExpression_Yield_Await
    : shiftExpression
    | relationalExpression_Yield_Await LessThan shiftExpression
    | relationalExpression_Yield_Await MoreThan shiftExpression
    | relationalExpression_Yield_Await LessThanEquals shiftExpression
    | relationalExpression_Yield_Await MoreThanEquals shiftExpression
    | relationalExpression_Yield_Await Instanceof shiftExpression
    ;

relationalExpression_In_Yield_Await
    : shiftExpression
    | relationalExpression_In_Yield_Await LessThan shiftExpression
    | relationalExpression_In_Yield_Await MoreThan shiftExpression
    | relationalExpression_In_Yield_Await LessThanEquals shiftExpression
    | relationalExpression_In_Yield_Await MoreThanEquals shiftExpression
    | relationalExpression_In_Yield_Await Instanceof shiftExpression
    | relationalExpression_Yield_Await In shiftExpression_Yield_Await
    ;

// 12.11: Equality Operators

equalityExpression
    : relationalExpression
    | equalityExpression Equals relationalExpression
    | equalityExpression NotEquals relationalExpression
    | equalityExpression IdentityEquals relationalExpression
    | equalityExpression IdentityNotEquals relationalExpression
    ;

equalityExpression_In
    : relationalExpression_In
    | equalityExpression_In Equals relationalExpression_In
    | equalityExpression_In NotEquals relationalExpression_In
    | equalityExpression_In IdentityEquals relationalExpression_In
    | equalityExpression_In IdentityNotEquals relationalExpression_In
    ;

equalityExpression_Yield
    : relationalExpression_Yield
    | equalityExpression_Yield Equals relationalExpression_Yield
    | equalityExpression_Yield NotEquals relationalExpression_Yield
    | equalityExpression_Yield IdentityEquals relationalExpression_Yield
    | equalityExpression_Yield IdentityNotEquals relationalExpression_Yield
    ;

equalityExpression_In_Yield
    : relationalExpression_In_Yield
    | equalityExpression_In_Yield Equals relationalExpression_In_Yield
    | equalityExpression_In_Yield NotEquals relationalExpression_In_Yield
    | equalityExpression_In_Yield IdentityEquals relationalExpression_In_Yield
    | equalityExpression_In_Yield IdentityNotEquals relationalExpression_In_Yield
    ;

equalityExpression_Await
    : relationalExpression_Await
    | equalityExpression_Await Equals relationalExpression_Await
    | equalityExpression_Await NotEquals relationalExpression_Await
    | equalityExpression_Await IdentityEquals relationalExpression_Await
    | equalityExpression_Await IdentityNotEquals relationalExpression_Await
    ;

equalityExpression_In_Await
    : relationalExpression_In_Await
    | equalityExpression_In_Await Equals relationalExpression_In_Await
    | equalityExpression_In_Await NotEquals relationalExpression_In_Await
    | equalityExpression_In_Await IdentityEquals relationalExpression_In_Await
    | equalityExpression_In_Await IdentityNotEquals relationalExpression_In_Await
    ;

equalityExpression_Yield_Await
    : relationalExpression_Yield_Await
    | equalityExpression_Yield_Await Equals relationalExpression_Yield_Await
    | equalityExpression_Yield_Await NotEquals relationalExpression_Yield_Await
    | equalityExpression_Yield_Await IdentityEquals relationalExpression_Yield_Await
    | equalityExpression_Yield_Await IdentityNotEquals relationalExpression_Yield_Await
    ;

equalityExpression_In_Yield_Await
    : relationalExpression_In_Yield_Await
    | equalityExpression_In_Yield_Await Equals relationalExpression_In_Yield_Await
    | equalityExpression_In_Yield_Await NotEquals relationalExpression_In_Yield_Await
    | equalityExpression_In_Yield_Await IdentityEquals relationalExpression_In_Yield_Await
    | equalityExpression_In_Yield_Await IdentityNotEquals relationalExpression_In_Yield_Await
    ;

// 12.12: Binary Bitwise Operators

bitwiseANDExpression
    : equalityExpression
    | bitwiseANDExpression BitAnd equalityExpression
    ;
bitwiseANDExpression_In
    : equalityExpression_In
    | bitwiseANDExpression_In BitAnd equalityExpression_In
    ;
bitwiseANDExpression_Yield
    : equalityExpression_Yield
    | bitwiseANDExpression_Yield BitAnd equalityExpression_Yield
    ;
bitwiseANDExpression_In_Yield
    : equalityExpression_In_Yield
    | bitwiseANDExpression_In_Yield BitAnd equalityExpression_In_Yield
    ;
bitwiseANDExpression_Await
    : equalityExpression_Await
    | bitwiseANDExpression_Await BitAnd equalityExpression_Await
    ;
bitwiseANDExpression_In_Await
    : equalityExpression_In_Await
    | bitwiseANDExpression_In_Await BitAnd equalityExpression_In_Await
    ;
bitwiseANDExpression_Yield_Await
    : equalityExpression_Yield_Await
    | bitwiseANDExpression_Yield_Await BitAnd equalityExpression_Yield_Await
    ;
bitwiseANDExpression_In_Yield_Await
    : equalityExpression_In_Yield_Await
    | bitwiseANDExpression_In_Yield_Await BitAnd equalityExpression_In_Yield_Await
    ;

bitwiseXORExpression
    : equalityExpression
    | bitwiseXORExpression BitXor equalityExpression
    ;
bitwiseXORExpression_In
    : equalityExpression_In
    | bitwiseXORExpression_In BitXor equalityExpression_In
    ;
bitwiseXORExpression_Yield
    : equalityExpression_Yield
    | bitwiseXORExpression_Yield BitXor equalityExpression_Yield
    ;
bitwiseXORExpression_In_Yield
    : equalityExpression_In_Yield
    | bitwiseXORExpression_In_Yield BitXor equalityExpression_In_Yield
    ;
bitwiseXORExpression_Await
    : equalityExpression_Await
    | bitwiseXORExpression_Await BitXor equalityExpression_Await
    ;
bitwiseXORExpression_In_Await
    : equalityExpression_In_Await
    | bitwiseXORExpression_In_Await BitXor equalityExpression_In_Await
    ;
bitwiseXORExpression_Yield_Await
    : equalityExpression_Yield_Await
    | bitwiseXORExpression_Yield_Await BitXor equalityExpression_Yield_Await
    ;
bitwiseXORExpression_In_Yield_Await
    : equalityExpression_In_Yield_Await
    | bitwiseXORExpression_In_Yield_Await BitXor equalityExpression_In_Yield_Await
    ;

bitwiseORExpression
    : equalityExpression
    | bitwiseORExpression BitOr equalityExpression
    ;
bitwiseORExpression_In
    : equalityExpression_In
    | bitwiseORExpression_In BitOr equalityExpression_In
    ;
bitwiseORExpression_Yield
    : equalityExpression_Yield
    | bitwiseORExpression_Yield BitOr equalityExpression_Yield
    ;
bitwiseORExpression_In_Yield
    : equalityExpression_In_Yield
    | bitwiseORExpression_In_Yield BitOr equalityExpression_In_Yield
    ;
bitwiseORExpression_Await
    : equalityExpression_Await
    | bitwiseORExpression_Await BitOr equalityExpression_Await
    ;
bitwiseORExpression_In_Await
    : equalityExpression_In_Await
    | bitwiseORExpression_In_Await BitOr equalityExpression_In_Await
    ;
bitwiseORExpression_Yield_Await
    : equalityExpression_Yield_Await
    | bitwiseORExpression_Yield_Await BitOr equalityExpression_Yield_Await
    ;
bitwiseORExpression_In_Yield_Await
    : equalityExpression_In_Yield_Await
    | bitwiseORExpression_In_Yield_Await BitOr equalityExpression_In_Yield_Await
    ;

// 12.13: Binary Logical Operators

logicalANDExpression
    : bitwiseORExpression
    | logicalANDExpression And bitwiseORExpression
    ;

logicalANDExpression_In
    : bitwiseORExpression_In
    | logicalANDExpression_In And bitwiseORExpression_In
    ;

logicalANDExpression_Yield
    : bitwiseORExpression_Yield
    | logicalANDExpression_Yield And bitwiseORExpression_Yield
    ;

logicalANDExpression_In_Yield
    : bitwiseORExpression_In_Yield
    | logicalANDExpression_In_Yield And bitwiseORExpression_In_Yield
    ;

logicalANDExpression_Await
    : bitwiseORExpression_Await
    | logicalANDExpression_Await And bitwiseORExpression_Await
    ;

logicalANDExpression_In_Await
    : bitwiseORExpression_In_Await
    | logicalANDExpression_In_Await And bitwiseORExpression_In_Await
    ;

logicalANDExpression_Yield_Await
    : bitwiseORExpression_Yield_Await
    | logicalANDExpression_Yield_Await And bitwiseORExpression_Yield_Await
    ;

logicalANDExpression_In_Yield_Await
    : bitwiseORExpression_In_Yield_Await
    | logicalANDExpression_In_Yield_Await And bitwiseORExpression_In_Yield_Await
    ;

logicalORExpression
    : bitwiseORExpression
    | logicalORExpression Or bitwiseORExpression
    ;

logicalORExpression_In
    : bitwiseORExpression_In
    | logicalORExpression_In Or bitwiseORExpression_In
    ;

logicalORExpression_Yield
    : bitwiseORExpression_Yield
    | logicalORExpression_Yield Or bitwiseORExpression_Yield
    ;

logicalORExpression_In_Yield
    : bitwiseORExpression_In_Yield
    | logicalORExpression_In_Yield Or bitwiseORExpression_In_Yield
    ;

logicalORExpression_Await
    : bitwiseORExpression_Await
    | logicalORExpression_Await Or bitwiseORExpression_Await
    ;

logicalORExpression_In_Await
    : bitwiseORExpression_In_Await
    | logicalORExpression_In_Await Or bitwiseORExpression_In_Await
    ;

logicalORExpression_Yield_Await
    : bitwiseORExpression_Yield_Await
    | logicalORExpression_Yield_Await Or bitwiseORExpression_Yield_Await
    ;

logicalORExpression_In_Yield_Await
    : bitwiseORExpression_In_Yield_Await
    | logicalORExpression_In_Yield_Await Or bitwiseORExpression_In_Yield_Await
    ;

// 12.14: Conditional Operator ( ? : )

conditionalExpression
    : logicalORExpression
    | logicalORExpression QuestionMark assignmentExpression_In Colon assignmentExpression
    ;

conditionalExpression_In
    : logicalORExpression_In
    | logicalORExpression_In QuestionMark assignmentExpression_In Colon assignmentExpression_In
    ;

conditionalExpression_Yield
    : logicalORExpression_Yield
    | logicalORExpression_Yield QuestionMark assignmentExpression_In_Yield Colon assignmentExpression_Yield
    ;

conditionalExpression_In_Yield
    : logicalORExpression_In_Yield
    | logicalORExpression_In_Yield QuestionMark assignmentExpression_In_Yield Colon assignmentExpression_In_Yield
    ;

conditionalExpression_Await
    : logicalORExpression_Await
    | logicalORExpression_Await QuestionMark assignmentExpression_In_Await Colon assignmentExpression_Await
    ;

conditionalExpression_In_Await
    : logicalORExpression_In_Await
    | logicalORExpression_In_Await QuestionMark assignmentExpression_In_Await Colon assignmentExpression_In_Await
    ;

conditionalExpression_Yield_Await
    : logicalORExpression_Yield_Await
    | logicalORExpression_Yield_Await QuestionMark assignmentExpression_In_Yield_Await Colon assignmentExpression_Yield_Await
    ;

conditionalExpression_In_Yield_Await
    : logicalORExpression_In_Yield_Await
    | logicalORExpression_In_Yield_Await QuestionMark assignmentExpression_In_Yield_Await Colon assignmentExpression_In_Yield_Await
    ;

// 12.15: Assignment Operators

assignmentOperator
    : MultiplyAssign
    | DivideAssign
    | ModulusAssign
    | PlusAssign
    | MinusAssign
    | LeftShiftArithmeticAssign
    | RightShiftArithmeticAssign
    | RightShiftLogicalAssign
    | BitAndAssign
    | BitXorAssign
    | BitOrAssign
    | PowerAssign
    ;

assignmentExpression
    : conditionalExpression
    | arrowFunction
    | asyncArrowFunction
    | leftHandSideExpression Assign assignmentExpression
    | leftHandSideExpression assignmentOperator assignmentExpression
    ;

assignmentExpression_In
    : conditionalExpression_In
    | arrowFunction_In
    | asyncArrowFunction_In
    | leftHandSideExpression Assign assignmentExpression_In
    | leftHandSideExpression assignmentOperator assignmentExpression_In
    ;

assignmentExpression_Yield
    : conditionalExpression_Yield
    | yieldExpression
    | arrowFunction_Yield
    | asyncArrowFunction_Yield
    | leftHandSideExpression Assign assignmentExpression_Yield
    | leftHandSideExpression assignmentOperator assignmentExpression_Yield
    ;

assignmentExpression_In_Yield
    : conditionalExpression_In_Yield
    | yieldExpression_In
    | arrowFunction_In_Yield
    | asyncArrowFunction_In_Yield
    | leftHandSideExpression Assign assignmentExpression_In_Yield
    | leftHandSideExpression assignmentOperator assignmentExpression_In_Yield
    ;

assignmentExpression_Await
    : conditionalExpression_Await
    | arrowFunction_Await
    | asyncArrowFunction_Await
    | leftHandSideExpression Assign assignmentExpression_Await
    | leftHandSideExpression assignmentOperator assignmentExpression_Await
    ;

assignmentExpression_In_Await
    : conditionalExpression_In_Await
    | arrowFunction_In_Await
    | asyncArrowFunction_In_Await
    | leftHandSideExpression Assign assignmentExpression_In_Await
    | leftHandSideExpression assignmentOperator assignmentExpression_In_Await
    ;

assignmentExpression_Yield_Await
    : conditionalExpression_Yield_Await
    | yieldExpression_Await
    | arrowFunction_Yield_Await
    | asyncArrowFunction_Yield_Await
    | leftHandSideExpression Assign assignmentExpression_Yield_Await
    | leftHandSideExpression assignmentOperator assignmentExpression_Yield_Await
    ;

assignmentExpression_In_Yield_Await
    : conditionalExpression_In_Yield_Await
    | yieldExpression_In_Await
    | arrowFunction_In_Yield_Await
    | asyncArrowFunction_In_Yield_Await
    | leftHandSideExpression Assign assignmentExpression_In_Yield_Await
    | leftHandSideExpression assignmentOperator assignmentExpression_In_Yield_Await
    ;

// 12.16: Comma Operator ( , )

expression
    : assignmentExpression
    | expression Comma assignmentExpression
    ;

expression_In
    : assignmentExpression_In
    | expression_In Comma assignmentExpression_In
    ;

expression_Yield
    : assignmentExpression_Yield
    | expression_Yield Comma assignmentExpression_Yield
    ;

expression_In_Yield
    : assignmentExpression_In_Yield
    | expression_In_Yield Comma assignmentExpression_In_Yield
    ;

expression_Await
    : assignmentExpression_Await
    | expression_Await Comma assignmentExpression_Await
    ;

expression_In_Await
    : assignmentExpression_In_Await
    | expression_In_Await Comma assignmentExpression_In_Await
    ;

expression_Yield_Await
    : assignmentExpression_Yield_Await
    | expression_Yield_Await Comma assignmentExpression_Yield_Await
    ;

expression_In_Yield_Await
    : assignmentExpression_In_Yield_Await
    | expression_In_Yield_Await Comma assignmentExpression_In_Yield_Await
    ;

// ************************************************************
// *** 13: ECMAScript Language: Statements and Declarations ***
// ************************************************************

statement
    : blockStatement
    | variableStatement
    | theEmptyStatement
    | expressionStatement
    | ifStatement
    | breakableStatement
    | continueStatement
    | breakStatement
    // | returnStatement
    | withStatement
    | labelledStatement
    | throwStatement
    | tryStatement
    | debuggerStatement
    ;

statement_Yield
    : blockStatement_Yield
    | variableStatement_Yield
    | theEmptyStatement
    | expressionStatement_Yield
    | ifStatement_Yield
    | breakableStatement_Yield
    | continueStatement_Yield
    | breakStatement_Yield
    // | returnStatement_Yield
    | withStatement_Yield
    | labelledStatement_Yield
    | throwStatement_Yield
    | tryStatement_Yield
    | debuggerStatement
    ;

statement_Await
    : blockStatement_Await
    | variableStatement_Await
    | theEmptyStatement
    | expressionStatement_Await
    | ifStatement_Await
    | breakableStatement_Await
    | continueStatement_Await
    | breakStatement_Await
    // | returnStatement_Await
    | withStatement_Await
    | labelledStatement_Await
    | throwStatement_Await
    | tryStatement_Await
    | debuggerStatement
    ;

statement_Yield_Await
    : blockStatement_Yield_Await
    | variableStatement_Yield_Await
    | theEmptyStatement
    | expressionStatement_Yield_Await
    | ifStatement_Yield_Await
    | breakableStatement_Yield_Await
    | continueStatement_Yield_Await
    | breakStatement_Yield_Await
    // | returnStatement_Yield_Await
    | withStatement_Yield_Await
    | labelledStatement_Yield_Await
    | throwStatement_Yield_Await
    | tryStatement_Yield_Await
    | debuggerStatement
    ;

statement_Return
    : blockStatement_Return
    | variableStatement
    | theEmptyStatement
    | expressionStatement
    | ifStatement_Return
    | breakableStatement_Return
    | continueStatement
    | breakStatement
    | returnStatement
    | withStatement_Return
    | labelledStatement_Return
    | throwStatement
    | tryStatement_Return
    | debuggerStatement
    ;

statement_Yield_Return
    : blockStatement_Yield_Return
    | variableStatement_Yield
    | theEmptyStatement
    | expressionStatement_Yield
    | ifStatement_Yield_Return
    | breakableStatement_Yield_Return
    | continueStatement_Yield
    | breakStatement_Yield
    | returnStatement_Yield
    | withStatement_Yield_Return
    | labelledStatement_Yield_Return
    | throwStatement_Yield
    | tryStatement_Yield_Return
    | debuggerStatement
    ;

statement_Await_Return
    : blockStatement_Await_Return
    | variableStatement_Await
    | theEmptyStatement
    | expressionStatement_Await
    | ifStatement_Await_Return
    | breakableStatement_Await_Return
    | continueStatement_Await
    | breakStatement_Await
    | returnStatement_Await
    | withStatement_Await_Return
    | labelledStatement_Await_Return
    | throwStatement_Await
    | tryStatement_Await_Return
    | debuggerStatement
    ;

statement_Yield_Await_Return
    : blockStatement_Yield_Await_Return
    | variableStatement_Yield_Await
    | theEmptyStatement
    | expressionStatement_Yield_Await
    | ifStatement_Yield_Await_Return
    | breakableStatement_Yield_Await_Return
    | continueStatement_Yield_Await
    | breakStatement_Yield_Await
    | returnStatement_Yield_Await
    | withStatement_Yield_Await_Return
    | labelledStatement_Yield_Await_Return
    | throwStatement_Yield_Await
    | tryStatement_Yield_Await_Return
    | debuggerStatement
    ;

declaration
    : hoistableDeclaration
    | classDeclaration
    | lexicalDeclaration_In
    ;

declaration_Yield
    : hoistableDeclaration_Yield
    | classDeclaration_Yield
    | lexicalDeclaration_In_Yield
    ;

declaration_Await
    : hoistableDeclaration_Await
    | classDeclaration_Await
    | lexicalDeclaration_In_Await
    ;

declaration_Yield_Await
    : hoistableDeclaration_Yield_Await
    | classDeclaration_Yield_Await
    | lexicalDeclaration_In_Yield_Await
    ;

hoistableDeclaration
    : functionDeclaration
    | generatorDeclaration
    | asyncFunctionDeclaration
    | asyncGeneratorDeclaration
    ;

hoistableDeclaration_Yield
    : functionDeclaration_Yield
    | generatorDeclaration_Yield
    | asyncFunctionDeclaration_Yield
    | asyncGeneratorDeclaration_Yield
    ;

hoistableDeclaration_Await
    : functionDeclaration_Await
    | generatorDeclaration_Await
    | asyncFunctionDeclaration_Await
    | asyncGeneratorDeclaration_Await
    ;

hoistableDeclaration_Yield_Await
    : functionDeclaration_Yield_Await
    | generatorDeclaration_Yield_Await
    | asyncFunctionDeclaration_Yield_Await
    | asyncGeneratorDeclaration_Yield_Await
    ;

hoistableDeclaration_Default
    : functionDeclaration_Default
    | generatorDeclaration_Default
    | asyncFunctionDeclaration_Default
    | asyncGeneratorDeclaration_Default
    ;

hoistableDeclaration_Yield_Default
    : functionDeclaration_Yield_Default
    | generatorDeclaration_Yield_Default
    | asyncFunctionDeclaration_Yield_Default
    | asyncGeneratorDeclaration_Yield_Default
    ;

hoistableDeclaration_Await_Default
    : functionDeclaration_Await_Default
    | generatorDeclaration_Await_Default
    | asyncFunctionDeclaration_Await_Default
    | asyncGeneratorDeclaration_Await_Default
    ;

hoistableDeclaration_Yield_Await_Default
    : functionDeclaration_Yield_Await_Default
    | generatorDeclaration_Yield_Await_Default
    | asyncFunctionDeclaration_Yield_Await_Default
    | asyncGeneratorDeclaration_Yield_Await_Default
    ;

breakableStatement
    : iterationStatement
    | switchStatement
    ;

breakableStatement_Yield
    : iterationStatement_Yield
    | switchStatement_Yield
    ;

breakableStatement_Await
    : iterationStatement_Await
    | switchStatement_Await
    ;

breakableStatement_Yield_Await
    : iterationStatement_Yield_Await
    | switchStatement_Yield_Await
    ;

breakableStatement_Return
    : iterationStatement_Return
    | switchStatement_Return
    ;

breakableStatement_Yield_Return
    : iterationStatement_Yield_Return
    | switchStatement_Yield_Return
    ;

breakableStatement_Await_Return
    : iterationStatement_Await_Return
    | switchStatement_Await_Return
    ;

breakableStatement_Yield_Await_Return
    : iterationStatement_Yield_Await_Return
    | switchStatement_Yield_Await_Return
    ;

// 13.2: Block

blockStatement
    : block
    ;

blockStatement_Yield
    : block_Yield
    ;

blockStatement_Await
    : block_Await
    ;

blockStatement_Yield_Await
    : block_Yield_Await
    ;

blockStatement_Return
    : block_Return
    ;

blockStatement_Yield_Return
    : block_Yield_Return
    ;

blockStatement_Await_Return
    : block_Await_Return
    ;

blockStatement_Yield_Await_Return
    : block_Yield_Await_Return
    ;

block
    : OpenBrace statementList? CloseBrace
    ;

block_Yield
    : OpenBrace statementList_Yield? CloseBrace
    ;

block_Await
    : OpenBrace statementList_Await? CloseBrace
    ;

block_Yield_Await
    : OpenBrace statementList_Yield_Await? CloseBrace
    ;

block_Return
    : OpenBrace statementList_Return? CloseBrace
    ;

block_Yield_Return
    : OpenBrace statementList_Yield_Return? CloseBrace
    ;

block_Await_Return
    : OpenBrace statementList_Await_Return? CloseBrace
    ;

block_Yield_Await_Return
    : OpenBrace statementList_Yield_Await_Return? CloseBrace
    ;

statementList
    : statementListItem+
    ;

statementList_Yield
    : statementListItem_Yield+
    ;

statementList_Await
    : statementListItem_Await+
    ;

statementList_Yield_Await
    : statementListItem_Yield_Await+
    ;

statementList_Return
    : statementListItem_Return+
    ;

statementList_Yield_Return
    : statementListItem_Yield_Return+
    ;

statementList_Await_Return
    : statementListItem_Await_Return+
    ;

statementList_Yield_Await_Return
    : statementListItem_Yield_Await_Return+
    ;

statementListItem
    : statement
    | declaration
    ;

statementListItem_Yield
    : statement_Yield
    | declaration_Yield
    ;

statementListItem_Await
    : statement_Await
    | declaration_Await
    ;

statementListItem_Yield_Await
    : statement_Yield_Await
    | declaration_Yield_Await
    ;

statementListItem_Return
    : statement_Return
    | declaration
    ;

statementListItem_Yield_Return
    : statement_Yield_Return
    | declaration_Yield
    ;

statementListItem_Await_Return
    : statement_Await_Return
    | declaration_Await
    ;

statementListItem_Yield_Await_Return
    : statement_Yield_Await_Return
    | declaration_Yield_Await
    ;

// 13.3: Declarations and the Variable Statement

lexicalDeclaration
    : letOrConst bindingList SemiColon
    ;

lexicalDeclaration_In
    : letOrConst bindingList_In SemiColon
    ;

lexicalDeclaration_Yield
    : letOrConst bindingList_Yield SemiColon
    ;

lexicalDeclaration_In_Yield
    : letOrConst bindingList_In_Yield SemiColon
    ;

lexicalDeclaration_Await
    : letOrConst bindingList_Await SemiColon
    ;

lexicalDeclaration_In_Await
    : letOrConst bindingList_In_Await SemiColon
    ;

lexicalDeclaration_Yield_Await
    : letOrConst bindingList_Yield_Await SemiColon
    ;

lexicalDeclaration_In_Yield_Await
    : letOrConst bindingList_In_Yield_Await SemiColon
    ;

letOrConst
    : Let
    | Const
    ;

bindingList
    : lexicalBinding
    | bindingList Comma lexicalBinding
    ;

bindingList_In
    : lexicalBinding_In
    | bindingList_In Comma lexicalBinding_In
    ;

bindingList_Yield
    : lexicalBinding_Yield
    | bindingList_Yield Comma lexicalBinding_Yield
    ;

bindingList_In_Yield
    : lexicalBinding_In_Yield
    | bindingList_In_Yield Comma lexicalBinding_In_Yield
    ;

bindingList_Await
    : lexicalBinding_Await
    | bindingList_Await Comma lexicalBinding_Await
    ;

bindingList_In_Await
    : lexicalBinding_In_Await
    | bindingList_In_Await Comma lexicalBinding_In_Await
    ;

bindingList_Yield_Await
    : lexicalBinding_Yield_Await
    | bindingList_Yield_Await Comma lexicalBinding_Yield_Await
    ;

bindingList_In_Yield_Await
    : lexicalBinding_In_Yield_Await
    | bindingList_In_Yield_Await Comma lexicalBinding_In_Yield_Await
    ;

lexicalBinding
    : bindingIdentifier initializer?
    | bindingPattern initializer
    ;

lexicalBinding_In
    : bindingIdentifier initializer_In?
    | bindingPattern initializer_In
    ;

lexicalBinding_Yield
    : bindingIdentifier_Yield initializer_Yield?
    | bindingPattern_Yield initializer_Yield
    ;

lexicalBinding_In_Yield
    : bindingIdentifier_Yield initializer_In_Yield?
    | bindingPattern_Yield initializer_In_Yield
    ;

lexicalBinding_Await
    : bindingIdentifier_Await initializer_Await?
    | bindingPattern_Await initializer_Await
    ;

lexicalBinding_In_Await
    : bindingIdentifier_Await initializer_In_Await?
    | bindingPattern_Await initializer_In_Await
    ;

lexicalBinding_Yield_Await
    : bindingIdentifier_Yield_Await initializer_Yield_Await?
    | bindingPattern_Yield_Await initializer_Yield_Await
    ;

lexicalBinding_In_Yield_Await
    : bindingIdentifier_Yield_Await initializer_In_Yield_Await?
    | bindingPattern_Yield_Await initializer_In_Yield_Await
    ;

// 13.3.2: Variable Statement

variableStatement
    : Var variableDeclarationList_In SemiColon
    ;

variableStatement_Yield
    : Var variableDeclarationList_In_Yield SemiColon
    ;

variableStatement_Await
    : Var variableDeclarationList_In_Await SemiColon
    ;

variableStatement_Yield_Await
    : Var variableDeclarationList_In_Yield_Await SemiColon
    ;

variableDeclarationList
    : variableDeclaration
    | variableDeclarationList Comma variableDeclaration
    ;

variableDeclarationList_In
    : variableDeclaration_In
    | variableDeclarationList_In Comma variableDeclaration_In
    ;

variableDeclarationList_Yield
    : variableDeclaration_Yield
    | variableDeclarationList_Yield Comma variableDeclaration_Yield
    ;

variableDeclarationList_In_Yield
    : variableDeclaration_In_Yield
    | variableDeclarationList_In_Yield Comma variableDeclaration_In_Yield
    ;

variableDeclarationList_Await
    : variableDeclaration_Await
    | variableDeclarationList_Await Comma variableDeclaration_Await
    ;

variableDeclarationList_In_Await
    : variableDeclaration_In_Await
    | variableDeclarationList_In_Await Comma variableDeclaration_In_Await
    ;

variableDeclarationList_Yield_Await
    : variableDeclaration_Yield_Await
    | variableDeclarationList_Yield_Await Comma variableDeclaration_Yield_Await
    ;

variableDeclarationList_In_Yield_Await
    : variableDeclaration_In_Yield_Await
    | variableDeclarationList_In_Yield_Await Comma variableDeclaration_In_Yield_Await
    ;

variableDeclaration
    : bindingIdentifier initializer?
    | bindingPattern initializer
    ;

variableDeclaration_In
    : bindingIdentifier initializer_In?
    | bindingPattern initializer_In
    ;

variableDeclaration_Yield
    : bindingIdentifier_Yield initializer_Yield?
    | bindingPattern_Yield initializer_Yield
    ;

variableDeclaration_In_Yield
    : bindingIdentifier_Yield initializer_In_Yield?
    | bindingPattern_Yield initializer_In_Yield
    ;

variableDeclaration_Await
    : bindingIdentifier_Await initializer_Await?
    | bindingPattern_Await initializer_Await
    ;

variableDeclaration_In_Await
    : bindingIdentifier_Await initializer_In_Await?
    | bindingPattern_Await initializer_In_Await
    ;

variableDeclaration_Yield_Await
    : bindingIdentifier_Yield_Await initializer_Yield_Await?
    | bindingPattern_Yield_Await initializer_Yield_Await
    ;

variableDeclaration_In_Yield_Await
    : bindingIdentifier_Yield_Await initializer_In_Yield_Await?
    | bindingPattern_Yield_Await initializer_In_Yield_Await
    ;

// 13.3.3: Destructuring Binding Patterns

bindingPattern
    : objectBindingPattern
    | arrayBindingPattern
    ;

bindingPattern_Yield
    : objectBindingPattern_Yield
    | arrayBindingPattern_Yield
    ;

bindingPattern_Await
    : objectBindingPattern_Await
    | arrayBindingPattern_Await
    ;

bindingPattern_Yield_Await
    : objectBindingPattern_Yield_Await
    | arrayBindingPattern_Yield_Await
    ;

objectBindingPattern
    : OpenBrace CloseBrace
    | OpenBrace bindingRestProperty CloseBrace
    | OpenBrace bindingPropertyList CloseBrace
    | OpenBrace bindingPropertyList Comma bindingRestProperty? CloseBrace
    ;

objectBindingPattern_Yield
    : OpenBrace CloseBrace
    | OpenBrace bindingRestProperty_Yield CloseBrace
    | OpenBrace bindingPropertyList_Yield CloseBrace
    | OpenBrace bindingPropertyList_Yield Comma bindingRestProperty_Yield? CloseBrace
    ;

objectBindingPattern_Await
    : OpenBrace CloseBrace
    | OpenBrace bindingRestProperty_Await CloseBrace
    | OpenBrace bindingPropertyList_Await CloseBrace
    | OpenBrace bindingPropertyList_Await Comma bindingRestProperty_Await? CloseBrace
    ;

objectBindingPattern_Yield_Await
    : OpenBrace CloseBrace
    | OpenBrace bindingRestProperty_Yield_Await CloseBrace
    | OpenBrace bindingPropertyList_Yield_Await CloseBrace
    | OpenBrace bindingPropertyList_Yield_Await Comma bindingRestProperty_Yield_Await? CloseBrace
    ;

arrayBindingPattern
    : OpenBracket elision? bindingRestElement? CloseBracket
    | OpenBracket bindingElementList CloseBracket
    | OpenBracket bindingElementList Comma elision? bindingRestElement? CloseBracket
    ;

arrayBindingPattern_Yield
    : OpenBracket elision? bindingRestElement_Yield? CloseBracket
    | OpenBracket bindingElementList_Yield CloseBracket
    | OpenBracket bindingElementList_Yield Comma elision? bindingRestElement_Yield? CloseBracket
    ;

arrayBindingPattern_Await
    : OpenBracket elision? bindingRestElement_Await? CloseBracket
    | OpenBracket bindingElementList_Await CloseBracket
    | OpenBracket bindingElementList_Await Comma elision? bindingRestElement_Await? CloseBracket
    ;

arrayBindingPattern_Yield_Await
    : OpenBracket elision? bindingRestElement_Yield_Await? CloseBracket
    | OpenBracket bindingElementList_Yield_Await CloseBracket
    | OpenBracket bindingElementList_Yield_Await Comma elision? bindingRestElement_Yield_Await? CloseBracket
    ;

bindingRestProperty
    : Ellipsis bindingIdentifier
    ;

bindingRestProperty_Yield
    : Ellipsis bindingIdentifier_Yield
    ;

bindingRestProperty_Await
    : Ellipsis bindingIdentifier_Await
    ;

bindingRestProperty_Yield_Await
    : Ellipsis bindingIdentifier_Yield_Await
    ;

bindingPropertyList
    : bindingProperty
    | bindingPropertyList Comma bindingProperty
    ;

bindingPropertyList_Yield
    : bindingProperty_Yield
    | bindingPropertyList_Yield Comma bindingProperty_Yield
    ;

bindingPropertyList_Await
    : bindingProperty_Await
    | bindingPropertyList_Await Comma bindingProperty_Await
    ;

bindingPropertyList_Yield_Await
    : bindingProperty_Yield_Await
    | bindingPropertyList_Yield_Await Comma bindingProperty_Yield_Await
    ;

bindingElementList
    : bindingElisionElement
    | bindingElementList Comma bindingElisionElement
    ;

bindingElementList_Yield
    : bindingElisionElement_Yield
    | bindingElementList_Yield Comma bindingElisionElement_Yield
    ;

bindingElementList_Await
    : bindingElisionElement_Await
    | bindingElementList_Await Comma bindingElisionElement_Await
    ;

bindingElementList_Yield_Await
    : bindingElisionElement_Yield_Await
    | bindingElementList_Yield_Await Comma bindingElisionElement_Yield_Await
    ;

bindingElisionElement
    : elision? bindingElement
    ;

bindingElisionElement_Yield
    : elision? bindingElement_Yield
    ;

bindingElisionElement_Await
    : elision? bindingElement_Await
    ;

bindingElisionElement_Yield_Await
    : elision? bindingElement_Yield_Await
    ;

bindingProperty
    : singleNameBinding
    | propertyName Colon bindingElement
    ;

bindingProperty_Yield
    : singleNameBinding_Yield
    | propertyName_Yield Colon bindingElement_Yield
    ;

bindingProperty_Await
    : singleNameBinding_Await
    | propertyName_Await Colon bindingElement_Await
    ;

bindingProperty_Yield_Await
    : singleNameBinding_Yield_Await
    | propertyName_Yield_Await Colon bindingElement_Yield_Await
    ;

bindingElement
    : singleNameBinding
    | bindingPattern initializer_In?
    ;

bindingElement_Yield
    : singleNameBinding_Yield
    | bindingPattern_Yield initializer_In_Yield?
    ;

bindingElement_Await
    : singleNameBinding_Await
    | bindingPattern_Await initializer_In_Await?
    ;

bindingElement_Yield_Await
    : singleNameBinding_Yield_Await
    | bindingPattern_Yield_Await initializer_In_Yield_Await?
    ;

singleNameBinding
    : bindingIdentifier initializer_In?
    ;

singleNameBinding_Yield
    : bindingIdentifier_Yield initializer_In_Yield?
    ;

singleNameBinding_Await
    : bindingIdentifier_Await initializer_In_Await?
    ;

singleNameBinding_Yield_Await
    : bindingIdentifier_Yield_Await initializer_In_Yield_Await?
    ;

bindingRestElement
    : Ellipsis bindingIdentifier
    | Ellipsis bindingPattern
    ;

bindingRestElement_Yield
    : Ellipsis bindingIdentifier_Yield
    | Ellipsis bindingPattern_Yield
    ;

bindingRestElement_Await
    : Ellipsis bindingIdentifier_Await
    | Ellipsis bindingPattern_Await
    ;

bindingRestElement_Yield_Await
    : Ellipsis bindingIdentifier_Yield_Await
    | Ellipsis bindingPattern_Yield_Await
    ;

// 13.4: Empty Statement

theEmptyStatement
    : SemiColon
    ;

// 13.5: Expression Statement

expressionStatement
    : expression_In { p.negativeLookahead(ECMAScriptLexerOpenBrace, ECMAScriptLexerFunction, ECMAScriptLexerAsyncFunction, ECMAScriptLexerClass, ECMAScriptLexerLetBracket) }? SemiColon
    ;

expressionStatement_Yield
    : expression_In_Yield { p.negativeLookahead(ECMAScriptLexerOpenBrace, ECMAScriptLexerFunction, ECMAScriptLexerAsyncFunction, ECMAScriptLexerClass, ECMAScriptLexerLetBracket) }? SemiColon
    ;

expressionStatement_Await
    : expression_In_Await { p.negativeLookahead(ECMAScriptLexerOpenBrace, ECMAScriptLexerFunction, ECMAScriptLexerAsyncFunction, ECMAScriptLexerClass, ECMAScriptLexerLetBracket) }? SemiColon
    ;

expressionStatement_Yield_Await
    : expression_In_Yield_Await { p.negativeLookahead(ECMAScriptLexerOpenBrace, ECMAScriptLexerFunction, ECMAScriptLexerAsyncFunction, ECMAScriptLexerClass, ECMAScriptLexerLetBracket) }? SemiColon
    ;

// 13.6: The if Statement

ifStatement
    : If OpenParen expression_In CloseParen statement Else statement
    | If OpenParen expression_In CloseParen statement
    ;

ifStatement_Yield
    : If OpenParen expression_In_Yield CloseParen statement_Yield Else statement_Yield
    | If OpenParen expression_In_Yield CloseParen statement_Yield
    ;

ifStatement_Await
    : If OpenParen expression_In_Await CloseParen statement_Await Else statement_Await
    | If OpenParen expression_In_Await CloseParen statement_Await
    ;

ifStatement_Yield_Await
    : If OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await Else statement_Yield_Await
    | If OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await
    ;

ifStatement_Return
    : If OpenParen expression_In CloseParen statement_Return Else statement_Return
    | If OpenParen expression_In CloseParen statement_Return
    ;

ifStatement_Yield_Return
    : If OpenParen expression_In_Yield CloseParen statement_Yield_Return Else statement_Yield_Return
    | If OpenParen expression_In_Yield CloseParen statement_Yield_Return
    ;

ifStatement_Await_Return
    : If OpenParen expression_In_Await CloseParen statement_Await_Return Else statement_Await_Return
    | If OpenParen expression_In_Await CloseParen statement_Await_Return
    ;

ifStatement_Yield_Await_Return
    : If OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await_Return Else statement_Yield_Await_Return
    | If OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    ;

// 13.7: Iteration Statements

iterationStatement
    : Do statement While OpenParen expression_In CloseParen SemiColon
    | While OpenParen expression_In CloseParen statement
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression? SemiColon expression_In? SemiColon expression_In? CloseParen statement
    | For OpenParen Var variableDeclarationList SemiColon expression_In? SemiColon expression_In? CloseParen statement
    | For OpenParen lexicalDeclaration expression_In? SemiColon expression_In? CloseParen statement
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression In expression_In CloseParen statement
    | For OpenParen Var forBinding In expression_In CloseParen statement
    | For OpenParen forDeclaration In expression_In CloseParen statement
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression Of assignmentExpression_In CloseParen statement
    | For OpenParen Var forBinding Of assignmentExpression_In CloseParen statement
    | For OpenParen forDeclaration Of assignmentExpression_In CloseParen statement
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression Of assignmentExpression_In CloseParen statement
    // | For Await OpenParen Var forBinding Of assignmentExpression_In CloseParen statement
    // | For Await OpenParen forDeclaration Of assignmentExpression_In CloseParen statement
    ;

iterationStatement_Yield
    : Do statement_Yield While OpenParen expression_In_Yield CloseParen SemiColon
    | While OpenParen expression_In_Yield CloseParen statement_Yield
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Yield? SemiColon expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield
    | For OpenParen Var variableDeclarationList_Yield SemiColon expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield
    | For OpenParen lexicalDeclaration_Yield expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Yield In expression_In_Yield CloseParen statement_Yield
    | For OpenParen Var forBinding_Yield In expression_In_Yield CloseParen statement_Yield
    | For OpenParen forDeclaration_Yield In expression_In_Yield CloseParen statement_Yield
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    | For OpenParen Var forBinding_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    | For OpenParen forDeclaration_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    // | For Await OpenParen Var forBinding_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    // | For Await OpenParen forDeclaration_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield
    ;

iterationStatement_Await
    : Do statement_Await While OpenParen expression_In_Await CloseParen SemiColon
    | While OpenParen expression_In_Await CloseParen statement_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Await? SemiColon expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await
    | For OpenParen Var variableDeclarationList_Await SemiColon expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await
    | For OpenParen lexicalDeclaration_Await expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Await In expression_In_Await CloseParen statement_Await
    | For OpenParen Var forBinding_Await In expression_In_Await CloseParen statement_Await
    | For OpenParen forDeclaration_Await In expression_In_Await CloseParen statement_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Await Of assignmentExpression_In_Await CloseParen statement_Await
    | For OpenParen Var forBinding_Await Of assignmentExpression_In_Await CloseParen statement_Await
    | For OpenParen forDeclaration_Await Of assignmentExpression_In_Await CloseParen statement_Await
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Await Of assignmentExpression_In_Await CloseParen statement_Await
    // | For Await OpenParen Var forBinding_Await Of assignmentExpression_In_Await CloseParen statement_Await
    // | For Await OpenParen forDeclaration_Await Of assignmentExpression_In_Await CloseParen statement_Await
    ;

iterationStatement_Yield_Await
    : Do statement_Yield_Await While OpenParen expression_In_Yield_Await CloseParen SemiColon
    | While OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Yield_Await? SemiColon expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await
    | For OpenParen Var variableDeclarationList_Yield_Await SemiColon expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await
    | For OpenParen lexicalDeclaration_Yield_Await expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen Var forBinding_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen forDeclaration_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen Var forBinding_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    | For OpenParen forDeclaration_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    // | For Await OpenParen Var forBinding_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    // | For Await OpenParen forDeclaration_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await
    ;

iterationStatement_Return
    : Do statement_Return While OpenParen expression_In CloseParen SemiColon
    | While OpenParen expression_In CloseParen statement_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression? SemiColon expression_In? SemiColon expression_In? CloseParen statement_Return
    | For OpenParen Var variableDeclarationList SemiColon expression_In? SemiColon expression_In? CloseParen statement_Return
    | For OpenParen lexicalDeclaration expression_In? SemiColon expression_In? CloseParen statement_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression In expression_In CloseParen statement_Return
    | For OpenParen Var forBinding In expression_In CloseParen statement_Return
    | For OpenParen forDeclaration In expression_In CloseParen statement_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression Of assignmentExpression_In CloseParen statement_Return
    | For OpenParen Var forBinding Of assignmentExpression_In CloseParen statement_Return
    | For OpenParen forDeclaration Of assignmentExpression_In CloseParen statement_Return
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression Of assignmentExpression_In CloseParen statement_Return
    // | For Await OpenParen Var forBinding Of assignmentExpression_In CloseParen statement_Return
    // | For Await OpenParen forDeclaration Of assignmentExpression_In CloseParen statement_Return
    ;

iterationStatement_Yield_Return
    : Do statement_Yield_Return While OpenParen expression_In_Yield CloseParen SemiColon
    | While OpenParen expression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Yield? SemiColon expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield_Return
    | For OpenParen Var variableDeclarationList_Yield SemiColon expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield_Return
    | For OpenParen lexicalDeclaration_Yield expression_In_Yield? SemiColon expression_In_Yield? CloseParen statement_Yield_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Yield In expression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen Var forBinding_Yield In expression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen forDeclaration_Yield In expression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen Var forBinding_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    | For OpenParen forDeclaration_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    // | For Await OpenParen Var forBinding_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    // | For Await OpenParen forDeclaration_Yield Of assignmentExpression_In_Yield CloseParen statement_Yield_Return
    ;

iterationStatement_Await_Return
    : Do statement_Await_Return While OpenParen expression_In_Await CloseParen SemiColon
    | While OpenParen expression_In_Await CloseParen statement_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Await? SemiColon expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await_Return
    | For OpenParen Var variableDeclarationList_Await SemiColon expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await_Return
    | For OpenParen lexicalDeclaration_Await expression_In_Await? SemiColon expression_In_Await? CloseParen statement_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Await In expression_In_Await CloseParen statement_Await_Return
    | For OpenParen Var forBinding_Await In expression_In_Await CloseParen statement_Await_Return
    | For OpenParen forDeclaration_Await In expression_In_Await CloseParen statement_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    | For OpenParen Var forBinding_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    | For OpenParen forDeclaration_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    // | For Await OpenParen Var forBinding_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    // | For Await OpenParen forDeclaration_Await Of assignmentExpression_In_Await CloseParen statement_Await_Return
    ;

iterationStatement_Yield_Await_Return
    : Do statement_Yield_Await_Return While OpenParen expression_In_Yield_Await CloseParen SemiColon
    | While OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? expression_Yield_Await? SemiColon expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await_Return
    | For OpenParen Var variableDeclarationList_Yield_Await SemiColon expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await_Return
    | For OpenParen lexicalDeclaration_Yield_Await expression_In_Yield_Await? SemiColon expression_In_Yield_Await? CloseParen statement_Yield_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLetBracket) }? leftHandSideExpression_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen Var forBinding_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen forDeclaration_Yield_Await In expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen Var forBinding_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    | For OpenParen forDeclaration_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    // | For Await OpenParen { p.negativeLookahead(ECMAScriptLexerLet) }? leftHandSideExpression_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    // | For Await OpenParen Var forBinding_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    // | For Await OpenParen forDeclaration_Yield_Await Of assignmentExpression_In_Yield_Await CloseParen statement_Yield_Await_Return
    ;

forDeclaration
    : letOrConst forBinding
    ;

forDeclaration_Yield
    : letOrConst forBinding_Yield
    ;

forDeclaration_Await
    : letOrConst forBinding_Await
    ;

forDeclaration_Yield_Await
    : letOrConst forBinding_Yield_Await
    ;

forBinding
    : bindingIdentifier
    | bindingPattern
    ;

forBinding_Yield
    : bindingIdentifier_Yield
    | bindingPattern_Yield
    ;

forBinding_Await
    : bindingIdentifier_Await
    | bindingPattern_Await
    ;

forBinding_Yield_Await
    : bindingIdentifier_Yield_Await
    | bindingPattern_Yield_Await
    ;

// 13.8: The continue Statement

continueStatement
    : Continue labelIdentifier? SemiColon
    ;

continueStatement_Yield
    : Continue labelIdentifier_Yield? SemiColon
    ;

continueStatement_Await
    : Continue labelIdentifier_Await? SemiColon
    ;

continueStatement_Yield_Await
    : Continue labelIdentifier_Yield_Await? SemiColon
    ;

// 13.9: The break Statement

breakStatement
    : Break labelIdentifier? SemiColon
    ;

breakStatement_Yield
    : Break labelIdentifier_Yield? SemiColon
    ;

breakStatement_Await
    : Break labelIdentifier_Await? SemiColon
    ;

breakStatement_Yield_Await
    : Break labelIdentifier_Yield_Await? SemiColon
    ;

// 13.10: The return Statement

returnStatement
    : Return expression_In? SemiColon
    ;

returnStatement_Yield
    : Return expression_In_Yield? SemiColon
    ;

returnStatement_Await
    : Return expression_In_Await? SemiColon
    ;

returnStatement_Yield_Await
    : Return expression_In_Yield_Await? SemiColon
    ;

// 13.11 The with Statement

withStatement
    : With OpenParen expression_In CloseParen statement
    ;

withStatement_Yield
    : With OpenParen expression_In_Yield CloseParen statement_Yield
    ;

withStatement_Await
    : With OpenParen expression_In_Await CloseParen statement_Await
    ;

withStatement_Yield_Await
    : With OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await
    ;

withStatement_Return
    : With OpenParen expression_In CloseParen statement_Return
    ;

withStatement_Yield_Return
    : With OpenParen expression_In_Yield CloseParen statement_Yield_Return
    ;

withStatement_Await_Return
    : With OpenParen expression_In_Await CloseParen statement_Await_Return
    ;

withStatement_Yield_Await_Return
    : With OpenParen expression_In_Yield_Await CloseParen statement_Yield_Await_Return
    ;

// 13.12: The switch Statement

switchStatement
    : Switch OpenParen expression_In CloseParen caseBlock
    ;

switchStatement_Yield
    : Switch OpenParen expression_In_Yield CloseParen caseBlock_Yield
    ;

switchStatement_Await
    : Switch OpenParen expression_In_Await CloseParen caseBlock_Await
    ;

switchStatement_Yield_Await
    : Switch OpenParen expression_In_Yield_Await CloseParen caseBlock_Yield_Await
    ;

switchStatement_Return
    : Switch OpenParen expression_In CloseParen caseBlock_Return
    ;

switchStatement_Yield_Return
    : Switch OpenParen expression_In_Yield CloseParen caseBlock_Yield_Return
    ;

switchStatement_Await_Return
    : Switch OpenParen expression_In_Await CloseParen caseBlock_Await_Return
    ;

switchStatement_Yield_Await_Return
    : Switch OpenParen expression_In_Yield_Await CloseParen caseBlock_Yield_Await_Return
    ;

caseBlock
    : OpenBrace caseClause* CloseBrace
    | OpenBrace caseClause* defaultClause caseClause* CloseBrace
    ;

caseBlock_Yield
    : OpenBrace caseClause_Yield* CloseBrace
    | OpenBrace caseClause_Yield* defaultClause_Yield caseClause_Yield* CloseBrace
    ;

caseBlock_Await
    : OpenBrace caseClause_Await* CloseBrace
    | OpenBrace caseClause_Await* defaultClause_Await caseClause_Await* CloseBrace
    ;

caseBlock_Yield_Await
    : OpenBrace caseClause_Yield_Await* CloseBrace
    | OpenBrace caseClause_Yield_Await* defaultClause_Yield_Await caseClause_Yield_Await* CloseBrace
    ;

caseBlock_Return
    : OpenBrace caseClause_Return* CloseBrace
    | OpenBrace caseClause_Return* defaultClause_Return caseClause_Return* CloseBrace
    ;

caseBlock_Yield_Return
    : OpenBrace caseClause_Yield_Return* CloseBrace
    | OpenBrace caseClause_Yield_Return* defaultClause_Yield_Return caseClause_Yield_Return* CloseBrace
    ;

caseBlock_Await_Return
    : OpenBrace caseClause_Await_Return* CloseBrace
    | OpenBrace caseClause_Await_Return* defaultClause_Await_Return caseClause_Await_Return* CloseBrace
    ;

caseBlock_Yield_Await_Return
    : OpenBrace caseClause_Yield_Await_Return* CloseBrace
    | OpenBrace caseClause_Yield_Await_Return* defaultClause_Yield_Await_Return caseClause_Yield_Await_Return* CloseBrace
    ;

caseClause
    : Case expression_In Colon statementList?
    ;

caseClause_Yield
    : Case expression_In_Yield Colon statementList_Yield?
    ;

caseClause_Await
    : Case expression_In_Await Colon statementList_Await?
    ;

caseClause_Yield_Await
    : Case expression_In_Yield_Await Colon statementList_Yield_Await?
    ;

caseClause_Return
    : Case expression_In Colon statementList_Return?
    ;

caseClause_Yield_Return
    : Case expression_In_Yield Colon statementList_Yield_Return?
    ;

caseClause_Await_Return
    : Case expression_In_Await Colon statementList_Await_Return?
    ;

caseClause_Yield_Await_Return
    : Case expression_In_Yield_Await Colon statementList_Yield_Await_Return?
    ;

defaultClause
    : Default Colon statementList?
    ;

defaultClause_Yield
    : Default Colon statementList_Yield?
    ;

defaultClause_Await
    : Default Colon statementList_Await?
    ;

defaultClause_Yield_Await
    : Default Colon statementList_Yield_Await?
    ;

defaultClause_Return
    : Default Colon statementList_Return?
    ;

defaultClause_Yield_Return
    : Default Colon statementList_Yield_Return?
    ;

defaultClause_Await_Return
    : Default Colon statementList_Await_Return?
    ;

defaultClause_Yield_Await_Return
    : Default Colon statementList_Yield_Await_Return?
    ;

// 13.13: Labelled Statements

labelledStatement
    : labelIdentifier Colon labelledItem
    ;

labelledStatement_Yield
    : labelIdentifier_Yield Colon labelledItem_Yield
    ;

labelledStatement_Await
    : labelIdentifier_Await Colon labelledItem_Await
    ;

labelledStatement_Yield_Await
    : labelIdentifier_Yield_Await Colon labelledItem_Yield_Await
    ;

labelledStatement_Return
    : labelIdentifier Colon labelledItem_Return
    ;

labelledStatement_Yield_Return
    : labelIdentifier_Yield Colon labelledItem_Yield_Return
    ;

labelledStatement_Await_Return
    : labelIdentifier_Await Colon labelledItem_Await_Return
    ;

labelledStatement_Yield_Await_Return
    : labelIdentifier_Yield_Await Colon labelledItem_Yield_Await_Return
    ;

labelledItem
    : statement
    | functionDeclaration
    ;

labelledItem_Yield
    : statement_Yield
    | functionDeclaration_Yield
    ;

labelledItem_Await
    : statement_Await
    | functionDeclaration_Await
    ;

labelledItem_Yield_Await
    : statement_Yield_Await
    | functionDeclaration_Yield_Await
    ;

labelledItem_Return
    : statement_Return
    | functionDeclaration
    ;

labelledItem_Yield_Return
    : statement_Yield_Return
    | functionDeclaration_Yield
    ;

labelledItem_Await_Return
    : statement_Await_Return
    | functionDeclaration_Await
    ;

labelledItem_Yield_Await_Return
    : statement_Yield_Await_Return
    | functionDeclaration_Yield_Await
    ;

// 13.14: The throw Statement

throwStatement
    : Throw expression_In SemiColon
    ;

throwStatement_Yield
    : Throw expression_In_Yield SemiColon
    ;

throwStatement_Await
    : Throw expression_In_Await SemiColon
    ;

throwStatement_Yield_Await
    : Throw expression_In_Yield_Await SemiColon
    ;

// 13.15: The try Statement

tryStatement
    : Try block catch_
    | Try block finally_
    | Try block catch_ finally_
    ;

tryStatement_Yield
    : Try block_Yield catch_Yield
    | Try block_Yield finally_Yield
    | Try block_Yield catch_Yield finally_Yield
    ;

tryStatement_Await
    : Try block_Await catch_Await
    | Try block_Await finally_Await
    | Try block_Await catch_Await finally_Await
    ;

tryStatement_Yield_Await
    : Try block_Yield_Await catch_Yield_Await
    | Try block_Yield_Await finally_Yield_Await
    | Try block_Yield_Await catch_Yield_Await finally_Yield_Await
    ;

tryStatement_Return
    : Try block_Return catch_Return
    | Try block_Return finally_Return
    | Try block_Return catch_Return finally_Return
    ;

tryStatement_Yield_Return
    : Try block_Yield_Return catch_Yield_Return
    | Try block_Yield_Return finally_Yield_Return
    | Try block_Yield_Return catch_Yield_Return finally_Yield_Return
    ;

tryStatement_Await_Return
    : Try block_Await_Return catch_Await_Return
    | Try block_Await_Return finally_Await_Return
    | Try block_Await_Return catch_Await_Return finally_Await_Return
    ;

tryStatement_Yield_Await_Return
    : Try block_Yield_Await_Return catch_Yield_Await_Return
    | Try block_Yield_Await_Return finally_Yield_Await_Return
    | Try block_Yield_Await_Return catch_Yield_Await_Return finally_Yield_Await_Return
    ;

catch_
    : Catch OpenParen catchParameter CloseParen block
    ;

catch_Yield
    : Catch OpenParen catchParameter_Yield CloseParen block_Yield
    ;

catch_Await
    : Catch OpenParen catchParameter_Await CloseParen block_Await
    ;

catch_Yield_Await
    : Catch OpenParen catchParameter_Yield_Await CloseParen block_Yield_Await
    ;

catch_Return
    : Catch OpenParen catchParameter CloseParen block_Return
    ;

catch_Yield_Return
    : Catch OpenParen catchParameter_Yield CloseParen block_Yield_Return
    ;

catch_Await_Return
    : Catch OpenParen catchParameter_Await CloseParen block_Await_Return
    ;

catch_Yield_Await_Return
    : Catch OpenParen catchParameter_Yield_Await CloseParen block_Yield_Await_Return
    ;

finally_
    : Finally block
    ;

finally_Yield
    : Finally block_Yield
    ;

finally_Await
    : Finally block_Await
    ;

finally_Yield_Await
    : Finally block_Yield_Await
    ;

finally_Return
    : Finally block_Return
    ;

finally_Yield_Return
    : Finally block_Yield_Return
    ;

finally_Await_Return
    : Finally block_Await_Return
    ;

finally_Yield_Await_Return
    : Finally block_Yield_Await_Return
    ;

catchParameter
    : bindingIdentifier
    | bindingPattern
    ;

catchParameter_Yield
    : bindingIdentifier_Yield
    | bindingPattern_Yield
    ;

catchParameter_Await
    : bindingIdentifier_Await
    | bindingPattern_Await
    ;

catchParameter_Yield_Await
    : bindingIdentifier_Yield_Await
    | bindingPattern_Yield_Await
    ;

// 13.16: The debugger Statement

debuggerStatement
    : Debugger SemiColon
    ;

// ******************************************************
// *** 14: ECMAScript Language: Functions and Classes ***
// ******************************************************

// 14.1: Function Definitions

functionDeclaration
    : Function bindingIdentifier OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Yield
    : Function bindingIdentifier_Yield OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Await
    : Function bindingIdentifier_Await OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Yield_Await
    : Function bindingIdentifier_Yield_Await OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Default
    : Function bindingIdentifier OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    | Function OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Yield_Default
    : Function bindingIdentifier_Yield OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    | Function OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Await_Default
    : Function bindingIdentifier_Await OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    | Function OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionDeclaration_Yield_Await_Default
    : Function bindingIdentifier_Yield_Await OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    | Function OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

functionExpression
    : Function bindingIdentifier? OpenParen formalParameters CloseParen OpenBrace functionBody CloseBrace
    ;

uniqueFormalParameters
    : formalParameters
    ;

uniqueFormalParameters_Yield
    : formalParameters_Yield
    ;

uniqueFormalParameters_Await
    : formalParameters_Await
    ;

uniqueFormalParameters_Yield_Await
    : formalParameters_Yield_Await
    ;

formalParameters
    : 
    | functionRestParameter
    | formalParameterList
    | formalParameterList Comma
    | formalParameterList Comma functionRestParameter
    ;

formalParameters_Yield
    : 
    | functionRestParameter_Yield
    | formalParameterList_Yield
    | formalParameterList_Yield Comma
    | formalParameterList_Yield Comma functionRestParameter_Yield
    ;

formalParameters_Await
    : 
    | functionRestParameter_Await
    | formalParameterList_Await
    | formalParameterList_Await Comma
    | formalParameterList_Await Comma functionRestParameter_Await
    ;

formalParameters_Yield_Await
    : 
    | functionRestParameter_Yield_Await
    | formalParameterList_Yield_Await
    | formalParameterList_Yield_Await Comma
    | formalParameterList_Yield_Await Comma functionRestParameter_Yield_Await
    ;

formalParameterList
    : formalParameter
    | formalParameterList Comma formalParameter
    ;

formalParameterList_Yield
    : formalParameter_Yield
    | formalParameterList_Yield Comma formalParameter_Yield
    ;

formalParameterList_Await
    : formalParameter_Await
    | formalParameterList_Await Comma formalParameter_Await
    ;

formalParameterList_Yield_Await
    : formalParameter_Yield_Await
    | formalParameterList_Yield_Await Comma formalParameter_Yield_Await
    ;

functionRestParameter
    : bindingRestElement
    ;

functionRestParameter_Yield
    : bindingRestElement_Yield
    ;

functionRestParameter_Await
    : bindingRestElement_Await
    ;

functionRestParameter_Yield_Await
    : bindingRestElement_Yield_Await
    ;

formalParameter
    : bindingElement
    ;

formalParameter_Yield
    : bindingElement_Yield
    ;

formalParameter_Await
    : bindingElement_Await
    ;

formalParameter_Yield_Await
    : bindingElement_Yield_Await
    ;

functionBody
    : functionStatementList
    ;

functionBody_Yield
    : functionStatementList_Yield
    ;

functionBody_Await
    : functionStatementList_Await
    ;

functionBody_Yield_Await
    : functionStatementList_Yield_Await
    ;

functionStatementList
    : statementList_Return?
    ;

functionStatementList_Yield
    : statementList_Yield_Return?
    ;

functionStatementList_Await
    : statementList_Await_Return?
    ;

functionStatementList_Yield_Await
    : statementList_Yield_Await_Return?
    ;

// 14.2: Arrow Function Definitions

arrowFunction
    : arrowParameters Arrow conciseBody
    ;

arrowFunction_In
    : arrowParameters Arrow conciseBody_In
    ;

arrowFunction_Yield
    : arrowParameters_Yield Arrow conciseBody
    ;

arrowFunction_In_Yield
    : arrowParameters_Yield Arrow conciseBody_In
    ;

arrowFunction_Await
    : arrowParameters_Await Arrow conciseBody
    ;

arrowFunction_In_Await
    : arrowParameters_Await Arrow conciseBody_In
    ;

arrowFunction_Yield_Await
    : arrowParameters_Yield_Await Arrow conciseBody
    ;

arrowFunction_In_Yield_Await
    : arrowParameters_Yield_Await Arrow conciseBody_In
    ;

arrowParameters
    : bindingIdentifier
    | coverParenthesizedExpressionAndArrowParameterList
    ;

arrowParameters_Yield
    : bindingIdentifier_Yield
    | coverParenthesizedExpressionAndArrowParameterList_Yield
    ;

arrowParameters_Await
    : bindingIdentifier_Await
    | coverParenthesizedExpressionAndArrowParameterList_Await
    ;

arrowParameters_Yield_Await
    : bindingIdentifier_Yield_Await
    | coverParenthesizedExpressionAndArrowParameterList_Yield_Await
    ;

conciseBody
    : { p.negativeLookahead(ECMAScriptLexerOpenBrace) }? assignmentExpression
    | OpenBrace functionBody CloseBrace
    ;

conciseBody_In
    : { p.negativeLookahead(ECMAScriptLexerOpenBrace) }? assignmentExpression_In
    | OpenBrace functionBody CloseBrace
    ;

// 14.3: Method Definitions

methodDefinition
    : propertyName OpenParen uniqueFormalParameters CloseParen OpenBrace functionBody CloseBrace
    | generatorMethod
    | asyncMethod
    | asyncGeneratorMethod
    | Get propertyName OpenParen CloseParen OpenBrace functionBody CloseBrace
    | Set propertyName OpenBrace propertySetParameterList CloseParen OpenBrace functionBody CloseBrace
    ;

methodDefinition_Yield
    : propertyName_Yield OpenParen uniqueFormalParameters CloseParen OpenBrace functionBody CloseBrace
    | generatorMethod_Yield
    | asyncMethod_Yield
    | asyncGeneratorMethod_Yield
    | Get propertyName_Yield OpenParen CloseParen OpenBrace functionBody CloseBrace
    | Set propertyName_Yield OpenBrace propertySetParameterList CloseParen OpenBrace functionBody CloseBrace
    ;

methodDefinition_Await
    : propertyName_Await OpenParen uniqueFormalParameters CloseParen OpenBrace functionBody CloseBrace
    | generatorMethod_Await
    | asyncMethod_Await
    | asyncGeneratorMethod_Await
    | Get propertyName_Await OpenParen CloseParen OpenBrace functionBody CloseBrace
    | Set propertyName_Await OpenBrace propertySetParameterList CloseParen OpenBrace functionBody CloseBrace
    ;

methodDefinition_Yield_Await
    : propertyName_Yield_Await OpenParen uniqueFormalParameters CloseParen OpenBrace functionBody CloseBrace
    | generatorMethod_Yield_Await
    | asyncMethod_Yield_Await
    | asyncGeneratorMethod_Yield_Await
    | Get propertyName_Yield_Await OpenParen CloseParen OpenBrace functionBody CloseBrace
    | Set propertyName_Yield_Await OpenBrace propertySetParameterList CloseParen OpenBrace functionBody CloseBrace
    ;

propertySetParameterList
    : formalParameter
    ;

// 14.4: Generator Function Definitions

generatorMethod
    : Multiply propertyName OpenParen uniqueFormalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorMethod_Yield
    : Multiply propertyName_Yield OpenParen uniqueFormalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorMethod_Await
    : Multiply propertyName_Await OpenParen uniqueFormalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorMethod_Yield_Await
    : Multiply propertyName_Yield_Await OpenParen uniqueFormalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration
    : Function Multiply bindingIdentifier OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Yield
    : Function Multiply bindingIdentifier_Yield OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Await
    : Function Multiply bindingIdentifier_Await OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Yield_Await
    : Function Multiply bindingIdentifier_Yield_Await OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Default
    : Function Multiply bindingIdentifier OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    | Function Multiply OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Yield_Default
    : Function Multiply bindingIdentifier_Yield OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    | Function Multiply OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Await_Default
    : Function Multiply bindingIdentifier_Await OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    | Function Multiply OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorDeclaration_Yield_Await_Default
    : Function Multiply bindingIdentifier_Yield_Await OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    | Function Multiply OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorExpression
    : Function Multiply bindingIdentifier_Yield? OpenParen formalParameters_Yield CloseParen OpenBrace generatorBody CloseBrace
    ;

generatorBody
    : functionBody_Yield
    ;

yieldExpression
    : Yield
    | Yield Multiply? assignmentExpression_Yield
    ;

yieldExpression_In
    : Yield
    | Yield Multiply? assignmentExpression_In_Yield
    ;

yieldExpression_Await
    : Yield
    | Yield Multiply? assignmentExpression_Yield_Await
    ;

yieldExpression_In_Await
    : Yield
    | Yield Multiply? assignmentExpression_In_Yield_Await
    ;

// 14.5: Async Generator Function Definitions

asyncGeneratorMethod
    : Async Multiply propertyName OpenParen uniqueFormalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorMethod_Yield
    : Async Multiply propertyName_Yield OpenParen uniqueFormalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorMethod_Await
    : Async Multiply propertyName_Await OpenParen uniqueFormalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorMethod_Yield_Await
    : Async Multiply propertyName_Yield_Await OpenParen uniqueFormalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration
    : Async Function Multiply bindingIdentifier OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Yield
    : Async Function Multiply bindingIdentifier_Yield OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Await
    : Async Function Multiply bindingIdentifier_Await OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Yield_Await
    : Async Function Multiply bindingIdentifier_Yield_Await OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Default
    : Async Function Multiply bindingIdentifier OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    | Async Function Multiply OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Yield_Default
    : Async Function Multiply bindingIdentifier_Yield OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    | Async Function Multiply OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Await_Default
    : Async Function Multiply bindingIdentifier_Await OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    | Async Function Multiply OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorDeclaration_Yield_Await_Default
    : Async Function Multiply bindingIdentifier_Yield_Await OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    | Async Function Multiply OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorExpression
    : Async Function Multiply bindingIdentifier_Yield_Await? OpenParen formalParameters_Yield_Await CloseParen OpenBrace asyncGeneratorBody CloseBrace
    ;

asyncGeneratorBody
    : functionBody_Yield_Await
    ;

// 14.6: Class Definitions

classDeclaration
    : Class bindingIdentifier classTail
    ;

classDeclaration_Yield
    : Class bindingIdentifier_Yield classTail_Yield
    ;

classDeclaration_Await
    : Class bindingIdentifier_Await classTail_Await
    ;

classDeclaration_Yield_Await
    : Class bindingIdentifier_Yield_Await classTail_Yield_Await
    ;

classDeclaration_Default
    : Class bindingIdentifier classTail
    | Class classTail
    ;

classDeclaration_Yield_Default
    : Class bindingIdentifier_Yield classTail_Yield
    | Class classTail_Yield
    ;

classDeclaration_Await_Default
    : Class bindingIdentifier_Await classTail_Await
    | Class classTail_Await
    ;

classDeclaration_Yield_Await_Default
    : Class bindingIdentifier_Yield_Await classTail_Yield_Await
    | Class classTail_Yield_Await
    ;

classExpression
    : Class bindingIdentifier? classTail
    ;

classExpression_Yield
    : Class bindingIdentifier_Yield? classTail_Yield
    ;

classExpression_Await
    : Class bindingIdentifier_Await? classTail_Await
    ;

classExpression_Yield_Await
    : Class bindingIdentifier_Yield_Await? classTail_Yield_Await
    ;

classTail
    : classHeritage? OpenBrace classBody? CloseBrace
    ;

classTail_Yield
    : classHeritage_Yield? OpenBrace classBody_Yield? CloseBrace
    ;

classTail_Await
    : classHeritage_Await? OpenBrace classBody_Await? CloseBrace
    ;

classTail_Yield_Await
    : classHeritage_Yield_Await? OpenBrace classBody_Yield_Await? CloseBrace
    ;

classHeritage
    : Extends leftHandSideExpression
    ;

classHeritage_Yield
    : Extends leftHandSideExpression_Yield
    ;

classHeritage_Await
    : Extends leftHandSideExpression_Await
    ;

classHeritage_Yield_Await
    : Extends leftHandSideExpression_Yield_Await
    ;

classBody
    : classElement+
    ;

classBody_Yield
    : classElement_Yield+
    ;

classBody_Await
    : classElement_Await+
    ;

classBody_Yield_Await
    : classElement_Yield_Await+
    ;

classElement
    : Static? methodDefinition
    | SemiColon
    ;

classElement_Yield
    : Static? methodDefinition_Yield
    | SemiColon
    ;

classElement_Await
    : Static? methodDefinition_Await
    | SemiColon
    ;

classElement_Yield_Await
    : Static? methodDefinition_Yield_Await
    | SemiColon
    ;

// 14.7: Async Function Definitions

asyncFunctionDeclaration
    : Async Function bindingIdentifier OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Yield
    : Async Function bindingIdentifier_Yield OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Await
    : Async Function bindingIdentifier_Await OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Yield_Await
    : Async Function bindingIdentifier_Yield_Await OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Default
    : Async Function bindingIdentifier OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    | Async Function OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Yield_Default
    : Async Function bindingIdentifier_Yield OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    | Async Function OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Await_Default
    : Async Function bindingIdentifier_Await OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    | Async Function OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionDeclaration_Yield_Await_Default
    : Async Function bindingIdentifier_Yield_Await OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    | Async Function OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionExpression
    : Async Function bindingIdentifier_Await? OpenParen formalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncMethod
    : Async propertyName OpenParen uniqueFormalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncMethod_Yield
    : Async propertyName_Yield OpenParen uniqueFormalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncMethod_Await
    : Async propertyName_Await OpenParen uniqueFormalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncMethod_Yield_Await
    : Async propertyName_Yield_Await OpenParen uniqueFormalParameters_Await CloseParen OpenBrace asyncFunctionBody CloseBrace
    ;

asyncFunctionBody
    : functionBody_Await
    ;

awaitExpression
    : Await unaryExpression_Await
    ;

awaitExpression_Yield
    : Await unaryExpression_Yield_Await
    ;

// 14.8: Async Arrow Function Definitions

asyncArrowFunction
    : Async asyncArrowBindingIdentifier Arrow asyncConciseBody
    | coverCallExpressionAndAsyncArrowHead Arrow asyncConciseBody
    ;

asyncArrowFunction_In
    : Async asyncArrowBindingIdentifier Arrow asyncConciseBody_In
    | coverCallExpressionAndAsyncArrowHead Arrow asyncConciseBody_In
    ;

asyncArrowFunction_Yield
    : Async asyncArrowBindingIdentifier_Yield Arrow asyncConciseBody
    | coverCallExpressionAndAsyncArrowHead_Yield Arrow asyncConciseBody
    ;

asyncArrowFunction_In_Yield
    : Async asyncArrowBindingIdentifier_Yield Arrow asyncConciseBody_In
    | coverCallExpressionAndAsyncArrowHead_Yield Arrow asyncConciseBody_In
    ;

asyncArrowFunction_Await
    : Async asyncArrowBindingIdentifier Arrow asyncConciseBody
    | coverCallExpressionAndAsyncArrowHead_Await Arrow asyncConciseBody
    ;

asyncArrowFunction_In_Await
    : Async asyncArrowBindingIdentifier Arrow asyncConciseBody_In
    | coverCallExpressionAndAsyncArrowHead_Await Arrow asyncConciseBody_In
    ;

asyncArrowFunction_Yield_Await
    : Async asyncArrowBindingIdentifier_Yield Arrow asyncConciseBody
    | coverCallExpressionAndAsyncArrowHead_Yield_Await Arrow asyncConciseBody
    ;

asyncArrowFunction_In_Yield_Await
    : Async asyncArrowBindingIdentifier_Yield Arrow asyncConciseBody_In
    | coverCallExpressionAndAsyncArrowHead_Yield_Await Arrow asyncConciseBody_In
    ;

asyncArrowBindingIdentifier
    : bindingIdentifier_Await
    ;

asyncArrowBindingIdentifier_Yield
    : bindingIdentifier_Yield_Await
    ;

coverCallExpressionAndAsyncArrowHead
    : memberExpression arguments
    ;

coverCallExpressionAndAsyncArrowHead_Yield
    : memberExpression_Yield arguments_Yield
    ;

coverCallExpressionAndAsyncArrowHead_Await
    : memberExpression_Await arguments_Await
    ;

coverCallExpressionAndAsyncArrowHead_Yield_Await
    : memberExpression_Yield_Await arguments_Yield_Await
    ;

// ***************************************************
// *** 15: ECMAScriptLanguage: Scripts and Modules ***
// ***************************************************

// 15.1: Scripts

script
    : scriptBody?
    ;

scriptBody
    : statementList
    ;

// 15.2: Modules

module
    : moduleBody?
    ;

moduleBody
    : moduleItem+
    ;

moduleItem
    : importDeclaration
    | exportDeclaration
    | statementListItem
    ;

// 15.2.2: Imports

importDeclaration
    : Import importClause fromClause SemiColon
    | Import moduleSpecifier SemiColon
    ;

importClause
    : importedDefaultBinding
    | nameSpaceImport
    | namedImports
    | importedDefaultBinding Comma nameSpaceImport
    | importedDefaultBinding Comma namedImports
    ;

importedDefaultBinding
    : importedBinding
    ;

nameSpaceImport
    : Multiply As importedBinding
    ;

namedImports
    : OpenBrace CloseBrace
    | OpenBrace importsList Comma? CloseBrace
    ;

fromClause
    : From moduleSpecifier
    ;

importsList
    : importSpecifier
    | importsList Comma importSpecifier
    ;

importSpecifier
    : importedBinding
    | IdentifierName As importedBinding
    ;

moduleSpecifier
    : StringLiteral
    ;

importedBinding
    : bindingIdentifier
    ;

// 15.2.3: Exports

exportDeclaration
    : Export Multiply fromClause SemiColon
    | Export exportClause fromClause? SemiColon
    | Export variableStatement
    | Export declaration
    | Export Default hoistableDeclaration_Default
    | Export Default classDeclaration_Default
    | Export Default { p.negativeLookahead(ECMAScriptLexerFunction, ECMAScriptLexerAsyncFunction, ECMAScriptLexerClass) }? assignmentExpression_In SemiColon
    ;

exportClause
    : OpenBrace CloseBrace
    | OpenBrace exportsList Comma? CloseBrace
    ;

exportsList
    : exportSpecifier
    | exportsList Comma exportSpecifier
    ;

exportSpecifier
    : IdentifierName
    | IdentifierName As IdentifierName
    ;

// A.4: Functions and Classes

asyncConciseBody
    : { p.negativeLookahead(ECMAScriptLexerOpenBrace) }? assignmentExpression_Await OpenBrace asyncFunctionBody CloseBrace
    ;

asyncConciseBody_In
    : { p.negativeLookahead(ECMAScriptLexerOpenBrace) }? assignmentExpression_In_Await OpenBrace asyncFunctionBody CloseBrace
    ;