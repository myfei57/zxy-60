package belt

import (
	"strings"

	"bagsort/internal/model"
)

func ValidateTransfer(cmd model.TransferCommand) bool {
	return strings.TrimSpace(cmd.FlightID) != "" && strings.TrimSpace(cmd.Carousel) != ""
}
