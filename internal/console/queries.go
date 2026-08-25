package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/model"
)

func (s *Server) handleSortLastChute(w http.ResponseWriter, r *http.Request) {
	bagID := chi.URLParam(r, "bag")
	chuteID, err := s.sorter.LastChute(bagID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"chute_id": chuteID})
}

func (s *Server) handleSortCount(w http.ResponseWriter, r *http.Request) {
	bagID := chi.URLParam(r, "bag")
	count, err := s.sorter.SortCount(bagID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]int{"count": count})
}

func (s *Server) handleAuditRecord(w http.ResponseWriter, r *http.Request) {
	var record model.SortRecord
	if !readBody(w, r, &record) {
		return
	}
	if err := s.audit.Record(record); err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, record)
}

func (s *Server) handleAuditForBag(w http.ResponseWriter, r *http.Request) {
	bagID := chi.URLParam(r, "bag")
	records, err := s.audit.ForBag(bagID)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, records)
}

func (s *Server) handleBeltReplay(w http.ResponseWriter, r *http.Request) {
	cmds, err := s.belt.ReplayTransfers()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, cmds)
}
