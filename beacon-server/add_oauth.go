package server

import (
	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

var passthroughHeaders = [...]string{
	"Authorization",
}

const xClaim = "X-Claim-"

func addOauth(c *gin.Context) {
	cfg := client.NewConfiguration()
	for _, h := range passthroughHeaders {
		auth := c.GetHeader(h)
		if auth != "" {
			cfg.AddDefaultHeader(h, auth)
		}
	}
	for k, v := range c.Request.Header {
		if len(xClaim) <= len(k) && k[:len(xClaim)] == xClaim {
			for _, v := range v {
				cfg.AddDefaultHeader(k, v)
			}
		}
	}

	storeClientGin(c, cfg)
	c.Next()
}
