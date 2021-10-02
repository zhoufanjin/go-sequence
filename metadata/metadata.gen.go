// sequence-metadata v0.4.0 f1264bc1871ed85b966466b348374ace021d3723
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/golang
// Do not edit by hand. Update your webrpc schema and re-generate.
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/0xsequence/go-sequence/lib/prototyp"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.4.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "f1264bc1871ed85b966466b348374ace021d3723"
}

//
// Types
//

type SwapType uint32

const (
	SwapType_UNKNOWN SwapType = 0
	SwapType_BUY     SwapType = 1
	SwapType_SELL    SwapType = 2
)

var SwapType_name = map[uint32]string{
	0: "UNKNOWN",
	1: "BUY",
	2: "SELL",
}

var SwapType_value = map[string]uint32{
	"UNKNOWN": 0,
	"BUY":     1,
	"SELL":    2,
}

func (x SwapType) String() string {
	return SwapType_name[uint32(x)]
}

func (x SwapType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(SwapType_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *SwapType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = SwapType(SwapType_value[j])
	return nil
}

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	AppVersion    string `json:"appVersion"`
}

type RuntimeStatus struct {
	HealthOK   bool           `json:"healthOK"`
	StartTime  time.Time      `json:"startTime"`
	Uptime     uint64         `json:"uptime"`
	Ver        string         `json:"ver"`
	Branch     string         `json:"branch"`
	CommitHash string         `json:"commitHash"`
	Checks     *RuntimeChecks `json:"checks"`
}

type RuntimeChecks struct {
}

type ContractInfo struct {
	ChainID    uint64                  `json:"chainId"`
	Address    string                  `json:"address"`
	Name       string                  `json:"name"`
	Type       string                  `json:"type"`
	Symbol     string                  `json:"symbol"`
	Decimals   uint64                  `json:"decimals"`
	LogoURI    string                  `json:"logoURI"`
	Extensions *ContractInfoExtensions `json:"extensions"`
}

type ContractInfoExtensions struct {
	Link          string `json:"link"`
	Description   string `json:"description"`
	OgImage       string `json:"ogImage"`
	OriginChainID uint64 `json:"originChainId"`
	OriginAddress string `json:"originAddress"`
}

type TokenMetadata struct {
	TokenID         string                   `json:"tokenId"`
	ContractAddress prototyp.Hash            `json:"contractAddress"`
	Name            string                   `json:"name"`
	Description     string                   `json:"description"`
	Image           string                   `json:"image"`
	Decimals        uint64                   `json:"decimals"`
	Properties      map[string]interface{}   `json:"properties"`
	ImageData       string                   `json:"image_data,omitempty"`
	ExternalUrl     string                   `json:"external_url,omitempty"`
	BackgroundColor string                   `json:"background_color,omitempty"`
	Attributes      []map[string]interface{} `json:"attributes"`
}

type GetNiftyswapUnitPricesRequest struct {
	SwapType *SwapType         `json:"swapType"`
	Ids      []prototyp.BigInt `json:"ids"`
	Amounts  []prototyp.BigInt `json:"amounts"`
}

type Metadata interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetTokenMetadata(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) ([]*TokenMetadata, error)
	GetTokenMetadataBatch(ctx context.Context, chainID string, contractTokenMap map[string][]string) (map[string][]*TokenMetadata, error)
	GetContractInfo(ctx context.Context, chainID string, contractAddress string) (*ContractInfo, error)
	GetContractInfoBatch(ctx context.Context, chainID string, contractAddresses []string) (map[string]*ContractInfo, error)
	SearchContractInfo(ctx context.Context, contractAddress string) ([]*ContractInfo, error)
	SearchContractInfoBatch(ctx context.Context, contractAddresses []string) (map[string][]*ContractInfo, error)
	GetNiftyswapTokenQuantity(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) (map[string]string, error)
	GetNiftyswapUnitPrices(ctx context.Context, chainID string, contractAddress string, req *GetNiftyswapUnitPricesRequest) (map[string]string, error)
}

var WebRPCServices = map[string][]string{
	"Metadata": {
		"Ping",
		"Version",
		"RuntimeStatus",
		"GetTokenMetadata",
		"GetTokenMetadataBatch",
		"GetContractInfo",
		"GetContractInfoBatch",
		"SearchContractInfo",
		"SearchContractInfoBatch",
		"GetNiftyswapTokenQuantity",
		"GetNiftyswapUnitPrices",
	},
}

