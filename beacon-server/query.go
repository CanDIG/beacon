package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetBeaconAlleleResponse(c *gin.Context) {
	query(c)
}

func PostBeaconAlleleResponse(c *gin.Context) {
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
	// setup context
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	// setup return
	status := BeaconAlleleResponse{
		ApiVersion: ApiVersion,
		BeaconId:   BeaconId,
	}

	// bind data
	var err error
	fail := func(code int) {
		status.Error = BeaconError{
			ErrorCode:    int32(code),
			ErrorMessage: err.Error(),
		}
		c.AbortWithStatusJSON(code, status)
	}
	var req BeaconAlleleRequest
	err = c.Bind(&req)
	if err != nil {
		fail(http.StatusBadRequest)
		return
	}

	status.Exists, status.DatasetAlleleResponses, err = internalRun(ctx, req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Data: status,
		Offered: []string{
			gin.MIMEYAML,
			gin.MIMEJSON,
		},
	})
	return
}

func internalRun(ctx context.Context, req BeaconAlleleRequest) (exists bool, out []BeaconDatasetAlleleResponse, err error) {
	defer func() {
		err = errors.WithStack(err)
	}()

	// argument cleanup
	if len(req.DatasetIds) == 0 {
		req.DatasetIds, err = allDatasetsStrings(ctx)
		if err != nil {
			return
		}
	}
	if req.IncludeDatasetResponses == "" {
		req.IncludeDatasetResponses = "NONE"
	}
	if req.StartMin == 0 {
		req.StartMin = int32(req.Start)
	} else if req.StartMax == 0 {
		err = errors.New("missing startMax parameter")
		return
	}
	if req.EndMax == 0 {
		req.EndMax = req.End
	}

	// build list of refsets
	refsets, err := allReferenceSets(ctx, req.AssemblyId)
	if _, ok := errors.Cause(err).(toomany); ok {
		refsets, err = nil, nil
	}
	if err != nil {
		return
	}
	refsetsmap := make(map[string]struct{})
	if refsets != nil {
		for _, refset := range refsets {
			refsetsmap[refset] = struct{}{}
		}
	} else {
		refsetsmap = nil
	}

	// setup variant checkers
	refvc, err := newVariantChecker(req.ReferenceBases)
	if err != nil {
		return
	}
	altvc, err := newVariantChecker(req.AlternateBases)
	if err != nil {
		return
	}

	// loop through datasets
	for _, dataset := range req.DatasetIds {
		var count int
		count, err = countVariants(ctx, dataset, refsetsmap, req, refvc, altvc)
		if err != nil {
			return
		}

		if count > 0 {
			exists = true
		}

		incl := false
		if req.IncludeDatasetResponses == "ALL" {
			incl = true
		} else if req.IncludeDatasetResponses == "HIT" && count > 0 {
			incl = true
		} else if req.IncludeDatasetResponses == "MISS" && count == 0 {
			incl = true
		}
		if incl {
			out = append(out, BeaconDatasetAlleleResponse{
				DatasetId:    dataset,
				VariantCount: int64(count),
				Exists:       count > 0,
				// todo: add more fields with metadata
			})
		}
	}

	return
}
