//go:build wireinject
// +build wireinject

package registory

import (
	domain "sekareco_srv/domain/infra"
	"sekareco_srv/interface/database"
	"sekareco_srv/interface/handler"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/interactor"
	"sekareco_srv/usecase/validator"

	"github.com/google/wire"
)

func InitializeDIContainer(sqlHandler infra.SqlHandler, txHandler infra.TxHandler, manager domain.TokenManager) *handler.AppContainer {
	wire.Build(
		handler.AppContainerProviderSet,
		validator.ValidatorProviderSet,
		interactor.InteractorProviderSet,
		database.DatabasePrividerSet,
	)
	return &handler.AppContainer{}
}
