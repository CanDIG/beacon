package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
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

type loggingRoundtripper struct{}

func (l loggingRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println(req)
	res, err := http.DefaultTransport.RoundTrip(req)
	fmt.Println(res)
	if err != nil {
		return res, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	fmt.Println(string(data))
	return res, err
}

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

	if os.Getenv("BEACON_VERBOSE") != "" {
		// log all http requests
		cfg.HTTPClient = &http.Client{
			Transport: loggingRoundtripper{},
		}
	}

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
