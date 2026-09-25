package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"

	"go.tradeforge.dev/fmp/encoder"
	"go.tradeforge.dev/fmp/model"
)

const (
	apiURL        = "https://financialmodelingprep.com"
	clientVersion = "v0.0.0"

	DefaultRetryCount    = 3
	DefaultClientTimeout = 300 * time.Second
)

func New(
	apiKey string,
	logger *slog.Logger,
) *Client {
	if logger == nil {
		logger = slog.Default()
	}
	red := redactor{apiKey: apiKey}

	c := resty.New()

	c.SetBaseURL(apiURL)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(DefaultClientTimeout)
	c.SetHeader("User-Agent", fmt.Sprintf("Tradeforge client/%v", clientVersion))
	c.SetHeader("Accept", "application/json")
	c.SetQueryParam(apiKeyQueryParam, apiKey)
	// Resty's default logger prints the full request URL, credential included,
	// to stderr on every retry and on the final failure. Replace it here,
	// before any request is built: Client.R() copies the logger into each
	// request, so a later swap would not reach requests already created.
	c.SetLogger(&restyLogger{logger: logger, redactor: red})

	return &Client{
		HTTP:     c,
		encoder:  encoder.New(),
		logger:   logger,
		redactor: red,
	}
}

// Client defines an HTTP client for the Polygon REST API.
type Client struct {
	HTTP     *resty.Client
	encoder  *encoder.Encoder
	logger   *slog.Logger
	redactor redactor
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, path string, params, response any, opts ...model.RequestOption) (*resty.Response, error) {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return nil, fmt.Errorf("encoding params: %w", err)
	}
	return c.CallURL(ctx, method, uri, response, opts...)
}

// CallURL makes an API call based on a request URI and options. The response is automatically unmarshaled.
func (c *Client) CallURL(ctx context.Context, method, uri string, response any, opts ...model.RequestOption) (*resty.Response, error) {
	options := mergeOptions(opts...)

	c.HTTP.SetTimeout(DefaultClientTimeout)
	req := c.HTTP.R().SetContext(ctx)
	if options.Body != nil {
		b, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		req.SetBody(b)
	}
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetResult(response).SetError(&model.ResponseError{})
	req.SetHeader("Content-Type", options.ContentType)

	res, err := req.Execute(method, uri)
	if err != nil {
		// The transport error carries the authenticated URL. Sanitize it
		// before it propagates into handler logs and Sentry.
		return nil, fmt.Errorf("failed to execute request: %w", c.redactor.sanitizeError(err))
	}
	if res.IsError() {
		if slices.Contains(options.IgnoredErrorStatusCodes, res.StatusCode()) {
			return res, nil
		}
		responseError := c.parseResponseError(res)
		if responseError != nil {
			c.logger.Error(
				"response error",
				slog.String("url", c.redactor.redact(uri)),
				slog.Int("status", responseError.StatusCode),
				slog.String("error message", responseError.ErrorMessage),
			)
		} else {
			c.logger.Error(
				"response error",
				slog.String("url", c.redactor.redact(uri)),
				slog.Int("status", res.StatusCode()),
				slog.String("error message", res.Status()),
				slog.String("response", c.redactor.redact(string(res.Body()))),
			)
		}
		return res, fmt.Errorf("service responded with an unexpected error code: %w", responseError)
	}

	if options.Trace {
		c.logger.Debug(
			"request",
			slog.String("url", c.redactor.redact(uri)),
			slog.Any("request headers", c.redactor.redactHeaders(req.Header)),
			slog.Any("response headers", c.redactor.redactHeaders(res.Header())),
		)
	}
	return res, nil
}

// parseResponseError builds the typed error for a non-2xx response. The
// upstream body is carried on it verbatim apart from credentials: an API that
// quotes the request URL back in its error payload would otherwise hand the
// key straight to every caller that renders the message.
func (c *Client) parseResponseError(res *resty.Response) *model.ResponseError {
	if res == nil {
		return nil
	}
	responseError := res.Error().(*model.ResponseError)
	responseError.StatusCode = res.StatusCode()
	responseError.ErrorMessage = c.redactor.redact(res.String())

	return responseError
}

func mergeOptions(opts ...model.RequestOption) *model.RequestOptions {
	options := &model.RequestOptions{
		ContentType: "application/json",
	}
	for _, o := range opts {
		o(options)
	}

	return options
}
