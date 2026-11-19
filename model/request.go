package model

import (
	"context"
	"net/http"
	"net/url"
)

type scheduledRetryContextKey struct{}

// WithScheduledRetries opts scheduled GET calls into bounded gateway recovery.
// Interactive calls should retain their own short request budget.
func WithScheduledRetries(ctx context.Context) context.Context {
	return context.WithValue(ctx, scheduledRetryContextKey{}, true)
}

// IsScheduledRetryContext reports whether scheduled gateway recovery is enabled.
func IsScheduledRetryContext(ctx context.Context) bool {
	enabled, _ := ctx.Value(scheduledRetryContextKey{}).(bool)
	return enabled
}

// RequestOptions are used to configure client calls.
type RequestOptions struct {
	// Body to pass with the request.
	Body any

	// Headers to apply to the request.
	Headers http.Header

	// Query params to apply to the request.
	QueryParams url.Values

	// ContentType to use in the header.
	ContentType string

	// IgnoredErrorStatusCodes are the error status codes that should be ignored.
	IgnoredErrorStatusCodes []int

	// Trace enables request tracing.
	Trace bool
}

// RequestOption changes the configuration of RequestOptions.
type RequestOption func(o *RequestOptions)

// Body sets a body as an option.
func Body(body any) RequestOption {
	return func(o *RequestOptions) {
		o.Body = body
	}
}

// Header sets a header as an option.
func Header(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(http.Header)
		}

		o.Headers.Add(key, value)
	}
}

// QueryParam sets a query param as an option.
func QueryParam(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.QueryParams == nil {
			o.QueryParams = make(url.Values)
		}

		o.QueryParams.Add(key, value)
	}
}

func WithContentType(contentType string) RequestOption {
	return func(o *RequestOptions) {
		o.ContentType = contentType
	}
}

func WithIgnoredErrorStatusCodes(codes ...int) RequestOption {
	return func(o *RequestOptions) {
		o.IgnoredErrorStatusCodes = codes
	}
}

// WithTrace enables or disables request tracing.
func WithTrace(trace bool) RequestOption {
	return func(o *RequestOptions) {
		o.Trace = trace
	}
}
