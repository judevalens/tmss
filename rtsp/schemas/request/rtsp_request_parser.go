// Code generated from rtsp_request.g4 by ANTLR 4.10.1. DO NOT EDIT.

package request // rtsp_request
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type rtsp_requestParser struct {
	*antlr.BaseParser
}

var rtsp_requestParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rtsp_requestParserInit() {
	staticData := &rtsp_requestParserStaticData
	staticData.literalNames = []string{
		"", "", "", "'://'", "", "':'", "'/'", "'.'", "", "", "", "", "", "",
		"' '", "'\\r'", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "RTSP_VERSION", "URI_SCHEME", "URI_DELIMETER", "IP", "HCOLON", "SLASH",
		"DOT", "LETTER", "TOKEN", "TEXT", "INT", "ID", "ALPHA", "SP", "CR",
		"LF", "CRLF", "CTL",
	}
	staticData.ruleNames = []string{
		"Request", "requestLine", "method", "response", "statusLine", "status_reason",
		"status_code", "header", "headerName", "headerValue", "rtspUri", "uriPath",
		"uriName", "uriHost", "eor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 140, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 50, 8, 2, 11, 2, 12, 2, 51,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 59, 8, 3, 10, 3, 12, 3, 62, 9, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 4, 5, 74, 8, 5,
		11, 5, 12, 5, 75, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 4, 8, 85, 8,
		8, 11, 8, 12, 8, 86, 1, 8, 5, 8, 90, 8, 8, 10, 8, 12, 8, 93, 9, 8, 1, 9,
		5, 9, 96, 8, 9, 10, 9, 12, 9, 99, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 5, 11, 108, 8, 11, 10, 11, 12, 11, 111, 9, 11, 1, 11,
		5, 11, 114, 8, 11, 10, 11, 12, 11, 117, 9, 11, 1, 12, 5, 12, 120, 8, 12,
		10, 12, 12, 12, 123, 9, 12, 1, 12, 1, 12, 4, 12, 127, 8, 12, 11, 12, 12,
		12, 128, 1, 13, 1, 13, 3, 13, 133, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14, 138,
		8, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 0, 6, 1, 0, 5, 11, 2, 0, 8, 9, 11, 11, 2, 0, 5, 11, 14, 14, 1,
		0, 7, 11, 1, 0, 8, 11, 2, 0, 8, 8, 11, 11, 139, 0, 30, 1, 0, 0, 0, 2, 42,
		1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10,
		73, 1, 0, 0, 0, 12, 77, 1, 0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 84, 1, 0, 0,
		0, 18, 97, 1, 0, 0, 0, 20, 100, 1, 0, 0, 0, 22, 105, 1, 0, 0, 0, 24, 121,
		1, 0, 0, 0, 26, 132, 1, 0, 0, 0, 28, 137, 1, 0, 0, 0, 30, 31, 3, 2, 1,
		0, 31, 37, 3, 28, 14, 0, 32, 33, 3, 14, 7, 0, 33, 34, 3, 28, 14, 0, 34,
		36, 1, 0, 0, 0, 35, 32, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0,
		0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41,
		3, 28, 14, 0, 41, 1, 1, 0, 0, 0, 42, 43, 3, 4, 2, 0, 43, 44, 5, 14, 0,
		0, 44, 45, 3, 20, 10, 0, 45, 46, 5, 14, 0, 0, 46, 47, 5, 1, 0, 0, 47, 3,
		1, 0, 0, 0, 48, 50, 5, 8, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 3, 8,
		4, 0, 54, 60, 3, 28, 14, 0, 55, 56, 3, 14, 7, 0, 56, 57, 3, 28, 14, 0,
		57, 59, 1, 0, 0, 0, 58, 55, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1,
		0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63,
		64, 3, 28, 14, 0, 64, 7, 1, 0, 0, 0, 65, 66, 5, 1, 0, 0, 66, 67, 5, 14,
		0, 0, 67, 68, 3, 12, 6, 0, 68, 69, 5, 14, 0, 0, 69, 70, 3, 10, 5, 0, 70,
		9, 1, 0, 0, 0, 71, 74, 7, 0, 0, 0, 72, 74, 5, 14, 0, 0, 73, 71, 1, 0, 0,
		0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 11, 1, 0, 0, 0, 77, 78, 5, 11, 0, 0, 78, 13, 1, 0, 0, 0,
		79, 80, 3, 16, 8, 0, 80, 81, 5, 5, 0, 0, 81, 82, 3, 18, 9, 0, 82, 15, 1,
		0, 0, 0, 83, 85, 7, 1, 0, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 91, 1, 0, 0, 0, 88, 90, 5, 14,
		0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 17, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 96, 7, 2, 0, 0,
		95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 19, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 101, 5, 2, 0, 0,
		101, 102, 5, 3, 0, 0, 102, 103, 3, 26, 13, 0, 103, 104, 3, 22, 11, 0, 104,
		21, 1, 0, 0, 0, 105, 109, 5, 6, 0, 0, 106, 108, 7, 3, 0, 0, 107, 106, 1,
		0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0,
		0, 110, 115, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 114, 3, 22, 11, 0,
		113, 112, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 23, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 120, 7,
		4, 0, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0,
		0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		126, 5, 7, 0, 0, 125, 127, 7, 5, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 25, 1, 0,
		0, 0, 130, 133, 5, 4, 0, 0, 131, 133, 3, 24, 12, 0, 132, 130, 1, 0, 0,
		0, 132, 131, 1, 0, 0, 0, 133, 27, 1, 0, 0, 0, 134, 138, 5, 16, 0, 0, 135,
		138, 5, 15, 0, 0, 136, 138, 5, 17, 0, 0, 137, 134, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 29, 1, 0, 0, 0, 14, 37, 51, 60,
		73, 75, 86, 91, 97, 109, 115, 121, 128, 132, 137,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// rtsp_requestParserInit initializes any static state used to implement rtsp_requestParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newrtsp_requestParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rtsp_requestParserInit() {
	staticData := &rtsp_requestParserStaticData
	staticData.once.Do(rtsp_requestParserInit)
}

// Newrtsp_requestParser produces a new parser instance for the optional input antlr.TokenStream.
func Newrtsp_requestParser(input antlr.TokenStream) *rtsp_requestParser {
	Rtsp_requestParserInit()
	this := new(rtsp_requestParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rtsp_requestParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "rtsp_request.g4"

	return this
}

// rtsp_requestParser tokens.
const (
	rtsp_requestParserEOF           = antlr.TokenEOF
	rtsp_requestParserRTSP_VERSION  = 1
	rtsp_requestParserURI_SCHEME    = 2
	rtsp_requestParserURI_DELIMETER = 3
	rtsp_requestParserIP            = 4
	rtsp_requestParserHCOLON        = 5
	rtsp_requestParserSLASH         = 6
	rtsp_requestParserDOT           = 7
	rtsp_requestParserLETTER        = 8
	rtsp_requestParserTOKEN         = 9
	rtsp_requestParserTEXT          = 10
	rtsp_requestParserINT           = 11
	rtsp_requestParserID            = 12
	rtsp_requestParserALPHA         = 13
	rtsp_requestParserSP            = 14
	rtsp_requestParserCR            = 15
	rtsp_requestParserLF            = 16
	rtsp_requestParserCRLF          = 17
	rtsp_requestParserCTL           = 18
)

// rtsp_requestParser rules.
const (
	rtsp_requestParserRULE_request       = 0
	rtsp_requestParserRULE_requestLine   = 1
	rtsp_requestParserRULE_method        = 2
	rtsp_requestParserRULE_response      = 3
	rtsp_requestParserRULE_statusLine    = 4
	rtsp_requestParserRULE_status_reason = 5
	rtsp_requestParserRULE_status_code   = 6
	rtsp_requestParserRULE_header        = 7
	rtsp_requestParserRULE_headerName    = 8
	rtsp_requestParserRULE_headerValue   = 9
	rtsp_requestParserRULE_rtspUri       = 10
	rtsp_requestParserRULE_uriPath       = 11
	rtsp_requestParserRULE_uriName       = 12
	rtsp_requestParserRULE_uriHost       = 13
	rtsp_requestParserRULE_eor           = 14
)

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) RequestLine() IRequestLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestLineContext)
}

