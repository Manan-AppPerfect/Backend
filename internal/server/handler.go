package server

import (
	"encoding/json"
	"net/http"

	"github.com/Manan-AppPerfect/Backend/internal/api/generated"
	"github.com/Manan-AppPerfect/Backend/internal/team/service"
)

type TeamHandler struct{
	service *service.Service
}

func NewTeamHandler(s *service.Service) api.ServerInterface {
	return &TeamHandler{
		service: s,
	}
}

func (h *TeamHandler) PostTeams(w http.ResponseWriter, r *http.Request) {
	
	//Decode Request
	var req api.CreateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call Service
	team, err := h.service.CreateTeam(
		r.Context(),
		req.Name, 
		req.Age,
		req.PlayerCount,
		req.Points,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Map domain -> Response 
	resp := api.Team{
		Id:          intPtr(int(team.ID)),
		Name:        stringPtr(team.Name),
		Age:         intPtr(team.Age),
		PlayerCount: intPtr(team.PlayerCount),
		Points:      intPtr(team.Points),
	}

	// Write Response
	json.NewEncoder(w).Encode(resp)
}


func (h *TeamHandler) GetTeamsId(w http.ResponseWriter, r *http.Request, id int) {
	
	team, err := h.service.GetTeamByID(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := api.Team{
		Id:          intPtr(int(team.ID)),
		Name:        stringPtr(team.Name),
		Age:         intPtr(team.Age),
		PlayerCount: intPtr(team.PlayerCount),
		Points:      intPtr(team.Points),
	}

	json.NewEncoder(w).Encode(resp)

}
// Helper Func 

func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}