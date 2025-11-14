package market

import (
    "context"
    "net/http"

    "go.tradeforge.dev/fmp/client/rest"
    "go.tradeforge.dev/fmp/model"
)

const (
    GetFMPArticlesPath      = "/stable/fmp-articles"
    GetGeneralNewsPath      = "/stable/news/general-latest"
    GetPressReleasesPath    = "/stable/news/press-releases-latest"
    GetLatestStockNewsPath  = "/stable/news/stock-latest"
    GetStockNewsPath        = "/stable/news/stock"
    GetLatestCryptoNewsPath = "/stable/news/crypto-latest"
    GetCryptoNewsPath       = "/stable/news/crypto"
    GetLatestForexNewsPath  = "/stable/news/forex-latest"
    GetForexNewsPath        = "/stable/news/forex"
)

type NewsClient struct {
    *rest.Client
}

func (nc *NewsClient) GetFMPArticles(ctx context.Context, params model.GetFMPArticlesParams, opts ...model.RequestOption) (model.GetFMPArticlesResponse, error) {
    var res model.GetFMPArticlesResponse
    _, err := nc.Call(ctx, http.MethodGet, GetFMPArticlesPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetGeneralNews(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetGeneralNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetPressReleases(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetPressReleasesPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetLatestStockNews(ctx context.Context, params model.GetLatestNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetLatestStockNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetStockNews(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetStockNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetLatestCryptoNews(ctx context.Context, params model.GetLatestNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetLatestCryptoNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetCryptoNews(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetCryptoNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetLatestForexNews(ctx context.Context, params model.GetLatestNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetLatestForexNewsPath, params, &res, opts...)
    return res, err
}

func (nc *NewsClient) GetForexNews(ctx context.Context, params model.GetNewsParams, opts ...model.RequestOption) (model.GetNewsResponse, error) {
    var res model.GetNewsResponse
    _, err := nc.Call(ctx, http.MethodGet, GetForexNewsPath, params, &res, opts...)
    return res, err
}
