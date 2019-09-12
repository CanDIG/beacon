package server

import (
	"context"
	"os"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

const apiClient = "github.com/CanDIG/beacon/apiclient"

var passthroughHeaders = [...]string{
	"Authorization",
}

const xClaim = "X-Claim-"

var CandigHost = os.Getenv("CANDIG_HOST")

func AddOauth(c *gin.Context) {
	cfg := client.NewConfiguration()
	for _, h := range passthroughHeaders {
		auth := c.GetHeader(h)
		if auth != "" {
			cfg.AddDefaultHeader(h, auth)
		}
	}
	for k, v := range c.Request.Header {
		if k[:len(xClaim)] == xClaim {
			for _, v := range v {
				cfg.AddDefaultHeader(k, v)
			}
		}
	}

	c.Set(apiClient, client.NewAPIClient(cfg))
	c.Next()
}

func getClient(ctx context.Context) *client.APIClient {
	return ctx.Value(apiClient).(*client.APIClient)
}
