package server

import (
	"context"
	"net/http"
	"os"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

const ApiVersion = "1.0.1"

var (
	CandigUrl  = os.Getenv("CANDIG_URL")
	BeaconName = os.Getenv("BEACON_NAME")
	BeaconId   = os.Getenv("BEACON_ID")
)

var Client = client.NewAPIClient(client.NewConfiguration())

func GetIndex(c *gin.Context) {
	var datasets []BeaconDataset
	listDatasets(c, func(ctx context.Context, d client.Ga4ghDataset) (err error) {
		datasets = append(datasets, BeaconDataset{
			Id:          d.GetId(),
			Name:        d.GetName(),
			Description: d.GetDescription(),
		})
		return
	})
	c.JSON(http.StatusOK, Beacon{
		Name:       BeaconName,
		Id:         BeaconId,
		ApiVersion: ApiVersion,
	})
}

func listDatasets(ctx context.Context, cb func(context.Context, client.Ga4ghDataset) error) (err error) {
	res, _, err := Client.MetadataServiceApi.SearchDatasets(ctx, client.Ga4ghSearchDatasetsRequest{})
	if err != nil {
		return
	}

	for len(res.GetDatasets()) != 0 {
		for _, d := range res.GetDatasets() {
			err = cb(ctx, d)
			if err != nil {
				return
			}
		}
		res, _, err = Client.MetadataServiceApi.SearchDatasets(ctx, client.Ga4ghSearchDatasetsRequest{
			PageToken: client.PtrString(res.GetNextPageToken()),
		})
	}

	return
}

func allDatasets(ctx context.Context) (datasets []BeaconDataset, err error) {
	err = listDatasets(ctx, func(ctx context.Context, d client.Ga4ghDataset) (err error) {
		datasets = append(datasets, BeaconDataset{
			Id:          d.GetId(),
			Name:        d.GetName(),
			Description: d.GetDescription(),
		})
		return
	})
	return
}

func GetQuery(c *gin.Context) {
	req := BeaconAlleleRequest{}

	err := c.Bind(&req)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, BeaconAlleleResponse{
		ApiVersion:    ApiVersion,
		AlleleRequest: req,
	})
}

func PostQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
