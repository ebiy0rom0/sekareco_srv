package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"

	"github.com/ebiy0rom0/errors"
)

type personRepository struct {
	infra.SqlHandler
}

func NewPersonRepository(h infra.SqlHandler) *personRepository {
	return &personRepository{h}
}

func (r *personRepository) Store(ctx context.Context, p model.Person) (int, error) {
	query := "INSERT INTO person (person_name, friend_code, is_compare)"
	query += " VALUES (?, ?, ?);"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.Execute(ctx, query, p.PersonName, p.FriendCode, p.IsCompare)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return int(newID64), nil
}

func (r *personRepository) GetByID(ctx context.Context, personID int) (model.Person, error) {
	query := "SELECT person_id, person_name, friend_code FROM person WHERE person_id = ?;"
	row := r.QueryRow(ctx, query, personID)

	var (
		personName string
		friendCode int
	)
	if err := row.Scan(&personID, &personName, &friendCode); err != nil {
		return model.Person{}, errors.New(err.Error())
	}

	user := model.Person{
		PersonID:   personID,
		PersonName: personName,
		FriendCode: friendCode,
	}
	return user, nil
}

var _ database.PersonRepository = (*personRepository)(nil)