func (s *RequestContext) AllEor() []IEorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEorContext); ok {
			len++
		}
	}

	tst := make([]IEorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEorContext); ok {
			tst[i] = t.(IEorContext)
			i++
		}
	}

	return tst
}

func (s *RequestContext) Eor(i int) IEorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEorContext)
}

func (s *RequestContext) AllHeader() []IHeaderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderContext); ok {
			len++
		}
	}

	tst := make([]IHeaderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderContext); ok {
			tst[i] = t.(IHeaderContext)
			i++
		}
	}

	return tst
}

func (s *RequestContext) Header(i int) IHeaderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *rtsp_requestParser) Request() (localctx IRequestContext) {
	this := p
	_ = this

	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, rtsp_requestParserRULE_request)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.RequestLine()
	}
	{
		p.SetState(31)
		p.Eor()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserINT))) != 0 {
		{
			p.SetState(32)
			p.Header()
		}
		{
			p.SetState(33)
			p.Eor()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Eor()
	}

	return localctx
}

// IRequestLineContext is an interface to support dynamic dispatch.
type IRequestLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestLineContext differentiates from other interfaces.
	IsRequestLineContext()
}

type RequestLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestLineContext() *RequestLineContext {
	var p = new(RequestLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_requestLine
	return p
}

func (*RequestLineContext) IsRequestLineContext() {}

func NewRequestLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestLineContext {
	var p = new(RequestLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_requestLine

	return p
}

func (s *RequestLineContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestLineContext) Method() IMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *RequestLineContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSP)
}

