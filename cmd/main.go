package main

import (
	"github.com/Manan-AppPerfect/Backend/internal/common/database"
	"github.com/Manan-AppPerfect/Backend/internal/gateway/user"
	"github.com/Manan-AppPerfect/Backend/internal/server"
	"github.com/Manan-AppPerfect/Backend/internal/team"
	"go.uber.org/fx"
)

// This is the conventional way of doing it without fx

// func main() {
// 	repo := repository.NewRepo()
// 	add := service.NewService(repo)

// 	fcBarcelona, err := add.CreateTeam(
// 		"1",
// 		"FC Barcelona",
// 		126,
// 	)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	add.AddPlayers("1", 32)
// 	add.AddPoints("1", 61)

// 	fmt.Println("Team Created: ", fcBarcelona)

// }


func main() {
	fx.New(
		// Register constructors
		// fx.Provide(
		// 	repository.NewRepo,
		// 	service.NewService,
		// ),

		// Trigger them
		database.DatabaseModule,
		team.TeamModule,
		server.ServerModule,
		user.GatewayModule,
	).Run()
}