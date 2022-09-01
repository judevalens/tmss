grammar common;
   OCTET           :  [\u0000-\u00FF] ; //any 8-bit sequence of data
   CHAR            :  [\u0001-\u007F] ; //any US-ASCII character (octets 1 - 127)
   UPALPHA         :  [\u0041-\u005A] ; //any US-ASCII uppercase letter "A".."Z"
   LOALPHA         :  [\u0061-\u007A] ; //any US-ASCII lowercase letter "a".."z"
   ALPHA           :  UPALPHA | LOALPHA;
   DIGIT           :  [\u0030-\u0039] ; //any US-ASCII digit "0".."9"
   CTL             :  [\u0000-\u001F] | [\u007F]  ; //any US-ASCII control character; (octets 0 - 31) and DEL (127)
   CR              :  [\u000D] ; //US-ASCII CR, carriage return (13)
   LF              :  [\u000A] ; //US-ASCII LF, linefeed (10)
   SP              :  [\u0020] ; //US-ASCII SP, space (32)
   HT              :  [\u0009] ; //US-ASCII HT, horizontal-tab (9)
   BACKSLASH       :  [\u005C] ; //US-ASCII backslash (92)
   CRLF            :  CR LF;
   LWS             :  CRLF? ( SP | HT )+ ; //Line-breaking whitespace
   SWS             :  LWS? ; //Separating whitespace
   HCOLON          :  ( SP | HT )* ':' SWS;
   TEXT            :  [\u0020-\u007E] | [\u0080-\u00FF] ; //any OCTET except CTLs
   TS_SEPCIALS       :  '(' | ')' | '<' | '>' | '@'
                       |  ',' | ';' | ':' | BACKSLASH | DQUOTE
                       |  '/' | '[' | ']'| '?' | '='
                       |  '{' | '}' | SP | HT;
   TOKEN           :  [\u0021\u0023-\u0027\u002D-\u002E\u0030-\u0039\u0041-\u005A\u005E-\u007A\u007C\u007E]; //1*<any CHAR except CTLs or tspecials>

   QDTEXT          : [\u0020-\u0021] | [\u0023-\u005B] | [\u005D-\u007E] | QUOTED_PAIR ; //No DQUOTE and no "\"
   QUOTED_PAIR     : '\\' | ( '\\' DQUOTE );
   CTEXT           :  [\u0020-\u0027] | [\u002A-\u007E] | [\u0080-\u00FF]; //any OCTET except CTLs, "(" and ")"

   GEN_VALUE       :  TEXT_UTF8char+;
   DQUOTE          : '"';


   SAFE            :  '$' | '-' | '_' | '.' | '+';
   EXTRA           :  '!' | '*' | '\'' | '(' | ')' | ',';
   RTSP_EXTRA      : '!' | '*' | '\'' | '(' | ')';

   HEX             : DIGIT | 'A'| 'B'| 'C' | 'D' | 'E'| 'F' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f';
   RESERVED        : ';' | '/' | '?' | ':' | '@' | '&' | '=';

   UNRESERVED      :  ALPHA | DIGIT | SAFE | EXTRA;
   RTSP_UNRESERVED :  ALPHA | DIGIT | SAFE | RTSP_EXTRA;


   BASE64_CHAR     :  ALPHA | DIGIT | '+' | '/';
   SLASH           :  SWS '/' SWS ; //slash
   EQUAL           :  SWS '=' SWS ; //equal
   LPAREN          :  SWS '(' SWS ; //left parenthesis
   RPAREN          :  SWS ')' SWS ; //right parenthesis
   COMMA           :  SWS ',' SWS ; //comma
   SEMI            :  SWS ';' SWS ; //semicolon
   COLON           :  SWS ':' SWS ; //colon
   MINUS           :  SWS '-' SWS ; //minus/dash
   LDQUOT          :  SWS DQUOTE ; //open double quotation mark
   RDQUOT          :  DQUOTE SWS ; //close double quotation mark
   RAQUOT          :  '>' SWS ; //right angle quote
   LAQUOT          :  SWS '<' ; //left angle quote
   TEXT_UTF8char   :  [\u0021-\u007E];
   POS_FLOAT        : DIGIT+ ('.' DIGIT+);
   FLOAT            : '-' POS_FLOAT;


   token           :  TOKEN+;
   quoted_string   :  ( DQUOTE QDTEXT* DQUOTE );
   base64          :  base64_unit* base64_pad?;
   base64_unit     :  BASE64_CHAR BASE64_CHAR BASE64_CHAR BASE64_CHAR;
   base64_pad      :  (BASE64_CHAR BASE64_CHAR EQUAL) | (BASE64_CHAR BASE64_CHAR BASE64_CHAR BASE64_CHAR BASE64_CHAR BASE64_CHAR EQUAL);
   generic_param   :  token ( EQUAL GEN_VALUE )?;
   lhex            : DIGIT | 'a' | 'b' | 'c' | 'd' | 'e' | 'f'; //lowercase "a-f" Hex
