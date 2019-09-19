package server

import (
	"context"
	"runtime"
	"strconv"
	"sync"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/pkg/errors"
)

func isvalidVariant(ctx context.Context, hasVariantSets bool, variant client.Ga4ghVariant, req BeaconAlleleRequest) (result bool, err error) {
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

	refvc, err := newVariantChecker(req.ReferenceBases)
	if err != nil {
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
		var altvc variantChecker
		altvc, err = newVariantChecker(req.ReferenceBases)
		if err != nil {
			return
		}

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

// number of goroutines per cpu if they block on the network
const sleepableFactor = 100

func isValidVariantPipeline(ctx context.Context, hasVariantSets bool, variantc <-chan client.Ga4ghVariant, req BeaconAlleleRequest) (<-chan int, <-chan error) {
	out := make(chan int)
	errc := make(chan error, 1)

	// only run as many routines as there are cpus
	isvalidcount := runtime.NumCPU()

	// increase the number of routines if they might get stalled
	// on waiting for the network
	if !hasVariantSets {
		isvalidcount *= sleepableFactor
	}

	var wg sync.WaitGroup
	wg.Add(isvalidcount)
	for i := 0; i < isvalidcount; i++ {
		go func() {
			defer wg.Done()
			for variant := range variantc {
				ok, err := isvalidVariant(ctx, hasVariantSets, variant, req)
				if err != nil {
					errc <- err
					return
				}
				if ok {
					select {
					case out <- 1:
					case <-ctx.Done():
						errc <- ctx.Err()
						return
					}
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
