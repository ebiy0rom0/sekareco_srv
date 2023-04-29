package handler

import (
	"sekareco_srv/interface/handler/auth"
	"sekareco_srv/interface/handler/health"
	"sekareco_srv/interface/handler/music"
	"sekareco_srv/interface/handler/person"
	"sekareco_srv/interface/handler/record"

	"github.com/google/wire"
)

type AppContainer struct {
	Auth   auth.Handler
	Health health.Handler
	Music  music.Handler
	Person person.Handler
	Record record.Handler
}

func NewAppContainer(
	authHandler auth.Handler,
	healthHandler health.Handler,
	musicHandler music.Handler,
	personHandler person.Handler,
	recordHandler record.Handler,
) *AppContainer {
	return &AppContainer{
		Auth:   authHandler,
		Health: healthHandler,
		Music:  musicHandler,
		Person: personHandler,
		Record: recordHandler,
	}
}

var AppContainerProviderSet = wire.NewSet(
	NewAppContainer,
	auth.AuthHandlerProviderSet,
	health.HealthHandlerProviderSet,
	music.MusicHandlerProviderSet,
	person.PersonHandlerProviderSet,
	record.RecordHandlerProviderSet,
)
