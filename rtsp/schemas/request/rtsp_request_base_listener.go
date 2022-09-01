// Code generated from rtsp_request.g4 by ANTLR 4.10.1. DO NOT EDIT.

package request // rtsp_request
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basertsp_requestListener is a complete listener for a parse tree produced by rtsp_requestParser.
type Basertsp_requestListener struct{}

var _ rtsp_requestListener = &Basertsp_requestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basertsp_requestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basertsp_requestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basertsp_requestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basertsp_requestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRequest is called when production Request is entered.
func (s *Basertsp_requestListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production Request is exited.
func (s *Basertsp_requestListener) ExitRequest(ctx *RequestContext) {}

// EnterRequestLine is called when production requestLine is entered.
func (s *Basertsp_requestListener) EnterRequestLine(ctx *RequestLineContext) {}

// ExitRequestLine is called when production requestLine is exited.
func (s *Basertsp_requestListener) ExitRequestLine(ctx *RequestLineContext) {}

// EnterMethod is called when production method is entered.
func (s *Basertsp_requestListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *Basertsp_requestListener) ExitMethod(ctx *MethodContext) {}

// EnterResponse is called when production response is entered.
func (s *Basertsp_requestListener) EnterResponse(ctx *ResponseContext) {}

// ExitResponse is called when production response is exited.
func (s *Basertsp_requestListener) ExitResponse(ctx *ResponseContext) {}

// EnterStatusLine is called when production statusLine is entered.
func (s *Basertsp_requestListener) EnterStatusLine(ctx *StatusLineContext) {}

// ExitStatusLine is called when production statusLine is exited.
func (s *Basertsp_requestListener) ExitStatusLine(ctx *StatusLineContext) {}

// EnterStatus_reason is called when production status_reason is entered.
func (s *Basertsp_requestListener) EnterStatus_reason(ctx *Status_reasonContext) {}

// ExitStatus_reason is called when production status_reason is exited.
func (s *Basertsp_requestListener) ExitStatus_reason(ctx *Status_reasonContext) {}

// EnterStatus_code is called when production status_code is entered.
func (s *Basertsp_requestListener) EnterStatus_code(ctx *Status_codeContext) {}

// ExitStatus_code is called when production status_code is exited.
func (s *Basertsp_requestListener) ExitStatus_code(ctx *Status_codeContext) {}

// EnterHeader is called when production header is entered.
func (s *Basertsp_requestListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *Basertsp_requestListener) ExitHeader(ctx *HeaderContext) {}

// EnterHeaderName is called when production headerName is entered.
func (s *Basertsp_requestListener) EnterHeaderName(ctx *HeaderNameContext) {}

// ExitHeaderName is called when production headerName is exited.
func (s *Basertsp_requestListener) ExitHeaderName(ctx *HeaderNameContext) {}

// EnterHeaderValue is called when production headerValue is entered.
func (s *Basertsp_requestListener) EnterHeaderValue(ctx *HeaderValueContext) {}

// ExitHeaderValue is called when production headerValue is exited.
func (s *Basertsp_requestListener) ExitHeaderValue(ctx *HeaderValueContext) {}

// EnterRtspUri is called when production rtspUri is entered.
func (s *Basertsp_requestListener) EnterRtspUri(ctx *RtspUriContext) {}

// ExitRtspUri is called when production rtspUri is exited.
func (s *Basertsp_requestListener) ExitRtspUri(ctx *RtspUriContext) {}

// EnterUriPath is called when production uriPath is entered.
func (s *Basertsp_requestListener) EnterUriPath(ctx *UriPathContext) {}

// ExitUriPath is called when production uriPath is exited.
func (s *Basertsp_requestListener) ExitUriPath(ctx *UriPathContext) {}

// EnterUriName is called when production uriName is entered.
func (s *Basertsp_requestListener) EnterUriName(ctx *UriNameContext) {}

// ExitUriName is called when production uriName is exited.
func (s *Basertsp_requestListener) ExitUriName(ctx *UriNameContext) {}

// EnterUriHost is called when production uriHost is entered.
func (s *Basertsp_requestListener) EnterUriHost(ctx *UriHostContext) {}

// ExitUriHost is called when production uriHost is exited.
func (s *Basertsp_requestListener) ExitUriHost(ctx *UriHostContext) {}

// EnterEor is called when production eor is entered.
func (s *Basertsp_requestListener) EnterEor(ctx *EorContext) {}

// ExitEor is called when production eor is exited.
func (s *Basertsp_requestListener) ExitEor(ctx *EorContext) {}
