package model

type FlightState string

const (
	FlightOpen     FlightState = "open"
	FlightBoarding FlightState = "boarding"
	FlightCutoff   FlightState = "cutoff"
	FlightClosed   FlightState = "closed"
)

type BagState string

const (
	BagCheckedIn BagState = "checked_in"
	BagInjected  BagState = "injected"
	BagSorted    BagState = "sorted"
	BagRechecked BagState = "rechecked"
	BagLoaded    BagState = "loaded"
	BagDiverted  BagState = "diverted"
)

type Flight struct {
	ID        string      `json:"id"`
	Code      string      `json:"code"`
	Terminal  string      `json:"terminal"`
	Carousel  string      `json:"carousel"`
	Chute     string      `json:"chute"`
	State     FlightState `json:"state"`
	Batch     uint64      `json:"batch"`
	CutoffAt  string      `json:"cutoff_at,omitempty"`
	ClosedAt  string      `json:"closed_at,omitempty"`
}

type Bag struct {
	ID       string   `json:"id"`
	Barcode  string   `json:"barcode"`
	Sequence uint64   `json:"sequence"`
	FlightID string   `json:"flight_id"`
	ChuteID  string   `json:"chute_id"`
	State    BagState `json:"state"`
}

type ChuteAssignment struct {
	FlightID  string `json:"flight_id"`
	ChuteID   string `json:"chute_id"`
	UpdatedAt string `json:"updated_at"`
}

type SortRecord struct {
	BagID    string `json:"bag_id"`
	FlightID string `json:"flight_id"`
	ChuteID  string `json:"chute_id"`
	Sequence uint64 `json:"sequence"`
	At       string `json:"at"`
}

type TransferCommand struct {
	ID        string `json:"id"`
	FlightID  string `json:"flight_id"`
	Carousel  string `json:"carousel"`
	Batch     uint64 `json:"batch"`
	CreatedAt string `json:"created_at"`
}

type LoadRecord struct {
	BagID    string `json:"bag_id"`
	Carousel string `json:"carousel"`
	Batch    uint64 `json:"batch"`
}

type SnapshotStats struct {
	Flights         int `json:"flights"`
	Bags            int `json:"bags"`
	ChuteAssignments int `json:"chute_assignments"`
	SortRecords     int `json:"sort_records"`
	TransferCommands int `json:"transfer_commands"`
}

type RouteDetail struct {
	BagID    string `json:"bag_id"`
	FlightID string `json:"flight_id"`
	ChuteID  string `json:"chute_id"`
}

type Sequence struct {
	Epoch uint64 `json:"epoch"`
	Next  uint64 `json:"next"`
}

type Snapshot struct {
	Flights          map[string]Flight  `json:"flights"`
	Bags             map[string]Bag     `json:"bags"`
	FlightMappings   map[string]string  `json:"flight_mappings"`
	ChuteAssignments map[string]string  `json:"chute_assignments"`
	TransferCommands []TransferCommand  `json:"transfer_commands"`
	Sequence         Sequence           `json:"sequence"`
	SortRecords      []SortRecord       `json:"sort_records"`
	// CommittedReads records the bag IDs whose barcode read has been committed.
	// Dispatch must not issue sort instructions for bags absent from this set.
	CommittedReads   map[string]bool    `json:"committed_reads"`
}

func NewSnapshot() *Snapshot {
	return &Snapshot{
		Flights:          map[string]Flight{},
		Bags:             map[string]Bag{},
		FlightMappings:   map[string]string{},
		ChuteAssignments: map[string]string{},
		TransferCommands: []TransferCommand{},
		SortRecords:      []SortRecord{},
		CommittedReads:   map[string]bool{},
	}
}
