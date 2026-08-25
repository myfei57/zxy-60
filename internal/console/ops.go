package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/model"
	"bagsort/internal/tag"
)

func (s *Server) handleStoreStats(w http.ResponseWriter, r *http.Request) {
	stats, err := s.store.Stats()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, stats)
}

func (s *Server) handleStoreClear(w http.ResponseWriter, r *http.Request) {
	if err := s.store.Clear(); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "cleared"})
}

func (s *Server) handleRouteDetail(w http.ResponseWriter, r *http.Request) {
	var bag model.Bag
	if !readBody(w, r, &bag) {
		return
	}
	detail, err := s.sorter.RouteDetail(bag)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, detail)
}

func (s *Server) handleCommitReadings(w http.ResponseWriter, r *http.Request) {
	var readings []tag.Reading
	if !readBody(w, r, &readings) {
		return
	}
	if err := s.reader.CommitBatch(readings); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]int{"committed": len(readings)})
}

func (s *Server) handleCancelCheckIn(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.desk.CancelCheckIn(id); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"status": "cancelled"})
}
