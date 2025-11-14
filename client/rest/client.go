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
	c := resty.New()

	c.SetBaseURL(apiURL)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(DefaultClientTimeout)
	c.SetHeader("User-Agent", fmt.Sprintf("Tradeforge client/%v", clientVersion))
	c.SetHeader("Accept", "application/json")
	c.SetQueryParam("apikey", apiKey)

	return &Client{
		HTTP:    c,
		encoder: encoder.New(),
		logger:  logger,
	}
}

// Client defines an HTTP client for the Polygon REST API.
type Client struct {
	HTTP    *resty.Client
	encoder *encoder.Encoder
	logger  *slog.Logger
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
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	if res.IsError() {
		if slices.Contains(options.IgnoredErrorStatusCodes, res.StatusCode()) {
			return res, nil
		}
		responseError := parseResponseError(res)
		if responseError != nil {
			c.logger.Error(
				"response error",
				slog.String("url", uri),
				slog.Int("status", responseError.StatusCode),
				slog.String("error message", responseError.ErrorMessage),
			)
		} else {
			c.logger.Error(
				"response error",
				slog.String("url", uri),
				slog.Int("status", res.StatusCode()),
				slog.String("error message", res.Status()),
				slog.String("response", string(res.Body())),
			)
		}
		return res, fmt.Errorf("service responded with an unexpected error code: %w", responseError)
	}

	if options.Trace {
		sanitizedHeaders := req.Header
		for k := range sanitizedHeaders {
			if k == "Authorization" {
				sanitizedHeaders[k] = []string{"REDACTED"}
			}
		}
		c.logger.Debug(
			"request",
			slog.String("url", uri),
			slog.Any("request headers", sanitizedHeaders),
			slog.Any("response headers", res.Header()),
		)
	}
	return res, nil
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

func parseResponseError(res *resty.Response) *model.ResponseError {
	if res == nil {
		return nil
	}
	responseError := res.Error().(*model.ResponseError)
	responseError.StatusCode = res.StatusCode()
	responseError.ErrorMessage = res.String()

	return responseError
}
