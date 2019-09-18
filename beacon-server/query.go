package server

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	client "github.com/CanDIG/beacon/candig-client"
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

// allVariants is a routine that streams every variant in a dataset
// within the given range into a provided channel, closing it when
// finished.
func allVariants(ctx context.Context, dataset string, req BeaconAlleleRequest, variantsets []string, ch chan<- client.Ga4ghVariant) (err error) {
	defer close(ch)

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

const MaxWorkingStringArray = 1_000

type toomany struct{}

func (t toomany) Error() string {
	return "Too many values to return."
}

func allReferenceSets(ctx context.Context, assId string) (out []string, err error) {
	res, _, err := getClient(ctx).ReferenceServiceApi.SearchReferenceSets(ctx, client.Ga4ghSearchReferenceSetsRequest{
		AssemblyId: &assId,
	})
	if err != nil {
		return
	}

	for len(res.GetReferenceSets()) != 0 {
		for _, rr := range res.GetReferenceSets() {
			select {
			default:
			case <-ctx.Done():
				err = ctx.Err()
				return
			}
			out = append(out, rr.GetId())

			// put a cap on the array we return
			if len(out) > MaxWorkingStringArray {
				return nil, toomany{}
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
					return nil, toomany{}
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
		var count int64
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
				VariantCount: count,
				Exists:       count > 0,
				// todo: add more fields with metadata
			})
		}
	}

	return
}

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

var variantCheckerChecker *regexp.Regexp

func init() {
	var err error
	variantCheckerChecker, err = regexp.Compile("^[ACGTN]*$")
	if err != nil {
		panic(err)
	}
}

type variantChecker struct {
	regex *regexp.Regexp
}

func newVariantChecker(expr string) (vc variantChecker, err error) {
	if !variantCheckerChecker.MatchString(expr) {
		err = errors.New("Invalid variant search string")
		return
	}

	expr = strings.ReplaceAll(expr, "N", ".")
	expr = "^" + expr + "$"

	vc.regex, err = regexp.Compile(expr)
	if err != nil {
		return
	}

	return
}

func (v *variantChecker) check(s string) bool {
	return v.regex.MatchString(s)
}
