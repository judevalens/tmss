grammar rtsp_request;
//RTSP_REQUEST_STATUS_LINE : REQUEST_LINE REQUEST_URI RTSP_VERSION NEW_LINE;
//REQUEST_URI: ;
HCOLON: ':';

request: request_line eor (header)* eor;

request_line: RTSP_METHOD SP RTSP_URI SP RTSP_VERSION;

header: header_name ':' header_value eor;

header_name: (ID | HEADER_NAME)*;
header_value: (HEADER_VALUE | HEADER_NAME | ID)*;

eor: LF | CR | (CRLF);

HEADER_NAME:  TOKEN (TOKEN)*?;
HEADER_VALUE: TEXT (TEXT)*?;




RTSP_METHOD: ('DESCRIBE' |'GET_PARAMETER' |'OPTIONS' |'PAUSE' |'PLAY' |'PLAY_NOTIFY' |'REDIRECT' |'SETUP' |'SET_PARAMETER' |'TEARDOWN');


//TS_SPECIALS:  ('(' | ')' | '<' | '>' | '@' | ',' | ';' | ':' | '\\' | '"' | '/' | '[' | ']' | '?' | '='| '{' | '}');
//TS_SPECIALS_S:  [()<>@,;:\\"/[\]?={}];

RTSP_VERSION: 'RTSP/' INT'.'INT;

//RTSP_URI
RTSP_URI: URI_SCHEME URI_DELIMETER URI_HOST URI_PATH;
URI_SCHEME: 'rtsp' | 'rtspu';
URI_HOST: IP | HOST;
URI_DELIMETER: '://';
URI_PATH: '/' (TEXT)*;
HOST: ID '.' ID;
IP: IP_PART '.' IP_PART '.' IP_PART;
IP_PART: DIGIT | (DIGIT DIGIT) | (DIGIT DIGIT DIGIT);

///RTSP_HEADER


//GENERICS
fragment LETTER: [a-zA-Z];
ALPHA: [a-zA-Z0-9];
fragment DIGIT: [0-9];
TOKEN: [\u0021\u0023-\u0027\u002D-\u002E\u0030-\u0039\u0041-\u005A\u005E-\u007A\u007C\u007E];
INT: DIGIT+;
ID: LETTER (LETTER|DIGIT)*;
SP: ' ';
CR: '\r';
LF: '\n';
CRLF: CR | LF | (CR LF);
CTL : [\u0000-\u001F];
TEXT: [\u0021-\u007E];