package validator

import (
	"sekareco_srv/usecase/inputport"

	"github.com/google/wire"
)

var ValidatorProviderSet = wire.NewSet(
	NewAuthValidator,
	NewPersonValidator,
	wire.Bind(new(inputport.AuthValidator), new(*authValidator)),
	wire.Bind(new(inputport.PersonValidator), new(*personValidator)),
)
