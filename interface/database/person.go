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

func NewPersonRepository(h infra.SqlHandler) database.PersonRepository {
	return &personRepository{h}
}

func (r *personRepository) Store(ctx context.Context, p model.Person) (int, error) {
	query := `
	INSERT INTO person (
		person_name, friend_code, is_compare
	) VALUES (
		:person_name, :friend_code, :is_compare
	);`

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.ExecNamedContext(ctx, query, p)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(newID64), nil
}

func (r *personRepository) GetByID(ctx context.Context, personID int) (model.Person, error) {
	query := `SELECT * FROM person WHERE person_id = $1;`

	var person model.Person
	if err := r.GetContext(ctx, &person, query, personID); err != nil {
		return model.Person{}, errors.WithStack(err)
	}
	return person, nil
}

func (r *personRepository) AddFriendCode(ctx context.Context, p model.Person) error {
	query := `UPDATE person SET friend_code = :friend_code WHERE person_id = ?;`

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	_, err := dao.UpdateNamedContext(ctx, query, p, p.PersonID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *personRepository) GetByFriendCode(ctx context.Context, code int) (model.Person, error) {
	query := `SELECT * FROM person WHERE friend_code = $1;`

	var person model.Person
	if err := r.GetContext(ctx, &person, query, code); err != nil {
		return model.Person{}, errors.WithStack(err)
	}
	return person, nil
}

var _ database.PersonRepository = (*personRepository)(nil)
