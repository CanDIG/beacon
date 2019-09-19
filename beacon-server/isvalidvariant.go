package server

import (
	"context"
	"strconv"
	"sync"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/pkg/errors"
)

func isvalidVariant(ctx context.Context, hasVariantSets bool, variant client.Ga4ghVariant, req BeaconAlleleRequest, refvc, altvc variantChecker) (result bool, err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	// If we didn't already filter out by assembly_id then we
	// gotta do it now.
	if !hasVariantSets {
		var res client.Ga4ghVariantSet
		res, _, err = getClient(ctx).VariantServiceApi.GetVariantSet(ctx, variant.GetVariantSetId())
		if err != nil {
			return
		}

		var ress client.Ga4ghReferenceSet
		ress, _, err = getClient(ctx).ReferenceServiceApi.GetReferenceSet(ctx, res.GetReferenceSetId())
		if err != nil {
			return
		}

		if ress.GetAssemblyId() != req.AssemblyId {
			return
		}
	}

	i, err := strconv.Atoi(*variant.Start)
	if err != nil {
		return
	}
	if req.StartMax != 0 && i > int(req.StartMax) {
		return
	}

	i, err = strconv.Atoi(*variant.End)
	if err != nil {
		return
	}
	if req.EndMin != 0 && i < int(req.EndMin) {
		return
	}

	if !refvc.check(variant.GetReferenceBases()) {
		return
	}

	// either we need a variant_type or we need alternate_bases
	if req.VariantType != "" {
		if variant.GetVariantType() != req.VariantType {
			return
		}
	} else {
		found := false
		for _, alt := range variant.GetAlternateBases() {
			if altvc.check(alt) {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	return true, nil
}

func isValidVariantPipeline(ctx context.Context, hasVariantSets bool, variantc <-chan client.Ga4ghVariant, req BeaconAlleleRequest, refvc, altvc variantChecker) (<-chan struct{}, <-chan error) {
	out := make(chan struct{})
	errc := make(chan error, 1)
	var wg sync.WaitGroup
	const isvalidcount = 20
	wg.Add(isvalidcount)
	for i := 0; i < isvalidcount; i++ {
		go func() {
			defer wg.Done()
			for variant := range variantc {
				ok, err := isvalidVariant(ctx, hasVariantSets, variant, req, refvc, altvc)
				if err != nil {
					errc <- err
					return
				}
				if ok {
					out <- struct{}{}
				}
			}
		}()
	}
	go func() {
		defer close(out)
		defer close(errc)
		wg.Wait()
	}()
	return out, errc
}
