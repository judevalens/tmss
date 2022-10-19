// Code generated from rtsp_request.g4 by ANTLR 4.10.1. DO NOT EDIT.

package request // rtsp_request
import "github.com/antlr/antlr4/runtime/Go/antlr"

// rtsp_requestListener is a complete listener for a parse tree produced by rtsp_requestParser.
type rtsp_requestListener interface {
	antlr.ParseTreeListener

	// EnterRequest is called when entering the Request production.
	EnterRequest(c *RequestContext)

	// EnterRequestLine is called when entering the requestLine production.
	EnterRequestLine(c *RequestLineContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterResponse is called when entering the response production.
	EnterResponse(c *ResponseContext)

	// EnterStatusLine is called when entering the statusLine production.
	EnterStatusLine(c *StatusLineContext)

	// EnterStatus_reason is called when entering the status_reason production.
	EnterStatus_reason(c *Status_reasonContext)

	// EnterStatus_code is called when entering the status_code production.
	EnterStatus_code(c *Status_codeContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterHeaderName is called when entering the headerName production.
	EnterHeaderName(c *HeaderNameContext)

	// EnterHeaderValue is called when entering the headerValue production.
	EnterHeaderValue(c *HeaderValueContext)

	// EnterRtspUri is called when entering the rtspUri production.
	EnterRtspUri(c *RtspUriContext)

	// EnterUriPath is called when entering the uriPath production.
	EnterUriPath(c *UriPathContext)

	// EnterUriName is called when entering the uriName production.
	EnterUriName(c *UriNameContext)

	// EnterUriHost is called when entering the uriHost production.
	EnterUriHost(c *UriHostContext)

	// EnterEor is called when entering the eor production.
	EnterEor(c *EorContext)

	// ExitRequest is called when exiting the Request production.
	ExitRequest(c *RequestContext)

	// ExitRequestLine is called when exiting the requestLine production.
	ExitRequestLine(c *RequestLineContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitResponse is called when exiting the response production.
	ExitResponse(c *ResponseContext)

	// ExitStatusLine is called when exiting the statusLine production.
	ExitStatusLine(c *StatusLineContext)

	// ExitStatus_reason is called when exiting the status_reason production.
	ExitStatus_reason(c *Status_reasonContext)

	// ExitStatus_code is called when exiting the status_code production.
	ExitStatus_code(c *Status_codeContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitHeaderName is called when exiting the headerName production.
	ExitHeaderName(c *HeaderNameContext)

	// ExitHeaderValue is called when exiting the headerValue production.
	ExitHeaderValue(c *HeaderValueContext)

	// ExitRtspUri is called when exiting the rtspUri production.
	ExitRtspUri(c *RtspUriContext)

	// ExitUriPath is called when exiting the uriPath production.
	ExitUriPath(c *UriPathContext)

	// ExitUriName is called when exiting the uriName production.
	ExitUriName(c *UriNameContext)

	// ExitUriHost is called when exiting the uriHost production.
	ExitUriHost(c *UriHostContext)

	// ExitEor is called when exiting the eor production.
	ExitEor(c *EorContext)
}
