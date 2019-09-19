package server

import (
	"context"
	"os"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

const (
	apiClient       = "github.com/CanDIG/beacon/apiclient"
	apiClientConfig = apiClient + "config"
)

var CandigHost = os.Getenv("CANDIG_HOST")

func storeClientGin(c *gin.Context, cfg *client.Configuration) {
	if cfg.Host == "" {
		cfg.Host = CandigHost
	}
	c.Set(apiClient, client.NewAPIClient(cfg))
	c.Set(apiClientConfig, cfg)
}

func storeClientContext(ctx context.Context, cfg *client.Configuration) context.Context {
	if cfg.Host == "" {
		cfg.Host = CandigHost
	}
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
