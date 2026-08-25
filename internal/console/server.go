package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"bagsort/internal/audit"
	"bagsort/internal/belt"
	"bagsort/internal/checkin"
	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/inject"
	"bagsort/internal/ns"
	"bagsort/internal/quota"
	"bagsort/internal/recheck"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
	"bagsort/internal/tag"
)

type Server struct {
	router   *chi.Mux
	store    *store.Store
	ns       *ns.Namespace
	book     *flight.Book
	chutes   *chute.Assigner
	sorter   *sorter.Sorter
	injector *inject.Injector
	reader   *tag.Reader
	recheck  *recheck.Processor
	recheckQ *recheck.Queue
	belt     *belt.Belt
	quota    *quota.Checker
	audit    *audit.Logger
	desk     *checkin.Desk
}

func New(dataDir string) (*Server, error) {
	st := store.New(dataDir, "bagsort.json")
	namespace := ns.New("T1")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	sortEngine := sorter.NewSorter(st, book, chutes)
	injector := inject.NewInjector(st, book, sortEngine)
	reader := tag.NewReader(st, namespace)
	quotaChecker := quota.NewChecker()
	auditLogger := audit.NewLogger(st)
	desk := checkin.NewDesk(reader, book, chutes, injector, quotaChecker)
	recheckQueue := recheck.NewQueue()
	recheckProcessor := recheck.NewProcessor(recheckQueue, book)
	beltEngine := belt.NewBelt(st, book)

	s := &Server{
		router:   chi.NewRouter(),
		store:    st,
		ns:       namespace,
		book:     book,
		chutes:   chutes,
		sorter:   sortEngine,
		injector: injector,
		reader:   reader,
		recheck:  recheckProcessor,
		recheckQ: recheckQueue,
		belt:     beltEngine,
		quota:    quotaChecker,
		audit:    auditLogger,
		desk:     desk,
	}
	s.router.Use(s.withRecovery)
	s.router.Use(s.withLogging)
	s.routes()
	return s, nil
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