func (s *RequestLineContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSP, i)
}

func (s *RequestLineContext) RtspUri() IRtspUriContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtspUriContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtspUriContext)
}

func (s *RequestLineContext) RTSP_VERSION() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserRTSP_VERSION, 0)
}

func (s *RequestLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterRequestLine(s)
	}
}

func (s *RequestLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitRequestLine(s)
	}
}

func (p *rtsp_requestParser) RequestLine() (localctx IRequestLineContext) {
	this := p
	_ = this

	localctx = NewRequestLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, rtsp_requestParserRULE_requestLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Method()
	}
	{
		p.SetState(43)
		p.Match(rtsp_requestParserSP)
	}
	{
		p.SetState(44)
		p.RtspUri()
	}
	{
		p.SetState(45)
		p.Match(rtsp_requestParserSP)
	}
	{
		p.SetState(46)
		p.Match(rtsp_requestParserRTSP_VERSION)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *MethodContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *rtsp_requestParser) Method() (localctx IMethodContext) {
	this := p
	_ = this

	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, rtsp_requestParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rtsp_requestParserLETTER {
		{
			p.SetState(48)
			p.Match(rtsp_requestParserLETTER)
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IResponseContext is an interface to support dynamic dispatch.
type IResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResponseContext differentiates from other interfaces.
	IsResponseContext()
}

type ResponseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponseContext() *ResponseContext {
	var p = new(ResponseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_response
	return p
}

func (*ResponseContext) IsResponseContext() {}

func NewResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponseContext {
	var p = new(ResponseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_response

	return p
}

func (s *ResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponseContext) StatusLine() IStatusLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatusLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatusLineContext)
}

func (s *ResponseContext) AllEor() []IEorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEorContext); ok {
			len++
		}
	}

	tst := make([]IEorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEorContext); ok {
			tst[i] = t.(IEorContext)
			i++
		}
	}

	return tst
}

func (s *ResponseContext) Eor(i int) IEorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEorContext)
}

func (s *ResponseContext) AllHeader() []IHeaderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderContext); ok {
			len++
		}
	}

	tst := make([]IHeaderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderContext); ok {
			tst[i] = t.(IHeaderContext)
			i++
		}
	}

	return tst
}

func (s *ResponseContext) Header(i int) IHeaderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *ResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterResponse(s)
	}
}

func (s *ResponseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitResponse(s)
	}
}

