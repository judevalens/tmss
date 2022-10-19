grammar header;
IP: INT '.' INT '.' INT;


DQUOTE: '"';
HT: [\u0009];
HCOLON : ( SP | HT )* ':' LWS?;
LWS: CRLF? ( SP | HT )+;
SEMI: LWS? ';' LWS? ; //semicolon
COMMA: LWS? ',' LWS?;
DOT: '.';
fragment DIGIT: [0-9];
HEX: DIGIT | 'A'| 'B'| 'C' | 'D' | 'E'| 'F' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f';
LETTER: [a-zA-Z];
TEXT: [\u0021-\u007E];
TOKEN: [\u0021\u0023-\u0027\u002D-\u002E\u0030-\u0039\u0041-\u005A\u005E-\u007A\u007C\u007E];
INT: DIGIT+;
ID: LETTER (LETTER|DIGIT)*?;
ALPHA: [a-zA-Z0-9];
SP: ' ';
CR: '\r';
LF: '\n';
CRLF: CR | LF | (CR LF);
CTL : [\u0000-\u001F];

transport: 'Transport' HCOLON transport_spect;
transport_spect: rtp_transport_id transport_parameter*;

rtp_transport_id: 'RTP/'profile lower_transport?;
profile: 'AVP' | 'SAVP' | 'AVPF' | 'SAVPF';
lower_transport: 'TCP' | 'UDP';
LHEX: 'a' | 'b' | 'c' | 'd' | 'e' | 'f'; //lowercase "a-f" Hex

transport_parameter:
    (SEMI ( 'unicast' | 'multicast') )
    | (SEMI 'interleaved' EQUAL channel ('-' channel)?)
    | (SEMI 'ttl' EQUAL ttl)
    | (SEMI 'layers' EQUAL INT)
    | (SEMI 'ssrc' EQUAL ssrc *(SLASH ssrc))
    | (SEMI 'mode' EQUAL mode_spec)
    | (SEMI 'dest_addr' EQUAL addr_list)
    | (SEMI 'src_addr' EQUAL addr_list)
    | (SEMI 'setup' EQUAL contrans_setup)
    | (SEMI 'connection' EQUAL contrans_con)
    | (SEMI 'RTCP-mux')
    | (SEMI 'MIKEY' EQUAL TEXT*)
    | (SEMI header);

ttl: INT;
channel: ttl;
ssrc: HEX+;
mode_spec: ( DQUOTE transport_mode (COMMA transport_mode)* DQUOTE );
transport_mode: 'PLAY';
addr_list: quoted_addr (SLASH quoted_addr)*;
quoted_addr: DQUOTE (uriHost (':' INT)?) | (':' INT);

header: headerName '=' headerValue;
headerName: (TOKEN | LETTER | INT)+ SP*;
headerValue: (SP | SLASH | DOT | LETTER | INT | TEXT | TOKEN| HCOLON)*;

uriName: (TEXT | INT | TOKEN | LETTER)* DOT (LETTER | INT)+;
uriHost: IP|uriName;

contrans_setup: 'active' | 'passive' | 'actpass';
contrans_con: 'new' | 'existing';
