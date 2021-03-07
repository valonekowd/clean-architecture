package errors

import (
	"errors"
)

var (
	ErrNoDataSource = errors.New("no data sources provided")
)
