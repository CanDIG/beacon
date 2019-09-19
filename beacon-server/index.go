package server

import (
	"context"
	"net/http"
	"time"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/gin-gonic/gin"
)

func listDatasets(ctx context.Context, cb func(context.Context, client.Ga4ghDataset) error) (err error) {
	res, _, err := getClient(ctx).MetadataServiceApi.SearchDatasets(ctx, client.Ga4ghSearchDatasetsRequest{})
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
		res, _, err = getClient(ctx).MetadataServiceApi.SearchDatasets(ctx, client.Ga4ghSearchDatasetsRequest{
			PageToken: client.PtrString(res.GetNextPageToken()),
		})
	}

	return
}

func getAssemblyId(ctx context.Context, dataId string) (assId string, err error) {
	vr, _, err := getClient(ctx).VariantServiceApi.SearchVariantSets(ctx, client.Ga4ghSearchVariantSetsRequest{
		DatasetId: client.PtrString(dataId),
		PageSize:  client.PtrInt32(1),
	})
	if err != nil {
		return
	}

	rr, _, err := getClient(ctx).ReferenceServiceApi.GetReferenceSet(ctx, vr.GetVariantSets()[0].GetReferenceSetId())
	if err != nil {
		return
	}

	assId = rr.GetAssemblyId()
	return
}

var DatasetOrigin = time.Date(2019, 9, 10, 11, 45, 00, 00, time.UTC)

func allDatasets(ctx context.Context) (datasets []BeaconDataset, err error) {
	err = listDatasets(ctx, func(ctx context.Context, d client.Ga4ghDataset) (err error) {
		assId, err := getAssemblyId(ctx, d.GetId())
		if err != nil {
			return
		}

		datasets = append(datasets, BeaconDataset{
			Id:             d.GetId(),
			Name:           d.GetName(),
			Description:    d.GetDescription(),
			AssemblyId:     assId,
			CreateDateTime: DatasetOrigin.String(), // all records created at once
			UpdateDateTime: time.Now().String(),    // they are always updated
		})
		return
	})
	return
}

var (
	ApiVersion        = "1.0.1"
	BeaconName        = "CanDig Beacon"
	BeaconId          = "ca.distributedgenomics.beacon"
	OrgId             = "CanDIG"
	OrgName           = "Canadian Distributed Genomics"
	OrgWelcomeUrl     = "https://www.distributedgenomics.ca"
	OrgContactUrl     = "mailto:info@distributedgenomics.ca"
	OrgLogoUrl        = "https://www.distributedgenomics.ca/img/logo_only.png"
	BeaconDescription = "Beacon interface to the CanDIG genomics database."
)

func GetBeacon(c *gin.Context) {
	datasets, err := allDatasets(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Beacon{
		WelcomeUrl: c.Request.URL.String(),
		Name:       BeaconName,
		Id:         BeaconId,
		ApiVersion: ApiVersion,
		Datasets:   datasets,
		Organization: BeaconOrganization{
			Id:         OrgId,
			Name:       OrgName,
			ContactUrl: OrgContactUrl,
			WelcomeUrl: OrgWelcomeUrl,
			LogoUrl:    OrgLogoUrl,
		},
		// todo: add more organizational metadata
	})
}
