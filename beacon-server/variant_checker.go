package server

import (
	"regexp"
	"strings"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
)

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
	if vcc, ok := variantCheckerCache.Get(expr); ok {
		vc = vcc.(variantChecker)
		return
	}

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

	variantCheckerCache.Add(expr, vc)

	return
}

func (v *variantChecker) check(s string) bool {
	return v.regex.MatchString(s)
}

var variantCheckerCache *lru.Cache

func init() {
	var err error
	variantCheckerCache, err = lru.New(1_000)
	if err != nil {
		panic(err)
	}
}