func (p *rtsp_requestParser) Response() (localctx IResponseContext) {
	this := p
	_ = this

	localctx = NewResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, rtsp_requestParserRULE_response)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.StatusLine()
	}
	{
		p.SetState(54)
		p.Eor()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserINT))) != 0 {
		{
			p.SetState(55)
			p.Header()
		}
		{
			p.SetState(56)
			p.Eor()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Eor()
	}

	return localctx
}

// IStatusLineContext is an interface to support dynamic dispatch.
type IStatusLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatusLineContext differentiates from other interfaces.
	IsStatusLineContext()
}

type StatusLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatusLineContext() *StatusLineContext {
	var p = new(StatusLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_statusLine
	return p
}

func (*StatusLineContext) IsStatusLineContext() {}

func NewStatusLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatusLineContext {
	var p = new(StatusLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_statusLine

	return p
}

func (s *StatusLineContext) GetParser() antlr.Parser { return s.parser }

func (s *StatusLineContext) RTSP_VERSION() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserRTSP_VERSION, 0)
}

func (s *StatusLineContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSP)
}

func (s *StatusLineContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSP, i)
}

func (s *StatusLineContext) Status_code() IStatus_codeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatus_codeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatus_codeContext)
}

func (s *StatusLineContext) Status_reason() IStatus_reasonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatus_reasonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatus_reasonContext)
}

func (s *StatusLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatusLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatusLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterStatusLine(s)
	}
}

func (s *StatusLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitStatusLine(s)
	}
}

func (p *rtsp_requestParser) StatusLine() (localctx IStatusLineContext) {
	this := p
	_ = this

	localctx = NewStatusLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, rtsp_requestParserRULE_statusLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(rtsp_requestParserRTSP_VERSION)
	}
	{
		p.SetState(66)
		p.Match(rtsp_requestParserSP)
	}
	{
		p.SetState(67)
		p.Status_code()
	}
	{
		p.SetState(68)
		p.Match(rtsp_requestParserSP)
	}
	{
		p.SetState(69)
		p.Status_reason()
	}

	return localctx
}

// IStatus_reasonContext is an interface to support dynamic dispatch.
type IStatus_reasonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatus_reasonContext differentiates from other interfaces.
	IsStatus_reasonContext()
}

type Status_reasonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatus_reasonContext() *Status_reasonContext {
	var p = new(Status_reasonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_status_reason
	return p
}

func (*Status_reasonContext) IsStatus_reasonContext() {}

func NewStatus_reasonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Status_reasonContext {
	var p = new(Status_reasonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_status_reason

	return p
}

func (s *Status_reasonContext) GetParser() antlr.Parser { return s.parser }

func (s *Status_reasonContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSP)
}

func (s *Status_reasonContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSP, i)
}

func (s *Status_reasonContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSLASH)
}

func (s *Status_reasonContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSLASH, i)
}

func (s *Status_reasonContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserDOT)
}

func (s *Status_reasonContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserDOT, i)
}

func (s *Status_reasonContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *Status_reasonContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *Status_reasonContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserINT)
}

func (s *Status_reasonContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, i)
}

func (s *Status_reasonContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTEXT)
}

func (s *Status_reasonContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTEXT, i)
}

func (s *Status_reasonContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTOKEN)
}

func (s *Status_reasonContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTOKEN, i)
}

func (s *Status_reasonContext) AllHCOLON() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserHCOLON)
}

func (s *Status_reasonContext) HCOLON(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserHCOLON, i)
}

func (s *Status_reasonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Status_reasonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Status_reasonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterStatus_reason(s)
	}
}

func (s *Status_reasonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitStatus_reason(s)
	}
}