//
// Client
//

const MetadataPathPrefix = "/rpc/Metadata/"

type metadataClient struct {
	client HTTPClient
	urls   [11]string
}

func NewMetadataClient(addr string, client HTTPClient) Metadata {
	prefix := urlBase(addr) + MetadataPathPrefix
	urls := [11]string{
		prefix + "Ping",
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "GetTokenMetadata",
		prefix + "GetTokenMetadataBatch",
		prefix + "GetContractInfo",
		prefix + "GetContractInfoBatch",
		prefix + "SearchContractInfo",
		prefix + "SearchContractInfoBatch",
		prefix + "GetNiftyswapTokenQuantity",
		prefix + "GetNiftyswapUnitPrices",
	}
	return &metadataClient{
		client: client,
		urls:   urls,
	}
}

func (c *metadataClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	return out.Ret0, err
}

func (c *metadataClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *metadataClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetTokenMetadata(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) ([]*TokenMetadata, error) {
	in := struct {
		Arg0 string   `json:"chainID"`
		Arg1 string   `json:"contractAddress"`
		Arg2 []string `json:"tokenIDs"`
	}{chainID, contractAddress, tokenIDs}
	out := struct {
		Ret0 []*TokenMetadata `json:"tokenMetadata"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetTokenMetadataBatch(ctx context.Context, chainID string, contractTokenMap map[string][]string) (map[string][]*TokenMetadata, error) {
	in := struct {
		Arg0 string              `json:"chainID"`
		Arg1 map[string][]string `json:"contractTokenMap"`
	}{chainID, contractTokenMap}
	out := struct {
		Ret0 map[string][]*TokenMetadata `json:"contractTokenMetadata"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetContractInfo(ctx context.Context, chainID string, contractAddress string) (*ContractInfo, error) {
	in := struct {
		Arg0 string `json:"chainID"`
		Arg1 string `json:"contractAddress"`
	}{chainID, contractAddress}
	out := struct {
		Ret0 *ContractInfo `json:"contractInfo"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetContractInfoBatch(ctx context.Context, chainID string, contractAddresses []string) (map[string]*ContractInfo, error) {
	in := struct {
		Arg0 string   `json:"chainID"`
		Arg1 []string `json:"contractAddresses"`
	}{chainID, contractAddresses}
	out := struct {
		Ret0 map[string]*ContractInfo `json:"contractInfoMap"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[6], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) SearchContractInfo(ctx context.Context, contractAddress string) ([]*ContractInfo, error) {
	in := struct {
		Arg0 string `json:"contractAddress"`
	}{contractAddress}
	out := struct {
		Ret0 []*ContractInfo `json:"contractInfoList"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[7], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) SearchContractInfoBatch(ctx context.Context, contractAddresses []string) (map[string][]*ContractInfo, error) {
	in := struct {
		Arg0 []string `json:"contractAddresses"`
	}{contractAddresses}
	out := struct {
		Ret0 map[string][]*ContractInfo `json:"contractInfoByChain"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[8], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetNiftyswapTokenQuantity(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) (map[string]string, error) {
	in := struct {
		Arg0 string   `json:"chainID"`
		Arg1 string   `json:"contractAddress"`
		Arg2 []string `json:"tokenIDs"`
	}{chainID, contractAddress, tokenIDs}
	out := struct {
		Ret0 map[string]string `json:"quantity"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[9], in, &out)
	return out.Ret0, err
}

func (c *metadataClient) GetNiftyswapUnitPrices(ctx context.Context, chainID string, contractAddress string, req *GetNiftyswapUnitPricesRequest) (map[string]string, error) {
	in := struct {
		Arg0 string                         `json:"chainID"`
		Arg1 string                         `json:"contractAddress"`
		Arg2 *GetNiftyswapUnitPricesRequest `json:"req"`
	}{chainID, contractAddress, req}
	out := struct {
		Ret0 map[string]string `json:"prices"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[10], in, &out)
	return out.Ret0, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("request failed", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return clientError("failed to read response body", err)
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return clientError("failed to unmarshal json response body", err)
		}
		if err = ctx.Err(); err != nil {
			return clientError("aborted because context was done", err)
		}
	}

	return nil
}

// errorFromResponse builds a webrpc Error from a non-200 HTTP response.
func errorFromResponse(resp *http.Response) Error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}

	var respErr ErrorPayload
	if err := json.Unmarshal(respBody, &respErr); err != nil {
		return clientError("failed unmarshal error response", err)
	}

	errCode := ErrorCode(respErr.Code)

	if HTTPStatusFromErrorCode(errCode) == 0 {
		return ErrorInternal("invalid code returned from server error response: %s", respErr.Code)
	}

	return &rpcErr{
		code:  errCode,
		msg:   respErr.Msg,
		cause: errors.New(respErr.Cause),
	}
}

func clientError(desc string, err error) Error {
	return WrapError(ErrInternal, err, desc)
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type ErrorPayload struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Cause  string `json:"cause,omitempty"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type Error interface {
	// Code is of the valid error codes
	Code() ErrorCode

	// Msg returns a human-readable, unstructured messages describing the error
	Msg() string

	// Cause is reason for the error
	Cause() error

	// Error returns a string of the form "webrpc error <Code>: <Msg>"
	Error() string

	// Error response payload
	Payload() ErrorPayload
}

func Errorf(code ErrorCode, msgf string, args ...interface{}) Error {
	msg := fmt.Sprintf(msgf, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code)}
}

func WrapError(code ErrorCode, cause error, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg, cause: cause}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code), cause: cause}
}

func Failf(format string, args ...interface{}) Error {
	return Errorf(ErrFail, format, args...)
}

func WrapFailf(cause error, format string, args ...interface{}) Error {
	return WrapError(ErrFail, cause, format, args...)
}

func ErrorNotFound(format string, args ...interface{}) Error {
	return Errorf(ErrNotFound, format, args...)
}

func ErrorInvalidArgument(argument string, validationMsg string) Error {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

func ErrorRequiredArgument(argument string) Error {
	return ErrorInvalidArgument(argument, "is required")
}

func ErrorInternal(format string, args ...interface{}) Error {
	return Errorf(ErrInternal, format, args...)
}

type ErrorCode string

const (
	// Unknown error. For example when handling errors raised by APIs that do not
	// return enough error information.
	ErrUnknown ErrorCode = "unknown"

	// Fail error. General failure error type.
	ErrFail ErrorCode = "fail"

	// Canceled indicates the operation was cancelled (typically by the caller).
	ErrCanceled ErrorCode = "canceled"

	// InvalidArgument indicates client specified an invalid argument. It
	// indicates arguments that are problematic regardless of the state of the
	// system (i.e. a malformed file name, required argument, number out of range,
	// etc.).
	ErrInvalidArgument ErrorCode = "invalid argument"

	// DeadlineExceeded means operation expired before completion. For operations
	// that change the state of the system, this error may be returned even if the
	// operation has completed successfully (timeout).
	ErrDeadlineExceeded ErrorCode = "deadline exceeded"

	// NotFound means some requested entity was not found.
	ErrNotFound ErrorCode = "not found"

	// BadRoute means that the requested URL path wasn't routable to a webrpc
	// service and method. This is returned by the generated server, and usually
	// shouldn't be returned by applications. Instead, applications should use
	// NotFound or Unimplemented.
	ErrBadRoute ErrorCode = "bad route"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	ErrAlreadyExists ErrorCode = "already exists"

	// PermissionDenied indicates the caller does not have permission to execute
	// the specified operation. It must not be used if the caller cannot be
	// identified (Unauthenticated).
	ErrPermissionDenied ErrorCode = "permission denied"

	// Unauthenticated indicates the request does not have valid authentication
	// credentials for the operation.
	ErrUnauthenticated ErrorCode = "unauthenticated"

	// ResourceExhausted indicates some resource has been exhausted, perhaps a
	// per-user quota, or perhaps the entire file system is out of space.
	ErrResourceExhausted ErrorCode = "resource exhausted"

	// FailedPrecondition indicates operation was rejected because the system is
	// not in a state required for the operation's execution. For example, doing
	// an rmdir operation on a directory that is non-empty, or on a non-directory
	// object, or when having conflicting read-modify-write on the same resource.
	ErrFailedPrecondition ErrorCode = "failed precondition"

	// Aborted indicates the operation was aborted, typically due to a concurrency
	// issue like sequencer check failures, transaction aborts, etc.
	ErrAborted ErrorCode = "aborted"

	// OutOfRange means operation was attempted past the valid range. For example,
	// seeking or reading past end of a paginated collection.
	//
	// Unlike InvalidArgument, this error indicates a problem that may be fixed if
	// the system state changes (i.e. adding more items to the collection).
	//
	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.
	// We recommend using OutOfRange (the more specific error) when it applies so
	// that callers who are iterating through a space can easily look for an
	// OutOfRange error to detect when they are done.
	ErrOutOfRange ErrorCode = "out of range"

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	ErrUnimplemented ErrorCode = "unimplemented"

	// Internal errors. When some invariants expected by the underlying system
	// have been broken. In other words, something bad happened in the library or
	// backend service. Do not confuse with HTTP Internal Server Error; an
	// Internal error could also happen on the client code, i.e. when parsing a
	// server response.
	ErrInternal ErrorCode = "internal"

	// Unavailable indicates the service is currently unavailable. This is a most
	// likely a transient condition and may be corrected by retrying with a
	// backoff.
	ErrUnavailable ErrorCode = "unavailable"

	// DataLoss indicates unrecoverable data loss or corruption.
	ErrDataLoss ErrorCode = "data loss"

	// ErrNone is the zero-value, is considered an empty error and should not be
	// used.
	ErrNone ErrorCode = ""
)

func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case ErrCanceled:
		return 408 // RequestTimeout
	case ErrUnknown:
		return 400 // Bad Request
	case ErrFail:
		return 422 // Unprocessable Entity
	case ErrInvalidArgument:
		return 400 // BadRequest
	case ErrDeadlineExceeded:
		return 408 // RequestTimeout
	case ErrNotFound:
		return 404 // Not Found
	case ErrBadRoute:
		return 404 // Not Found
	case ErrAlreadyExists:
		return 409 // Conflict
	case ErrPermissionDenied:
		return 403 // Forbidden
	case ErrUnauthenticated:
		return 401 // Unauthorized
	case ErrResourceExhausted:
		return 403 // Forbidden
	case ErrFailedPrecondition:
		return 412 // Precondition Failed
	case ErrAborted:
		return 409 // Conflict
	case ErrOutOfRange:
		return 400 // Bad Request
	case ErrUnimplemented:
		return 501 // Not Implemented
	case ErrInternal:
		return 500 // Internal Server Error
	case ErrUnavailable:
		return 503 // Service Unavailable
	case ErrDataLoss:
		return 500 // Internal Server Error
	case ErrNone:
		return 200 // OK
	default:
		return 0 // Invalid!
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if rpcErr, ok := err.(Error); ok {
		if rpcErr.Code() == code {
			return true
		}
	}
	return false
}

func IsValidErrorCode(code ErrorCode) bool {
	return HTTPStatusFromErrorCode(code) != 0
}

type rpcErr struct {
	code  ErrorCode
	msg   string
	cause error
}

func (e *rpcErr) Code() ErrorCode {
	return e.code
}

func (e *rpcErr) Msg() string {
	return e.msg
}

func (e *rpcErr) Cause() error {
	return e.cause
}

func (e *rpcErr) Error() string {
	if e.cause != nil && e.cause.Error() != "" {
		if e.msg != "" {
			return fmt.Sprintf("webrpc %s error: %s -- %s", e.code, e.cause.Error(), e.msg)
		} else {
			return fmt.Sprintf("webrpc %s error: %s", e.code, e.cause.Error())
		}
	} else {
		return fmt.Sprintf("webrpc %s error: %s", e.code, e.msg)
	}
}

func (e *rpcErr) Payload() ErrorPayload {
	statusCode := HTTPStatusFromErrorCode(e.Code())
	errPayload := ErrorPayload{
		Status: statusCode,
		Code:   string(e.Code()),
		Msg:    e.Msg(),
		Error:  e.Error(),
	}
	if e.Cause() != nil {
		errPayload.Cause = e.Cause().Error()
	}
	return errPayload
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)
