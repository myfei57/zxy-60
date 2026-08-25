package model

import "strings"

func ValidateFlight(f Flight) error {
	if strings.TrimSpace(f.ID) == "" {
		return errEmptyID
	}
	if strings.TrimSpace(f.Code) == "" {
		return errEmptyCode
	}
	return nil
}

func ValidateBag(b Bag) error {
	if strings.TrimSpace(b.ID) == "" {
		return errEmptyID
	}
	if strings.TrimSpace(b.Barcode) == "" {
		return errEmptyBarcode
	}
	return nil
}

func IsOpenFlightState(state FlightState) bool {
	return state == FlightOpen || state == FlightBoarding
}

func IsActiveBagState(state BagState) bool {
	switch state {
	case BagCheckedIn, BagInjected, BagSorted, BagRechecked:
		return true
	default:
		return false
	}
}
