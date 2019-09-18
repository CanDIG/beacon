package server

import (
	"context"
	"strconv"

	client "github.com/CanDIG/beacon/candig-client"
	"github.com/pkg/errors"
)

func countVariants(ctx context.Context, dataset string, refsetsmap map[string]struct{}, req BeaconAlleleRequest, refvc, altvc variantChecker) (count int64, err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	// build list of variantsets
	variantsets, err := allVariantSets(ctx, dataset, refsetsmap)
	if _, ok := errors.Cause(err).(toomany); ok {
		variantsets, err = nil, nil
	}
	if err != nil {
		return
	}

	// iterate through variants
	varch := make(chan client.Ga4ghVariant)
	go func() {
		defer close(varch)
		err = allVariants(ctx, dataset, req, variantsets, varch)
	}()

	for variant := range varch {
		var ok bool
		ok, err = isvalidVariant(ctx, variantsets != nil, variant, req, refvc, altvc)
		if err != nil {
			return
		}

		if ok {
			count++

			// short circuit to avoid iterating over all
			// the data
			if req.IncludeDatasetResponses == "NONE" {
				return
			}
		}
	}

	return
}

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
