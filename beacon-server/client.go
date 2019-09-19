package server

import (
	"context"
	"net/url"
	"os"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

const (
	apiClient       = "github.com/CanDIG/beacon/apiclient"
	apiClientConfig = apiClient + "config"
)

var candigURL = os.Getenv("CANDIG_URL")

func storeURL(cfg *client.Configuration) {
	u, err := url.Parse(candigURL)
	if err != nil {
		panic(err)
	}

	cfg.Scheme = u.Scheme
	cfg.Host = u.Host
	cfg.BasePath = u.EscapedPath()
}

func storeClientGin(c *gin.Context, cfg *client.Configuration) {
	storeURL(cfg)
	c.Set(apiClient, client.NewAPIClient(cfg))
	c.Set(apiClientConfig, cfg)
}

func storeClientContext(ctx context.Context, cfg *client.Configuration) context.Context {
	storeURL(cfg)
	ctx = context.WithValue(ctx, apiClient, client.NewAPIClient(cfg))
	ctx = context.WithValue(ctx, apiClientConfig, cfg)
	return ctx
}

func getClient(ctx context.Context) *client.APIClient {
	return ctx.Value(apiClient).(*client.APIClient)
}

func getClientConfig(ctx context.Context) *client.Configuration {
	return ctx.Value(apiClientConfig).(*client.Configuration)
}
