package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"

	"github.com/pkg/errors"
)

type personRepository struct {
	infra.SqlHandler
}

func NewPersonRepository(h infra.SqlHandler) *personRepository {
	return &personRepository{h}
}

func (r *personRepository) Store(ctx context.Context, p model.Person) (personID int, err error) {
	query := "INSERT INTO person (person_name, friend_code, is_compare)"
	query += " VALUES (?, ?, ?);"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.Execute(ctx, query, p.PersonName, p.FriendCode, p.IsCompare)
	if err != nil {
		err = errors.Wrap(err, "failed to execute store person")
		return
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return
	}

	personID = int(newID64)
	return
}

func (r *personRepository) GetByID(ctx context.Context, personID int) (user model.Person, err error) {
	query := "SELECT person_id, person_name, friend_code FROM person WHERE person_id = ?;"
	row := r.QueryRow(ctx, query, personID)

	var (
		personName string
		friendCode int
	)
	err = row.Scan(&personID, &personName, &friendCode)
	if err != nil {
		return
	}

	user.PersonID = personID
	user.PersonName = personName
	user.FriendCode = friendCode
	return
}

var _ database.PersonRepository = (*personRepository)(nil)
