package server

import (
	"context"

	"github.com/pkg/errors"
)

func countVariants(ctx context.Context, dataset string, refsetsmap map[string]struct{}, req BeaconAlleleRequest) (count int, err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	var errcs []<-chan error

	// build list of variantsets
	variantsets, err := allVariantSets(ctx, dataset, refsetsmap)
	if _, ok := errors.Cause(err).(toomany); ok {
		variantsets, err = nil, nil
	}
	if err != nil {
		return
	}

	// iterate through variants
	varch, errc := allVariantsPipeline(ctx, dataset, req, variantsets)
	errcs = append(errcs, errc)

	// filter out the ones that are valid
	validc, errc := isValidVariantPipeline(ctx, variantsets != nil, varch, req)
	errcs = append(errcs, errc)

	for i := range validc {
		count += i

		// short circuit to avoid iterating over all
		// the data
		if req.IncludeDatasetResponses == "NONE" {
			return
		}
	}

	err = waitPipeline(errcs...)
	return
}
