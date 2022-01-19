package main

type METHOD string

const (
	Get     METHOD = "GET"     // RFC 7231, 4.3.1
	Head           = "HEAD"    // RFC 7231, 4.3.2
	Post           = "POST"    // RFC 7231, 4.3.3
	Put            = "PUT"     // RFC 7231, 4.3.4
	Patch          = "PATCH"   // RFC 5789
	Delete         = "DELETE"  // RFC 7231, 4.3.5
	Connect        = "CONNECT" // RFC 7231, 4.3.6
	Options        = "OPTIONS" // RFC 7231, 4.3.7
	Trace          = "TRACE"   // RFC 7231, 4.3.8
)

func (m METHOD) convert() string {
	return string(m)
}

type STATUS int

func (s STATUS) convert() int {
	return int(s)
}

const (
	Continue           STATUS = 100 // RFC 7231, 6.2.1
	SwitchingProtocols        = 101 // RFC 7231, 6.2.2
	Processing                = 102 // RFC 2518, 10.1

	OK                   = 200 // RFC 7231, 6.3.1
	Created              = 201 // RFC 7231, 6.3.2
	Accepted             = 202 // RFC 7231, 6.3.3
	NonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	NoContent            = 204 // RFC 7231, 6.3.5
	ResetContent         = 205 // RFC 7231, 6.3.6
	PartialContent       = 206 // RFC 7233, 4.1
	MultiStatus          = 207 // RFC 4918, 11.1
	AlreadyReported      = 208 // RFC 5842, 7.1
	IMUsed               = 226 // RFC 3229, 10.4.1

	MultipleChoices   = 300 // RFC 7231, 6.4.1
	MovedPermanently  = 301 // RFC 7231, 6.4.2
	Found             = 302 // RFC 7231, 6.4.3
	SeeOther          = 303 // RFC 7231, 6.4.4
	NotModified       = 304 // RFC 7232, 4.1
	UseProxy          = 305 // RFC 7231, 6.4.5
	_                 = 306 // RFC 7231, 6.4.6 (Unused)
	TemporaryRedirect = 307 // RFC 7231, 6.4.7
	PermanentRedirect = 308 // RFC 7538, 3

	BadRequest                   = 400 // RFC 7231, 6.5.1
	Unauthorized                 = 401 // RFC 7235, 3.1
	PaymentRequired              = 402 // RFC 7231, 6.5.2
	Forbidden                    = 403 // RFC 7231, 6.5.3
	NotFound                     = 404 // RFC 7231, 6.5.4
	MethodNotAllowed             = 405 // RFC 7231, 6.5.5
	NotAcceptable                = 406 // RFC 7231, 6.5.6
	ProxyAuthRequired            = 407 // RFC 7235, 3.2
	RequestTimeout               = 408 // RFC 7231, 6.5.7
	Conflict                     = 409 // RFC 7231, 6.5.8
	Gone                         = 410 // RFC 7231, 6.5.9
	LengthRequired               = 411 // RFC 7231, 6.5.10
	PreconditionFailed           = 412 // RFC 7232, 4.2
	RequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	RequestURITooLong            = 414 // RFC 7231, 6.5.12
	UnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	RequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	ExpectationFailed            = 417 // RFC 7231, 6.5.14
	Teapot                       = 418 // RFC 7168, 2.3.3
	UnprocessableEntity          = 422 // RFC 4918, 11.2
	Locked                       = 423 // RFC 4918, 11.3
	FailedDependency             = 424 // RFC 4918, 11.4
	UpgradeRequired              = 426 // RFC 7231, 6.5.15
	PreconditionRequired         = 428 // RFC 6585, 3
	TooManyRequests              = 429 // RFC 6585, 4
	RequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	UnavailableForLegalReasons   = 451 // RFC 7725, 3

	InternalServerError           = 500 // RFC 7231, 6.6.1
	NotImplemented                = 501 // RFC 7231, 6.6.2
	BadGateway                    = 502 // RFC 7231, 6.6.3
	ServiceUnavailable            = 503 // RFC 7231, 6.6.4
	GatewayTimeout                = 504 // RFC 7231, 6.6.5
	HTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	VariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	InsufficientStorage           = 507 // RFC 4918, 11.5
	LoopDetected                  = 508 // RFC 5842, 7.2
	NotExtended                   = 510 // RFC 2774, 7
	NetworkAuthenticationRequired = 511 // RFC 6585, 6
)

type Settings struct {

	// Enables automatic recovering from panic while executing handlers by
	// answering with HTTP status code 500 and logging error message without
	// stopping service
	AutoRecover bool // default false

	// ServerName for sending in response headers
	ServerName string // default ""

	// Maximum request body size
	MaxRequestBodySize int // default 4 * 1024 * 1024

	// Maximum number of route params count
	MaxRouteParams int // default 1024

	// Max request url length
	MaxRequestURLLength int // default 2048

	// Enable case sensitive routing
	CaseInSensitive bool
}
