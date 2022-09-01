package request

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"net/url"
)

func ParseRtspReq(req string) (listener *Parser, parseErr error) {
	input := antlr.NewInputStream(req)
	lexer := Newrtsp_requestLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := Newrtsp_requestParser(stream)
	listener = &Parser{
		InputStream:              input,
		Basertsp_requestListener: Basertsp_requestListener{},
	}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	defer func() {
		if r := recover(); r != nil {
			err := errors.New(fmt.Sprintf("Could not parse Request\nerr: %v", r))
			fmt.Printf("%v\n", err)
			listener = &Parser{}
			parseErr = err
		}
	}()
	tree := parser.Request()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener, parseErr
}

type Response struct {
	Version string
	status  string
	Reason  string
	Headers map[string]string
	Body    []byte
}
type Request struct {
	Method  string
	Url     *url.URL
	Version string
	Headers map[string]string
	Body    []byte
}
type Parser struct {
	Basertsp_requestListener
	*antlr.InputStream
	errors   []error
	Request  Request
	Response Response
}

func (r *Parser) EnterRequest(c *RequestContext) {
	r.errors = make([]error, 0)
	r.Request = Request{
		Headers: map[string]string{},
	}
}

func (r *Parser) EnterRequestLine(c *RequestLineContext) {
	fmt.Println(c.GetText())
}
func (r *Parser) EnterHeader(c *HeaderContext) {
	r.Request.Headers[c.HeaderName().GetText()] = c.HeaderValue().GetText()
}

func (r *Parser) EnterRtspUri(c *RtspUriContext) {
	newUrl, err := url.Parse(c.GetText())
	if err != nil {
		c.GetParser().SetParserRuleContext(c.BaseParserRuleContext)
		c.GetParser().NotifyErrorListeners(err.Error(), c.parser.GetCurrentToken(), antlr.NewBaseRecognitionException(err.Error(), c.parser, r.InputStream, c))
	}
	r.Request.Url = newUrl
}

func (r *Parser) EnterUriPath(c *UriPathContext) {

}

func (r *Parser) EnterUriName(c *UriNameContext) {

}

func (r *Parser) EnterUriHost(c *UriHostContext) {

}

func (r *Parser) EnterEor(c *EorContext) {

}

func (r *Parser) ExitRequest(c *RequestContext) {
}

func (r *Parser) ExitRequestLine(c *RequestLineContext) {
	r.Request.Method = c.Method().GetText()
	r.Request.Version = c.RTSP_VERSION().GetText()
}

func (r *Parser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := errors.New(msg)
	r.errors = append(r.errors, err)
	fmt.Printf("new error from lister....\n")
	panic(err)
}

func (r *Parser) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (r *Parser) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

func (r *Parser) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {

}
