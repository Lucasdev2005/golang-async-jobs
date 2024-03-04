package enums

const (
	HttpStatusOK               = 200
	HttpStatusCreated          = 201
	HttpStatusAccepted         = 202
	HttpStatusNonAuthoritative = 203
	HttpStatusNoContent        = 204
	HttpStatusResetContent     = 205
	HttpStatusPartialContent   = 206
	HttpStatusMultiStatus      = 207
	HttpStatusAlreadyReported  = 208
	HttpStatusIMUsed           = 226

	InternalServerError           = 500
	NotImplemented                = 501
	BadGateway                    = 502
	ServiceUnavailable            = 503
	GatewayTimeout                = 504
	HTTPVersionNotSupported       = 505
	VariantAlsoNegotiates         = 506
	InsufficientStorage           = 507
	LoopDetected                  = 508
	NotExtended                   = 510
	NetworkAuthenticationRequired = 511

	BadRequest                   = 400
	Unauthorized                 = 401
	PaymentRequired              = 402
	Forbidden                    = 403
	NotFound                     = 404
	MethodNotAllowed             = 405
	NotAcceptable                = 406
	ProxyAuthRequired            = 407
	RequestTimeout               = 408
	Conflict                     = 409
	Gone                         = 410
	LengthRequired               = 411
	PreconditionFailed           = 412
	RequestEntityTooLarge        = 413
	RequestURITooLong            = 414
	UnsupportedMediaType         = 415
	RequestedRangeNotSatisfiable = 416
	ExpectationFailed            = 417
	Teapot                       = 418
	UnprocessableEntity          = 422
	Locked                       = 423
	FailedDependency             = 424
	UpgradeRequired              = 426
	PreconditionRequired         = 428
	TooManyRequests              = 429
	RequestHeaderFieldsTooLarge  = 431
	UnavailableForLegalReasons   = 451
)