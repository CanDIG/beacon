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

	// not sure why an empty url results in silence from
	// everything but oh well
	if u == nil {
		panic("No CANDIG_URL")
	}

	cfg.Scheme = u.Scheme
	cfg.Host = u.Host
	cfg.BasePath = u.EscapedPath()

	// add header to disable federation and expose the underlying
	// openapi code as needed
	cfg.AddDefaultHeader("X-No-Federation", "yes")

	// strip slashes from the end because our client doesn't like
	// them at all
	for len(cfg.BasePath) > 0 && cfg.BasePath[len(cfg.BasePath)-1] == '/' {
		cfg.BasePath = cfg.BasePath[:len(cfg.BasePath)-1]
	}
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
