package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/model"
)

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	flights, err := s.book.List()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	bags, err := s.store.ListBags()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	assignments, err := s.store.ListChuteAssignments()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	openCount := 0
	for _, f := range flights {
		if model.IsOpenFlightState(f.State) {
			openCount++
		}
	}
	activeCount := 0
	for _, b := range bags {
		if model.IsActiveBagState(b.State) {
			activeCount++
		}
	}
	writeJSON(w, 200, map[string]any{
		"flight_count":  len(flights),
		"open_flights":  openCount,
		"bag_count":     len(bags),
		"active_bags":   activeCount,
		"chute_count":   len(assignments),
		"terminal":      s.ns.Terminal(),
	})
}

func (s *Server) handleFindFlightByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	f, ok, err := s.book.FindByCode(code)
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

func (s *Server) handleOpenFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := s.book.OpenFlights()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, flights)
}

func (s *Server) handleFlightAuditSummary(w http.ResponseWriter, r *http.Request) {
	flightID := chi.URLParam(r, "flight")
	summary, err := s.audit.FlightSummary(flightID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, summary)
}