func (p *rtsp_requestParser) Status_reason() (localctx IStatus_reasonContext) {
	this := p
	_ = this

	localctx = NewStatus_reasonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, rtsp_requestParserRULE_status_reason)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserHCOLON)|(1<<rtsp_requestParserSLASH)|(1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT)|(1<<rtsp_requestParserSP))) != 0) {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case rtsp_requestParserHCOLON, rtsp_requestParserSLASH, rtsp_requestParserDOT, rtsp_requestParserLETTER, rtsp_requestParserTOKEN, rtsp_requestParserTEXT, rtsp_requestParserINT:
			{
				p.SetState(71)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserHCOLON)|(1<<rtsp_requestParserSLASH)|(1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		case rtsp_requestParserSP:
			{
				p.SetState(72)
				p.Match(rtsp_requestParserSP)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatus_codeContext is an interface to support dynamic dispatch.
type IStatus_codeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatus_codeContext differentiates from other interfaces.
	IsStatus_codeContext()
}

type Status_codeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatus_codeContext() *Status_codeContext {
	var p = new(Status_codeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_status_code
	return p
}

func (*Status_codeContext) IsStatus_codeContext() {}

func NewStatus_codeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Status_codeContext {
	var p = new(Status_codeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_status_code

	return p
}

func (s *Status_codeContext) GetParser() antlr.Parser { return s.parser }

func (s *Status_codeContext) INT() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, 0)
}

func (s *Status_codeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Status_codeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Status_codeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterStatus_code(s)
	}
}

func (s *Status_codeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitStatus_code(s)
	}
}

func (p *rtsp_requestParser) Status_code() (localctx IStatus_codeContext) {
	this := p
	_ = this

	localctx = NewStatus_codeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, rtsp_requestParserRULE_status_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(rtsp_requestParserINT)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) HeaderName() IHeaderNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderNameContext)
}

func (s *HeaderContext) HCOLON() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserHCOLON, 0)
}

func (s *HeaderContext) HeaderValue() IHeaderValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderValueContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *rtsp_requestParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, rtsp_requestParserRULE_header)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.HeaderName()
	}
	{
		p.SetState(80)
		p.Match(rtsp_requestParserHCOLON)
	}
	{
		p.SetState(81)
		p.HeaderValue()
	}

	return localctx
}

// IHeaderNameContext is an interface to support dynamic dispatch.
type IHeaderNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderNameContext differentiates from other interfaces.
	IsHeaderNameContext()
}

type HeaderNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderNameContext() *HeaderNameContext {
	var p = new(HeaderNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_headerName
	return p
}

func (*HeaderNameContext) IsHeaderNameContext() {}

func NewHeaderNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderNameContext {
	var p = new(HeaderNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_headerName

	return p
}

func (s *HeaderNameContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderNameContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSP)
}

func (s *HeaderNameContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSP, i)
}

func (s *HeaderNameContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTOKEN)
}

func (s *HeaderNameContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTOKEN, i)
}

func (s *HeaderNameContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *HeaderNameContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *HeaderNameContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserINT)
}

func (s *HeaderNameContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, i)
}

func (s *HeaderNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterHeaderName(s)
	}
}

func (s *HeaderNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitHeaderName(s)
	}
}

func (p *rtsp_requestParser) HeaderName() (localctx IHeaderNameContext) {
	this := p
	_ = this

	localctx = NewHeaderNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, rtsp_requestParserRULE_headerName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserINT))) != 0) {
		{
			p.SetState(83)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserINT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == rtsp_requestParserSP {
		{
			p.SetState(88)
			p.Match(rtsp_requestParserSP)
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderValueContext is an interface to support dynamic dispatch.
type IHeaderValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderValueContext differentiates from other interfaces.
	IsHeaderValueContext()
}

type HeaderValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderValueContext() *HeaderValueContext {
	var p = new(HeaderValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_headerValue
	return p
}

func (*HeaderValueContext) IsHeaderValueContext() {}

func NewHeaderValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderValueContext {
	var p = new(HeaderValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_headerValue

	return p
}

func (s *HeaderValueContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderValueContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSP)
}

func (s *HeaderValueContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSP, i)
}

func (s *HeaderValueContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserSLASH)
}

func (s *HeaderValueContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSLASH, i)
}

func (s *HeaderValueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserDOT)
}

func (s *HeaderValueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserDOT, i)
}

func (s *HeaderValueContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *HeaderValueContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *HeaderValueContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserINT)
}

func (s *HeaderValueContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, i)
}

func (s *HeaderValueContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTEXT)
}

func (s *HeaderValueContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTEXT, i)
}

func (s *HeaderValueContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTOKEN)
}

func (s *HeaderValueContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTOKEN, i)
}

func (s *HeaderValueContext) AllHCOLON() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserHCOLON)
}

