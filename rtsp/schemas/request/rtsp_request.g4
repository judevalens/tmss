grammar rtsp_request;
//RTSP_REQUEST_STATUS_LINE : REQUEST_LINE REQUEST_URI RTSP_VERSION NEW_LINE;
//REQUEST_URI: ;


//TS_SPECIALS:  ('(' | ')' | '<' | '>' | '@' | ',' | ';' | ':' | '\\' | '"' | '/' | '[' | ']' | '?' | '='| '{' | '}');
//TS_SPECIALS_S:  [()<>@,;:\\"/[\]?={}];

RTSP_VERSION: 'RTSP/' INT'.'INT;

//RTSP_URI

URI_SCHEME: 'rtsp' | 'rtspu';
URI_DELIMETER: '://';

IP: INT '.' INT '.' INT;


///RTSP_HEADER



//BASE SYNTAX
HCOLON: ':';
SLASH: '/';
DOT: '.';

fragment DIGIT: [0-9];
LETTER: [a-zA-Z];
TOKEN: [\u0021\u0023-\u0027\u002D-\u002E\u0030-\u0039\u0041-\u005A\u005E-\u007A\u007C\u007E];
TEXT: [\u0021-\u007E];
INT: DIGIT+;
ID: LETTER (LETTER|DIGIT)*?;
ALPHA: [a-zA-Z0-9];
SP: ' ';
CR: '\r';
LF: '\n';
CRLF: CR | LF | (CR LF);
CTL : [\u0000-\u001F];


//PARSER RULES
request: requestLine eor (header eor)* eor;
requestLine: method SP rtspUri SP RTSP_VERSION;
method: LETTER+;

response: statusLine eor (header eor)* eor;
statusLine: RTSP_VERSION SP status_code SP status_reason;
status_reason: ((SLASH | DOT | LETTER | INT | TEXT | TOKEN| HCOLON) | SP)+;
status_code: INT;

header: headerName ':' headerValue;
headerName: (TOKEN | LETTER | INT)+ SP*;
headerValue: (SP | SLASH | DOT | LETTER | INT | TEXT | TOKEN| HCOLON)*;
//uri_rules
rtspUri: URI_SCHEME URI_DELIMETER uriHost uriPath;
uriPath: (SLASH (LETTER | INT | TEXT | TOKEN | DOT)*) uriPath*;
uriName: (TEXT | INT | TOKEN | LETTER)* DOT (LETTER | INT)+;
uriHost: IP|uriName;


eor: LF | CR | (CRLF);
