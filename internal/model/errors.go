package model

import "errors"

var (
	errEmptyID      = errors.New("id must not be empty")
	errEmptyCode    = errors.New("code must not be empty")
	errEmptyBarcode = errors.New("barcode must not be empty")
)
