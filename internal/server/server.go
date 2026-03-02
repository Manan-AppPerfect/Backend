package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Manan-AppPerfect/Backend/internal/api/generated"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	
	Lifecycle fx.Lifecycle
	Handler api.ServerInterface
}

func NewHTTPServer(p Params) {

	mux := chi.NewRouter()

	api.HandlerFromMux(p.Handler, mux)

	server := &http.Server{
		Addr: ":8000",
		Handler: mux,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting server at 8000: ")

			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Println("Server error:", err)
				}
			}()
			
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping the server... ")
			return server.Shutdown(ctx)
		},
	})
}