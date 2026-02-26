package repository

import "github.com/Manan-AppPerfect/Backend/internal/team/models"

// Repository.go handles the data storage
/* 	It is responsible for:
	Saving data
	Retrieving data
It hides where data is stored.
*/

type Repo struct {
	data map[string]*models.Team
}

func NewRepo() *Repo {
	return &Repo{
		data: make(map[string]*models.Team),
	}
}

func (r *Repo) Save(team *models.Team) {
	r.data[team.ID] = team
}

func (r *Repo) Get(id string) *models.Team {
	return r.data[id]
}