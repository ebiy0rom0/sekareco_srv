package interactor

import (
	"sekareco_srv/usecase/inputport"

	"github.com/google/wire"
)

var InteractorProviderSet = wire.NewSet(
	NewAuthInteractor,
	NewMusicInteractor,
	NewPersonInteractor,
	NewRecordInteractor,
	wire.Bind(new(inputport.AuthInputport), new(*authInteractor)),
	wire.Bind(new(inputport.MusicInputport), new(*musicInteractor)),
	wire.Bind(new(inputport.PersonInputport), new(*personInteractor)),
	wire.Bind(new(inputport.RecordInputport), new(*recordInteractor)),
)
