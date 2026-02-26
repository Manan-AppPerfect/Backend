package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Manan-AppPerfect/Backend/internal/team/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	
	Lifecycle fx.Lifecycle
	Service *service.Service
}

func NewHTTPServer(p Params) {

	mux := http.NewServeMux()

	mux.HandleFunc("/create-team", func(w http.ResponseWriter, r *http.Request) {
		fcBarcelona, err := p.Service.CreateTeam("1", "FC Barcelona", 126)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		fmt.Fprintf(w, "Team created: %v\n", fcBarcelona)
	})

	server := &http.Server{
		Addr: ":8000",
		Handler: mux,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting server at 8000: ")

			go server.ListenAndServe()
			
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping the server... ")
			return server.Shutdown(ctx)
		},
	})
}