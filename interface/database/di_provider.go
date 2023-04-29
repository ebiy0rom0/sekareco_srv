package database

import (
	"sekareco_srv/usecase/database"

	"github.com/google/wire"
)

var DatabasePrividerSet = wire.NewSet(
	NewLoginRepository,
	NewMusicRepository,
	NewPersonRepository,
	NewRecordRepository,
	NewTransaction,
	wire.Bind(new(database.LoginRepository), new(*loginRepository)),
	wire.Bind(new(database.MusicRepository), new(*musicRepository)),
	wire.Bind(new(database.PersonRepository), new(*personRepository)),
	wire.Bind(new(database.RecordRepository), new(*recordRepository)),
	wire.Bind(new(database.SqlTransaction), new(*tx)),
)