func (s *HeaderValueContext) HCOLON(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserHCOLON, i)
}

func (s *HeaderValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterHeaderValue(s)
	}
}

func (s *HeaderValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitHeaderValue(s)
	}
}

func (p *rtsp_requestParser) HeaderValue() (localctx IHeaderValueContext) {
	this := p
	_ = this

	localctx = NewHeaderValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, rtsp_requestParserRULE_headerValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserHCOLON)|(1<<rtsp_requestParserSLASH)|(1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT)|(1<<rtsp_requestParserSP))) != 0 {
		{
			p.SetState(94)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserHCOLON)|(1<<rtsp_requestParserSLASH)|(1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT)|(1<<rtsp_requestParserSP))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRtspUriContext is an interface to support dynamic dispatch.
type IRtspUriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRtspUriContext differentiates from other interfaces.
	IsRtspUriContext()
}

type RtspUriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRtspUriContext() *RtspUriContext {
	var p = new(RtspUriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_rtspUri
	return p
}

func (*RtspUriContext) IsRtspUriContext() {}

func NewRtspUriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RtspUriContext {
	var p = new(RtspUriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_rtspUri

	return p
}

func (s *RtspUriContext) GetParser() antlr.Parser { return s.parser }

func (s *RtspUriContext) URI_SCHEME() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserURI_SCHEME, 0)
}

func (s *RtspUriContext) URI_DELIMETER() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserURI_DELIMETER, 0)
}

func (s *RtspUriContext) UriHost() IUriHostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUriHostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUriHostContext)
}

func (s *RtspUriContext) UriPath() IUriPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUriPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUriPathContext)
}

func (s *RtspUriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RtspUriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RtspUriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterRtspUri(s)
	}
}

func (s *RtspUriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitRtspUri(s)
	}
}

func (p *rtsp_requestParser) RtspUri() (localctx IRtspUriContext) {
	this := p
	_ = this

	localctx = NewRtspUriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, rtsp_requestParserRULE_rtspUri)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(rtsp_requestParserURI_SCHEME)
	}
	{
		p.SetState(101)
		p.Match(rtsp_requestParserURI_DELIMETER)
	}
	{
		p.SetState(102)
		p.UriHost()
	}
	{
		p.SetState(103)
		p.UriPath()
	}

	return localctx
}

// IUriPathContext is an interface to support dynamic dispatch.
type IUriPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUriPathContext differentiates from other interfaces.
	IsUriPathContext()
}

type UriPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUriPathContext() *UriPathContext {
	var p = new(UriPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_uriPath
	return p
}

func (*UriPathContext) IsUriPathContext() {}

func NewUriPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UriPathContext {
	var p = new(UriPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_uriPath

	return p
}

func (s *UriPathContext) GetParser() antlr.Parser { return s.parser }

func (s *UriPathContext) SLASH() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserSLASH, 0)
}

func (s *UriPathContext) AllUriPath() []IUriPathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUriPathContext); ok {
			len++
		}
	}

	tst := make([]IUriPathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUriPathContext); ok {
			tst[i] = t.(IUriPathContext)
			i++
		}
	}

	return tst
}

func (s *UriPathContext) UriPath(i int) IUriPathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUriPathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUriPathContext)
}

func (s *UriPathContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *UriPathContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *UriPathContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserINT)
}

func (s *UriPathContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, i)
}

func (s *UriPathContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTEXT)
}

func (s *UriPathContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTEXT, i)
}

func (s *UriPathContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTOKEN)
}

func (s *UriPathContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTOKEN, i)
}

func (s *UriPathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserDOT)
}

func (s *UriPathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserDOT, i)
}

func (s *UriPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UriPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterUriPath(s)
	}
}

func (s *UriPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitUriPath(s)
	}
}

