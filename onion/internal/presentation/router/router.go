package router

import (
	"context"
	"fmt"
	"net/http"
	usecase "onion/internal/usecase/interfaces"
)

func New(port int, logger usecase.Logger) *Router {
	return &Router{
		mux:    http.NewServeMux(),
		port:   port,
		logger: logger,
	}
}

type Router struct {
	mux    *http.ServeMux
	port   int
	logger usecase.Logger
}

func (r *Router) Run() error {
	routerWithMiddleware := LoggingMiddleware(r.logger, r.mux)
	r.logger.Info(context.Background(), fmt.Sprintf("Server is running on port %d", r.port))
	addr := fmt.Sprintf(":%d", r.port)
	if err := http.ListenAndServe(addr, routerWithMiddleware); err != nil {
		return err
	}
	return nil
}

func (r *Router) AddFetchArticlesHandler(uc usecase.FetchArticlesUsecase) {
	const path = "GET /articles"
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if err := uc.Run(req.Context()); err != nil {
			r.logger.Error(req.Context(), "Failed to fetch articles", fmt.Sprintf("%+v", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
	r.logger.Info(context.Background(), fmt.Sprintf("Added handler for %s", path))
}
