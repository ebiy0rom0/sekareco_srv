//go:build wireinject
// +build wireinject

package registory

import (
	"sekareco_srv/interface/handler"
	"sekareco_srv/interface/handler/auth"
	"sekareco_srv/interface/handler/health"
	"sekareco_srv/interface/handler/music"
	"sekareco_srv/interface/handler/person"
	"sekareco_srv/interface/handler/record"
	"sekareco_srv/interface/infra"

	"github.com/google/wire"
)

func InitializeDIContainer(sqlHandler infra.SqlHandler) *handler.AppContainer {
	wire.Build(
		handler.NewAppContainer,
		auth.NewAuthHandler,
		health.NewHealthHandler,
		music.NewMusicHandler,
		person.NewPersonHandler,
		record.NewRecordHandler,
	)
	return &handler.AppContainer{}
}
