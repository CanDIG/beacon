package server

import (
	"context"
	"net/http"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

func GetQuery(c *gin.Context) {
	query(c)
}

func PostQuery(c *gin.Context) {
	query(c)
}

func allDatasetsStrings(ctx context.Context) (out []string, err error) {
	r, err := allDatasets(ctx)
	if err != nil {
		return
	}

	out = make([]string, len(r))
	for i := range r {
		out[i] = r[i].Id
	}
	return
}

func query(c *gin.Context) {
	status := BeaconAlleleResponse{
		ApiVersion: ApiVersion,
		BeaconId:   BeaconId,
	}
	// cleanup
	var err error
	fail := func(code uint) {
		status.Error = BeaconError{
			ErrorCode:    int32(code),
			ErrorMessage: err.Error(),
		}
		c.AbortWithStatusJSON(int(code), status)
	}

	var req BeaconAlleleRequest
	err = c.Bind(&req)
	if err != nil {
		fail(400)
		return
	}

	if len(req.DatasetIds) == 0 {
		req.DatasetIds, err = allDatasetsStrings(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	if req.IncludeDatasetResponses == "" {
		req.IncludeDatasetResponses = "NONE"
	}

	var out []BeaconDatasetAlleleResponse
	if req.IncludeDatasetResponses != "NONE" {
		components := []client.Ga4ghComponent{
			client.Ga4ghComponent{
				Id:       client.PtrString("A"),
				Variants: &client.Ga4ghSearchVariantsRequest{},
			},
		}
		for _, d := range req.DatasetIds {
			res, _, err := getClient(c).SearchServiceApi.GetCount(c, client.Ga4ghSearchQueryRequest{
				DatasetId:  client.PtrString(d),
				Components: &components,
				// Logic:      &logic,
				// Results:    &results,
			})
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			out = append(out, BeaconDatasetAlleleResponse{
				DatasetId:    d,
				Exists:       true,
				VariantCount: int64(len(res.GetVariants())), // XXX
			})
		}

	}

	status.DatasetAlleleResponses = out
	c.JSON(http.StatusOK, status)
	return
}
