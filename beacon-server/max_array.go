package server

import "github.com/pkg/errors"

const MaxWorkingStringArray = 1_000

var toomany = errors.New("Too many elements.")
