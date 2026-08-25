package console

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/belt"
	"bagsort/internal/model"
)

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func readBody(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid json"})
		return false
	}
	return true
}

func (s *Server) handleListFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := s.book.List()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, flights)
}

func (s *Server) handleRegisterFlight(w http.ResponseWriter, r *http.Request) {
	var f model.Flight
	if !readBody(w, r, &f) {
		return
	}
	if err := s.book.Register(f); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, f)
}

func (s *Server) handleCloseFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.book.Close(id, s.injector); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "closed"})
}

func (s *Server) handleSetCutoff(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload struct {
		Cutoff string `json:"cutoff"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.book.SetCutoff(id, payload.Cutoff); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleSetCarousel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload struct {
		Carousel string `json:"carousel"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.book.SetCarousel(id, payload.Carousel); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleSetBatch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload struct {
		Batch uint64 `json:"batch"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.book.SetBatch(id, payload.Batch); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleCheckIn(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Barcode  string `json:"barcode"`
		FlightID string `json:"flight_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	bag, err := s.desk.CheckIn(payload.Barcode, payload.FlightID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, bag)
}

func (s *Server) handleScheduleChange(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		BagID    string `json:"bag_id"`
		FlightID string `json:"flight_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.book.ScheduleChange(payload.BagID, payload.FlightID); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleBagFlight(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	flightID, err := s.book.BagFlight(id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"flight_id": flightID})
}

func (s *Server) handleAssignChute(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		ChuteID  string `json:"chute_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.sorter.AssignChute(payload.FlightID, payload.ChuteID); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleReassignChute(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		ChuteID  string `json:"chute_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.chutes.Reassign(payload.FlightID, payload.ChuteID); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleCurrentChute(w http.ResponseWriter, r *http.Request) {
	flight := chi.URLParam(r, "flight")
	chuteID, err := s.sorter.CurrentChute(flight)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"chute_id": chuteID})
}

func (s *Server) handleSort(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Bag     model.Bag `json:"bag"`
		ChuteID string    `json:"chute_id"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.injector.Push(payload.Bag, payload.ChuteID); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleRetry(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	if err := s.injector.Retry(bag); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleInTransit(w http.ResponseWriter, r *http.Request) {
	flight := chi.URLParam(r, "flight")
	writeJSON(w, 200, map[string]int{"in_transit": s.injector.InTransit(flight)})
}

func (s *Server) handleEnqueueRecheck(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	s.recheckQ.Enqueue(bag)
	writeJSON(w, 201, map[string]int{"len": s.recheckQ.Len()})
}

func (s *Server) handleNextRecheck(w http.ResponseWriter, r *http.Request) {
	bag, ok := s.recheck.Next()
	if !ok {
		writeJSON(w, 404, map[string]string{"error": "empty"})
		return
	}
	writeJSON(w, 200, bag)
}

func (s *Server) handleRecheckDeadline(w http.ResponseWriter, r *http.Request) {
	flight := chi.URLParam(r, "flight")
	deadline, err := s.recheck.Deadline(flight)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"deadline": deadline})
}

func (s *Server) handleBeltSwitch(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Carousel string `json:"carousel"`
		Batch    uint64 `json:"batch"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if err := s.belt.Switch(payload.Carousel, payload.Batch); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleBeltTransfers(w http.ResponseWriter, r *http.Request) {
	cmds, err := s.belt.TransferCommands()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, cmds)
}

func (s *Server) handleAppendTransfer(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		Carousel string `json:"carousel"`
		Batch    uint64 `json:"batch"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	if !belt.ValidateTransfer(model.TransferCommand{FlightID: payload.FlightID, Carousel: payload.Carousel, Batch: payload.Batch}) {
		writeJSON(w, 400, map[string]string{"error": "invalid transfer"})
		return
	}
	if err := s.belt.AppendTransfer(payload.FlightID, payload.Carousel, payload.Batch); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, map[string]string{"status": "ok"})
}

func (s *Server) handleBeltPlace(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	s.belt.Place(bag)
	writeJSON(w, 201, map[string]int{"on_carousel": len(s.belt.OnCarousel())})
}

func (s *Server) handleBeltRotate(w http.ResponseWriter, r *http.Request) {
	if err := s.belt.Rotate(); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleBeltLoads(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.belt.Loads())
}

func (s *Server) handleSetQuota(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FlightID string `json:"flight_id"`
		Limit    int    `json:"limit"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	s.quota.SetLimit(payload.FlightID, payload.Limit)
	writeJSON(w, 200, map[string]string{"status": "ok"})
}

func (s *Server) handleQuota(w http.ResponseWriter, r *http.Request) {
	flight := chi.URLParam(r, "flight")
	writeJSON(w, 200, map[string]int{
		"limit":     s.quota.Limit(flight),
		"available": s.quota.Available(flight, 0),
		"allowed":   boolInt(s.quota.Allowed(flight, 0)),
	})
}

func (s *Server) handleAudit(w http.ResponseWriter, r *http.Request) {
	records, err := s.audit.List()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, records)
}

func (s *Server) handleAuditCount(w http.ResponseWriter, r *http.Request) {
	count, err := s.audit.Count()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]int{"count": count})
}

func (s *Server) handleVerdict(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Epoch uint64 `json:"epoch"`
		Seq   uint64 `json:"seq"`
	}
	if !readBody(w, r, &payload) {
		return
	}
	fresh := s.sorter.Verdict(payload.Epoch, payload.Seq)
	s.sorter.MarkSeen(payload.Epoch, payload.Seq)
	writeJSON(w, 200, map[string]bool{"fresh": fresh})
}
