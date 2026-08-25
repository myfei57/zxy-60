package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/checkin"
	"bagsort/internal/model"
)

func (s *Server) handleStorePath(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]string{"path": s.store.Path()})
}

func (s *Server) handleListBags(w http.ResponseWriter, r *http.Request) {
	bags, err := s.book.ListBags()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, bags)
}

func (s *Server) handleGetBag(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bag, ok, err := s.book.GetBag(id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, 200, bag)
}

func (s *Server) handlePutBag(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	if err := s.book.PutBag(bag); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, bag)
}

func (s *Server) handleGetFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	f, ok, err := s.store.GetFlight(id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, 200, f)
}

func (s *Server) handlePutFlight(w http.ResponseWriter, r *http.Request) {
	var f model.Flight
	if !readBody(w, r, &f) {
		return
	}
	if err := s.store.PutFlight(f); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, f)
}

func (s *Server) handleValidateFlight(w http.ResponseWriter, r *http.Request) {
	var f model.Flight
	if !readBody(w, r, &f) {
		return
	}
	if err := s.book.Validate(f); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "valid"})
}

func (s *Server) handleOpenFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.book.Open(id); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "open"})
}

func (s *Server) handleBoardFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.book.Board(id); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "boarding"})
}

func (s *Server) handleCutoffFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.book.Cutoff(id); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "cutoff"})
}

func (s *Server) handleFlightBatch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	batch, err := s.book.Batch(id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]uint64{"batch": batch})
}

func (s *Server) handleListAssignments(w http.ResponseWriter, r *http.Request) {
	assignments, err := s.store.ListChuteAssignments()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, assignments)
}

func (s *Server) handlePutAssignment(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		ChuteID  string `json:"chute_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.store.PutChuteAssignment(payload.FlightID, payload.ChuteID); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleGetAssignment(w http.ResponseWriter, r *http.Request) {
	flightID := chi.URLParam(r, "flight")
	chuteID, ok, err := s.store.GetChuteAssignment(flightID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, 200, map[string]string{"chute_id": chuteID})
}

func (s *Server) handleListChutes(w http.ResponseWriter, r *http.Request) {
	assignments, err := s.chutes.List()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, assignments)
}

func (s *Server) handleValidateChute(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ChuteID string `json:"chute_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	writeJSON(w, 200, map[string]bool{"valid": s.chutes.Validate(payload.ChuteID)})
}

func (s *Server) handleDivert(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	if err := s.sorter.Divert(bag); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "diverted"})
}

func (s *Server) handleSortRecords(w http.ResponseWriter, r *http.Request) {
	records, err := s.sorter.ListRecords()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, records)
}

func (s *Server) handleBeltCurrent(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{
		"carousel": s.belt.Carousel(),
		"batch":    s.belt.CurrentBatch(),
	})
}

func (s *Server) handleBeltPlaceBatch(w http.ResponseWriter, r *http.Request) {
	var bags []model.Bag
	if !readBody(w, r, &bags) {
		return
	}
	s.belt.PlaceBatch(bags)
	writeJSON(w, 201, map[string]int{"on_carousel": len(s.belt.OnCarousel())})
}

func (s *Server) handleAuditSummary(w http.ResponseWriter, r *http.Request) {
	summary, err := s.audit.Summary()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, summary)
}

func (s *Server) handleAuditLast(w http.ResponseWriter, r *http.Request) {
	record, ok, err := s.audit.Last()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "empty"})
		return
	}
	writeJSON(w, 200, record)
}

func (s *Server) handleQuotaConsume(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		Count    int    `json:"count"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	writeJSON(w, 200, map[string]bool{"allowed": s.quota.Consume(payload.FlightID, payload.Count)})
}

func (s *Server) handleQuotaRelease(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		Count    int    `json:"count"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	s.quota.Release(payload.FlightID, payload.Count)
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleQuotaLimits(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.quota.Limits())
}

func (s *Server) handleRecheckPeek(w http.ResponseWriter, r *http.Request) {
	bag, ok := s.recheckQ.Peek()
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "empty"})
		return
	}
	writeJSON(w, 200, bag)
}

func (s *Server) handleRecheckItems(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.recheckQ.Items())
}

func (s *Server) handleTagSequence(w http.ResponseWriter, r *http.Request) {
	seq, err := s.reader.LastSequence()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, seq)
}

func (s *Server) handleTagBatch(w http.ResponseWriter, r *http.Request) {
	var barcodes []string
	if !readBody(w, r, &barcodes) {
		return
	}
	readings, err := s.reader.ReadBatch(barcodes)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, readings)
}

func (s *Server) handleFindBag(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bag, ok, err := s.desk.FindBag(id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, 200, bag)
}

func (s *Server) handleCheckInBatch(w http.ResponseWriter, r *http.Request) {
	var entries []checkin.Entry
	if !readBody(w, r, &entries) {
		return
	}
	bags, err := s.desk.CheckInBatch(entries)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, bags)
}

func (s *Server) handleStoreSequence(w http.ResponseWriter, r *http.Request) {
	seq, err := s.store.Sequence()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, seq)
}

func (s *Server) handlePutStoreSequence(w http.ResponseWriter, r *http.Request) {
	var seq model.Sequence
	if !readBody(w, r, &seq) {
		return
	}
	if err := s.store.PutSequence(seq); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, seq)
}

func (s *Server) handleStoreFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := s.store.ListFlights()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, flights)
}

func (s *Server) handleStoreRecords(w http.ResponseWriter, r *http.Request) {
	records, err := s.store.ListSortRecords()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, records)
}
