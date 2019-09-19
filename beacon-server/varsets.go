package server

import (
	"context"

	client "github.com/CanDIG/beacon/candig-client"
)

func allVariantSets(ctx context.Context, dataId string, refsetmap map[string]struct{}) (variantsets []string, err error) {
	// we can't filter them out so don't bother returning them
	if refsetmap == nil {
		return
	}

	res, _, err := getClient(ctx).VariantServiceApi.SearchVariantSets(ctx, client.Ga4ghSearchVariantSetsRequest{
		DatasetId: &dataId,
	})
	if err != nil {
		return
	}

	for len(res.GetVariantSets()) != 0 {
		for _, vs := range res.GetVariantSets() {
			if _, ok := refsetmap[vs.GetReferenceSetId()]; ok {
				variantsets = append(variantsets, vs.GetId())

				// cap the array we return
				if len(variantsets) > MaxWorkingStringArray {
					return nil, toomany
				}
			}
		}

		res, _, err = getClient(ctx).VariantServiceApi.SearchVariantSets(ctx, client.Ga4ghSearchVariantSetsRequest{
			DatasetId: &dataId,
			PageToken: client.PtrString(res.GetNextPageToken()),
		})
		if err != nil {
			return
		}
	}

	return
}