func (p *rtsp_requestParser) UriPath() (localctx IUriPathContext) {
	this := p
	_ = this

	localctx = NewUriPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, rtsp_requestParserRULE_uriPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(rtsp_requestParserSLASH)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT))) != 0 {
		{
			p.SetState(106)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserDOT)|(1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(112)
				p.UriPath()
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IUriNameContext is an interface to support dynamic dispatch.
type IUriNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUriNameContext differentiates from other interfaces.
	IsUriNameContext()
}

type UriNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUriNameContext() *UriNameContext {
	var p = new(UriNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_uriName
	return p
}

func (*UriNameContext) IsUriNameContext() {}

func NewUriNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UriNameContext {
	var p = new(UriNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_uriName

	return p
}

func (s *UriNameContext) GetParser() antlr.Parser { return s.parser }

func (s *UriNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserDOT, 0)
}

func (s *UriNameContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTEXT)
}

func (s *UriNameContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTEXT, i)
}

func (s *UriNameContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserINT)
}

func (s *UriNameContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserINT, i)
}

func (s *UriNameContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserTOKEN)
}

func (s *UriNameContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserTOKEN, i)
}

func (s *UriNameContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(rtsp_requestParserLETTER)
}

func (s *UriNameContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLETTER, i)
}

func (s *UriNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UriNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterUriName(s)
	}
}

func (s *UriNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitUriName(s)
	}
}

func (p *rtsp_requestParser) UriName() (localctx IUriNameContext) {
	this := p
	_ = this

	localctx = NewUriNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, rtsp_requestParserRULE_uriName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT))) != 0 {
		{
			p.SetState(118)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rtsp_requestParserLETTER)|(1<<rtsp_requestParserTOKEN)|(1<<rtsp_requestParserTEXT)|(1<<rtsp_requestParserINT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(rtsp_requestParserDOT)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rtsp_requestParserLETTER || _la == rtsp_requestParserINT {
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !(_la == rtsp_requestParserLETTER || _la == rtsp_requestParserINT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUriHostContext is an interface to support dynamic dispatch.
type IUriHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUriHostContext differentiates from other interfaces.
	IsUriHostContext()
}

type UriHostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUriHostContext() *UriHostContext {
	var p = new(UriHostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_uriHost
	return p
}

func (*UriHostContext) IsUriHostContext() {}

func NewUriHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UriHostContext {
	var p = new(UriHostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_uriHost

	return p
}

func (s *UriHostContext) GetParser() antlr.Parser { return s.parser }

func (s *UriHostContext) IP() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserIP, 0)
}

func (s *UriHostContext) UriName() IUriNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUriNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUriNameContext)
}

func (s *UriHostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriHostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UriHostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterUriHost(s)
	}
}

func (s *UriHostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitUriHost(s)
	}
}

func (p *rtsp_requestParser) UriHost() (localctx IUriHostContext) {
	this := p
	_ = this

	localctx = NewUriHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, rtsp_requestParserRULE_uriHost)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rtsp_requestParserIP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(rtsp_requestParserIP)
		}

	case rtsp_requestParserDOT, rtsp_requestParserLETTER, rtsp_requestParserTOKEN, rtsp_requestParserTEXT, rtsp_requestParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.UriName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEorContext is an interface to support dynamic dispatch.
type IEorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEorContext differentiates from other interfaces.
	IsEorContext()
}

type EorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEorContext() *EorContext {
	var p = new(EorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rtsp_requestParserRULE_eor
	return p
}

func (*EorContext) IsEorContext() {}

func NewEorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EorContext {
	var p = new(EorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rtsp_requestParserRULE_eor

	return p
}

func (s *EorContext) GetParser() antlr.Parser { return s.parser }

func (s *EorContext) LF() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserLF, 0)
}

func (s *EorContext) CR() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserCR, 0)
}

func (s *EorContext) CRLF() antlr.TerminalNode {
	return s.GetToken(rtsp_requestParserCRLF, 0)
}

func (s *EorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.EnterEor(s)
	}
}

func (s *EorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rtsp_requestListener); ok {
		listenerT.ExitEor(s)
	}
}

func (p *rtsp_requestParser) Eor() (localctx IEorContext) {
	this := p
	_ = this

	localctx = NewEorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, rtsp_requestParserRULE_eor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rtsp_requestParserLF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(rtsp_requestParserLF)
		}

	case rtsp_requestParserCR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(rtsp_requestParserCR)
		}

	case rtsp_requestParserCRLF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.Match(rtsp_requestParserCRLF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
