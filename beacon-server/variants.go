package server

import (
	"context"

	client "github.com/CanDIG/beacon/candig-client"
)

// allVariants is a routine that streams every variant in a dataset
// within the given range into a provided channel, closing it when
// finished.
func allVariants(ctx context.Context, dataset string, req BeaconAlleleRequest, variantsets []string, ch chan<- client.Ga4ghVariant) (err error) {
	// omit end if it's the default 0
	var end *string
	if req.EndMax > 0 {
		end = client.PtrString(string(req.EndMax))
	}

	res, _, err := getClient(ctx).VariantServiceApi.SearchVariants(ctx, client.Ga4ghSearchVariantsRequest{
		DatasetId:     &dataset,
		Start:         client.PtrString(string(req.StartMin)), // start will default to 0 anyway
		End:           end,
		ReferenceName: client.PtrString(string(req.ReferenceName)),
		VariantSetIds: &variantsets,
	})
	if err != nil {
		return
	}

	for res.GetVariants() != nil {
		for _, v := range res.GetVariants() {
			select {
			case ch <- v:
			case <-ctx.Done():
				err = ctx.Err()
				return
			}
		}

		// omit end if it's the default 0
		var end *string
		if req.EndMax > 0 {
			end = client.PtrString(string(req.EndMax))
		}

		res, _, err = getClient(ctx).VariantServiceApi.SearchVariants(ctx, client.Ga4ghSearchVariantsRequest{
			DatasetId:     &dataset,
			Start:         client.PtrString(string(req.StartMin)),
			End:           end,
			ReferenceName: client.PtrString(string(req.ReferenceName)),
			PageToken:     client.PtrString(res.GetNextPageToken()),
		})
		if err != nil {
			return
		}
	}

	return
}

func allVariantsPipeline(ctx context.Context, dataset string, req BeaconAlleleRequest, variantsets []string) (<-chan client.Ga4ghVariant, <-chan error) {
	out := make(chan client.Ga4ghVariant)
	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		defer close(out)
		errc <- allVariants(ctx, dataset, req, variantsets, out)
	}()

	return out, errc
}
