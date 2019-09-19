package server

import (
	"context"

	client "github.com/CanDIG/beacon/candig-client"
)

func allReferenceSets(ctx context.Context, assId string) (out []string, err error) {
	res, _, err := getClient(ctx).ReferenceServiceApi.SearchReferenceSets(ctx, client.Ga4ghSearchReferenceSetsRequest{
		AssemblyId: &assId,
	})
	if err != nil {
		return
	}

	for len(res.GetReferenceSets()) != 0 {
		for _, rr := range res.GetReferenceSets() {
			out = append(out, rr.GetId())

			// put a cap on the array we return
			if len(out) > MaxWorkingStringArray {
				return nil, toomany
			}
		}

		res, _, err = getClient(ctx).ReferenceServiceApi.SearchReferenceSets(ctx, client.Ga4ghSearchReferenceSetsRequest{
			AssemblyId: &assId,
			PageToken:  client.PtrString(res.GetNextPageToken()),
		})
		if err != nil {
			return
		}
	}

	return
}
